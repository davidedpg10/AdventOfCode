package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var StringIntMap = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	partOne()
	partTwo()
}

func partOne() {
	fileImport, err := os.ReadFile("input.txt")
	if err != nil {
		panic(fmt.Errorf("Error reading file: %w", err))
	}
	calibrationDocument := string(fileImport)

	var calibrationValues []int
	for _, v := range strings.Split(calibrationDocument, "\n") {
		intSlice := ParseNumbersFromString(v)
		if i := FirstLastIntsFromSlice(intSlice); i != -1 {
			calibrationValues = append(calibrationValues, i)
		}
	}
	fmt.Println("Part 1: Total value: ", SumIntSlice(calibrationValues))
}

func partTwo() {
	fileImport, err := os.ReadFile("input.txt")
	if err != nil {
		panic(fmt.Errorf("Error reading file: %w", err))
	}
	calibrationDocument := string(fileImport)

	var calibrationValues []int
	for _, v := range strings.Split(calibrationDocument, "\n") {
		intSlice := ParseNumberFromStringByWords(v)
		if i := FirstLastIntsFromSlice(intSlice); i != -1 {
			calibrationValues = append(calibrationValues, i)
		}
	}
	fmt.Println("Part 2: Total value: ", SumIntSlice(calibrationValues))
}

func ParseNumbersFromString(str string) []int {
	var ints []int
	for _, v := range str {
		if i, err := strconv.Atoi(string(v)); err == nil {
			ints = append(ints, i)
		}
	}
	return ints
}

func ParseNumberFromStringByWords(str string) []int {
	// Map will be populated as follows:
	// Field 1: Index where number was found
	// Field 2: Number found as word
	numMap := make(map[int]int)
	var ints []int

	// identifies spelled out numbers, records the index in which they were found (+1 to avoid endless loops),
	// recursively does this by starting the next index operation at the lastIndex location (to find repeat spelled out numbers)
	// Exits loop at the first failure to find spelled out number
	for numberName, num := range StringIntMap {
		lastIndex := 0
		for lastIndex >= 0 {
			idx := strings.Index(strings.ToLower(str[lastIndex:]), strings.ToLower(numberName))
			if idx != -1 {
				numMap[idx+lastIndex] = num
				lastIndex = idx + lastIndex + 1
			} else {
				lastIndex = idx
			}
		}
	}
	// Identifying digit numbers
	for idx, v := range str {
		if i, err := strconv.Atoi(string(v)); err == nil {
			numMap[idx] = i
		}
	}

	// Separating and sorting the indexes
	var indexes []int
	for i := range numMap {
		indexes = append(indexes, i)
	}
	sort.Ints(indexes)

	// Aggregating for return
	for _, v := range indexes {
		ints = append(ints, numMap[v])
	}
	return ints
}

func FirstLastIntsFromSlice(intSlice []int) int {
	if len(intSlice) > 0 {
		v1 := intSlice[0]
		v2 := intSlice[len(intSlice)-1]
		computed, err := strconv.Atoi(fmt.Sprintf("%d%d", v1, v2))
		if err != nil {
			panic(err)
		}
		return computed
	}
	return -1
}

func SumIntSlice(intSlice []int) int {
	sum := 0
	for _, v := range intSlice {
		sum = sum + v
	}
	return sum
}
