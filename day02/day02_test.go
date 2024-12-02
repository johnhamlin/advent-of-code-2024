package day02

import (
	"fmt"
	"testing"
)

func Test_PartOne(t *testing.T) {
	got := PartOne()
	want := 407

	if got != want {
		fmt.Println("got", got, "want", want)
	}
}
