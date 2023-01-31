package models

import "linkedInLearning/tempService4/data"

type cities struct {
	cityMap map[string]CityTemp
}

type Cities interface {
	ListAll() []CityTemp
}

// NewCities initialises the Cities data structure by calling the
// ReadData method to read information from file.
func NewCities(reader data.DataReader) (Cities, error) {
	d, err := reader.ReadData()
	if err != nil {
		return nil, err
	}
	cmap := make(map[string]CityTemp)
	for _, r := range d {
		city := NewCity(r.Name, r.TempC, r.HasBeach, r.HasMountain)
		city.TempF()
		cmap[r.Id] = city
	}

	return &cities{
		cityMap: cmap,
	}, nil
}

// ListAll returns a slice of all the cities.
func (c cities) ListAll() []CityTemp {
	var cs []CityTemp
	for _, rc := range c.cityMap {
		cs = append(cs, rc)
	}
	return cs
}
