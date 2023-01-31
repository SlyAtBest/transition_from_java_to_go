package models

var (
	// hot enough for the beach
	beachVacationThreshold float64 = 22
	// cold enough to snow
	skiVacationThreshold float64 = -2
)

type City struct {
	name        string
	tempC       float64 // temperature in Celcius
	hasBeach    bool
	hasMountain bool
}

// NewCity creates a new City instance with the given name and Celsius temperated
func NewCity(n string, t float64, hasBeach bool, hasMountain bool) *City {
	return &City{
		name:        n,
		tempC:       t,
		hasBeach:    hasBeach,
		hasMountain: hasMountain,
	}
}

// Name returns the city name
func (c City) Name() string {
	return c.name
}

func (c City) TempC() float64 {
	return c.tempC
}

// TempF returns the city temperature converted to Fahrenheit
func (c City) TempF() float64 {
	return (c.tempC * 9 / 5) + 32
}

func (c City) BeachVacationReady() bool {
	return c.hasBeach && c.tempC > beachVacationThreshold
}

func (c City) SkiVacationReady() bool {
	return c.hasMountain && c.tempC < skiVacationThreshold
}
