package iteration

import "testing"

const RepeatCount = 5
const RepeatChar = "a"

func TestRepeat(t *testing.T) {
	repeated := Repeat(RepeatChar, RepeatCount)
	expected := ""
	for i := 0; i < RepeatCount; i++ {
		expected += RepeatChar
	}

	if repeated != expected {
		t.Errorf("expected '%q' but got '%q", expected, repeated)
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat(RepeatChar, RepeatCount)
	}
}
