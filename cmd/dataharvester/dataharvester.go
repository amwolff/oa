package main

import (
	"os"

	"github.com/fsnotify/fsnotify"
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

	pflag.String("loglevel", "debug", "Logging level")
	pflag.String("logDir", "/tmp", "Logs directory")
	pflag.Bool("forceColors", true, "Force colors when printing to stdout")

	configFile := pflag.String("config", "", "A config file to load")

	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	if len(*configFile) > 0 {
		viper.AddConfigPath(*configFile)
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
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

func main() {

}
