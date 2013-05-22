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
