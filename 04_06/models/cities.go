package models

import (
	"linkedInLearning/tempService5/data"
	"sort"
)

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
		cmap[r.Id] = NewCity(r.Name, r.TempC, r.HasBeach, r.HasMountain)
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
	sortAlphabetically(cs)
	return cs
}

// sortAlphabetically is a helper method that
// sorts a slice of cities alphabetically by name
func sortAlphabetically(cs []CityTemp) {
	sort.Slice(cs, func(i, j int) bool {
		return cs[i].Name() < cs[j].Name()
	})
}
