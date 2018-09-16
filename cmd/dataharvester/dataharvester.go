package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/amwolff/oa/pkg/common"
	"github.com/amwolff/oa/pkg/dataharvest"
	"github.com/amwolff/oa/pkg/municommodels"
	"github.com/davecgh/go-spew/spew"
	"github.com/fsnotify/fsnotify"
	"github.com/gocraft/dbr"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type config struct {
	ClientName string
	ClientUA   string
	ClientURL  string

	DbHost string
	DbPort int
	DbName string
	DbUser string
	DbPass string

	LogLevel    string
	LogDir      string
	ForceColors bool

	SentrySecret string
}

func loadConfig(log logrus.FieldLogger) (cfg config) {
	viper.Set("Verbose", true)
	viper.Set("LogFile", os.Stderr)

	pflag.String("clientName", "dataharvester", "Web Service Client name")
	pflag.String("clientUA", "oaservice/1.0 (+https://ssdip.bip.gov.pl/artykuly/art-61-konstytucji-rp_75.html)", "Web Service Client UAString")
	pflag.String("clientURL", "http://sip.zdzit.olsztyn.eu/PublicService.asmx", "Web Service URL")

	pflag.String("dbHost", "localhost", "Database host")
	pflag.Int("dbPort", 5432, "Database port")
	pflag.String("dbName", "oadb", "Database name")
	pflag.String("dbUser", "data_service", "Database user")
	pflag.String("dbPass", "data_service", "Database password")

	pflag.String("loglevel", "debug", "Logging level")
	pflag.String("logDir", "/tmp", "Logs directory")
	pflag.Bool("forceColors", true, "Force colors when printing to stdout")

	pflag.String("sentrySecret", "", "Secret string for Raven")

	configFile := pflag.String("config", "", "A config file to load")

	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	if len(*configFile) > 0 {
		viper.AddConfigPath(*configFile)
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			// TODO(amw): maybe unmarshal 2nd time
			log.Warnf("Config file changed: %s", e.Name)
		})
		if err := viper.ReadInConfig(); err != nil {
			log.WithError(err).Fatal("Cannot read config")
		}
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.WithError(err).Fatal("Cannot unmarshal config")
	}

	return
}

func calibrate(client *municommodels.WebServiceClient, log logrus.FieldLogger,
	cookies []http.Cookie, available *municommodels.GetRouteAndVariantsResponse) (bool, error) {

	for _, r := range available.GetRouteAndVariantsResult.L {
		payload := municommodels.CNRGetVehicles{
			R: r.Number,
			D: r.Direction,
		}
		log.Infof("calibrate: trying %s (%s)", payload.R, payload.D)

		durationPool := 30 * time.Second
		var actual, previous float64
		for durationPool > 0 {
			if previous != 0 && previous != actual {
				log.Debugf("calibrate: %f != %f", previous, actual)
				return true, nil
			}
			previous = actual

			now := time.Now()
			ctx, canc := context.WithTimeout(context.Background(), durationPool)
			v, err := client.CallCNRGetVehicles(ctx, cookies, payload)
			if err != nil {
				canc()
				return false, err
			}
			canc()

			s := v.CNRGetVehiclesResult.Sanitized
			if s == nil {
				break
			}

			actual = s[0].Szerokosc

			time.Sleep(time.Second)
			durationPool -= time.Since(now)
		}
	}
	return false, nil
}

func insertRoutesChunk(dbS dbr.SessionRunner, log logrus.FieldLogger,
	chunk *municommodels.GetRouteAndVariantsResponse, t time.Time) {

	log = log.WithField("func", "insertRoutesChunk")

	if len(chunk.GetRouteAndVariantsResult.L) > 0 {
		if err := dataharvest.InsertGetRouteAndVariantsResponseIntoDb(dbS, chunk, t); err != nil {
			log.WithError(err).Fatal("InsertGetRouteAndVariantsResponseIntoDb")
		}
		log.Infof("Inserted %d vehicles", len(chunk.GetRouteAndVariantsResult.L))
		return
	}
	log.Error("Zero-length data chunk")
}

func insertVehiclesChunk(dbS dbr.SessionRunner, log logrus.FieldLogger,
	chunk []*municommodels.CNRGetVehiclesResponse, t time.Time) {

	log = log.WithField("func", "insertVehiclesChunk")

	if len(chunk) > 0 {
		if err := dataharvest.InsertCNRGetVehiclesResponsesIntoDb(dbS, chunk, t); err != nil {
			log.WithError(err).Error("InsertCNRGetVehiclesResponsesIntoDb")
		}

		var cnt int
		for _, c := range chunk {
			cnt += len(c.CNRGetVehiclesResult.Sanitized)
		}
		log.Infof("Inserted %d vehicles", cnt)
		return
	}
	log.Warn("Skip inserting zero-length data chunk")
}

// func serviceShutdown(callback func()) {
// 	s := make(chan os.Signal, 1)
// 	signal.Notify(s, os.Interrupt)
// 	signal.Notify(s, os.Kill)
// 	signal.Notify(s, syscall.SIGTERM)
// 	go func() {
// 		<-s
// 		callback()
// 		os.Exit(0)
// 	}()
// }

var (
	BuildTimeCommitMD5 string
	BuildTimeTime      string
	BuildTimeIsDev     = "false"
)

func main() {
	isDev, err := strconv.ParseBool(BuildTimeIsDev)
	if err != nil {
		panic(fmt.Sprintf("Cannot parse BuildTimeIsDev: %v", err))
	}

	log := logrus.New()
	initFields := logrus.Fields{
		"commit-md5": BuildTimeCommitMD5,
		"build-time": BuildTimeTime,
		"is-dev":     isDev,
	}
	if isDev {
		log.SetLevel(logrus.DebugLevel)
	}

	log.WithFields(initFields).Info("dataharvester greeting")

	cfg := loadConfig(log)
	log.Infof("Loaded config: %s", spew.Sdump(cfg))

	client := municommodels.NewWebServiceClient(log, cfg.ClientName, cfg.ClientUA, cfg.ClientURL)
	sessionCookies := []http.Cookie{{Name: "ASP.NET_SessionId", Value: "m2yxf41efvby2tj4z5m1esiy"}} // TODO(amw): make it configurable

	dsn := common.GetDsn(cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName)
	log.Debugf("DSN: %s", dsn)

	dbConn, err := dbr.Open("postgres", dsn, &dbr.NullEventReceiver{})
	if err != nil {
		log.WithError(err).Fatal("Cannot connect to database")
	}
	dbSess := dbConn.NewSession(&dbr.NullEventReceiver{})

	tz, err := time.LoadLocation("Europe/Warsaw")
	if err != nil {
		log.WithError(err).Fatal("Cannot parse location")
	}

	log.Info("Initialization completed")
	routes := &municommodels.GetRouteAndVariantsResponse{}
	for {
		ctx, canc := context.WithTimeout(context.Background(), time.Second)
		routes, err = client.CallGetRouteAndVariants(ctx, sessionCookies, municommodels.GetRouteAndVariants{})
		if err != nil {
			// canc()
			log.WithError(err).Fatal("CallGetRouteAndVariants")
		}
		canc()

		now := time.Now().In(tz)

		insertRoutesChunk(dbSess, log, routes, now)

		if ok, err := calibrate(client, log, sessionCookies, routes); !ok {
			log.WithError(err).Fatal("Calibration unsuccessful")
		}
		log.Info("Calibration completed")

		for now.Before(time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, tz)) {
			durationPool := 21 * time.Second
			var vehiclesChunk []*municommodels.CNRGetVehiclesResponse
			for _, r := range routes.GetRouteAndVariantsResult.L {
				now = time.Now().In(tz)

				payload := municommodels.CNRGetVehicles{
					R: r.Number,
					D: r.Direction,
				}

				ctx, canc := context.WithTimeout(context.Background(), durationPool)
				vehicles, err := client.CallCNRGetVehicles(ctx, sessionCookies, payload)
				if err != nil {
					canc()
					if vehicles != nil {
						log.WithError(err).Warn("Results' sanitation unsuccessful")
						// TODO(amw): rescue data
					}
					log.WithError(err).Error("CallCNRGetVehicles")
					continue
				}
				canc()

				if len(vehicles.CNRGetVehiclesResult.Sanitized) > 0 {
					vehiclesChunk = append(vehiclesChunk, vehicles)
				}

				durationPool -= time.Since(now)
			}

			go insertVehiclesChunk(dbSess, log, vehiclesChunk, now)

			log.Infof("Will now wait %v", durationPool)
			time.Sleep(durationPool)
			now = time.Now().In(tz)
		}
	}
}
