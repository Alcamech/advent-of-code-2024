package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readFileAndParse(fileName string) []string {
	var input []string
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return input
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	return input
}

func cleanInputArray(inputArr []string) []string {
	reInstructions := regexp.MustCompile(`do\(\)|don't\(\)|mul\(\d{1,3},\d{1,3}\)`)

	isEnabled := true
	var cleaned []string

	for _, str := range inputArr {
		instructions := reInstructions.FindAllString(str, -1)
		fmt.Printf("Found instructions: %v\n", instructions)

		for _, instruction := range instructions {
			if instruction == "do()" {
				fmt.Println("Found do(): Enabling mul instructions.")
				isEnabled = true
			} else if instruction == "don't()" {
				fmt.Println("Found don't(): Disabling mul instructions.")
				isEnabled = false
			} else if strings.HasPrefix(instruction, "mul(") && isEnabled {
				fmt.Printf("Valid mul instruction: %s\n", instruction)
				cleaned = append(cleaned, instruction)
			}
		}
	}

	return cleaned
}

func processCleanedInput(inputArr []string) []int {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	var processed []int

	for _, mul := range inputArr {
		fmt.Printf("Processing mul: %v\n", mul)
		match := re.FindStringSubmatch(mul)
		if len(match) < 3 {
			fmt.Printf("Invalid mul: %s\n", mul)
			continue
		}

		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])

		product := x * y
		processed = append(processed, product)
	}

	return processed
}

func sumArray(inputArr []int) int {
	sum := 0
	for _, num := range inputArr {
		sum += num
	}
	return sum
}

func main() {
	input := readFileAndParse("input.txt")
	fmt.Printf("Read input: %v\n", input)

	cleanedInput := cleanInputArray(input)
	fmt.Printf("Cleaned input: %v\n", cleanedInput)

	muls := processCleanedInput(cleanedInput)
	fmt.Println("Processed muls:", muls)

	total := sumArray(muls)
	fmt.Println("Summed mul inputs:", total)
}
