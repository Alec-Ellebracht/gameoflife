package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	// validate that we have 2 inputs
	numInputs := len(os.Args)
	if numInputs != 3 {
		log.Fatal(fmt.Errorf("Must provide 2 integers"))
	}

	// get args from the command line
	strCols, strRows := os.Args[1], os.Args[2]

	// covert string args to nums
	cols, err := strconv.Atoi(strCols)
	if err != nil {
		fmt.Printf("%T, %v\n", cols, cols)
		log.Fatal(err)
	}

	rows, err := strconv.Atoi(strRows)
	if err != nil {
		fmt.Printf("%T, %v\n", rows, rows)
		log.Fatal(err)
	}

	// cols, rows := 145, 61

	grid := makeGrid(cols, rows)

	for {
		printGeneration(grid)
		grid = nextGeneration(grid)
	}
}

func nextGeneration(grid [][]int) [][]int {

	next := makeDupe(grid)

	for i, row := range grid {
		for j, col := range row {

			state := col

			neighbors := countNeighbors(grid, i, j)

			// rules of life
			if state == 1 && (neighbors == 2 || neighbors == 3) {

				next[i][j] = 1

			} else if state == 0 && neighbors == 3 {

				next[i][j] = 1

			} else {

				next[i][j] = 0

			}
		}
	}

	return next
}

// make a dupe of a [][]int so that we dont
// accidentally mutate values
func makeDupe(grid [][]int) [][]int {

	duplicate := make([][]int, len(grid))

	for i := range grid {

		duplicate[i] = make([]int, len(grid[i]))
		copy(duplicate[i], grid[i])
	}

	return duplicate
}

func isEdge(i, j int, grid [][]int, row []int) bool {
	return (i == 0 || i == len(grid)-1 || j == 0 || j == len(row)-1)
}

// returns the number of neighbors when give coords
func countNeighbors(grid [][]int, x, y int) int {

	rows := len(grid)
	cols := len(grid[0])

	var sum int

	for i := -1; i < 2; i++ {

		for j := -1; j < 2; j++ {

			row := (x + i + rows) % rows
			col := (y + j + cols) % cols
			sum += grid[row][col]
		}
	}

	// dont count yourself (sum - me)
	return (sum - grid[x][y])
}

// returns a grid with the given num rows & cols
func makeGrid(cols int, rows int) [][]int {

	grid := make([][]int, rows)

	for i := 0; i < rows; i++ {

		grid[i] = make([]int, cols)

		for j := 0; j < cols; j++ {

			grid[i][j] = rand.Intn(2)
		}
	}

	return grid
}

// prints the provided grid to the terminal
// replacing 1's with * and 0's with a space
func printGeneration(grid [][]int) {

	var gen string

	for _, row := range grid {
		for _, col := range row {

			switch col {
			case 1:
				gen += fmt.Sprintf("%v", "???")
			case 0:
				gen += fmt.Sprintf("%v", "???")
			}
		}
		gen += fmt.Sprintf("\n")
	}

	fmt.Print("\x0c")
	fmt.Println(gen)

	time.Sleep(time.Second / 30)
}
