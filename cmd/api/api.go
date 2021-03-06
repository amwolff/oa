package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/NYTimes/gziphandler"
	"github.com/amwolff/oa/pkg/common"
	"github.com/davecgh/go-spew/spew"
	"github.com/fsnotify/fsnotify"
	"github.com/getsentry/raven-go"
	"github.com/gocraft/dbr"
	"github.com/gocraft/dbr/dialect"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type config struct {
	DbHost string
	DbPort int
	DbName string
	DbUser string
	DbPass string

	Addr       string
	HealthPath string

	LogLevel    string
	LogDir      string
	ForceColors bool

	SentryDSN         string
	SentryEnvironment string
}

func loadConfig(log logrus.FieldLogger) (cfg config) {
	viper.Set("Verbose", true)
	viper.Set("LogFile", os.Stderr)

	pflag.String("dbHost", "localhost", "Database host")
	pflag.Int("dbPort", 5432, "Database port")
	pflag.String("dbName", "oadb", "Database name")
	pflag.String("dbUser", "data_service", "Database user")
	pflag.String("dbPass", "data_service", "Database password")

	pflag.String("addr", ":8080", "Port to listen at")
	pflag.String("healthPath", "/healthz", "")

	pflag.String("logLevel", "debug", "Logging level")
	pflag.String("logDir", "/tmp", "Logs directory")
	pflag.Bool("forceColors", true, "Force colors when printing to stdout")

	pflag.String("sentrySecret", "", "Secret string for Raven")
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

func build(stmt *dbr.SelectStmt) (string, error) {
	d := dialect.PostgreSQL

	buf := dbr.NewBuffer()
	if err := stmt.Build(d, buf); err != nil {
		return "", err
	}
	q, err := dbr.InterpolateForDialect(buf.String(), buf.Value(), d)
	if err != nil {
		return "", err
	}

	return q, nil
}

func buildStaticQueries() (map[string]string, error) {
	m := make(map[string]string)
	{
		sel := dbr.
			Select(
				"ts",
				"number",
			).
			Distinct().
			From("olsztyn_static.routes").
			Where("ts = (SELECT ts FROM olsztyn_static.routes ORDER BY ts DESC LIMIT 1)").
			OrderBy("number")

		q, err := build(sel)
		if err != nil {
			return nil, err
		}
		m["Routes"] = q
	}
	{
		sel := dbr.
			Select(
				"ts",
				"numer_lini",
				"id_kursu",
				"szerokosc",
				"dlugosc",
				"odchylenie",
				"opis_tabl",
				"nast_num_lini",
				"nast_id_kursu",
				"nast_opis_tabl",
				"wektor",
			).
			From("olsztyn_live.vehicles").
			Where("ts = (SELECT ts FROM olsztyn_live.vehicles ORDER BY ts DESC LIMIT 1)").
			OrderBy("numer_lini")

		q, err := build(sel)
		if err != nil {
			return nil, err
		}
		m["Vehicles"] = q
	}
	return m, nil
}

// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Last-Modified#Syntax
const lastModifiedTimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT"

type routesResponse struct {
	Timestamp time.Time `json:"-"     db:"ts"`
	Route     string    `json:"route" db:"number"`
}

func endpointRoutes(dbC *dbr.Connection, q, org string, log logrus.FieldLogger) http.HandlerFunc {
	log = log.WithField("handler", "AvailableRoutes")

	return raven.RecoveryHandler(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", org)
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		var resp []routesResponse
		if err := dbC.NewSession(nil).SelectBySql(q).LoadOne(&resp); err != nil {
			http.Error(w, "", http.StatusServiceUnavailable)
			raven.CaptureErrorAndWait(err, map[string]string{"func": "LoadOne"})
			log.WithError(err).WithField("func", "LoadOne").Error("Cannot execute query")
			return
		}
		w.Header().Set("Last-Modified", resp[0].Timestamp.UTC().Format(lastModifiedTimeFormat))
		if err := json.NewEncoder(w).Encode(&resp); err != nil {
			http.Error(w, "", http.StatusServiceUnavailable)
			raven.CaptureErrorAndWait(err, map[string]string{"func": "Encode"})
			log.WithError(err).WithField("func", "Encode").Error("Cannot encode fetched")
			return
		}
	})
}

type vehiclesResponse struct {
	Timestamp time.Time `json:"-" db:"ts"`

	Azimuth         float64 `json:"azimuth"          db:"wektor"`
	Description     string  `json:"description"      db:"opis_tabl"`
	Latitude        float64 `json:"latitude"         db:"szerokosc"`
	Longitude       float64 `json:"longitude"        db:"dlugosc"`
	NextDescription string  `json:"next_description" db:"nast_opis_tabl"`
	NextRoute       string  `json:"next_route"       db:"nast_num_lini"`
	NextTripID      int     `json:"next_trip_id"     db:"nast_id_kursu"`
	Route           string  `json:"route"            db:"numer_lini"`
	TripID          int     `json:"trip_id"          db:"id_kursu"`
	Variance        int     `json:"variance"         db:"odchylenie"`
}

func endpointVehicles(dbC *dbr.Connection, q, org string, log logrus.FieldLogger) http.HandlerFunc {
	log = log.WithField("handler", "VehiclesData")

	return raven.RecoveryHandler(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", org)
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		var resp []vehiclesResponse
		if err := dbC.NewSession(nil).SelectBySql(q).LoadOne(&resp); err != nil {
			http.Error(w, "", http.StatusServiceUnavailable)
			raven.CaptureErrorAndWait(err, map[string]string{"func": "LoadOne"})
			log.WithError(err).WithField("func", "LoadOne").Error("Cannot execute query")
			return
		}
		w.Header().Set("Last-Modified", resp[0].Timestamp.UTC().Format(lastModifiedTimeFormat))
		if err := json.NewEncoder(w).Encode(&resp); err != nil {
			http.Error(w, "", http.StatusServiceUnavailable)
			raven.CaptureErrorAndWait(err, map[string]string{"func": "Encode"})
			log.WithError(err).WithField("func", "Encode").Error("Cannot encode fetched")
			return
		}
	})
}

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

	log.WithFields(initFields).Info("data API service greeting")

	cfg := loadConfig(log)
	log.Infof("Loaded config: %s", spew.Sdump(cfg))

	// tz, err := time.LoadLocation("Europe/Warsaw")
	// if err != nil {
	// 	log.WithError(err).Fatal("Cannot parse location")
	// }

	queries, err := buildStaticQueries()
	if err != nil {
		log.WithError(err).Panic("Cannot build static queries")
	}

	raven.SetDefaultLoggerName("API")
	raven.SetEnvironment(cfg.SentryEnvironment)
	raven.SetRelease(BuildTimeCommitMD5)
	if err := raven.SetDSN(cfg.SentryDSN); err != nil {
		log.WithError(err).Fatal("raven.SetDSN")
	}

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

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { http.Error(w, "Invalid Request", http.StatusBadRequest) })
	mux.HandleFunc(cfg.HealthPath, func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "OK") })
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) { return })

	org := "https://autobusy.olsztyn.pl"
	if isDev {
		org = "*"
	}
	mux.Handle("/Routes", gziphandler.GzipHandler(endpointRoutes(dbConn, queries["Routes"], org, log)))
	mux.Handle("/Vehicles", gziphandler.GzipHandler(endpointVehicles(dbConn, queries["Vehicles"], org, log)))

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
