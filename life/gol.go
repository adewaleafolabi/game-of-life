package life

import (
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
	Board [][]bool
}

func NewGame(rows, columns int) (Game, error) {
	game := Game{}
	if rows <= 0 {
		return game, fmt.Errorf("the numbers of rows must be greater than 1")
	}
	if columns <= 0 {
		return game, fmt.Errorf("the numbers of columns must be greater than 1")
	}
	board := make([][]bool, rows)

	for i := 0; i < rows; i++ {
		board[i] = make([]bool, columns)
		for j := 0; j < columns; j++ {
			board[i][j] = deadOrAlive()
		}

	}
	game.Board = board
	return game, nil
}

//Randomly determine the state of the cell
func deadOrAlive() bool {
	rand.Seed(time.Now().UTC().UnixNano())
	if rand.Intn(2) > 0 {
		return true
	}
	return false
}

//Return cell using the row and column combination
func (g *Game) getCell(row int, column int) bool {
	return g.Board[row][column]
}

//Check cell is within the boundaries of the array
func (g *Game) isValidCell(row int, column int) bool {
	if row >= 0 && row <= len(g.Board)-1 && column >= 0 && column <= len(g.Board[0])-1 {
		return true
	}
	return false
}

//Call reanimation jutsu. Make cell alive again
func (g *Game) reanimateCell(row int, column int) {
	if g.isValidCell(row, column) {
		g.Board[row][column] = true
	}
}

//Kill cell, set value to false
func (g *Game) killCell(row int, column int) {
	if g.isValidCell(row, column) {
		g.Board[row][column] = false
	}
}

//Kill all the board cells
func (g *Game) killAllCells() {
	for i := 0; i < len(g.Board); i++ {
		for j := 0; j < len(g.Board[0]); j++ {
			g.Board[i][j] = false
		}
	}
}

//Get the count of the alive neighbours.
// Neighbour cells are checked if they are valid first
func (g *Game) getAliveNeighboursCount(row int, column int) int {

	count := 0
	//Middle Right
	if g.isValidCell(row, column+1) && g.getCell(row, column+1) {
		count++
	}

	//Middle Left
	if g.isValidCell(row, column-1) && g.getCell(row, column-1) {
		count++
	}

	//Upper Right
	if g.isValidCell(row-1, column+1) && g.getCell(row-1, column+1) {
		count++
	}
	//Upper Middle
	if g.isValidCell(row-1, column) && g.getCell(row-1, column) {
		count++
	}
	//Upper Left
	if g.isValidCell(row-1, column-1) && g.getCell(row-1, column-1) {
		count++
	}

	//Lower Right
	if g.isValidCell(row+1, column+1) && g.getCell(row+1, column+1) {
		count++
	}
	//Lower Middle
	if g.isValidCell(row+1, column) && g.getCell(row+1, column) {
		count++
	}
	//Lower Left
	if g.isValidCell(row+1, column-1) && g.getCell(row+1, column-1) {
		count++
	}

	return count
}

//Changes the cell state based on the rules of the game
// the return value does not represent the state of action.
// the return value only helps to tell if the game should continue at
//the end of the iteration.
//it basically informs the game if there is at least one live cell
//this was done to avoid having to loop through the list the second time to check
//if there's a living cell.
func (g *Game) evolveCell(row, column int) bool {
	continueGame:=false
	cellIsAlive := g.getCell(row, column)
	aliveNeighboursCount := g.getAliveNeighboursCount(row, column)
	fmt.Println(fmt.Sprintf("cell[%d,%d] %v AliveNeighbourCount:%d", row, column, cellIsAlive, aliveNeighboursCount))
	if cellIsAlive && aliveNeighboursCount < 2 {
		g.killCell(row, column)
	}
	if cellIsAlive && (aliveNeighboursCount == 2 || aliveNeighboursCount == 3) {
		g.reanimateCell(row, column)
		continueGame = true
	}

	if cellIsAlive && aliveNeighboursCount > 3 {
		g.killCell(row, column)
	}

	if (!cellIsAlive) && aliveNeighboursCount == 3 {
		g.reanimateCell(row, column)
		continueGame = true
	}
	return continueGame
}

func (g *Game) RunSimulation(){
		for {
			i, j := 0, 0
			continueGame := false
			for i = 0; i < len(g.Board)-1; i++ {
				for j = 0; j < len(g.Board[0])-1; j++ {
					continueGame = g.evolveCell(i, j)
				}
			}
			if continueGame {
				i, j = 0, 0
			} else {
				fmt.Println("simulation complete. No more moves available !!!")
				break
			}
		}
}