package zdzit

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/jlaffaye/ftp"
	"io"
	"sort"
	"strconv"
	"strings"
)

type BusStop struct {
	Number		string
	Name       	string
	StreetName 	string
	GetLatLong	struct {
		Unsanitized 	string
		Latitude 		float64
		Longitude 		float64
	}
}

func GetLatestFile(url string) (*ftp.Response, error) {

	conn, err := ftp.Connect(url)
	if err != nil {
		return nil, err
	}
	defer conn.Quit()

	if err := conn.Login("anonymous", "anonymous"); err != nil {
		return nil, err
	}

	var files []string
	entries, _ := conn.List("")

	for _, entry := range entries {
		date := entry.Name
		files = append(files, date)
	}

	sort.Strings(files)

	data, err := conn.Retr(files[len(files)-1])
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	data.Read(buf.Bytes())

	return data, nil
}

func ParseBusStop(url string) ([]BusStop, error) {

	var busStops []BusStop

	data, err := GetLatestFile(url)
	if err != nil {
		return nil, err
	}

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = ';'
		return r
	})

	if err := gocsv.UnmarshalWithoutHeaders(data, &busStops); err != nil {
		fmt.Println(err)
	}

	for i, stop := range busStops {
		splittedLatLong := strings.Split(stop.GetLatLong.Unsanitized, ",")

		f1, err := strconv.ParseFloat(splittedLatLong[0], 64)
		if err != nil {
			f1 = 0
		}

		splittedLatLong = append(splittedLatLong, "")

		f2, err := strconv.ParseFloat(splittedLatLong[1], 64)
		if err != nil {
			f2 = 0
		}

		busStops[i].GetLatLong.Latitude		= f1
		busStops[i].GetLatLong.Longitude	= f2

	}

	return busStops[1:], nil 	// we need to ignore names of the columns
}