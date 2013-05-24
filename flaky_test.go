package flaky

import "testing"

func TestFloatToLisp(t *testing.T) {
	const in, out = 56.3, "(/,5630,1000)"
	flake := FloatToFlake(56.3)
	x := flake.ShowLisp()
	if x != out {
		t.Errorf("Float %v to lisp %s, expected %s", in, x, out)
	}
}

func TestFloatToFlakyToFloat(t *testing.T) {
	const in, out = 14.56, 1.456
	flake := FloatToFlake(14.56)
	x := flake.Float64()
	if x != out {
		t.Errorf("Float %v to falky, back to float %s, expected %s", in, x, out)
	}
}
