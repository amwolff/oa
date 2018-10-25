package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	"github.com/fsnotify/fsnotify"
	"github.com/getsentry/raven-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type config struct {
	Addr              string
	ServeFrom         string
	LogLevel          string
	LogDir            string
	ForceColors       bool
	SentryDSN         string
	SentryEnvironment string
}

func loadConfig(log logrus.FieldLogger) (cfg config) {
	viper.Set("Verbose", true)
	viper.Set("LogFile", os.Stderr)

	pflag.String("addr", ":8080", "Port to listen at")
	pflag.String("serveFrom", "", "Directory to serve from")

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

var (
	BuildTimeCommitMD5 string
	BuildTimeTime      string
	BuildTimeIsDev     string
)

func mainHandler(serveDir string, log logrus.FieldLogger) http.HandlerFunc {
	fs := http.FileServer(http.Dir(serveDir))
	return raven.RecoveryHandler(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
		log.Infof("Served for %s", r.UserAgent())
	})
}

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

	log.WithFields(initFields).Info("dirserver service greeting")

	cfg := loadConfig(log)
	log.Infof("Loaded config: %s", spew.Sdump(cfg))

	raven.SetDefaultLoggerName("dirserver")
	raven.SetEnvironment(cfg.SentryEnvironment)
	raven.SetRelease(BuildTimeCommitMD5)
	raven.SetDSN(cfg.SentryDSN)

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "OK") })
	mux.Handle("/", mainHandler(cfg.ServeFrom, log))

	srv := http.Server{
		Addr:    cfg.Addr,
		Handler: mux,
	}
	log.Info("Initialization completed")

	log.Infof("Begin listening on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		raven.CaptureErrorAndWait(err, nil)
		log.Error(err)
	}
}
