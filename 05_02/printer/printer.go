package printer

import (
	"fmt"
	"linkedInLearning/tempService9/models"
	"os"
	"text/tabwriter"
)

type Printer struct {
	w *tabwriter.Writer
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
func (p *Printer) CityDetails(c models.CityTemp, q models.CityQuery) {
	fmt.Fprintf(p.w, "%v\t%v\t%v\t%v\t%v\n", c.Name(), c.TempC(q), c.TempF(q), c.BeachVacationReady(q), c.SkiVacationReady(q))
}

// Cleanup closes up the printer instance
func (p *Printer) Cleanup() {
	p.w.Flush()
}
