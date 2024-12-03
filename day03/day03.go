package day03

import (
	"github.com/johnhamlin/advent-of-code-2024/tree/main/utils"
	"log"
	"os"
	"regexp"
	"strconv"
)

func PartOne() (sum int) {
	content, err := os.ReadFile(utils.GetInputPath())
	if err != nil {
		log.Fatal(err)
	}
	input := string(content)

	var validOp = regexp.MustCompile(`mul\(\d+,\d+\)`)
	var digits = regexp.MustCompile(`\d+`)

	matches := validOp.FindAllString(input, -1)

	for _, m := range matches {
		operands := digits.FindAllString(m, -1)
		x, _ := strconv.Atoi(operands[0])
		y, _ := strconv.Atoi(operands[1])
		sum += x * y

	}

	return sum
}
