package pflagdecimal

import (
	"github.com/shopspring/decimal"
)

type Decimal struct {
	decimal.Decimal
}

// Implement pflag.Value interface

// Type implements pflag.Value.Type().
// Returns "Decimal".
func (*Decimal) Type() string {
	return "Decimal"
}

// String implements pflag.Value.String().
func (d *Decimal) String() string {
	return d.Decimal.String()
}

// Set implements pflag.Value.Set().
func (d *Decimal) Set(value string) error {
	parsed, err := decimal.NewFromString(value)
	if err != nil {
		return err
	}
	d.Decimal = parsed
	return nil
}
