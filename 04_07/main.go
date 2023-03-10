package main

import (
	"flag"
	"fmt"
	"linkedInLearning/tempService6/data"

	"linkedInLearning/tempService6/models"
	"linkedInLearning/tempService6/printer"
)

func main() {
	fmt.Printf("Welcome to the LinkedIn Learning Temperature Service!\n\n")
	beachReady := flag.Bool("beach", false, "Display only beach ready destinations")
	skiReady := flag.Bool("ski", false, "Display only ski ready destinations")
	flag.Parse()

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

	// filter cities
	cs := cities.Filter(*beachReady, *skiReady)
	//print all the cities
	for _, c := range cs {
		p.CityDetails(c)
	}
}
