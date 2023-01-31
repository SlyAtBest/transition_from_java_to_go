package models

import "linkedInLearning/tempService5/data"

type Cities struct {
	cityMap map[string]*City
}

type DataReader interface {
	ReadData() ([]data.Response, error)
}

// NewCities initialises the Cities data structure by calling the
// ReadData method to read information from file.
func NewCities(reader DataReader) (*Cities, error) {
	d, err := reader.ReadData()
	if err != nil {
		return nil, err
	}
	cmap := make(map[string]*City)
	for _, r := range d {
		cmap[r.Id] = NewCity(r.Name, r.TempC, r.HasBeach, r.HasMountain)
	}

	return &Cities{
		cityMap: cmap,
	}, nil
}

// ListAll returns a slice of all the cities.
func (c Cities) ListAll() []*City {
	var cs []*City
	for _, rc := range c.cityMap {
		cs = append(cs, rc)
	}
	return cs
}
