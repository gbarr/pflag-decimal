package main

import (
	"fmt"

	pfdecimal "github.com/gbarr/pflag-decimal"
	"github.com/spf13/pflag"
)

func main() {
	var dec pfdecimal.Decimal
	pflag.VarP(&dec, "value", "v", "decimal value")
	pflag.Parse()
	fmt.Println("value:", dec.String())
}
