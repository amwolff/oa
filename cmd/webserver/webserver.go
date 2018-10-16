package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	srv := http.Server{
		Addr:    ":8081",
		Handler: http.FileServer(http.Dir("/Users/amw/Desktop/Go/src/github.com/amwolff/oa/pkg/frontend")),
	}
	logrus.Infof("Begin listening on port %s", srv.Addr)
	logrus.Error(srv.ListenAndServe())
}
