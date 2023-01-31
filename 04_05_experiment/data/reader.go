package data

import (
	"encoding/json"
	"io/ioutil"
)

type Response struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	HasBeach    bool    `json:"hasBeach"`
	HasMountain bool    `json:"hasMountain"`
	TempC       float64 `json:"tempC"`
}

type Reader struct {
	path string
}

// NewReader initialises a DataReader
func NewReader() *Reader {
	return &Reader{
		path: "./data/cities.json",
	}
}

// ReadData is a helper method to read the file at
// the given path and return a response array.
func (r *Reader) ReadData() ([]Response, error) {
	file, err := ioutil.ReadFile(r.path)
	if err != nil {
		return nil, err
	}

	var data []Response
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
