package models

type city struct {
	name  string
	tempC float64 // temperature in Celcius
}

// CityTemp interface defines all things city temperature related
type CityTemp interface {
	Name() string
	TempC() float64
	TempF() float64
}

// NewCity creates a new City instance with the given name and Celsius temperated
func NewCity(n string, t float64) CityTemp {
	return &city{
		name:  n,
		tempC: t,
	}
}

// Name returns the city name
func (c city) Name() string {
	return c.name
}

func (c city) TempC() float64 {
	return c.tempC
}

// TempF returns the city temperature converted to Fahrenheit
func (c city) TempF() float64 {
	return (c.tempC * 9 / 5) + 32
}
