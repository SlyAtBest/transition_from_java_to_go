package printer

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type Printer struct {
	w *tabwriter.Writer
}

// CityTemp interface defines all things city temperature related
type CityTemp interface {
	Name() string
	TempC() float64
	TempF() float64
	BeachVacationReady() bool
	SkiVacationReady() bool
}

// New initialises a new Printer instance
func New() *Printer {
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 3, ' ', tabwriter.TabIndent)
	return &Printer{
		w: w,
	}
}

// CityHeader prints out the table header for the city table
func (p *Printer) CityHeader() {
	fmt.Fprintln(p.w, "Name\tTempC\tTempF\tBeach Vacation Ready?\tSki vacation ready?")
}

// CityDetails prints out the given city
func (p *Printer) CityDetails(c CityTemp) {
	fmt.Fprintf(p.w, "%v\t%v\t%v\t%v\t%v\n", c.Name(), c.TempC(), c.TempF(), c.BeachVacationReady(), c.SkiVacationReady())
}

// Cleanup closes up the printer instance
func (p *Printer) Cleanup() {
	p.w.Flush()
}
