package sudoku

import "math"

type Sudoku struct {
	Sectors []Sector
	Columns int
	Lines   int
}

type Sector struct {
	Tiles       []Tile
	Line        int
	Column      int
	StartLine   int
	StartColumn int
}

type Tile struct {
	Value  int
	Line   int
	Column int
}

type Track struct {
	Number int
	Tiles  []*Tile
}

type Line Track
type Column Track

func GenerateSudoku(lines, cols int) Sudoku {
	var sudoku Sudoku = Sudoku{
		Columns: cols,
		Lines:   lines,
	}

	// Create sectors
	sudoku.CreateSectors()

	// Generate numbers

	return sudoku
}

func (s *Sudoku) CreateSectors() {
	var sectorsLength = s.Lines * s.Columns

	for i := 0; i < s.Lines; i++ {
		for j := 0; j < s.Columns; j++ {
			s.Sectors = append(s.Sectors, Sector{
				Tiles:       []Tile{},
				Line:        i,
				Column:      j,
				StartLine:   i * s.Lines,
				StartColumn: j * s.Columns,
			})
		}
	}

	// Fill sectors with numbers
	for i := range s.Sectors {
		s.fillSector(&s.Sectors[i], sectorsLength)
	}
}

func (s *Sudoku) GetSectorByCoord(line, col int) *Sector {
	for i := range s.Sectors {
		if s.Sectors[i].Column == col && s.Sectors[i].Line == line {
			return &s.Sectors[i]
		}
	}

	return nil
}

func (s *Sudoku) GetSectorByTileCoord(line, col int) *Sector {
	x := float64(line) / float64(s.Lines)
	y := float64(col) / float64(s.Columns)

	sectorLine := math.Floor(x)
	sectorColumn := math.Floor(y)

	return s.GetSectorByCoord(int(sectorLine), int(sectorColumn))
}

func (s *Sudoku) GetLine(lineNumber int) (line *Line) {
	line = &Line{}

	if lineNumber < 0 || lineNumber >= s.Lines {
		return
	}

	for i := range s.Sectors {
		for j := range s.Sectors[i].Tiles {
			if s.Sectors[i].Tiles[j].Line == lineNumber+1 {
				line.Tiles = append(line.Tiles, &s.Sectors[i].Tiles[j])
			}
		}
	}

	return
}

func (s *Sudoku) GetColumn(colNumber int) (column *Column) {
	column = &Column{}

	if colNumber < 0 || colNumber >= s.Columns {
		return
	}

	for i := range s.Sectors {
		for j := range s.Sectors[i].Tiles {
			if s.Sectors[i].Tiles[j].Column == colNumber+1 {
				column.Tiles = append(column.Tiles, &s.Sectors[i].Tiles[j])
			}
		}
	}

	return
}

func (s *Sudoku) fillSector(sector *Sector, size int) {
	var number = 1

	for i := 0; i < s.Lines; i++ {
		for j := 0; j < s.Columns; j++ {
			sector.Tiles = append(sector.Tiles, Tile{
				Line:   sector.StartLine + i,
				Column: sector.StartColumn + j,
			})

			number++
		}
	}
}

func (s *Sudoku) generateNumber(line, col int) (number int) {
	// Get numbers from sector
	sectorTiles := s.GetSectorByTileCoord(line, col).Tiles

	// Get numbers from line
	lineTiles := 

	// Get numbers from collumn

	return
}

func (s *Sector) getTileByCoord(line, col int) *Tile {
	for i := range s.Tiles {
		if s.Tiles[i].Line == line && s.Tiles[i].Column == col {
			return &s.Tiles[i]
		}
	}

	return nil
}
