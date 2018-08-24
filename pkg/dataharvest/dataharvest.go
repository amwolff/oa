package dataharvest

import (
	"time"

	"github.com/amwolff/oa/pkg/municommodels"
	"github.com/gocraft/dbr"
	"github.com/pkg/errors"
)

func InsertCNRGetVehiclesResponsesIntoDb(dbSess dbr.SessionRunner,
	vehicles []municommodels.CNRGetVehiclesResponse, fetchTime time.Time) error {

	q := dbSess.InsertInto("olsztyn.live").Columns(
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
		"raw",
	)

	for _, v := range vehicles {
		for i, s := range v.CNRGetVehiclesResult.Sanitized {
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
				v.CNRGetVehiclesResult.Unsanitized[i],
			)
		}
	}

	if _, err := q.Exec(); err != nil {
		return errors.Wrap(err, "cannot insert vehicles dump into the database")
	}

	return nil
}
