package day01

import (
	"bufio"
	"fmt"
	"github.com/johnhamlin/advent-of-code-2024/tree/main/utils"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func PartOne() {
	var left, right [1000]int

	file, err := os.Open(utils.GetInputPath())
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		numbers := strings.Fields(scanner.Text())

		left[i], err = strconv.Atoi(numbers[0])
		if err != nil {
			log.Fatal(err)
		}

		right[i], err = strconv.Atoi(numbers[1])
		if err != nil {
			log.Fatal(err)
		}
	}

	sort.Ints(left[:])
	sort.Ints(right[:])

	distance := 0
	for i := range left {
		distance += int(math.Abs(float64(left[i] - right[i])))
	}
	fmt.Println(distance)

}

func PartTwo() {
	var left [1000]int

	rightFreq := make(map[int]int)

	file, err := os.Open(utils.GetInputPath())
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		numbers := strings.Fields(scanner.Text())

		left[i], err = strconv.Atoi(numbers[0])
		if err != nil {
			log.Fatal(err)
		}

		right, err := strconv.Atoi(numbers[1])
		if err != nil {
			log.Fatal(err)
		}

		rightFreq[right]++
	}

	similarityScore := 0
	for _, n := range left {
		similarityScore += n * rightFreq[n]
	}

	fmt.Println(similarityScore)
}
