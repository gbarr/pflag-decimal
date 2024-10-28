package pflagdecimal

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/spf13/pflag"
)

func TestDecimal(t *testing.T) {
	dec := new(Decimal)

	if dec.Type() != "Decimal" {
		t.Errorf("unexpected type name %q", dec.Type())
	}

	flag := &pflag.Flag{
		Name:      "decimal",
		Shorthand: "d",
		Usage:     "decimal value",
		Value:     dec,
	}

	pfs := pflag.NewFlagSet("test", pflag.ContinueOnError)
	pfs.AddFlag(flag)

	err := pfs.Parse([]string{"-d", "123.45"})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	exp, _ := decimal.NewFromString("123.45")
	if !dec.Equal(exp) {
		t.Errorf("unexpected value: %s", dec.Decimal)
	}

	if dec.String() != "123.45" {
		t.Errorf("unexpected value: %s", dec)
	}

	err2 := pfs.Parse([]string{"--decimal", "invalid"})
	if err2 == nil {
		t.Error("expected error, got nil")
	}
}
