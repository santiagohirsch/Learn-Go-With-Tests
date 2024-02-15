package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("A")
	expected := "AAAAA"

	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", repeated, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
