package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func findEquationsToTarget(numbers []int, target int) bool {
	if len(numbers) == 0 {
		return false
	}
	var isEquationEqualToTarget bool
	// debugging: var dfs func(index int, currentValue int, path string)
	var dfs func(index int, currentValue int)
	dfs = func(index int, currentValue int) {
		// fmt.Printf("Index: %d, Current Value: %d, Path: %s\n", index, currentValue, path)
		if isEquationEqualToTarget {
			return
		}
		if index >= len(numbers) {
			if currentValue == target {
				// fmt.Printf("Target matched! Path: %s = %d\n", path, target)
				isEquationEqualToTarget = true
			}
			return
		}

		// dfs(index+1, currentValue+numbers[index],
		// fmt.Sprintf("%s + %d", path, numbers[index]))
		dfs(index+1, currentValue+numbers[index])
		if isEquationEqualToTarget {
			return
		}

		// dfs(index+1, currentValue*numbers[index],
		// 	fmt.Sprintf("%s * %d", path, numbers[index]))
		dfs(index+1, currentValue*numbers[index])
		if isEquationEqualToTarget {
			return
		}

		// part 2
		concatStr := strconv.Itoa(currentValue) + strconv.Itoa(numbers[index])
		concatVal, _ := strconv.Atoi(concatStr)
		// dfs(index+1, concatVal,
		// 	fmt.Sprintf("%s || %d", path, numbers[index]))
		dfs(index+1, concatVal)
	}

	// dfs(1, numbers[0], fmt.Sprintf("%d", numbers[0]))
	dfs(1, numbers[0])
	return isEquationEqualToTarget
}

func main() {
	start := time.Now()
	file, _ := os.Open("input.txt")
	calibrationResult := 0
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		parts := strings.Split(scanner.Text(), ":")
		target, _ := strconv.Atoi(parts[0])
		numbers := []int{}
		for _, v := range strings.Fields(parts[1]) {
			n, _ := strconv.Atoi(v)
			numbers = append(numbers, n)
		}

		if findEquationsToTarget(numbers, target) {
			calibrationResult += target
		}
	}

	duration := time.Since(start)
	fmt.Printf("Time taken: %v\n", duration)
	fmt.Printf("sum: %d\n", calibrationResult)
}
