package pkg

import "testing"

func TestOne(t *testing.T) {
	if One() != 1 {
		t.Error("expecting 1")
	}
}
