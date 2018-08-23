package main

import (
	"context"
	"net/http"
	"time"

	"github.com/amwolff/oa/pkg/municommodels"
	"github.com/sirupsen/logrus"
)

var reqTimeout = 10 * time.Second

func main() {
	logger := logrus.New()
	sessionCookies := []http.Cookie{{Name: "ASP.NET_SessionId", Value: "kx0qshioeixcqio2le5dbvw4"}}

	wsc := municommodels.NewWebServiceClient(
		logger,
		"Test",
		"PZI TARAN",
		"http://sip.zdzit.olsztyn.eu/PublicService.asmx",
	)

	routesReq := municommodels.GetRouteAndVariants{}
	ctx, canc := context.WithTimeout(context.Background(), (5 * time.Second))
	routesResp, err := wsc.CallGetRouteAndVariants(ctx, sessionCookies, routesReq)
	if err != nil {
		canc()
		logger.WithError(err).Fatal("CallGetRouteAndVariants")
	}
	canc()

	var vehiclesResps []*municommodels.CNRGetVehiclesResponse
	for _, l := range routesResp.GetRouteAndVariantsResult.L {
		vehiclesReq := municommodels.CNRGetVehicles{
			R: l.Number,
			D: l.Direction,
		}
		logger.Infof("Requesting for Number: %s Direction: %s", l.Number, l.Direction)
		ctx, canc := context.WithTimeout(context.Background(), (5 * time.Second))
		vehiclesResp, err := wsc.CallCNRGetVehicles(ctx, sessionCookies, vehiclesReq)
		if err != nil {
			canc()
			logger.WithError(err).Fatal("CallCNRGetVehicles")
		}
		vehiclesResps = append(vehiclesResps, vehiclesResp)
	}

	for _, v := range vehiclesResps {
		for _, s := range v.CNRGetVehiclesResult.Sanitized {
			logger.Printf("Numer: %d Kurs: %d Linia: %s Cel: %s", s.Nb, s.IdKursu, s.NumerLini, s.OpisTabl)
		}
	}
}
