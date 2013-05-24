package flaky

import (
	"math"
	"strconv"
)

type Flake struct {
	Operation int //with terms defines type: a fraction, product?
	Terms     []Flake
}

func (f *Flake) ShowLisp() string {
	if len(f.Terms) == 0 {
		return strconv.Itoa(f.Operation)
	} else {
		// then we must recall all the underlying types
		header := string('(')
		switch f.Operation {
		case 0:
			header += string("*")
		case 1:
			header += string("/")
		case 2:
			header += string("+")
		}

		for i := range f.Terms {
			header += "," + f.Terms[i].ShowLisp()
		}
		header += ")"
		return header
	}
}

func FloatToFlake(x float64) Flake {
	order := int(math.Floor(math.Log10(x)))
	nominator := Flake{Operation: int(math.Floor(x / (math.Pow10(order - 3))))}
	denominator := Flake{Operation: 1000}

	out := Flake{1, []Flake{nominator, denominator}}
	return out
}

func (f Flake) String() string {
	return f.ShowLisp()
}

func (f Flake) Float64() float64 {
	if len(f.Terms) == 0 {
		return float64(f.Operation)
	} else {
		var value float64
		switch f.Operation {
		case 0:
			value = 1.
			for i := range f.Terms {
				value = value * f.Terms[i].Float64()
			}
		case 1:
			value = f.Terms[0].Float64() / f.Terms[1].Float64()
		case 2:
			value = 0.
			for i := range f.Terms {
				value = value + f.Terms[i].Float64()
			}
		}
		return value
	}
}
