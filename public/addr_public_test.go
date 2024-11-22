package addrPublic_test

import (
	"testing"
	addrPublic "testing/public"
)

func TestAddNumber(t *testing.T) {
	result := addrPublic.AddNumbers(1, 2)
	if result != 3 {
		t.Fatalf("incorrect result: expected (%d), got (%d)", 3, result)
	}
}
