package zdzit

import (
	"encoding/csv"
	"net/http"
	"strconv"
	"strings"
)

type GetBusStopsResponse struct {
	GetBusStopsResult struct {
		Parsed []BusStop
	}
}

type BusStop struct {
	number        string
	name          string
	street_name   string
	coordinates_X float64
	coordinates_Y float64
}

//TODO: Can't retrieve data from ftp yet
func ReadCSVFromURL(url string) ([][]string, error) {

	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	r := csv.NewReader(response.Body)
	r.Comma = ';'

	data, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r * GetBusStopsResponse) ParseData(file [][]string) error {

	for _, line := range file {

		splitted_cords := strings.Split(line[3],",")
		f1, err := strconv.ParseFloat(splitted_cords[0], 64)

		//TODO: not exactly a great idea
		if err != nil {
			f1 = 0
		}

		if len(splitted_cords) <= 1 {
			continue
		}

		f2, err := strconv.ParseFloat(splitted_cords[1], 64)

		//same
		if err != nil {
			f2 = 0
		}

		bs := BusStop{
			number:        line[0],
			name:          line[1],
			street_name:   line[2],
			coordinates_X: f1,
			coordinates_Y: f2,
		}

		r.GetBusStopsResult.Parsed = append(r.GetBusStopsResult.Parsed, bs)
	}

	return nil
}



