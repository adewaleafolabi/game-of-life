package life

import (
	"fmt"
	"testing"
)

func TestNewGameInvalidRowsAndColumns(t *testing.T) {
	game, err := NewGame(-1, 5)
	if game.Board != nil {
		t.Errorf("game board must be nil")
	}

	if err == nil {
		t.Errorf("error cannot be empty")
	}

	if err.Error() != `the numbers of rows must be greater than 1` {
		t.Errorf(`the error should be 'the numbers of rows must be greater than 1'`)
	}

	game, err = NewGame(5, 0)
	if game.Board != nil {
		t.Errorf("game board must be nil")
	}

	if err == nil {
		t.Errorf("error cannot be empty")
	}

	if err.Error() != `the numbers of columns must be greater than 1` {
		t.Errorf(`the error should be 'the numbers of columns must be greater than 1'`)
	}
}

func TestReanimateAndKillCell(t *testing.T) {
	game, _ := NewGame(4, 5)
	game.reanimateCell(0, 1)

	if game.getCell(0, 1) != true {
		t.Errorf("cell should not be dead")
	}

	game.killCell(0, 1)

	if game.getCell(0, 1) != false {
		t.Errorf("cell should be dead")
	}
}

func TestIsCellValid(t *testing.T) {
	game, _ := NewGame(4, 5)
	if !game.isValidCell(0, 4) {
		t.Errorf("cell should be valid")
	}

	if game.isValidCell(4, 9) != false {
		t.Errorf("cell should be invalid")
	}
}

func TestKillAllCells(t *testing.T) {
	game, _ := NewGame(4, 5)
	game.killAllCells()
	fmt.Println(game.getCell(3, 4))
	if game.getCell(3, 4) != false {
		t.Errorf("all cells should be dead")
	}
	if game.getCell(0, 3) != false {
		t.Errorf("all cells should be dead")
	}
	if game.getCell(1, 3) != false {
		t.Errorf("all cells should be dead")
	}
}

func TestGetAliveNeighbourCount(t *testing.T) {
	game, _ := NewGame(4, 5)
	game.killAllCells()
	if game.getAliveNeighboursCount(0, 0) != 0 {
		t.Errorf("active neighbour count must be zero")
	}

	game.reanimateCell(0, 0)
	game.reanimateCell(0, 1)
	game.reanimateCell(0, 2)
	if game.getAliveNeighboursCount(1, 1) != 3 {
		t.Errorf("active neighbour count must be three")
	}

	game.reanimateCell(1, 0)
	game.reanimateCell(1, 2)

	if game.getAliveNeighboursCount(1, 1) != 5 {
		t.Errorf("active neighbour count must be five")
	}
	game.reanimateCell(2, 0)
	game.reanimateCell(2, 1)
	game.reanimateCell(2, 2)
	game.reanimateCell(2, 3)
	game.reanimateCell(2, 4)

	if game.getAliveNeighboursCount(1, 1) != 8 {
		t.Errorf("active neighbour count must be eight")
	}
}

func TestRule1CellAliveWithLessThanTwoLiveNeighbours(t *testing.T) {

	game, _ := NewGame(4, 5)
	game.killAllCells()

	game.reanimateCell(0, 0)
	game.reanimateCell(1, 1)
	game.evolveCell(1, 1)
	if game.getCell(1, 1) == true {
		t.Errorf("cell must be dead based on rule 1")
	}

}

func TestRule2CellAliveWith2Or3LiveNeighbours(t *testing.T) {
	game, _ := NewGame(4, 5)
	game.killAllCells()

	//Reanimate starting point
	game.reanimateCell(1, 1)
	//Reanimate Neighbours
	game.reanimateCell(0, 0)
	game.reanimateCell(0, 1)
	//Evolve starting point
	game.evolveCell(1, 1)
	if game.getCell(1, 1) != true {
		t.Errorf("cell must be alive based on rule 2")
	}

	//Kill all cells
	game.killAllCells()

	//Reanimate starting point
	game.reanimateCell(1, 1)
	//Reanimate Neighbours
	game.reanimateCell(0, 0)
	game.reanimateCell(0, 1)
	game.reanimateCell(0, 2)

	//Evolve starting point
	game.evolveCell(1, 1)
	if game.getCell(1, 1) != true {
		t.Errorf("cell must be alive based on rule 2")
	}
}

func TestRule3CellAliveWithMoreThan3LiveNeighbours(t *testing.T) {

	game, _ := NewGame(4, 5)
	game.killAllCells()

	//Reanimate starting point
	game.reanimateCell(1, 1)
	//Reanimate Neighbours
	game.reanimateCell(0, 0)
	game.reanimateCell(0, 1)
	game.reanimateCell(0, 2)
	game.reanimateCell(1, 2)
	//Evolve starting point
	game.evolveCell(1, 1)
	if game.getCell(1, 1) != false {
		t.Errorf("cell must be dead based on rule 3")
	}
}

func TestRule4CellDeadWithExactly3LiveNeighbours(t *testing.T) {

	game, _ := NewGame(4, 5)
	game.killAllCells()

	//Reanimate Neighbours
	game.reanimateCell(0, 0)
	game.reanimateCell(0, 1)
	//Evolve starting point
	game.evolveCell(1, 1)
	if game.getCell(1, 1) != false {
		t.Errorf("cell must be dead based on rule 4")
	}

	game.killCell(1, 1)
	game.reanimateCell(0, 2)

	game.evolveCell(1, 1)
	if game.getCell(1, 1) != true {
		t.Errorf("cell must be alive based on rule 4")
	}
}
