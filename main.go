package main

import (
	"fmt"
	"math/rand"
	"time"
)

const maxRows = 5
const maxColumns = 5

var grid [maxRows][maxColumns]bool

//Randomly determine the state of the cell
func deadOrAlive() bool {
	rand.Seed(time.Now().UTC().UnixNano())
	if rand.Intn(2) > 0 {
		return true
	}
	return false
}

//Call reanimation jutsu. Make cell alive again
func reanimate(row int, column int) {
	grid[row][column] = true
}

//Kill cell, set value to false
func kill(row int, column int) {
	grid[row][column] = false
}

//Return cell using the row and column combination
func getCell(row int, column int) bool {
	return grid[row][column]
}

//Check cell is within the boundaries of the array
func isValidCell(row int, column int) bool {
	if row >= 0 && row <= maxRows-1 && column >= 0 && column <= maxColumns-1 {
		return true
	}
	return false
}

//Get the count of the alive neighbours.
// Neighbour cells are checked if they are valid first
func getAliveNeighboursCount(row int, column int) int {

	count := 0
	//Middle Right
	if isValidCell(row, column+1) && getCell(row, column+1) {
		count++
	}

	//Middle Left
	if isValidCell(row, column-1) && getCell(row, column-1) {
		count++
	}

	//Upper Right
	if isValidCell(row-1, column+1) && getCell(row-1, column+1) {
		count++
	}
	//Upper Middle
	if isValidCell(row-1, column) && getCell(row-1, column) {
		count++
	}
	//Upper Left
	if isValidCell(row-1, column-1) && getCell(row-1, column-1) {
		count++
	}

	//Lower Right
	if isValidCell(row+1, column+1) && getCell(row+1, column+1) {
		count++
	}
	//Lower Middle
	if isValidCell(row+1, column) && getCell(row+1, column) {
		count++
	}
	//Lower Left
	if isValidCell(row+1, column-1) && getCell(row+1, column-1) {
		count++
	}

	return count
}

func main() {

	//Initialize the grid
	for i := 0; i < maxRows; i++ {
		for j := 0; j < maxColumns; j++ {
			grid[i][j] = deadOrAlive()
		}
		fmt.Println(grid[i])
	}

	for {
		i, j := 0, 0
		continueGame := false
		for i = 0; i < maxRows; i++ {
			for j = 0; j < maxColumns; j++ {
				cellIsAlive := getCell(i, j)
				aliveNeighboursCount := getAliveNeighboursCount(i, j)
				fmt.Println(fmt.Sprintf("cell[%d,%d] %v AliveNeighbourCount:%d", i, j, cellIsAlive, aliveNeighboursCount))
				if cellIsAlive && aliveNeighboursCount < 2 {
					kill(i, j)
				}
				if cellIsAlive && (aliveNeighboursCount == 2 || aliveNeighboursCount == 3) {
					reanimate(i, j)
					continueGame = true
				}

				if cellIsAlive && aliveNeighboursCount > 3 {
					kill(i, j)
				}

				if (!cellIsAlive) && aliveNeighboursCount == 3 {
					reanimate(i, j)
					continueGame = true
				}
			}
		}
		if continueGame {
			i, j = 0, 0
		} else {
			break
		}
	}

}
