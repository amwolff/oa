package dataharvest

import (
	"time"

	"github.com/amwolff/oa/pkg/feeds/zdzit"
	"github.com/amwolff/oa/pkg/municommodels"
	"github.com/gocraft/dbr"
	"github.com/pkg/errors"
)

func InsertGetRouteAndVariantsResponseIntoDb(dbSess dbr.SessionRunner,
	routes *municommodels.GetRouteAndVariantsResponse, fetchTime time.Time) error {

	q := dbSess.InsertInto("olsztyn_static.routes").Columns(
		"ts",

		"number",
		"description",
		"description2",
		"variant",
		"transport",
		"direction",
	)

	for _, l := range routes.GetRouteAndVariantsResult.L {
		q.Values(
			fetchTime,

			l.Number,
			l.Description,
			l.Description2,
			l.Variant,
			l.Transport,
			l.Direction,
		)
	}

	if _, err := q.Exec(); err != nil {
		return errors.Wrap(err, "cannot insert routes dump into the database")
	}

	return nil
}

func InsertCNRGetVehiclesResponsesIntoDb(dbSess dbr.SessionRunner,
	vehicles []*municommodels.CNRGetVehiclesResponse, fetchTime time.Time) error {

	q := dbSess.InsertInto("olsztyn_live.vehicles").Columns(
		"ts",

		"nr_radia",
		"nb",
		"numer_lini",
		"war_trasy",
		"kierunek",
		"id_kursu",
		"lp_przyst",
		"droga_plan",
		"droga_wyko",
		"dlugosc",
		"szerokosc",
		"prev_dlugosc",
		"prev_szerokosc",
		"odchylenie",
		"odchylenie_str",
		"stan",
		"plan_godz_rozp",
		"nast_id_kursu",
		"nast_plan_godz_rozp",
		"nast_num_lini",
		"nast_war_trasy",
		"nast_kierunek",
		"ile_sek_do_odjazdu",
		"typ_pojazdu",
		"transport",
		"cechy",
		"opis_tabl",
		"nast_opis_tabl",
		"wektor",
	)

	for _, v := range vehicles {
		for _, s := range v.CNRGetVehiclesResult.Sanitized {
			q.Values(
				fetchTime,

				s.NrRadia,
				s.Nb,
				s.NumerLini,
				s.WarTrasy,
				s.Kierunek,
				s.IdKursu,
				s.LpPrzyst,
				s.DrogaPlan,
				s.DrogaWyko,
				s.Dlugosc,
				s.Szerokosc,
				s.PrevDlugosc,
				s.PrevSzerokosc,
				s.Odchylenie,
				s.OdchylenieStr,
				s.Stan,
				s.PlanGodzRozp,
				s.NastIdKursu,
				s.NastPlanGodzRozp,
				s.NastNumLini,
				s.NastWarTrasy,
				s.NastKierunek,
				s.IleSekDoOdjazdu,
				s.TypPojazdu,
				s.Transport,
				s.Cechy,
				s.OpisTabl,
				s.NastOpisTabl,
				s.Wektor,
			)
		}
	}

	if _, err := q.Exec(); err != nil {
		return errors.Wrap(err, "cannot insert vehicles dump into the database")
	}

	return nil
}

func InsertBusStopsIntoDb(dbSess dbr.SessionRunner, stops []zdzit.BusStop, fetchTime time.Time) error {

	q := dbSess.InsertInto("olsztyn_static.stops").Columns(
		"ts",

		"number",
		"name",
		"street_name",
		"latitude",
		"longitude",
	)

	for _, s := range stops {
		q.Values(
			fetchTime,

			s.Number,
			s.Name,
			s.StreetName,
			s.LatLng.Latitude,
			s.LatLng.Longitude,
		)
	}

	if _, err := q.Exec(); err != nil {
		return errors.Wrap(err, "cannot insert bus stops dump into the database")
	}

	return nil
}
