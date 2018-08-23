package municommodels

import (
	"encoding/xml"
	"reflect"
	"testing"
)

func TestSanitize(t *testing.T) {
	d := []byte(`<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><CNR_GetVehiclesResponse xmlns="http://PublicService/"><CNR_GetVehiclesResult><VL xmlns=""><p>[2008, 2008, "111  ", "0", "T", 6430, 3, 320, 176, 20.48651, 53.75529, 20.48642, 53.75467, 53, "00:00:53", 1, "14:36", 6431,"15:36","111  ","2","P",0, "A", "NK        ","GUTKOWO","NAG&amp;#211;RKI", 206.47]</p></VL></CNR_GetVehiclesResult></CNR_GetVehiclesResponse></soap:Body></soap:Envelope>`)

	var e Envelope
	if err := xml.Unmarshal(d, &e); err != nil {
		t.Fatalf("xml.Unmarshal returned an error: %v", err)
	}

	var r CNRGetVehiclesResponse
	if err := xml.Unmarshal(e.Body.InnerXML, &r); err != nil {
		t.Fatalf("xml.Unmarshal returned an error: %v", err)
	}

	if err := r.Sanitize(); err != nil {
		t.Fatalf("Sanitize returned an error: %v", err)
	}
	r0 := r.CNRGetVehiclesResult.Sanitized[0]

	if err := r.DeprecatedSanitize(); err != nil {
		t.Fatalf("DeprecatedSanitize returned an error: %v", err)
	}
	r1 := r.CNRGetVehiclesResult.Sanitized[1]

	if !reflect.DeepEqual(r0, r1) {
		t.Fatal("Sanitize and DeprecatedSanitize produce different output")
	} else {
		t.Log("OK")
	}
}
