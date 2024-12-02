package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Absolute difference calculation for integers
func absDiff(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}

/*
* determine which reports are safe
* levels should be all increasing or all decreasing and any two adjacent levels
* should differ by at least one or at most three
 */
func isReportSafe(report []int) bool {
	if len(report) < 2 {
		return true
	}

	isDecreasing := true
	isIncreasing := true
	isAbsDiffBetweenOneAndThree := true

	for i := range report[:len(report)-1] {
		if report[i] >= report[i+1] {
			isDecreasing = false
		}

		if report[i] <= report[i+1] {
			isIncreasing = false
		}

		adjDiff := absDiff(report[i], report[i+1])
		if adjDiff < 1 || adjDiff > 3 {
			isAbsDiffBetweenOneAndThree = false
		}
	}

	return (isIncreasing || isDecreasing) && isAbsDiffBetweenOneAndThree
}

func problemDampenerModule(report []int) bool {
	isNewReportSafe := false
	for i := range report[:] {
		newReport := append([]int{}, report[:i]...)
		newReport = append(newReport, report[i+1:]...)
		isNewReportSafe = isReportSafe(newReport)
		fmt.Printf("newReport: %v, Safe:%v\n", newReport, isNewReportSafe)

		if isNewReportSafe {
			return true
		}
	}

	return false
}

// Reads the input file and parses the numbers into a matrix
func readFileAndParse(fileName string) ([][]int, error) {
	var matrix [][]int
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		digits := strings.Fields(line)

		var row []int
		for _, digit := range digits {
			num, err := strconv.Atoi(digit)
			if err != nil {
				fmt.Println("Error parsing number:", digit)
				continue
			}

			row = append(row, num)
		}

		matrix = append(matrix, row)
	}

	return matrix, nil
}

func main() {
	reportMatrix, err := readFileAndParse("input.txt")
	if err != nil {
		fmt.Println("error parsing file input.txt")
	}

	safeCount := 0
	for i, report := range reportMatrix {
		isSafe := isReportSafe(report)
		if !isSafe {
			isSafe = problemDampenerModule(report)
		}
		if isSafe {
			safeCount++
		}

		fmt.Printf("Report %d: %v, Safe:%v, SafeCount:%v\n", i, report, isSafe, safeCount)
	}

	fmt.Println("Safe Count:", safeCount)
}
