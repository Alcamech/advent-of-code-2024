package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInputVertically(fileName string) []string {
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
		for len(input) < len(line) {
			input = append(input, "")
		}

		for i, char := range line {
			input[i] += string(char)
		}
	}

	return input
}

func readInputHorizontally(fileName string) []string {
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

func readInputToMatrix(fileName string) [][]rune {
	var input [][]rune
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return input
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		input = append(input, row)
	}

	return input
}

func getDiagonals(matrix [][]rune) ([]string, []string) {
	rows := len(matrix)
	fmt.Println("matrix length:", rows)
	cols := len(matrix[0])

	mainDiagonals := make([][]rune, rows+cols-1)
	antiDiagonals := make([][]rune, rows+cols-1)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			mainIndex := i - j + cols - 1
			mainDiagonals[mainIndex] = append(mainDiagonals[mainIndex], matrix[i][j])

			antiIndex := i + j
			antiDiagonals[antiIndex] = append(antiDiagonals[antiIndex], matrix[i][j])
		}
	}

	mainDiagonalStrings := make([]string, len(mainDiagonals))
	antiDiagonalStrings := make([]string, len(antiDiagonals))

	for i, diag := range mainDiagonals {
		mainDiagonalStrings[i] = string(diag) // Convert []rune to string
	}
	for i, diag := range antiDiagonals {
		antiDiagonalStrings[i] = string(diag) // Convert []rune to string
	}

	return mainDiagonalStrings, antiDiagonalStrings
}

func countHorizontalXmas(input []string, sequence string) int {
	sequenceLen := len(sequence)
	totalCount := 0

	for lineIndex, line := range input {
		fmt.Println("line:", line)
		hXmasCount := 0
		for i := 0; i < len(line)-sequenceLen; i++ {
			window := line[i : i+sequenceLen]
			if window == sequence {
				fmt.Println("found match!")
				hXmasCount++
			}
			fmt.Printf("window: %v:, lineIndex: %v, hXmasCount: %v\n", window, lineIndex, hXmasCount)
		}
		totalCount += hXmasCount
	}

	return totalCount
}

func printRuneMatrix(runeMatrix [][]rune) {
	for _, row := range runeMatrix {
		for _, r := range row {
			fmt.Printf("%c", r)
		}
		fmt.Println()
	}
}

func printRuneArray(runeArray []rune) {
	for _, row := range runeArray {
		fmt.Printf("%c", row)
	}
}

func main() {
	vInput := readInputVertically("input.txt")
	hInput := readInputHorizontally("input.txt")
	mInput := readInputToMatrix("input.txt")
	dInput, adInput := getDiagonals(mInput)
	// printRuneMatrix(mInput)
	// fmt.Println("Main Diagonals as Strings:")
	// for i, diag := range dInput{
	// 	fmt.Printf("Diagonal %d: %s\n", i, diag)
	// }
	//
	// fmt.Println("\nAnti-Diagonals as Strings:")
	// for i, diag := range adInput{
	// 	fmt.Printf("Anti-Diagonal %d: %s\n", i, diag)
	// }
	horizontalXmasCount := countHorizontalXmas(hInput, "XMAS")
	horizontalBackwardsXmasCount := countHorizontalXmas(hInput, "SAMX")
	verticalXmasCount := countHorizontalXmas(vInput, "XMAS")
	verticalBackwardsXmasCount := countHorizontalXmas(vInput, "SAMX")
	diagonalXmasCount := countHorizontalXmas(dInput, "XMAS")
	diagonalBackwardsXmasCount := countHorizontalXmas(dInput, "SAMX")
	antiDiagonalXmasCount := countHorizontalXmas(adInput, "XMAS")
	antiDiagonalBackwardsXmasCount := countHorizontalXmas(adInput, "SAMX")
	fmt.Printf("XMAS Horizontal Found: %v Times\n", horizontalXmasCount)
	fmt.Printf("XMAS Backwards Horizontal Found: %v Times\n", horizontalBackwardsXmasCount)
	fmt.Printf("XMAS Vertical Found: %v Times\n", verticalXmasCount)
	fmt.Printf("XMAS Backwards Vertical Found: %v Times\n", verticalBackwardsXmasCount)
	fmt.Printf("XMAS Diagonal Found: %v Times\n", diagonalXmasCount)
	fmt.Printf("XMAS Backwards Diagonal Found: %v Times\n", diagonalBackwardsXmasCount)
	fmt.Printf("XMAS Anti-Diagonal Found: %v Times\n", antiDiagonalXmasCount)
	fmt.Printf("XMAS Backwards Anti-Diagonal Found: %v Times\n", antiDiagonalBackwardsXmasCount)
}
