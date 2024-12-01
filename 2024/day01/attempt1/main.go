package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

// Reads the input file and parses the numbers into two slices
func readFileAndParse(fileName string) ([]int, []int, error) {
	var leftNums, rightNums []int
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		leftNum, err1 := strconv.Atoi(parts[0])
		rightNum, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Error parsing numbers:", line)
		}

		leftNums = append(leftNums, leftNum)
		rightNums = append(rightNums, rightNum)
	}

	return leftNums, rightNums, nil
}

// sorts the slices and caclulates the accumulated difference
func calculateDifferences(leftNums, rightNums []int) int {
	sort.Ints(leftNums)
	sort.Ints(rightNums)

	result := 0
	for i := 0; i < len(leftNums) && i < len(rightNums); i++ {
		diff := absDiff(leftNums[i], rightNums[i])
		result += diff
		fmt.Printf("left: %d, right: %d, difference: %d, accumulated result: %d\n", leftNums[i], rightNums[i], diff, result)
	}

	return result
}

// calculates the accumulated similarity score
func calculateSimilartyScore(leftNums, rightNums []int) int {
	rightCount := make(map[int]int)
	for _, num := range rightNums {
		rightCount[num]++
	}

	result := 0
	for _, leftNum := range leftNums {
		count := rightCount[leftNum]
		fmt.Printf("Value %d appears %d times in rightNumbers\n", leftNum, count)

		simScore := leftNum * count
		result += simScore
	}

	return result
}

func main() {
	leftNums, rightNums, err := readFileAndParse("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	diffResult := calculateDifferences(leftNums, rightNums)
	fmt.Printf("Final accumulated diff result: %d\n", diffResult)

	simScoreResult := calculateSimilartyScore(leftNums, rightNums)
	fmt.Printf("Final accumulated simScore result: %d\n", simScoreResult)
}
