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
	"github.com/gocraft/dbr"
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

	Addr string

	LogLevel    string
	LogDir      string
	ForceColors bool

	SentrySecret string
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

type vehicleResponse struct {
	Timestamp       time.Time `json:"-"                db:"ts"`
	VehicleID       int       `json:"vehicle_id"       db:"nb"`
	VehicleType     string    `json:"vehicle_type"     db:"typ_pojazdu"`
	Route           string    `json:"route"            db:"numer_lini"`
	TripID          int       `json:"trip_id"          db:"id_kursu"`
	Latitude        float64   `json:"latitude"         db:"szerokosc"`
	Longitude       float64   `json:"longitude"        db:"dlugosc"`
	LastLatitude    float64   `json:"last_latitude"    db:"prev_szerokosc"`
	LastLongitude   float64   `json:"last_longitude"   db:"prev_dlugosc"`
	Variance        int       `json:"variance"         db:"odchylenie"`
	Description     string    `json:"description"      db:"opis_tabl"`
	NextRoute       string    `json:"next_route"       db:"nast_num_lini"`
	NextTripID      int       `json:"next_trip_id"     db:"nast_id_kursu"`
	DepartureIn     int       `json:"departure_in"     db:"ile_sek_do_odjazdu"`
	NextDescription string    `json:"next_description" db:"nast_opis_tabl"`
	Vector          float64   `json:"vector"           db:"wektor"`
}

func endpointVehiclesData(dbS dbr.SessionRunner, log logrus.FieldLogger) http.HandlerFunc {
	// The database call might be WAY more optimized, but I'm leaving this the
	// way it is since we don't anticipate 10k calls /s kind of traffic.
	q := dbS.
		Select(
			"ts",
			"nb",
			"typ_pojazdu",
			"numer_lini",
			"id_kursu",
			"szerokosc",
			"dlugosc",
			"prev_szerokosc",
			"prev_dlugosc",
			"odchylenie",
			"opis_tabl",
			"nast_num_lini",
			"nast_id_kursu",
			"ile_sek_do_odjazdu",
			"nast_opis_tabl",
			"wektor",
		).
		From("olsztyn_live.vehicles").
		Where("ts = (SELECT ts FROM olsztyn_live.vehicles ORDER BY id DESC LIMIT 1)").
		OrderBy("id")

	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Last-Modified#Syntax
	const lastModifiedTimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT"

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:63342")
		var resp []vehicleResponse
		if err := q.LoadOne(&resp); err != nil {
			http.Error(w, "", http.StatusServiceUnavailable)
			log.WithError(err).WithField("func", "LoadOne").Error("Cannot execute query")
			return
		}
		if len(resp) > 0 {
			w.Header().Set("Last-Modified", resp[0].Timestamp.UTC().Format(lastModifiedTimeFormat))
		}
		if err := json.NewEncoder(w).Encode(&resp); err != nil {
			http.Error(w, "", http.StatusServiceUnavailable)
			log.WithError(err).WithField("func", "Encode").Error("Cannot encode fetched")
			return
		}
	}
}

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

	dsn := common.GetDsn(cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName)
	log.Debugf("DSN: %s", dsn)

	dbConn, err := dbr.Open("postgres", dsn, &dbr.NullEventReceiver{})
	if err != nil {
		log.WithError(err).Fatal("Cannot connect to database")
	}
	dbSess := dbConn.NewSession(&dbr.NullEventReceiver{})

	// tz, err := time.LoadLocation("Europe/Warsaw")
	// if err != nil {
	// 	log.WithError(err).Fatal("Cannot parse location")
	// }

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "OK") })
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) { return })
	mux.Handle("/", gziphandler.GzipHandler(endpointVehiclesData(dbSess, log)))

	srv := http.Server{
		Addr:    cfg.Addr,
		Handler: mux,
	}
	log.Info("Initialization completed")

	log.Infof("Begin listening on port %s", srv.Addr)
	log.Error(srv.ListenAndServe())
}
