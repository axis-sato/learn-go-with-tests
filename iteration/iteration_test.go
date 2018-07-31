package iteration

import (
	"testing"
	"strings"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i :=0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func BenchmarkCompare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Compare("foo", "foo")
	}
}

func BenchmarkToUpper(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.ToUpper("foobar")
	}
}
