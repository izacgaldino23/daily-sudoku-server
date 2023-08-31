package sudoku

type Sudoku struct {
	Sectors []Sector
	Columns int
	Lines   int
}

type Sector struct {
	Tiles  []Tile
	Line   int
	Column int
}

type Tile struct {
	Value  int
	Line   int
	Column int
}

type Faixa struct {
	Number int
	Tiles  []*Tile
}

type Line Faixa
type Column Faixa

func GenerateSudoku() Sudoku {
	var sudoku Sudoku

	return sudoku
}

func (s *Sudoku) GetSectorByCoord(line, col int) *Sector {
	for i := range s.Sectors {
		if s.Sectors[i].Column == col && s.Sectors[i].Line == line {
			return &s.Sectors[i]
		}
	}

	return nil
}

func (s *Sudoku) GetLine(lineNumber int) (line *Line) {
	line = &Line{}

	if lineNumber < 0 || lineNumber > s.Lines {
		return
	}

	for i := range s.Sectors {
		for j := range s.Sectors[i].Tiles {
			if s.Sectors[i].Tiles[j].Line == lineNumber {
				line.Tiles = append(line.Tiles, &s.Sectors[i].Tiles[j])
			}
		}
	}

	return
}

func (s *Sudoku) GetColumn(colNumber int) (column *Column) {
	column = &Column{}

	if colNumber < 0 || colNumber > s.Columns {
		return
	}

	for i := range s.Sectors {
		for j := range s.Sectors[i].Tiles {
			if s.Sectors[i].Tiles[j].Column == colNumber {
				column.Tiles = append(column.Tiles, &s.Sectors[i].Tiles[j])
			}
		}
	}

	return
}
