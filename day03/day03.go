package day03

import (
	"github.com/johnhamlin/advent-of-code-2024/tree/main/utils"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
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

func PartTwo() (sum int) {
	content, err := os.ReadFile(utils.GetInputPath())
	if err != nil {
		log.Fatal(err)
	}
	input := string(content)

	var opsToExclude = regexp.MustCompile(`don't\(\).*?do\(\)`)
	var validOp = regexp.MustCompile(`mul\(\d+,\d+\)`)
	var digits = regexp.MustCompile(`\d+`)

	// Removing line breaks makes the regex simpler
	oneLineInput := strings.ReplaceAll(strings.ReplaceAll(input, "\r\n", ""), "\n", "")
	sanitized := opsToExclude.ReplaceAllString(oneLineInput, ``)
	matches := validOp.FindAllString(sanitized, -1)

	for _, m := range matches {
		operands := digits.FindAllString(m, -1)
		x, _ := strconv.Atoi(operands[0])
		y, _ := strconv.Atoi(operands[1])
		sum += x * y

	}

	return sum
}
