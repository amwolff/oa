package municommodels

import (
	"encoding/json"
	"encoding/xml"
	"html"
	"reflect"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// CNRGetVehicles represents SOAP request body of an CNRGetVehicles endpoint.
type CNRGetVehicles struct {
	// reference:
	// PublicService.CNR_GetVehicles(String s, String r, String v, String d, String nb, String tp)
	//
	// I got the signature of this function from a panic stacktrace.

	XMLName xml.Name `xml:"http://PublicService/ CNR_GetVehicles"`

	// Fields names are reverse engineered (again) from the biggest (ugliest?)
	// JS on the site.
	//
	// It goes like this:
	// s:  some kind of token
	// r:  "linie"
	// v:  "war_trasy" ~ "warunki trasy"
	// d:  "kierunki"
	// nb: "nb" ~ numer boczny"
	// tp: "typy"

	S  string `xml:"s"`
	R  string `xml:"r"`
	V  string `xml:"v"`
	D  string `xml:"d"`
	Nb string `xml:"nb"`
	Tp string `xml:"tp"`
}

type CNRGetVehiclesResponse struct {
	XMLName xml.Name `xml:"http://PublicService/ CNR_GetVehiclesResponse"`

	CNRGetVehiclesResult struct {
		Unsanitized []string                        `xml:"p,omitempty"`
		Sanitized   []SanitizedCNRGetVehiclesResult `xml:"-"`
	} `xml:"CNR_GetVehiclesResult>VL,omitempty"`
}

type SanitizedCNRGetVehiclesResult struct {
	// Fields names are reverse engineered (again) from the biggest (ugliest?)
	// JS on the site.
	NrRadia          int
	Nb               int
	NumerLini        string
	WarTrasy         string
	Kierunek         string
	IdKursu          int
	LpPrzyst         int
	DrogaPlan        int
	DrogaWyko        int
	Dlugosc          float64
	Szerokosc        float64
	PrevDlugosc      float64
	PrevSzerokosc    float64
	Odchylenie       int
	OdchylenieStr    string
	Stan             int
	PlanGodzRozp     string
	NastIdKursu      int
	NastPlanGodzRozp string
	NastNumLini      string
	NastWarTrasy     string
	NastKierunek     string
	IleSekDoOdjazdu  int
	TypPojazdu       string
	Transport        string
	Cechy            string
	OpisTabl         string
	NastOpisTabl     string
	Wektor           float64
	// DataRj: "",
	// WersjaRj: "",
	// Zajezdnia: 0,
	// TypDnia: "",
	// NrKursowk: "",
	// KodLinii: 0,
	// NrEw: "",
	// Nim: "",
	// Telefon: "",
	// IdPrzyst1: 0,
	// IdPrzyst2: 0,
	// Awaria: 0,
	// DataCzasOTR: "",
	// DataCzasOPP: "",
	// SkrCzynn: "",
	// CzynnFCol: "",
	// CzynnBCol: "",
	// IdWersja: 0,
	// Uwagi: "",
	// NazwaSrg: "",
	// NazwaPrz1: "",
	// NazwaPrz2: "",
	// NastPlanGodzRozp: "",
	// Obciazenie: 0,
	// DataCzasStatus: "",
	// JestStatus: 0
}

// Deprecated: _Sanitize is deprecated. Leaving for testing purposes. Please use
// Sanitize.
func (r *CNRGetVehiclesResponse) _Sanitize() error {
	// Long story short: result from this method is insanely ugly and impossible
	// to use without proper sanitation. This function is going to extract each
	// paragraph and extract its data into the additional structure in the
	// response struct into fields, which names were extracted from the reverse
	// engineered script.
	for _, u := range r.CNRGetVehiclesResult.Unsanitized {
		// Step one: drop brackets
		droppedBrackets := strings.TrimPrefix(strings.TrimSuffix(u, "]"), "[")

		// Step two: split fields
		splitted := strings.Split(droppedBrackets, ",")

		// Step three: remove rest of the trailing characters
		var splittedSanitized []string
		for _, s := range splitted {
			sanitized := strings.TrimSpace(
				strings.TrimSuffix(
					strings.TrimSpace(
						strings.TrimPrefix(
							strings.TrimSpace(s), `"`)), `"`))

			splittedSanitized = append(splittedSanitized, sanitized)
		}

		// Step four: extract, convert, append
		i0, err := strconv.Atoi(splittedSanitized[0])
		if err != nil {
			return err
		}
		i1, err := strconv.Atoi(splittedSanitized[1])
		if err != nil {
			return err
		}
		i5, err := strconv.Atoi(splittedSanitized[5])
		if err != nil {
			return err
		}
		i6, err := strconv.Atoi(splittedSanitized[6])
		if err != nil {
			return err
		}
		i7, err := strconv.Atoi(splittedSanitized[7])
		if err != nil {
			return err
		}
		i8, err := strconv.Atoi(splittedSanitized[8])
		if err != nil {
			return err
		}
		f9, err := strconv.ParseFloat(splittedSanitized[9], 64)
		if err != nil {
			return err
		}
		f10, err := strconv.ParseFloat(splittedSanitized[10], 64)
		if err != nil {
			return err
		}
		f11, err := strconv.ParseFloat(splittedSanitized[11], 64)
		if err != nil {
			return err
		}
		f12, err := strconv.ParseFloat(splittedSanitized[12], 64)
		if err != nil {
			return err
		}
		i13, err := strconv.Atoi(splittedSanitized[13])
		if err != nil {
			return err
		}
		i15, err := strconv.Atoi(splittedSanitized[15])
		if err != nil {
			return err
		}
		i17, err := strconv.Atoi(splittedSanitized[17])
		if err != nil {
			return err
		}
		i22, err := strconv.Atoi(splittedSanitized[22])
		if err != nil {
			return err
		}

		var s25 string
		if len(splittedSanitized) > 25 {
			s25 = html.UnescapeString(splittedSanitized[25])
		}

		var s26 string
		if len(splittedSanitized) > 26 {
			s26 = html.UnescapeString(splittedSanitized[26])
		}

		var f27 float64 = -1
		if len(splittedSanitized) > 27 {
			f27, err = strconv.ParseFloat(splittedSanitized[27], 64)
			if err != nil {
				return err
			}
		}

		a := SanitizedCNRGetVehiclesResult{
			NrRadia:          i0,
			Nb:               i1,
			NumerLini:        html.UnescapeString(splittedSanitized[2]),
			WarTrasy:         html.UnescapeString(splittedSanitized[3]),
			Kierunek:         html.UnescapeString(splittedSanitized[4]),
			IdKursu:          i5,
			LpPrzyst:         i6,
			DrogaPlan:        i7,
			DrogaWyko:        i8,
			Dlugosc:          f9,
			Szerokosc:        f10,
			PrevDlugosc:      f11,
			PrevSzerokosc:    f12,
			Odchylenie:       i13,
			OdchylenieStr:    html.UnescapeString(splittedSanitized[14]),
			Stan:             i15,
			PlanGodzRozp:     html.UnescapeString(splittedSanitized[16]),
			NastIdKursu:      i17,
			NastPlanGodzRozp: html.UnescapeString(splittedSanitized[18]),
			NastNumLini:      html.UnescapeString(splittedSanitized[19]),
			NastWarTrasy:     html.UnescapeString(splittedSanitized[20]),
			NastKierunek:     html.UnescapeString(splittedSanitized[21]),
			IleSekDoOdjazdu:  i22,
			TypPojazdu:       html.UnescapeString(splittedSanitized[23]),
			Transport:        html.UnescapeString(splittedSanitized[23]),
			Cechy:            html.UnescapeString(splittedSanitized[24]),
			OpisTabl:         s25,
			NastOpisTabl:     s26,
			Wektor:           f27,
		}
		r.CNRGetVehiclesResult.Sanitized = append(r.CNRGetVehiclesResult.Sanitized, a)
	}

	return nil
}

func sanitize(u string) (SanitizedCNRGetVehiclesResult, error) {
	var (
		s      SanitizedCNRGetVehiclesResult
		fields []interface{}
	)

	if err := json.Unmarshal([]byte(u), &fields); err != nil {
		return s, errors.Wrap(err, "cannot unmarshal as JSON")
	}

	f0, ok := fields[0].(float64)
	if !ok {
		return s, errors.Errorf("cannot assert %v (type: %v) as type float64", fields[0],
			reflect.TypeOf(f0))
	}

	f1, ok := fields[1].(float64)
	if !ok {
		return s, errors.Errorf("cannot assert %v (type: %v) as type float64", fields[1],
			reflect.TypeOf(f1))
	}

	f5, ok := fields[5].(float64)
	if !ok {
		return s, errors.Errorf("cannot assert %v (type: %v) as type float64", fields[5],
			reflect.TypeOf(f5))
	}

	f6, ok := fields[6].(float64)
	if !ok {
		return s, errors.Errorf("cannot assert %v (type: %v) as type float64", fields[6],
			reflect.TypeOf(f6))
	}

	f7, ok := fields[7].(float64)
	if !ok {
		return s, errors.Errorf("cannot assert %v (type: %v) as type float64", fields[7],
			reflect.TypeOf(f7))
	}

	f8, ok := fields[8].(float64)
	if !ok {
		return s, errors.Errorf("cannot assert %v (type: %v) as type float64", fields[8],
			reflect.TypeOf(f8))
	}

	f9, ok := fields[9].(float64)
	if !ok {
		return s, errors.Errorf("cannot assert %v (type: %v) as type float64", fields[9],
			reflect.TypeOf(f9))
	}

	f10, ok := fields[10].(float64)
	if !ok {
		return s, errors.Errorf("cannot assert %v (type: %v) as type float64", fields[10],
			reflect.TypeOf(f10))
	}

	f11, ok := fields[11].(float64)
	if !ok {
		return s, errors.Errorf("cannot assert %v (type: %v) as type float64", fields[11],
			reflect.TypeOf(f11))
	}

	f12, ok := fields[12].(float64)
	if !ok {
		return s, errors.Errorf("cannot assert %v (type: %v) as type float64", fields[12],
			reflect.TypeOf(f12))
	}

	f13, ok := fields[13].(float64)
	if !ok {
		return s, errors.Errorf("cannot assert %v (type: %v) as type float64", fields[13],
			reflect.TypeOf(f13))
	}

	f15, ok := fields[15].(float64)
	if !ok {
		return s, errors.Errorf("cannot assert %v (type: %v) as type float64", fields[15],
			reflect.TypeOf(f15))
	}

	f17, ok := fields[17].(float64)
	if !ok {
		return s, errors.Errorf("cannot assert %v (type: %v) as type float64", fields[17],
			reflect.TypeOf(f17))
	}

	f22, ok := fields[22].(float64)
	if !ok {
		return s, errors.Errorf("cannot assert %v (type: %v) as type float64", fields[22],
			reflect.TypeOf(f22))
	}

	var f27 float64 = -1
	if len(fields) > 27 {
		f27, ok = fields[27].(float64)
		if !ok {
			return s, errors.Errorf("cannot assert %v (type: %v) as type float64", fields[27],
				reflect.TypeOf(f27))
		}
	}

	s2, ok := fields[2].(string)
	if !ok {
		return s, errors.Errorf("cannot assert %v (type: %v) as type string", fields[2],
			reflect.TypeOf(s2))
	}

	s3, ok := fields[3].(string)
	if !ok {
		return s, errors.Errorf("cannot assert %v (type: %v) as type string", fields[3],
			reflect.TypeOf(s3))
	}

	s4, ok := fields[4].(string)
	if !ok {
		return s, errors.Errorf("cannot assert %v (type: %v) as type string", fields[4],
			reflect.TypeOf(s4))
	}

	s14, ok := fields[14].(string)
	if !ok {
		return s, errors.Errorf("cannot assert %v (type: %v) as type string", fields[14],
			reflect.TypeOf(s14))
	}

	s16, ok := fields[16].(string)
	if !ok {
		return s, errors.Errorf("cannot assert %v (type: %v) as type string", fields[16],
			reflect.TypeOf(s16))
	}

	s18, ok := fields[18].(string)
	if !ok {
		return s, errors.Errorf("cannot assert %v (type: %v) as type string", fields[18],
			reflect.TypeOf(s18))
	}

	s19, ok := fields[19].(string)
	if !ok {
		return s, errors.Errorf("cannot assert %v (type: %v) as type string", fields[19],
			reflect.TypeOf(s19))
	}

	s20, ok := fields[20].(string)
	if !ok {
		return s, errors.Errorf("cannot assert %v (type: %v) as type string", fields[20],
			reflect.TypeOf(s20))
	}

	s21, ok := fields[21].(string)
	if !ok {
		return s, errors.Errorf("cannot assert %v (type: %v) as type string", fields[21],
			reflect.TypeOf(s21))
	}

	s23, ok := fields[23].(string)
	if !ok {
		return s, errors.Errorf("cannot assert %v (type: %v) as type string", fields[23],
			reflect.TypeOf(s23))
	}
	s23 = html.UnescapeString(strings.TrimSpace(s23))

	s24, ok := fields[24].(string)
	if !ok {
		return s, errors.Errorf("cannot assert %v (type: %v) as type string", fields[24],
			reflect.TypeOf(s24))
	}

	var s25 string
	if len(fields) > 25 {
		s25, ok = fields[25].(string)
		if !ok {
			return s, errors.Errorf("cannot assert %v (type: %v) as type string", fields[25],
				reflect.TypeOf(s25))
		}
		s25 = html.UnescapeString(strings.TrimSpace(s25))
	}

	var s26 string
	if len(fields) > 26 {
		s26, ok = fields[26].(string)
		if !ok {
			return s, errors.Errorf("cannot assert %v (type: %v) as type string", fields[26],
				reflect.TypeOf(s26))
		}
		s26 = html.UnescapeString(strings.TrimSpace(s26))
	}

	// These casts are generally safe
	s = SanitizedCNRGetVehiclesResult{
		NrRadia:          int(f0),
		Nb:               int(f1),
		NumerLini:        html.UnescapeString(strings.TrimSpace(s2)),
		WarTrasy:         html.UnescapeString(strings.TrimSpace(s3)),
		Kierunek:         html.UnescapeString(strings.TrimSpace(s4)),
		IdKursu:          int(f5),
		LpPrzyst:         int(f6),
		DrogaPlan:        int(f7),
		DrogaWyko:        int(f8),
		Dlugosc:          f9,
		Szerokosc:        f10,
		PrevDlugosc:      f11,
		PrevSzerokosc:    f12,
		Odchylenie:       int(f13),
		OdchylenieStr:    html.UnescapeString(strings.TrimSpace(s14)),
		Stan:             int(f15),
		PlanGodzRozp:     html.UnescapeString(strings.TrimSpace(s16)),
		NastIdKursu:      int(f17),
		NastPlanGodzRozp: html.UnescapeString(strings.TrimSpace(s18)),
		NastNumLini:      html.UnescapeString(strings.TrimSpace(s19)),
		NastWarTrasy:     html.UnescapeString(strings.TrimSpace(s20)),
		NastKierunek:     html.UnescapeString(strings.TrimSpace(s21)),
		IleSekDoOdjazdu:  int(f22),
		TypPojazdu:       s23,
		Transport:        s23,
		Cechy:            html.UnescapeString(strings.TrimSpace(s24)),
		OpisTabl:         s25,
		NastOpisTabl:     s26,
		Wektor:           f27,
	}

	return s, nil
}

func (r *CNRGetVehiclesResponse) Sanitize() error {
	// Long story short: result from this method is insanely ugly and impossible
	// to use without proper sanitation. This function is going to extract each
	// paragraph and extract its data into the additional structure in the
	// response struct into fields, which names were extracted from the reverse
	// engineered script.
	for _, u := range r.CNRGetVehiclesResult.Unsanitized {
		s, err := sanitize(u)
		if err != nil {
			return errors.Wrap(err, "sanitize")
		}
		r.CNRGetVehiclesResult.Sanitized = append(r.CNRGetVehiclesResult.Sanitized, s)
	}

	return nil
}
