package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	partOne()

}

func partOne() {
	var safeCount int
	file, err := os.Open("./2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	//	Loop
	scanner := bufio.NewScanner(file)

LineLoop:
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

		// descending
		if nums[0] > nums[len(nums)-1] {
			for i := 1; i < len(nums); i++ {
				diff := nums[i-1] - nums[i]
				if diff <= 0 || diff > 3 {
					//	need to break outer loops too
					continue LineLoop
				}
			}
			//	ascending
		} else {
			for i := 1; i < len(nums); i++ {
				diff := nums[i] - nums[i-1]
				if diff <= 0 || diff > 3 {
					//	need to break outer loops too
					continue LineLoop
				}
			}
		}
		safeCount++
	}

	fmt.Println("Part One:", safeCount)

}
