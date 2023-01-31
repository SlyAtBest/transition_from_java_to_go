package main

import (
	"fmt"
	"linkedInLearning/tempService5/data"

	"linkedInLearning/tempService5/models"
	"linkedInLearning/tempService5/printer"
)

func main() {
	fmt.Printf("Welcome to the LinkedIn Learning Temperature Service!\n\n")

	// create cities
	cities, err := models.NewCities(data.NewReader())
	if err != nil {
		fmt.Println("Fatal error occurred: ", err)
		return
	}
	// initialise printer and defer cleanup
	p := printer.New()
	defer p.Cleanup()
	p.CityHeader()

	//print all the cities
	for _, c := range cities.ListAll() {
		p.CityDetails(c)
	}
}
