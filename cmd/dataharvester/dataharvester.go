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
	"github.com/getsentry/raven-go"
	"github.com/gocraft/dbr"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type config struct {
	ClientName   string
	ClientUA     string
	ClientURL    string
	ClientCookie string

	DbHost string
	DbPort int
	DbName string
	DbUser string
	DbPass string

	LogLevel    string
	LogDir      string
	ForceColors bool

	SentryDSN         string
	SentryEnvironment string
}

func loadConfig(log logrus.FieldLogger) (cfg config) {
	viper.Set("Verbose", true)
	viper.Set("LogFile", os.Stderr)

	pflag.String("clientName", "dataharvester", "Web Service Client name")
	pflag.String("clientUA", "oaservice/1.0 (+https://ssdip.bip.gov.pl/artykuly/art-61-konstytucji-rp_75.html)", "Web Service Client UAString")
	pflag.String("clientURL", "http://sip.zdzit.olsztyn.eu/PublicService.asmx", "Web Service URL")
	pflag.String("clientCookie", "o41zqttsghzoiufjjxd01r5t", "ASP.NET_SessionId cookie")

	pflag.String("dbHost", "localhost", "Database host")
	pflag.Int("dbPort", 5432, "Database port")
	pflag.String("dbName", "oadb", "Database name")
	pflag.String("dbUser", "data_service", "Database user")
	pflag.String("dbPass", "data_service", "Database password")

	pflag.String("logLevel", "debug", "Logging level")
	pflag.String("logDir", "/tmp", "Logs directory")
	pflag.Bool("forceColors", true, "Force colors when printing to stdout")

	pflag.String("sentryDSN", "", "Secret string for Raven")
	pflag.String("sentryEnvironment", "", "")

	configFile := pflag.String("config", "", "A config file to load")

	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	if len(*configFile) > 0 {
		viper.SetConfigFile(*configFile)
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			// TODO(amwolff): maybe unmarshal 2nd time
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

func insertRoutesChunk(dbC *dbr.Connection, log logrus.FieldLogger,
	chunk *municommodels.GetRouteAndVariantsResponse, t time.Time) {

	log = log.WithField("func", "insertRoutesChunk")
	errTag := map[string]string{"func": "insertRoutesChunk"}

	dbS := dbC.NewSession(&dbr.NullEventReceiver{})
	if len(chunk.GetRouteAndVariantsResult.L) > 0 {
		if err := dataharvest.InsertGetRouteAndVariantsResponseIntoDb(dbS, chunk, t); err != nil {
			raven.CaptureErrorAndWait(err, errTag)
			log.WithError(err).Fatal("InsertGetRouteAndVariantsResponseIntoDb")
		}
		log.WithField("ts", t).Infof("Inserted %d routes", len(chunk.GetRouteAndVariantsResult.L))
		return
	}
	err := errors.New("Zero-length data chunk")
	log.Error(err)
	raven.CaptureErrorAndWait(err, errTag)
}

func insertVehiclesChunk(dbC *dbr.Connection, log logrus.FieldLogger,
	chunk []*municommodels.CNRGetVehiclesResponse, t time.Time) {

	log = log.WithField("func", "insertVehiclesChunk")
	errTag := map[string]string{"func": "insertVehiclesChunk"}

	dbS := dbC.NewSession(&dbr.NullEventReceiver{})
	if len(chunk) > 0 {
		if err := dataharvest.InsertCNRGetVehiclesResponsesIntoDb(dbS, chunk, t); err != nil {
			log.WithError(err).Error("InsertCNRGetVehiclesResponsesIntoDb")
			raven.CaptureErrorAndWait(err, errTag)
			return
		}

		var cnt int
		for _, c := range chunk {
			cnt += len(c.CNRGetVehiclesResult.Sanitized)
		}
		log.WithField("ts", t).Infof("Inserted %d vehicles", cnt)
		return
	}

	_msg := "Skip inserting zero-length data chunk"
	log.Warn(_msg)
	raven.CaptureMessageAndWait(_msg, errTag)
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
	BuildTimeIsDev     string
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

	tz, err := time.LoadLocation("Europe/Warsaw")
	if err != nil {
		log.WithError(err).Fatal("Cannot parse location")
	}

	raven.SetDefaultLoggerName("dataharvester")
	raven.SetEnvironment(cfg.SentryEnvironment)
	raven.SetRelease(BuildTimeCommitMD5)
	raven.SetDSN(cfg.SentryDSN)

	client := municommodels.NewWebServiceClient(log, cfg.ClientName, cfg.ClientUA, cfg.ClientURL)
	sessionCookies := []http.Cookie{{Name: "ASP.NET_SessionId", Value: cfg.ClientCookie}}

	dsn := common.GetDsn(cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName)
	log.Debugf("DSN: %s", dsn)

	dbConn, err := dbr.Open("postgres", dsn, &dbr.NullEventReceiver{})
	if err != nil {
		log.WithError(err).Fatal("Cannot connect to database")
	}
	dbConn.SetConnMaxLifetime(10 * time.Minute)

	if err := common.WaitForPostgres(dbConn, 10, log); err != nil {
		log.WithError(err).Fatal("Cannot connect to database")
	}

	log.Info("Initialization completed")

	coldStart := true
	routes := &municommodels.GetRouteAndVariantsResponse{}
	for {
		now := time.Now().In(tz)
		maintenanceStart := time.Date(now.Year(), now.Month(), now.Day(), 1, 59, 54, 0, tz)
		maintenanceEnd := time.Date(now.Year(), now.Month(), now.Day(), 3, 59, 54, 0, tz)

		if !coldStart && now.After(maintenanceStart) && now.Before(maintenanceEnd) {
			d := time.Until(maintenanceEnd)
			log.Infof("Will now wait %v", d)
			time.Sleep(d)
		}

		ctx, canc := context.WithTimeout(context.Background(), time.Minute)
		routes, err = client.CallGetRouteAndVariants(ctx, sessionCookies, municommodels.GetRouteAndVariants{})
		if err != nil {
			// canc()
			raven.CaptureErrorAndWait(err, map[string]string{"call-to": "CallGetRouteAndVariants"})
			log.WithError(err).Fatal("CallGetRouteAndVariants")
		}
		canc()

		now = time.Now().In(tz)

		calibrationTime := time.Date(now.Year(), now.Month(), now.Day(), (now.Hour() + 1), 59, 54, 0, tz)
		if coldStart {
			coldStart = false
		}

		insertRoutesChunk(dbConn, log, routes, now)

		log.Info("Calibration completed")

		for now.Before(calibrationTime) {
			durationPool := 6 * time.Second
			var vehiclesChunk []*municommodels.CNRGetVehiclesResponse
			for _, r := range routes.GetRouteAndVariantsResult.L {
				now = time.Now().In(tz)

				if durationPool <= 0 {
					break
				}

				payload := municommodels.CNRGetVehicles{
					R: r.Number,
					// TODO(amwolff): we should probably also query by 'V' to avoid data duplication
					D: r.Direction,
				}

				ctx, canc := context.WithTimeout(context.Background(), durationPool)
				vehicles, err := client.CallCNRGetVehicles(ctx, sessionCookies, payload)
				if err != nil {
					canc()
					if vehicles != nil {
						log.WithError(err).Warn("Results' sanitation unsuccessful")
						// TODO(amwolff): rescue data in case of sanitation error
					}
					raven.CaptureError(err, map[string]string{"call-to": "CallCNRGetVehicles"})
					log.WithError(err).Error("CallCNRGetVehicles")
					durationPool -= time.Since(now)
					continue
				}
				canc()

				if len(vehicles.CNRGetVehiclesResult.Sanitized) > 0 {
					vehiclesChunk = append(vehiclesChunk, vehicles)
				}

				durationPool -= time.Since(now)
			}

			go insertVehiclesChunk(dbConn, log, vehiclesChunk, now)

			log.Infof("Will now wait %v", durationPool)
			time.Sleep(durationPool)
			now = time.Now().In(tz)
		}
		log.Info("Will soon perform calibration")
	}
}
