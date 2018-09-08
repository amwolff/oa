package zdzit

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/jlaffaye/ftp"
	"strconv"
	"strings"
	"time"
)

type BusStop struct {
	Number     string
	Name       string
	StreetName string
	Latitude   float64
	Longitude  float64
}

var AllBusStops []BusStop


func ReadCSVFromFTP(url string) ([][]string, error) {

	conn, err := ftp.Dial(url)

	if err != nil {
		fmt.Println(err)
	}

	defer conn.Quit()

	if err := conn.Login("anonymous", "anonymous"); err != nil {
		fmt.Println(err)
	}

	// LOOKING FOR LATEST CSV FILE

	timeNow := time.Now()
	timeNow = timeNow.UTC()
	offset := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	min := timeNow.Sub(offset)
	var NewestCSV string

	entries, _ := conn.List("")

	for _, entry := range entries {
		date := entry.Time
		k := timeNow.Sub(date)

		if k < min {
			min = k
			if entry.Name[0] == 'g' {
				NewestCSV = entry.Name
			}
		}
	}

	// READING DATA FROM CSV

	data, err := conn.Retr(NewestCSV)

	if err != nil {
		fmt.Println(err)
	}

	buf := new(bytes.Buffer)
	data.Read(buf.Bytes())

	if err != nil {
		fmt.Println(err)
	}

	defer data.Close()

	r := csv.NewReader(data)
	r.Comma = ';'
	file, err := r.ReadAll()

	if err != nil {
		fmt.Println(err)
	}

	return file, nil
}

func ParseData(file [][]string) error {

	for _, line := range file {
		splittedCords := strings.Split(line[3],",")
		f1, err := strconv.ParseFloat(splittedCords[0], 64)

		//TODO: not exactly a great idea
		if err != nil {
			f1 = 0
		}

		if len(splittedCords) <= 1 {
			continue
		}

		f2, err := strconv.ParseFloat(splittedCords[1], 64)

		//same
		if err != nil {
			f2 = 0
		}

		bs := BusStop{
			Number:     line[0],
			Name:       line[1],
			StreetName: line[2],
			Latitude:   f1,
			Longitude:  f2,
		}

		AllBusStops = append(AllBusStops, bs)
	}

	return nil
}

func GetBusStops() ([]BusStop, error) {

	url := "helios.zdzit.olsztyn.eu:21"
	data, err := ReadCSVFromFTP(url)

	if err != nil {
		fmt.Println("error")
	}

	ParseData(data)
	return AllBusStops, nil
}



