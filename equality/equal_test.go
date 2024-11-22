package equal_test

import (
	"testing"
	equal "testing/equality"

	"github.com/google/go-cmp/cmp"
)

func TestCreatePerson(t *testing.T) {
	expected := equal.Person{
		Name: "Dennis",
		Age:  37,
	}
	result := equal.CreatePerson("Dennis", 37)
	comparer := cmp.Comparer(func(x, y equal.Person) bool { return x.Name == y.Name && x.Age == y.Age })

	if diff := cmp.Diff(expected, result, comparer); diff != "" {
		t.Error(diff)
	}
}
