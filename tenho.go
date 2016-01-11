package main

import (
	"flag"
	"github.com/hoshinotsuyoshi/tenho-go/src/tenho"
)

func main() {
	// opt-parse
	c := flag.Bool("no-chitoitsu", false, "Not apply chitoitsu")
	k := flag.Bool("no-kokushi", false, "Not apply kokushi")
	n := flag.Bool("no-normal", false, "Not apply normal")
	u := flag.Int("output-per-trial", 10000, "Output per n trials")
	s := flag.Int64("seed", 0, "Spcify a seed")

	flag.Parse()

	o := tenho.OptionStruct{
		NoChitoitsu:    *c,
		NoKokushi:      *k,
		NoNormal:       *n,
		OutputPerTrial: *u,
		Seed:           *s,
	}

	tenho.Start(o)
}
