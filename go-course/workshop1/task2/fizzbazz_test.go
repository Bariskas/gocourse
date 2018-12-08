package task2

import "testing"

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		in int
		want string
	}{
		{1, "1"},
		{2, "1 2"},
		{3, "1 2 Fizz"},
		{5, "1 2 Fizz 4 Buzz"},
		{15, "1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 Fizz Buzz"},
	}

	for _, c := range cases {
		got := FizzBuzz(c.in)
		if got != c.want {
			t.Errorf("Reverse (%d) == %q, want %q", c.in, got, c.want)
		}
	}
}

func BenchmarkFizzBuzz50(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		FizzBuzz(50)
	}
}

func BenchmarkFizzBuzz49(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		FizzBuzz(49)
	}
}

func BenchmarkFizzBuzz20(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		FizzBuzz(20)
	}
}

func BenchmarkFizzBuzz100(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		FizzBuzz(100)
	}
}