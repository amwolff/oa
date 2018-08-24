package municommodelsreference

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type CNRGetVehicles struct {
	XMLName xml.Name `xml:"http://PublicService/ CNR_GetVehicles"`

	S  string `xml:"s,omitempty"`
	R  string `xml:"r,omitempty"`
	V  string `xml:"v,omitempty"`
	D  string `xml:"d,omitempty"`
	Nb string `xml:"nb,omitempty"`
	Tp string `xml:"tp,omitempty"`
}

type CNRGetVehiclesResponse struct {
	XMLName xml.Name `xml:"http://PublicService/ CNR_GetVehiclesResponse"`

	CNRGetVehiclesResult struct {
	} `xml:"CNR_GetVehiclesResult,omitempty"`
}

type CNRGetVehicleNumbers struct {
	XMLName xml.Name `xml:"http://PublicService/ CNR_GetVehicleNumbers"`

	S           string `xml:"s,omitempty"`
	Nrlini      string `xml:"nr_lini,omitempty"`
	Przewoznicy string `xml:"przewoznicy,omitempty"`
	Nrzaj       string `xml:"nr_zaj,omitempty"`
	Kursow      string `xml:"kursow,omitempty"`
	Czynnosci   string `xml:"czynnosci,omitempty"`
	Idprz       string `xml:"id_prz,omitempty"`
	Stan        string `xml:"stan,omitempty"`
	Odch        string `xml:"odch,omitempty"`
	Nb          string `xml:"nb,omitempty"`
	Typy        string `xml:"typy,omitempty"`
	Wartrasy    string `xml:"war_trasy,omitempty"`
	Idkursu     string `xml:"id_kursu,omitempty"`
	Typyawarii  string `xml:"typy_awarii,omitempty"`
}

type CNRGetVehicleNumbersResponse struct {
	XMLName xml.Name `xml:"http://PublicService/ CNR_GetVehicleNumbersResponse"`

	CNRGetVehicleNumbersResult struct {
	} `xml:"CNR_GetVehicleNumbersResult,omitempty"`
}

type CNRGetVehicleTypes struct {
	XMLName xml.Name `xml:"http://PublicService/ CNR_GetVehicleTypes"`

	S string `xml:"s,omitempty"`
}

type CNRGetVehicleTypesResponse struct {
	XMLName xml.Name `xml:"http://PublicService/ CNR_GetVehicleTypesResponse"`

	CNRGetVehicleTypesResult struct {
	} `xml:"CNR_GetVehicleTypesResult,omitempty"`
}

type CNRRouteVariants struct {
	XMLName xml.Name `xml:"http://PublicService/ CNR_RouteVariants"`

	S        string  `xml:"s,omitempty"`
	Actual   string  `xml:"actual,omitempty"`
	Vt       string  `xml:"vt,omitempty"`
	N        string  `xml:"n,omitempty"`
	Idwersja float64 `xml:"id_wersja,omitempty"`
}

type CNRRouteVariantsResponse struct {
	XMLName xml.Name `xml:"http://PublicService/ CNR_RouteVariantsResponse"`

	CNRRouteVariantsResult struct {
	} `xml:"CNR_RouteVariantsResult,omitempty"`
}

type CNRGetVehicleDepartures struct {
	XMLName xml.Name `xml:"http://PublicService/ CNR_GetVehicleDepartures"`

	S        string  `xml:"s,omitempty"`
	Idwersja float64 `xml:"id_wersja,omitempty"`
	Nbwozu   string  `xml:"nb_wozu,omitempty"`
	Idkursu  string  `xml:"id_kursu,omitempty"`
	Nrlini   string  `xml:"nr_lini,omitempty"`
	Wartrasy string  `xml:"war_trasy,omitempty"`
	Idprzyst string  `xml:"id_przyst,omitempty"`
}

type CNRGetVehicleDeparturesResponse struct {
	XMLName xml.Name `xml:"http://PublicService/ CNR_GetVehicleDeparturesResponse"`

	CNRGetVehicleDeparturesResult struct {
	} `xml:"CNR_GetVehicleDeparturesResult,omitempty"`
}

type GetTimeline struct {
	XMLName xml.Name `xml:"http://PublicService/ GetTimeline"`

	S        string  `xml:"s,omitempty"`
	Idwersja float64 `xml:"id_wersja,omitempty"`
	Krs      string  `xml:"krs,omitempty"`
	Idp      string  `xml:"idp,omitempty"`
}

type GetTimelineResponse struct {
	XMLName xml.Name `xml:"http://PublicService/ GetTimelineResponse"`

	GetTimelineResult struct {
	} `xml:"GetTimelineResult,omitempty"`
}

type CNRGetRealDepartures struct {
	XMLName xml.Name `xml:"http://PublicService/ CNR_GetRealDepartures"`

	Idwersja float64 `xml:"id_wersja,omitempty"`
	S        string  `xml:"s,omitempty"`
	Id       string  `xml:"id,omitempty"`
	Tp       string  `xml:"tp,omitempty"`
}

type CNRGetRealDeparturesResponse struct {
	XMLName xml.Name `xml:"http://PublicService/ CNR_GetRealDeparturesResponse"`

	CNRGetRealDeparturesResult struct {
	} `xml:"CNR_GetRealDeparturesResult,omitempty"`
}

type GetPlanedDeparutres struct {
	XMLName xml.Name `xml:"http://PublicService/ GetPlanedDeparutres"`

	S        string  `xml:"s,omitempty"`
	Idwersja float64 `xml:"id_wersja,omitempty"`
	Id       string  `xml:"id,omitempty"`
	Tpd      string  `xml:"tpd,omitempty"`
	Nrl      string  `xml:"nrl,omitempty"`
	God      string  `xml:"god,omitempty"`
	Gdo      string  `xml:"gdo,omitempty"`
	Tpp      string  `xml:"tpp,omitempty"`
	Kierunek string  `xml:"kierunek,omitempty"`
}

type GetPlanedDeparutresResponse struct {
	XMLName xml.Name `xml:"http://PublicService/ GetPlanedDeparutresResponse"`

	GetPlanedDeparutresResult struct {
	} `xml:"GetPlanedDeparutresResult,omitempty"`
}

type GetPlanedDeparutresFullInfo struct {
	XMLName xml.Name `xml:"http://PublicService/ GetPlanedDeparutresFullInfo"`

	S        string  `xml:"s,omitempty"`
	Idwersja float64 `xml:"id_wersja,omitempty"`
	Id       string  `xml:"id,omitempty"`
	Krs      string  `xml:"krs,omitempty"`
}

type GetPlanedDeparutresFullInfoResponse struct {
	XMLName xml.Name `xml:"http://PublicService/ GetPlanedDeparutresFullInfoResponse"`

	GetPlanedDeparutresFullInfoResult struct {
	} `xml:"GetPlanedDeparutresFullInfoResult,omitempty"`
}

type GetNote struct {
	XMLName xml.Name `xml:"http://PublicService/ GetNote"`

	S        string  `xml:"s,omitempty"`
	Idwersja float64 `xml:"id_wersja,omitempty"`
	Id       string  `xml:"id,omitempty"`
}

type GetNoteResponse struct {
	XMLName xml.Name `xml:"http://PublicService/ GetNoteResponse"`

	GetNoteResult struct {
	} `xml:"GetNoteResult,omitempty"`
}

type GetStopsOnStreet struct {
	XMLName xml.Name `xml:"http://PublicService/ GetStopsOnStreet"`

	S        string  `xml:"s,omitempty"`
	Idwersja float64 `xml:"id_wersja,omitempty"`
	Idul     int32   `xml:"id_ul,omitempty"`
}

type GetStopsOnStreetResponse struct {
	XMLName xml.Name `xml:"http://PublicService/ GetStopsOnStreetResponse"`

	GetStopsOnStreetResult struct {
	} `xml:"GetStopsOnStreetResult,omitempty"`
}

type GetStopsByNumber struct {
	XMLName xml.Name `xml:"http://PublicService/ GetStopsByNumber"`

	S        string  `xml:"s,omitempty"`
	Idwersja float64 `xml:"id_wersja,omitempty"`
	Nr       string  `xml:"nr,omitempty"`
}

type GetStopsByNumberResponse struct {
	XMLName xml.Name `xml:"http://PublicService/ GetStopsByNumberResponse"`

	GetStopsByNumberResult struct {
	} `xml:"GetStopsByNumberResult,omitempty"`
}

type GetStopsByName struct {
	XMLName xml.Name `xml:"http://PublicService/ GetStopsByName"`

	S        string  `xml:"s,omitempty"`
	Idwersja float64 `xml:"id_wersja,omitempty"`
	Q        string  `xml:"q,omitempty"`
	Limit    string  `xml:"limit,omitempty"`
	Transp   string  `xml:"transp,omitempty"`
}

type GetStopsByNameResponse struct {
	XMLName xml.Name `xml:"http://PublicService/ GetStopsByNameResponse"`

	GetStopsByNameResult struct {
	} `xml:"GetStopsByNameResult,omitempty"`
}

type GetStopsDropDown struct {
	XMLName xml.Name `xml:"http://PublicService/ GetStopsDropDown"`

	S     string `xml:"s,omitempty"`
	Q     string `xml:"q,omitempty"`
	Limit string `xml:"limit,omitempty"`
}

type GetStopsDropDownResponse struct {
	XMLName xml.Name `xml:"http://PublicService/ GetStopsDropDownResponse"`

	GetStopsDropDownResult struct {
	} `xml:"GetStopsDropDownResult,omitempty"`
}

type GetGoogleStops struct {
	XMLName xml.Name `xml:"http://PublicService/ GetGoogleStops"`

	S        string  `xml:"s,omitempty"`
	Idwersja float64 `xml:"id_wersja,omitempty"`
}

type GetGoogleStopsResponse struct {
	XMLName xml.Name `xml:"http://PublicService/ GetGoogleStopsResponse"`

	GetGoogleStopsResult struct {
	} `xml:"GetGoogleStopsResult,omitempty"`
}

type DajGrafyGoogle struct {
	XMLName xml.Name `xml:"http://PublicService/ DajGrafyGoogle"`

	S         string  `xml:"s,omitempty"`
	Idwersja  float64 `xml:"id_wersja,omitempty"`
	Numerlini string  `xml:"numer_lini,omitempty"`
	Wartrasy  string  `xml:"war_trasy,omitempty"`
	Cdata     string  `xml:"cdata,omitempty"`
	Idkursu   string  `xml:"id_kursu,omitempty"`
}

type DajGrafyGoogleResponse struct {
	XMLName xml.Name `xml:"http://PublicService/ DajGrafyGoogleResponse"`

	DajGrafyGoogleResult struct {
	} `xml:"DajGrafyGoogleResult,omitempty"`
}

type DajGrafyGoogleKlient struct {
	XMLName xml.Name `xml:"http://PublicService/ DajGrafyGoogleKlient"`

	S         string  `xml:"s,omitempty"`
	Idwersja  float64 `xml:"id_wersja,omitempty"`
	Numerlini string  `xml:"numer_lini,omitempty"`
	Wartrasy  string  `xml:"war_trasy,omitempty"`
	Cdata     string  `xml:"cdata,omitempty"`
	Idkursu   string  `xml:"id_kursu,omitempty"`
}

type DajGrafyGoogleKlientResponse struct {
	XMLName xml.Name `xml:"http://PublicService/ DajGrafyGoogleKlientResponse"`

	DajGrafyGoogleKlientResult struct {
	} `xml:"DajGrafyGoogleKlientResult,omitempty"`
}

type GetDayTypes struct {
	XMLName xml.Name `xml:"http://PublicService/ GetDayTypes"`

	Idwersja float64 `xml:"id_wersja,omitempty"`
}

type GetDayTypesResponse struct {
	XMLName xml.Name `xml:"http://PublicService/ GetDayTypesResponse"`

	GetDayTypesResult struct {
	} `xml:"GetDayTypesResult,omitempty"`
}

type DajCentrumMapy struct {
	XMLName xml.Name `xml:"http://PublicService/ DajCentrumMapy"`

	S string `xml:"s,omitempty"`
}

type DajCentrumMapyResponse struct {
	XMLName xml.Name `xml:"http://PublicService/ DajCentrumMapyResponse"`

	DajCentrumMapyResult struct {
	} `xml:"DajCentrumMapyResult,omitempty"`
}

type GetRoutes struct {
	XMLName xml.Name `xml:"http://PublicService/ GetRoutes"`

	S        string  `xml:"s,omitempty"`
	Idwersja float64 `xml:"id_wersja,omitempty"`
	Tp       string  `xml:"tp,omitempty"`
}

type GetRoutesResponse struct {
	XMLName xml.Name `xml:"http://PublicService/ GetRoutesResponse"`

	GetRoutesResult struct {
	} `xml:"GetRoutesResult,omitempty"`
}

type GetStreets struct {
	XMLName xml.Name `xml:"http://PublicService/ GetStreets"`

	S string `xml:"s,omitempty"`
	N string `xml:"n,omitempty"`
}

type GetStreetsResponse struct {
	XMLName xml.Name `xml:"http://PublicService/ GetStreetsResponse"`

	GetStreetsResult struct {
	} `xml:"GetStreetsResult,omitempty"`
}

type SearchConnection struct {
	XMLName xml.Name `xml:"http://PublicService/ SearchConnection"`

	Px1      float64 `xml:"px1,omitempty"`
	Py1      float64 `xml:"py1,omitempty"`
	Px2      float64 `xml:"px2,omitempty"`
	Py2      float64 `xml:"py2,omitempty"`
	Rozp     string  `xml:"rozp,omitempty"`
	Data     string  `xml:"data,omitempty"`
	Ileprzes int32   `xml:"ile_przes,omitempty"`
	Tryb     int32   `xml:"tryb,omitempty"`
	Lang     string  `xml:"lang,omitempty"`
	Kodbledu uint16  `xml:"kod_bledu,omitempty"`
}

type SearchConnectionResponse struct {
	XMLName xml.Name `xml:"http://PublicService/ SearchConnectionResponse"`

	SearchConnectionResult struct {
	} `xml:"SearchConnectionResult,omitempty"`

	Kodbledu uint16 `xml:"kod_bledu,omitempty"`
}

type GetTicketTerminals struct {
	XMLName xml.Name `xml:"http://PublicService/ GetTicketTerminals"`
}

type GetTicketTerminalsResponse struct {
	XMLName xml.Name `xml:"http://PublicService/ GetTicketTerminalsResponse"`

	GetTicketTerminalsResult struct {
	} `xml:"GetTicketTerminalsResult,omitempty"`
}

type PublicServiceSoap struct {
	client *SOAPClient
}

func NewPublicServiceSoap(url string, tls bool, auth *BasicAuth) *PublicServiceSoap {
	if url == "" {
		url = "http://sip.zdzit.olsztyn.eu/PublicService.asmx"
	}
	client := NewSOAPClient(url, tls, auth)

	return &PublicServiceSoap{
		client: client,
	}
}

func (service *PublicServiceSoap) CNRGetVehicles(request *CNRGetVehicles) (*CNRGetVehiclesResponse, error) {
	response := new(CNRGetVehiclesResponse)
	err := service.client.Call("http://PublicService/CNR_GetVehicles", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *PublicServiceSoap) CNRGetVehicleNumbers(request *CNRGetVehicleNumbers) (*CNRGetVehicleNumbersResponse, error) {
	response := new(CNRGetVehicleNumbersResponse)
	err := service.client.Call("http://PublicService/CNR_GetVehicleNumbers", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *PublicServiceSoap) CNRGetVehicleTypes(request *CNRGetVehicleTypes) (*CNRGetVehicleTypesResponse, error) {
	response := new(CNRGetVehicleTypesResponse)
	err := service.client.Call("http://PublicService/CNR_GetVehicleTypes", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *PublicServiceSoap) CNRRouteVariants(request *CNRRouteVariants) (*CNRRouteVariantsResponse, error) {
	response := new(CNRRouteVariantsResponse)
	err := service.client.Call("http://PublicService/CNR_RouteVariants", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *PublicServiceSoap) CNRGetVehicleDepartures(request *CNRGetVehicleDepartures) (*CNRGetVehicleDeparturesResponse, error) {
	response := new(CNRGetVehicleDeparturesResponse)
	err := service.client.Call("http://PublicService/CNR_GetVehicleDepartures", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *PublicServiceSoap) GetTimeline(request *GetTimeline) (*GetTimelineResponse, error) {
	response := new(GetTimelineResponse)
	err := service.client.Call("http://PublicService/GetTimeline", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *PublicServiceSoap) CNRGetRealDepartures(request *CNRGetRealDepartures) (*CNRGetRealDeparturesResponse, error) {
	response := new(CNRGetRealDeparturesResponse)
	err := service.client.Call("http://PublicService/CNR_GetRealDepartures", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *PublicServiceSoap) GetPlanedDeparutres(request *GetPlanedDeparutres) (*GetPlanedDeparutresResponse, error) {
	response := new(GetPlanedDeparutresResponse)
	err := service.client.Call("http://PublicService/GetPlanedDeparutres", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *PublicServiceSoap) GetPlanedDeparutresFullInfo(request *GetPlanedDeparutresFullInfo) (*GetPlanedDeparutresFullInfoResponse, error) {
	response := new(GetPlanedDeparutresFullInfoResponse)
	err := service.client.Call("http://PublicService/GetPlanedDeparutresFullInfo", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *PublicServiceSoap) GetNote(request *GetNote) (*GetNoteResponse, error) {
	response := new(GetNoteResponse)
	err := service.client.Call("http://PublicService/GetNote", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *PublicServiceSoap) GetStopsOnStreet(request *GetStopsOnStreet) (*GetStopsOnStreetResponse, error) {
	response := new(GetStopsOnStreetResponse)
	err := service.client.Call("http://PublicService/GetStopsOnStreet", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *PublicServiceSoap) GetStopsByNumber(request *GetStopsByNumber) (*GetStopsByNumberResponse, error) {
	response := new(GetStopsByNumberResponse)
	err := service.client.Call("http://PublicService/GetStopsByNumber", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *PublicServiceSoap) GetStopsByName(request *GetStopsByName) (*GetStopsByNameResponse, error) {
	response := new(GetStopsByNameResponse)
	err := service.client.Call("http://PublicService/GetStopsByName", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *PublicServiceSoap) GetStopsDropDown(request *GetStopsDropDown) (*GetStopsDropDownResponse, error) {
	response := new(GetStopsDropDownResponse)
	err := service.client.Call("http://PublicService/GetStopsDropDown", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *PublicServiceSoap) GetGoogleStops(request *GetGoogleStops) (*GetGoogleStopsResponse, error) {
	response := new(GetGoogleStopsResponse)
	err := service.client.Call("http://PublicService/GetGoogleStops", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *PublicServiceSoap) DajGrafyGoogle(request *DajGrafyGoogle) (*DajGrafyGoogleResponse, error) {
	response := new(DajGrafyGoogleResponse)
	err := service.client.Call("http://PublicService/DajGrafyGoogle", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *PublicServiceSoap) DajGrafyGoogleKlient(request *DajGrafyGoogleKlient) (*DajGrafyGoogleKlientResponse, error) {
	response := new(DajGrafyGoogleKlientResponse)
	err := service.client.Call("http://PublicService/DajGrafyGoogleKlient", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *PublicServiceSoap) GetDayTypes(request *GetDayTypes) (*GetDayTypesResponse, error) {
	response := new(GetDayTypesResponse)
	err := service.client.Call("http://PublicService/GetDayTypes", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *PublicServiceSoap) DajCentrumMapy(request *DajCentrumMapy) (*DajCentrumMapyResponse, error) {
	response := new(DajCentrumMapyResponse)
	err := service.client.Call("http://PublicService/DajCentrumMapy", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *PublicServiceSoap) GetRoutes(request *GetRoutes) (*GetRoutesResponse, error) {
	response := new(GetRoutesResponse)
	err := service.client.Call("http://PublicService/GetRoutes", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *PublicServiceSoap) GetStreets(request *GetStreets) (*GetStreetsResponse, error) {
	response := new(GetStreetsResponse)
	err := service.client.Call("http://PublicService/GetStreets", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *PublicServiceSoap) SearchConnection(request *SearchConnection) (*SearchConnectionResponse, error) {
	response := new(SearchConnectionResponse)
	err := service.client.Call("http://PublicService/SearchConnection", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *PublicServiceSoap) GetTicketTerminals(request *GetTicketTerminals) (*GetTicketTerminalsResponse, error) {
	response := new(GetTicketTerminalsResponse)
	err := service.client.Call("http://PublicService/GetTicketTerminals", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type PublicServiceHttpGet struct {
	client *SOAPClient
}

func NewPublicServiceHttpGet(url string, tls bool, auth *BasicAuth) *PublicServiceHttpGet {
	if url == "" {
		url = ""
	}
	client := NewSOAPClient(url, tls, auth)

	return &PublicServiceHttpGet{
		client: client,
	}
}

func (service *PublicServiceHttpGet) CNRGetVehicles(request *CNRGetVehicles) (*CNRGetVehiclesResponse, error) {
	response := new(CNRGetVehiclesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *PublicServiceHttpGet) CNRGetVehicleNumbers(request *CNRGetVehicleNumbers) (*CNRGetVehicleNumbersResponse, error) {
	response := new(CNRGetVehicleNumbersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *PublicServiceHttpGet) CNRGetVehicleTypes(request *CNRGetVehicleTypes) (*CNRGetVehicleTypesResponse, error) {
	response := new(CNRGetVehicleTypesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *PublicServiceHttpGet) GetStopsDropDown(request *GetStopsDropDown) (*GetStopsDropDownResponse, error) {
	response := new(GetStopsDropDownResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *PublicServiceHttpGet) DajCentrumMapy(request *DajCentrumMapy) (*DajCentrumMapyResponse, error) {
	response := new(DajCentrumMapyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *PublicServiceHttpGet) GetStreets(request *GetStreets) (*GetStreetsResponse, error) {
	response := new(GetStreetsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *PublicServiceHttpGet) GetTicketTerminals() (*GetTicketTerminalsResponse, error) {
	response := new(GetTicketTerminalsResponse)
	err := service.client.Call("", nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type PublicServiceHttpPost struct {
	client *SOAPClient
}

func NewPublicServiceHttpPost(url string, tls bool, auth *BasicAuth) *PublicServiceHttpPost {
	if url == "" {
		url = ""
	}
	client := NewSOAPClient(url, tls, auth)

	return &PublicServiceHttpPost{
		client: client,
	}
}

func (service *PublicServiceHttpPost) CNRGetVehicles(request *CNRGetVehicles) (*CNRGetVehiclesResponse, error) {
	response := new(CNRGetVehiclesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *PublicServiceHttpPost) CNRGetVehicleNumbers(request *CNRGetVehicleNumbers) (*CNRGetVehicleNumbersResponse, error) {
	response := new(CNRGetVehicleNumbersResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *PublicServiceHttpPost) CNRGetVehicleTypes(request *CNRGetVehicleTypes) (*CNRGetVehicleTypesResponse, error) {
	response := new(CNRGetVehicleTypesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *PublicServiceHttpPost) GetStopsDropDown(request *GetStopsDropDown) (*GetStopsDropDownResponse, error) {
	response := new(GetStopsDropDownResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *PublicServiceHttpPost) DajCentrumMapy(request *DajCentrumMapy) (*DajCentrumMapyResponse, error) {
	response := new(DajCentrumMapyResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *PublicServiceHttpPost) GetStreets(request *GetStreets) (*GetStreetsResponse, error) {
	response := new(GetStreetsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *PublicServiceHttpPost) GetTicketTerminals() (*GetTicketTerminalsResponse, error) {
	response := new(GetTicketTerminalsResponse)
	err := service.client.Call("", nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

var timeout = time.Duration(30 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`

	Body SOAPBody
}

type SOAPHeader struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`

	Header interface{}
}

type SOAPBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`

	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

	Code   string `xml:"faultcode,omitempty"`
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`
	Detail string `xml:"detail,omitempty"`
}

type BasicAuth struct {
	Login    string
	Password string
}

type SOAPClient struct {
	url  string
	tls  bool
	auth *BasicAuth
}

func (b *SOAPBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

func (f *SOAPFault) Error() string {
	return f.String
}

func NewSOAPClient(url string, tls bool, auth *BasicAuth) *SOAPClient {
	return &SOAPClient{
		url:  url,
		tls:  tls,
		auth: auth,
	}
}

func (s *SOAPClient) Call(soapAction string, request, response interface{}) error {
	envelope := SOAPEnvelope{
	// Header:        SoapHeader{},
	}

	envelope.Body.Content = request
	buffer := new(bytes.Buffer)

	encoder := xml.NewEncoder(buffer)
	// encoder.Indent("  ", "    ")

	if err := encoder.Encode(envelope); err != nil {
		return err
	}

	if err := encoder.Flush(); err != nil {
		return err
	}

	log.Println(buffer.String())

	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return err
	}
	if s.auth != nil {
		req.SetBasicAuth(s.auth.Login, s.auth.Password)
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	if soapAction != "" {
		req.Header.Add("SOAPAction", soapAction)
	}

	req.Header.Set("User-Agent", "PZI TARAN")
	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: s.tls,
		},
		Dial: dialTimeout,
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	rawbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if len(rawbody) == 0 {
		log.Println("empty response")
		return nil
	}

	log.Println(string(rawbody))
	respEnvelope := new(SOAPEnvelope)
	respEnvelope.Body = SOAPBody{Content: response}
	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fault
	}

	return nil
}
