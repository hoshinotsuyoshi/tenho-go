package main

import (
	"flag"
	"github.com/hoshinotsuyoshi/tenho-go/src/tenho"
)

func main() {
	flgNoKokushi := flag.Bool("no-kokushi", false, "Not apply kokushi")
	flgNoChitoitsu := flag.Bool("no-chitoitsu", false, "Not apply chitoitsu")
	flgNoNormal := flag.Bool("no-normal", false, "Not apply normal")

	flag.Parse()

	tenho.Start(*flgNoKokushi, *flgNoChitoitsu, *flgNoNormal)
}
