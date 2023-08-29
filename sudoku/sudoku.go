package sudoku

type Sudoku []Sectors

type Sectors struct {
	Tiles  []Tile
	Line   int
	Column int
}

type Tile struct {
	Value  int
	Line   int
	Column int
}

func GenerateSudoku() Sudoku {
	var sudoku Sudoku

	return sudoku
}
