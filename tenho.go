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

	flag.Parse()

	o := tenho.OptionStruct{
		NoChitoitsu: *c,
		NoKokushi:   *k,
		NoNormal:    *n,
	}

	tenho.Start(o)
}
