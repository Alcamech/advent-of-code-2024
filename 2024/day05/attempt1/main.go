package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func indexOf(slice []string, value string) int {
	for p, v := range slice {
		if v == value {
			return p
		}
	}
	return -1
}

func isBefore(arr []string, num1 string, num2 string) bool {
	for _, num := range arr {
		if num == num1 {
			fmt.Printf("arr: %v, %v < %v is true\n", arr, num1, num2)
			return true
		}
		if num == num2 {
			fmt.Printf("arr: %v, %v < %v is false\n", arr, num1, num2)
			return false
		}
	}
	return false
}

func getRuleParts(rule string) (string, string) {
	parts := strings.Split(rule, "|")

	return parts[0], parts[1]
}

func readRulesAndUpdates(filename string) ([]string, []string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("error reading file:", filename)
		return nil, nil
	}

	defer file.Close()

	var rules, updates []string
	scanner := bufio.NewScanner(file)
	target := &rules

	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println("line:", line)
		if line == "" {
			target = &updates
			continue
		}

		*target = append(*target, line)
	}

	return rules, updates
}

func findPageOrderingRules(pageNumbers []string, rules []string) []string {
	var pageOrderingRules []string

	for _, rule := range rules {
		left, right := getRuleParts(rule)

		if contains(pageNumbers, left) && contains(pageNumbers, right) {
			pageOrderingRules = append(pageOrderingRules, rule)
		}
	}

	return pageOrderingRules
}

func fixUpdateAgainstRule(pageNumbers []string, num1 string, num2 string) bool {
	idx1 := indexOf(pageNumbers, num1)
	idx2 := indexOf(pageNumbers, num2)

	if idx1 > idx2 {
		pageNumbers[idx1], pageNumbers[idx2] = pageNumbers[idx2], pageNumbers[idx1]
		return true
	}

	return false
}

func fixUpdateUsingRules(pageNumbers []string, rules []string) {
	change := true

	for change {
		change = false
		for _, rule := range rules {
			left, right := getRuleParts(rule)
			if fixUpdateAgainstRule(pageNumbers, left, right) {
				change = true
			}
		}
	}
}

func checkUpdateAgainstRules(pageNumbers []string, rules []string) bool {
	isUpdateCorrect := true

	for _, rule := range rules {
		left, right := getRuleParts(rule)
		isUpdateCorrect = isBefore(pageNumbers, left, right)
		if !isUpdateCorrect {
			return false
		}
	}

	return isUpdateCorrect
}

func main() {
	filename := "input.txt"
	rules, updates := readRulesAndUpdates(filename)
	// fmt.Println("rules:", rules)
	// fmt.Println("updates:", updates)
	resultPart1 := 0
	resultPart2 := 0

	for _, update := range updates {
		pageNumbers := strings.Split(update, ",")
		pageOrderingRules := findPageOrderingRules(pageNumbers, rules)
		fmt.Println("update:", update)
		fmt.Println("pageOrderingRules:", pageOrderingRules)
		isUpdateCorrect := checkUpdateAgainstRules(pageNumbers, pageOrderingRules)
		fmt.Printf("update %v correct ordering = %v\n", update, isUpdateCorrect)

		if isUpdateCorrect {
			middleIndex := len(pageNumbers) / 2
			fmt.Printf("update ar: %v, arr len: %v, middle: %v", pageNumbers, len(pageNumbers), middleIndex)
			middlePage := pageNumbers[middleIndex]
			middleNumber, _ := strconv.Atoi(middlePage)
			resultPart1 += middleNumber
		} else {
			fixUpdateUsingRules(pageNumbers, pageOrderingRules)
			middleIndex := len(pageNumbers) / 2
			middlePage := pageNumbers[middleIndex]
			middleNumber, _ := strconv.Atoi(middlePage)
			fmt.Println("correct update middle:", middleNumber)
			resultPart2 += middleNumber
			fmt.Println("correct update:", pageNumbers)
		}

		fmt.Println(" ")
	}

	fmt.Println("accumulated middle page numbers from correctly-ordered updates:", resultPart1)
	fmt.Println("accumulated middle page numbers from corected-ordered updates:", resultPart2)
}
