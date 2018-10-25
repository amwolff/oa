package common

import (
	"fmt"
	"time"

	"github.com/gocraft/dbr"
	"github.com/sirupsen/logrus"
)

func GetDsn(user, pass, host string, port int, name string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, pass, host, port, name)
}

func WaitForPostgres(dbC *dbr.Connection, n int, log logrus.FieldLogger) error {
	var err error
	for i := 0; i < n; i++ {
		if err = dbC.Ping(); err != nil {
			log.WithError(err).Debug("postgres still down")
			time.Sleep(time.Second)
		} else {
			log.Debug("postgres is up")
			return nil
		}
	}
	return err
}
