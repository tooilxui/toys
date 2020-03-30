package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	var pi float64 = 3.14159
	value, exp := floatToDecimal(pi)
	fmt.Printf("float=%f,value=%d, exp=%d", pi, value, exp)
}

func floatToDecimal(f float64) (value int64, exp int32) {
	d := decimal.NewFromFloat(f)
	exp = d.Exponent()
	value = d.Mul(decimal.New(1, -d.Exponent())).IntPart()
	return
}
