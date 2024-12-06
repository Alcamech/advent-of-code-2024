package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInputToMatrix(filename string) [][]rune {
	var input [][]rune
	file, err := os.Open(filename)
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

func findGuardIndex(matrix [][]rune, guard rune) (int, int) {
	for i, row := range matrix {
		for j, val := range row {
			if val == guard {
				return i, j
			}
		}
	}

	return -1, -1
}

func determineDirectionAfterObstacle(direction string) string {
	if direction == "up" {
		return "right"
	}
	if direction == "right" {
		return "down"
	}
	if direction == "down" {
		return "left"
	}

	if direction == "left" {
		return "up"
	}

	return direction
}

func doMovementGivenDirection(direction string, x, y int) (int, int) {
	if direction == "up" {
		x--
	}
	if direction == "right" {
		y++
	}
	if direction == "down" {
		x++
	}
	if direction == "left" {
		y--
	}
	return x, y
}

type Position struct {
	X, Y int
}

func simulateGuardMovement(matrix [][]rune, startX, startY int) int {
	visited := make(map[Position]bool)
	direction := "up"
	x, y := startX, startY
	rows := len(matrix)
	cols := len(matrix[0])

	for {
		if x < 0 || x >= rows || y < 0 || y >= cols {
			fmt.Println("Out of bounds! Guard left the area!.")
			break
		}

		currentPos := matrix[x][y]
		position := Position{X: x, Y: y}
		visited[position] = true

		var nextPos rune

		switch direction {
		case "up":
			if x-1 < 0 {
				fmt.Println("went out of bounds!")
				break
			}
			nextPos = matrix[x-1][y]
		case "down":
			if x+1 >= rows {
				fmt.Println("went out of bounds!")
				break
			}
			nextPos = matrix[x+1][y]
		case "left":
			if y-1 < 0 {
				fmt.Println("went out of bounds!")
				break
			}
			nextPos = matrix[x][y-1]
		case "right":
			if y+1 >= cols {
				fmt.Println("went out of bounds!")
				break
			}
			nextPos = matrix[x][y+1]
		default:
			panic("unrecognized direction")
		}

		fmt.Printf("Value at (%d, %d): %s, ahead: %s, direction %s\n",
			x, y, string(currentPos), string(nextPos), direction)

		if nextPos == '#' {
			fmt.Println("found obstacle!")
			direction = determineDirectionAfterObstacle(direction)
			fmt.Println("turning !", direction)
		}

		x, y = doMovementGivenDirection(direction, x, y)
	}

	fmt.Println("guard left!")

	return len(visited)
}

type Location struct {
	Pos       Position
	Direction string
}

func doesGuardLoop(matrix [][]rune, startX, startY int) bool {
	visited := make(map[Location]bool)
	x, y := startX, startY
	direction := "up"
	rows := len(matrix)
	cols := len(matrix[0])

	for {
		loc := Location{
			Pos:       Position{X: x, Y: y},
			Direction: direction,
		}

		if visited[loc] {
			return true
		}

		visited[loc] = true

		nextX, nextY := x, y
		switch direction {
		case "up":
			nextX--
		case "down":
			nextX++
		case "left":
			nextY--
		case "right":
			nextY++
		}

		if nextX < 0 || nextX >= rows || nextY < 0 || nextY >= cols {
			return false
		}

		if matrix[nextX][nextY] == '#' {
			direction = determineDirectionAfterObstacle(direction)
		} else {
			x, y = nextX, nextY
		}
	}
}

func countLoopCausingPositions(matrix [][]rune, guardX, guardY int) int {
	rows := len(matrix)
	cols := len(matrix[0])
	loopCausingPositions := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '#' || (i == guardX && j == guardY) {
				continue
			}

			matrix[i][j] = '#'

			if doesGuardLoop(matrix, guardX, guardY) {
				loopCausingPositions++
			}

			matrix[i][j] = '.'
		}
	}

	return loopCausingPositions
}

func main() {
	filename := "input.txt"
	matrix := readInputToMatrix(filename)

	fmt.Println("matrix:")
	for _, row := range matrix {
		fmt.Println(string(row))
	}

	row, col := findGuardIndex(matrix, '^')
	fmt.Printf("found guard index at %v, %v:\n", row, col)
	fmt.Println("guard:", string(matrix[row][col]))

	// Part 1: Count distinct positions
	distinctPosCount := simulateGuardMovement(matrix, row, col)
	fmt.Println("guard distinct positions:", distinctPosCount)

	// Part 2: Count loop-causing positions
	loopCausingPositions := countLoopCausingPositions(matrix, row, col)
	fmt.Printf("Number of loop-causing positions: %d\n", loopCausingPositions)
}
