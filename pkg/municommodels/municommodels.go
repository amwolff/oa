package municommodels

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Envelope represents SOAP response envelope
type Envelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`

	// Body represents SOAP response body
	Body struct {
		// InnerXML is actual response body to be unmarshaled
		InnerXML []byte `xml:",innerxml"`
	} `xml:"Body"`
}

// GetRouteAndVariants represents SOAP request body of an GetRouteAndVariants
// endpoint.
type GetRouteAndVariants struct {
	// reference:
	// PublicService.GetRouteAndVariants(String s, Decimal id_wersja, String q)
	//
	// I got the signature of this function from a panic stacktrace.

	XMLName xml.Name `xml:"http://PublicService/ GetRouteAndVariants"`

	S        string `xml:"s"`
	Idwersja int    `xml:"id_wersja"`
	Q        string `xml:"q"`

	// Note on "id_wersja" tag:
	// For some reason requests without this field or with this field not equal
	// to zero are recognized as malformed.
}

type GetRouteAndVariantsResponse struct {
	XMLName xml.Name `xml:"http://PublicService/ GetRouteAndVariantsResponse"`

	GetRouteAndVariantsResult struct {
		L []struct {
			// Fields names are reverse engineered from the biggest (ugliest?)
			// JS on the site.
			Number       string `xml:"l,attr,omitempty"`
			Description  string `xml:"o,attr,omitempty"`
			Description2 string `xml:"o2,attr,omitempty"`
			Variant      int    `xml:"wt,attr,omitempty"`
			Transport    string `xml:"t,attr,omitempty"`
			Direction    string `xml:"d,attr,omitempty"`
		} `xml:"L,omitempty"`
	} `xml:"GetRouteAndVariantsResult>R,omitempty"`
}

// WebServiceClient abstracts communication with PublicService.asmx @ PIS
type WebServiceClient struct {
	client        *http.Client
	dataTmpl      *template.Template
	log           logrus.FieldLogger
	name, ua, url string
}

func NewWebServiceClient(logger logrus.FieldLogger, name, UA, URL string) *WebServiceClient {
	c := &WebServiceClient{
		client: http.DefaultClient,
		log:    logger,
		name:   name,
		ua:     UA,
		url:    URL,
	}

	c.log.WithField("web-service-client", c.name)

	t, err := template.New("data-binary").Parse(`<?xml version='1.0' encoding='utf-8'?>
	<soap:Envelope xmlns:soap='http://schemas.xmlsoap.org/soap/envelope/'>
    <soap:Body>{{ . }}</soap:Body>
	</soap:Envelope>`)
	if err != nil {
		c.log.WithError(err).Fatal("Cannot parse template")
	}
	c.dataTmpl = t

	return c
}

func (c *WebServiceClient) UnmarshalSOAP(data []byte, v interface{}) error {
	b := bytes.NewReader(data)
	d := xml.NewDecoder(b)

	// TODO(amw): could also unmarshal SOAP fault's in case of error

	var e Envelope
	if err := d.Decode(&e); err != nil {
		return err
	}

	b.Reset(e.Body.InnerXML)

	if err := d.Decode(&v); err != nil {
		return err
	}

	return nil
}

func (c *WebServiceClient) call(ctx context.Context, cookies []http.Cookie, method string,
	data interface{}) ([]byte, error) {

	reqBody, err := xml.Marshal(data)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := c.dataTmpl.Execute(&buf, string(reqBody)); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.url, &buf)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	// Example response headers:
	// HTTP/1.1 200 OK
	// Cache-Control: no-cache
	// Pragma: no-cache
	// Content-Type: text/xml; charset=utf-8
	// Content-Encoding: gzip
	// Expires: -1
	// Vary: Accept-Encoding
	// Server: Microsoft-IIS/8.0
	// X-AspNet-Version: 4.0.30319
	// X-Powered-By: ASP.NET
	// Date: Fri, 17 Aug 2018 20:38:36 GMT
	// Content-Length: 507

	// req.Header.Set("Accept", "*/*")
	// req.Header.Set("Accept-Encoding", "gzip")
	// req.Header.Set("Accept-Language", "pl-PL,pl;q=0.9,en-US;q=0.8,en;q=0.7")
	// req.Header.Set("Age", "0")
	// req.Header.Set("Cache-Control", "no-cache")
	// req.Header.Set("Connection", "keep-alive")
	// req.Header.Set("Content-Length", "0")
	req.Header.Set("Content-Type", "text/xml;charset=UTF-8")
	// req.Header.Set("Host", "sip.zdzit.olsztyn.eu")
	// req.Header.Set("Origin", "http://sip.zdzit.olsztyn.eu")
	// req.Header.Set("Pragma", "no-cache")
	// req.Header.Set("Referer", "http://sip.zdzit.olsztyn.eu/")
	req.Header.Set("SOAPAction", fmt.Sprintf("http://PublicService/%s", method))
	req.Header.Set("User-Agent", c.ua)

	for _, c := range cookies {
		req.AddCookie(&c)
	}

	c.log.Debugf("Call%s: req.URL=%v", method, req.URL)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("server returned %d (%s) != 200", resp.StatusCode, resp.Status)
	}

	return respBody, nil
}

func (c *WebServiceClient) CallGetRouteAndVariants(ctx context.Context, cookies []http.Cookie,
	data GetRouteAndVariants) (*GetRouteAndVariantsResponse, error) {

	b, err := c.call(ctx, cookies, "GetRouteAndVariants", data)
	if err != nil {
		return nil, err
	}

	ret := &GetRouteAndVariantsResponse{}
	if err := c.UnmarshalSOAP(b, ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (c *WebServiceClient) CallCNRGetVehicles(ctx context.Context, cookies []http.Cookie,
	data CNRGetVehicles) (*CNRGetVehiclesResponse, error) {

	b, err := c.call(ctx, cookies, "CNR_GetVehicles", data)
	if err != nil {
		return nil, err
	}

	ret := &CNRGetVehiclesResponse{}
	if err := c.UnmarshalSOAP(b, ret); err != nil {
		return nil, err
	}

	if err := ret.Sanitize(); err != nil {
		return nil, err
	}

	return ret, nil
}
