package main

import (
	"fmt"

	"linkedInLearning/tempService2/models"
	"linkedInLearning/tempService2/printer"
)

func main() {
	fmt.Printf("Welcome to the LinkedIn Learning Temperature Service!\n\n")

	// initialise printer and defer cleanup
	p := printer.New()
	defer p.Cleanup()
	p.CityHeader()

	// setup 3 cities
	lon := models.NewCity("London", 7.5)
	ams := models.NewCity("Amsterdam", 11)
	nyc := models.NewCity("New York", -3)

	// print all the cities
	p.CityDetails(lon)
	p.CityDetails(ams)
	p.CityDetails(nyc)
}
