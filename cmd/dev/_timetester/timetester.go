package main

import (
	"context"
	"net/http"
	"time"

	"github.com/amwolff/oa/pkg/municommodels"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	sessionCookies := []http.Cookie{{Name: "ASP.NET_SessionId", Value: "zwtadx2ambmxdzgbkwlxwdfp"}}

	wsc := municommodels.NewWebServiceClient(
		logger,
		"Test",
		"PZI TARAN",
		"http://sip.zdzit.olsztyn.eu/PublicService.asmx",
	)

	ctx, canc := context.WithTimeout(context.Background(), time.Second)
	routesResp, err := wsc.CallGetRouteAndVariants(ctx, sessionCookies, municommodels.GetRouteAndVariants{})
	if err != nil {
		canc()
		logger.WithError(err).Fatal("CallGetRouteAndVariants")
	}
	canc()

	for _, r := range routesResp.GetRouteAndVariantsResult.L {
		logger.Infof("Change to request new vehicles @ route %s (%s)", r.Number, r.Description)

		vehiclesReq := municommodels.CNRGetVehicles{
			R: r.Number,
			D: r.Direction,
		}

		var previous *municommodels.CNRGetVehiclesResponse
		for i := 0; i < 120; i++ {
			ctx, canc := context.WithTimeout(context.Background(), time.Second)
			vehiclesResp, err := wsc.CallCNRGetVehicles(ctx, sessionCookies, vehiclesReq)
			if err != nil {
				canc()
				logger.WithError(err).Error("CallCNRGetVehicles")
				break
			}
			canc()

			if previous != nil {
				if previous.CNRGetVehiclesResult.Sanitized[0].Odchylenie == vehiclesResp.
					CNRGetVehiclesResult.Sanitized[0].Odchylenie && previous.CNRGetVehiclesResult.
					Sanitized[0].Szerokosc == vehiclesResp.CNRGetVehiclesResult.Sanitized[0].
					Szerokosc && previous.CNRGetVehiclesResult.Sanitized[0].Dlugosc == vehiclesResp.
					CNRGetVehiclesResult.Sanitized[0].Dlugosc {

					logger.Infof("Vehicle: %d Route: %s LatLng: %f %f Variance: %d",
						vehiclesResp.CNRGetVehiclesResult.Sanitized[0].Nb,
						vehiclesResp.CNRGetVehiclesResult.Sanitized[0].NumerLini,
						vehiclesResp.CNRGetVehiclesResult.Sanitized[0].Dlugosc,
						vehiclesResp.CNRGetVehiclesResult.Sanitized[0].Szerokosc,
						vehiclesResp.CNRGetVehiclesResult.Sanitized[0].Odchylenie,
					)
				} else {
					logger.Warnf("Vehicle: %d Route: %s LatLng: %f %f PrevLatLng: %f %f Variance: %d",
						vehiclesResp.CNRGetVehiclesResult.Sanitized[0].Nb,
						vehiclesResp.CNRGetVehiclesResult.Sanitized[0].NumerLini,
						vehiclesResp.CNRGetVehiclesResult.Sanitized[0].Dlugosc,
						vehiclesResp.CNRGetVehiclesResult.Sanitized[0].Szerokosc,
						vehiclesResp.CNRGetVehiclesResult.Sanitized[0].PrevDlugosc,
						vehiclesResp.CNRGetVehiclesResult.Sanitized[0].PrevSzerokosc,
						vehiclesResp.CNRGetVehiclesResult.Sanitized[0].Odchylenie,
					)
				}
			}

			previous = vehiclesResp
			time.Sleep(5 * time.Second)
		}
	}
}
