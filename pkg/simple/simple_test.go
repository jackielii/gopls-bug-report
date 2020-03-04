package simple

import (
	"testing"
)

func TestSimple(t *testing.T) {
	x := 42
	y := 105
	g := Gcd(x, y)
	expected := 21
	if expected != g {
		t.Fatalf("expecting %d got %d", expected, g)
	}
}
