package day02

import (
	"bufio"
	"fmt"
	"github.com/johnhamlin/advent-of-code-2024/tree/main/utils"
	"log"
	"os"
	"strconv"
	"strings"
)

func PartOne() (safeCount int) {
	file, err := os.Open(utils.GetInputPath())
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//	Loop
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//	Parse the data into fields
		fields := strings.Fields(scanner.Text())
		if len(fields) <= 1 {
			continue
		}

		// Parse fields to ints
		nums := make([]int, len(fields))
		for i, f := range fields {
			num, err := strconv.Atoi(f)
			if err != nil {
				fmt.Println(err)
				return
			}
			nums[i] = num
		}

		differ := getDiffer(nums[0], nums[len(nums)-1])

		if isSafe(nums, differ) {
			safeCount++
		}
	}

	return safeCount

}

func PartTwo() (safeCount int) {
	file, err := os.Open(utils.GetInputPath())
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//	Loop
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		//	Parse the data into fields
		fields := strings.Fields(scanner.Text())
		if len(fields) <= 1 {
			continue
		}

		// Parse fields to ints
		nums := make([]int, len(fields))
		for i, f := range fields {
			num, err := strconv.Atoi(f)
			if err != nil {
				fmt.Println(err)
				return
			}
			nums[i] = num
		}

		differ := getDiffer(nums[0], nums[len(nums)-1])

		// If it's safe as is, add it and move on
		if isSafe(nums, differ) {
			safeCount++
			continue
		}

		//	Try removing each number one at a time
		for i := 0; i < len(nums); i++ {
			newNums := make([]int, len(nums)-1)
			copy(newNums, nums[:i])
			copy(newNums[i:], nums[i+1:])

			if isSafe(newNums, differ) {
				safeCount++
				break // Found a working solution, no need to try more
			}
		}

	}

	return safeCount

}

// Returns a callback function that returns the distance between two numbers based on whether
// the series is descending or ascending.
func getDiffer(first, last int) func(i, j int) int {
	// Descending list
	if first > last {
		return func(i, j int) int {
			return i - j
		}
	}
	// Ascending list
	return func(i, j int) int {
		return j - i
	}
}

func isSafe(nums []int, differ func(x, y int) int) bool {
	for i := 1; i < len(nums); i++ {
		diff := differ(nums[i-1], nums[i])
		if diff <= 0 || diff > 3 {
			return false
		}
	}
	return true
}
