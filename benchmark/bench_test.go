package benchmark

import (
	"fmt"
	"testing"
)

var blackhole int

func TestFileLen(t *testing.T) {
	result, err := FileLen("testdata/text.txt", 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 1542 {
		t.Fatalf("incorrect result: expected (%d), got (%d)", 1542, result)
	}
}

func BenchmarkFileLen(b *testing.B) {
	for _, v := range []int{1, 10, 100, 1000, 10000, 100000} {
		b.Run(fmt.Sprintf("FileLen-%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result, err := FileLen("testdata/text.txt", v)
				if err != nil {
					b.Fatalf("unexpected error: %v", err)
				}
				blackhole = result
			}
		})
	}
}

//Keep in mind that when a buffer is bigger than the file the function slows down due to some extra allocations
// BenchmarkFileLen/FileLen-1-8                 214           6616914 ns/op             697 B/op          4 allocs/op
// BenchmarkFileLen/FileLen-10-8               2066            909531 ns/op             712 B/op          4 allocs/op
// BenchmarkFileLen/FileLen-100-8              9265            212575 ns/op             808 B/op          4 allocs/op
// BenchmarkFileLen/FileLen-1000-8            14826            121979 ns/op            1720 B/op          4 allocs/op
// BenchmarkFileLen/FileLen-10000-8            7730            146785 ns/op           10936 B/op          4 allocs/op
// BenchmarkFileLen/FileLen-100000-8          12384            155962 ns/op          107192 B/op          4 allocs/op
