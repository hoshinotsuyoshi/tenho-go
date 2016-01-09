package main

import (
	"flag"
	"github.com/hoshinotsuyoshi/tenho-go/src/tenho"
)

func main() {
	o := tenho.OptionStruct{
		NoKokushi:   flag.Bool("no-kokushi", false, "Not apply kokushi"),
		NoChitoitsu: flag.Bool("no-chitoitsu", false, "Not apply chitoitsu"),
		NoNormal:    flag.Bool("no-normal", false, "Not apply normal"),
	}

	flag.Parse()

	tenho.Start(o)
}
