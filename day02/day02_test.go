package day02

import (
	"fmt"
	"testing"
)

func TestPartOne(t *testing.T) {
	got := PartOne()
	want := 407

	if got != want {
		fmt.Println("got", got, "want", want)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PartTwo()
	}
}
