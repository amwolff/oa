package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/amwolff/oa/pkg/dataharvest"
	"github.com/amwolff/oa/pkg/municommodels"
	"github.com/davecgh/go-spew/spew"
	"github.com/fsnotify/fsnotify"
	"github.com/gocraft/dbr"
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
	pflag.String("clientName", "dataharvester/1.0 (+https://ssdip.bip.gov.pl/artykuly/art-61-konstytucji-rp_75.html)", "Web Service Client UAString")
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

func getDsn(user, pass, host string, port int, name string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, pass, host, port, name)
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
	BuildTimeIsDev     = "true"
)

func main() {
	isDev, err := strconv.ParseBool(BuildTimeIsDev)
	if err != nil {
		panic(fmt.Sprintf("Cannot parse BuildTimeIsDev: %v", err))
	}

	log := logrus.New()
	logFields := logrus.Fields{
		"commit-md5": BuildTimeCommitMD5,
		"build-time": BuildTimeTime,
		"is-dev":     isDev,
	}
	if isDev {
		log.SetLevel(logrus.DebugLevel)
	}

	log.WithFields(logFields).Info("dataharvester greeting")

	cfg := loadConfig(log)
	log.Infof("Loaded config: %s", spew.Sdump(cfg))

	client := municommodels.NewWebServiceClient(
		log,
		cfg.ClientName,
		cfg.ClientUA,

		cfg.ClientURL,
	)
	sessionCookies := []http.Cookie{{Name: "TODO(amw): make it configurable", Value: "TODO(amw): make it configurable"}}

	dsn := getDsn(cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName)
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
			log.WithError(err).Error("CallGetRouteAndVariants")
		}
		canc()

		now := time.Now().In(tz)
		dataharvest.InsertGetRouteAndVariantsResponseIntoDb(dbSess, routes, now)

		// calibration...
		log.Info("Calibration completed")

		for now.Before(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, tz)) {

			// normal fetches...

			time.Sleep(20 * time.Second) // minus requests' time...
			now = time.Now().In(tz)
		}
	}
}
