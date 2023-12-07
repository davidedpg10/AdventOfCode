package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Calibration document begin
	fileImport, err := os.ReadFile("input.txt")
	calibrationDocument := string(fileImport)
	//fmt.Println(calibrationDocument)
	if err != nil {
		panic(fmt.Errorf("Error reading file: %w", err))
	}

	// Calibration document end
	var calibrationValues []int
	var sum int = 0
	for _, v := range strings.Split(calibrationDocument, "\n") {
		fmt.Println("Line: ", v)
		intArray := GetInts(v)
		fmt.Println("Ints returned: ", intArray)
		if len(intArray) > 0 {
			v1 := intArray[0]
			v2 := intArray[len(intArray)-1]
			computed, err := strconv.Atoi(fmt.Sprintf("%d%d", v1, v2))
			if err != nil {
				panic(err)
			}
			calibrationValues = append(calibrationValues, computed)
		}
	}
	fmt.Printf("%#v\n", calibrationValues)
	for _, v := range calibrationValues {
		sum = sum + v
	}
	fmt.Println("Total value: ", sum)
}

func GetInts(str string) []int {
	var ints []int
	for _, v := range str {
		if i, err := strconv.Atoi(string(v)); err == nil {
			ints = append(ints, i)
		}
	}
	return ints
}
