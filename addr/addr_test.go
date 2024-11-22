package addr

import (
	"testing"
)

func Test_Number(t *testing.T) {
	result := addNumbers(1, 2)
	if result != 5 {
		t.Fatalf("incorrect result: expected (%d), got (%d)", 3, result)
	}
}
