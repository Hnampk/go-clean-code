package chapter2

// GetThem: bad - meaningless naming
//	- what kind of things are in theList?
//	- what is the significance of the zeroth subscript of an item in theList?
//	- what is the significance of the value 4?
//	- how would I use the list being returned?
func GetThem() [][]int {
	theList := [][]int{}
	list1 := [][]int{}

	for _, x := range theList {
		if x[0] == 4 {
			list1 = append(list1, x)
		}
	}
	return list1
}

// GetFlaggedCells: better - just rename the variables and function name
func GetFlaggedCells() [][]int {
	STATUS_VALUE := 0
	FLAGGED := 4
	gameBoard := [][]int{}
	flaggedCells := [][]int{}

	for _, cell := range gameBoard {
		if cell[STATUS_VALUE] == FLAGGED {
			flaggedCells = append(flaggedCells, cell)
		}
	}

	return flaggedCells
}

// ****Add a simple struct for cells. Keep the same algorithm****

type cell struct {
	values []int
}

const (
	STATUS_VALUE = 0
	FLAGGED      = 4
)

func (c *cell) isFlagged() bool {
	return c.values[STATUS_VALUE] == FLAGGED
}

func BetterGetFlaggedCells() []cell {
	gameBoard := []cell{}
	flaggedCells := []cell{}

	for _, cell := range gameBoard {
		if cell.isFlagged() {
			flaggedCells = append(flaggedCells, cell)
		}
	}

	return flaggedCells
}
