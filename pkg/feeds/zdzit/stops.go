package zdzit

import (
	"encoding/csv"
	"github.com/gocarina/gocsv"
	"github.com/jlaffaye/ftp"
	"io"
	"sort"
	"strconv"
	"strings"
)

type BusStop struct {
	Number     string
	Name       string
	StreetName string
	LatLng     struct {
		Unsanitized string
		Latitude    float64
		Longitude   float64
	}
}

func getLatestFile(url string) (*ftp.Response, error) {

	conn, err := ftp.Connect(url)
	if err != nil {
		return nil, err
	}
	defer conn.Quit()

	if err := conn.Login("anonymous", "anonymous"); err != nil {
		return nil, err
	}

	var files []string
	entries, err := conn.List("")
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		date := entry.Name
		files = append(files, date)
	}

	sort.Strings(files)

	return conn.Retr(files[len(files)-1])
}

func ParseBusStop(url string) ([]BusStop, error) {

	var busStops []BusStop

	data, err := getLatestFile(url)
	if err != nil {
		return nil, err
	}

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = ';'
		return r
	})

	if err := gocsv.UnmarshalWithoutHeaders(data, &busStops); err != nil {
		return nil, err
	}

	for i, busStop := range busStops[1:] {
		splittedLatLong := strings.Split(busStop.LatLng.Unsanitized, ",")
		if len(splittedLatLong) != 2 {
			busStops[i].LatLng.Latitude = 0
			busStops[i].LatLng.Longitude = 0
			continue
		}

		f1, err := strconv.ParseFloat(splittedLatLong[0], 64)
		if err != nil {
			f1 = 0
		}

		f2, err := strconv.ParseFloat(splittedLatLong[1], 64)
		if err != nil {
			f2 = 0
		}

		busStops[i].LatLng.Latitude = f1
		busStops[i].LatLng.Longitude = f2

	}

	return busStops[1:], nil // we need to ignore columns description
}
