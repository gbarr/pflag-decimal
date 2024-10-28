# pflag-decimal #
[![Go Reference](https://pkg.go.dev/badge/github.com/gbarr/pflag-decimal.svg)](https://pkg.go.dev/github.com/gbarr/pflag-decimal)
[![codecov](https://codecov.io/github/gbarr/pflag-decimal/graph/badge.svg?token=WCV5JUZHFY)](https://codecov.io/github/gbarr/pflag-decimal)
[![Go ReportCard](https://goreportcard.com/badge/gbarr/pflag-decimal)](http://goreportcard.com/report/gbarr/pflag-decimal)
[![golangci-lint](https://github.com/gbarr/pflag-decimal/actions/workflows/lint.yml/badge.svg)](https://github.com/gbarr/pflag-decimal/actions/workflows/lint.yml)

[`pflag-decimal`](https://github.com/gbarr/pflag-decimal) implements a Golang [`pflag.Value`](https://pkg.go.dev/github.com/spf13/pflag#Value) interface decimal values.
Combining [github.com/spf13/pflag](https://pkg.go.dev/github.com/spf13/pflag) with [github.com/shopspring/decimal](https://pkg.go.dev/github.com/shopspring/decimal).
This facilitiates command-line argument handling of date parameters such  `--value=123.45`.

## Documentation ##

```
go get github.com/gbarr/pflag-decimal
```

https://pkg.go.dev/github.com/gbarr/pflag-decimal

## Example ##

```go
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
```
----

## Credits and License

Copyright (c) 2024 [Graham Barr](https://github.com/gbarr).

Released under the [MIT License](https://en.wikipedia.org/wiki/MIT_License), see [LICENSE](./LICENSE).
