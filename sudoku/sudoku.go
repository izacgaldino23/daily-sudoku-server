package sudoku

import (
	"github.com/izacgaldino23/daily-sudoku-server/utils"
)

type Sudoku struct {
	Sectors           []Sector
	Columns           int
	Lines             int
	SectorLineCount   int
	SectorColumnCount int
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
	Sector *Sector
}

type Track struct {
	Number int
	Tiles  []*Tile
}

type Line Track
type Column Track

func GenerateSudoku(lines, cols, sectorLines, sectorCols int) Sudoku {
	var sudoku Sudoku = Sudoku{
		Columns:           cols,
		Lines:             lines,
		SectorLineCount:   sectorLines,
		SectorColumnCount: sectorCols,
	}

	// Create sectors
	sudoku.CreateSectors()

	// Try generate sudoku numbers
	sudoku.generateAllNumbers(0, 0, sudoku.Lines*sudoku.SectorLineCount, sudoku.Columns*sudoku.SectorColumnCount)

	return sudoku
}

func (s *Sudoku) CreateSectors() {
	var sectorsLength = s.SectorLineCount * s.SectorColumnCount

	for i := 0; i < s.Lines; i++ {
		for j := 0; j < s.Columns; j++ {
			s.Sectors = append(s.Sectors, Sector{
				Tiles:       []Tile{},
				Line:        i,
				Column:      j,
				StartLine:   i * s.SectorLineCount,
				StartColumn: j * s.SectorColumnCount,
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
	for i := range s.Sectors {
		for j := range s.Sectors[i].Tiles {
			if s.Sectors[i].Tiles[j].Column == col && s.Sectors[i].Tiles[j].Line == line {
				return &s.Sectors[i]
			}
		}
	}

	return nil
}

func (s *Sudoku) GetTileByCoord(line, col int) *Tile {
	for i := range s.Sectors {
		for j := range s.Sectors[i].Tiles {
			if s.Sectors[i].Tiles[j].Column == col && s.Sectors[i].Tiles[j].Line == line {
				return &s.Sectors[i].Tiles[j]
			}
		}
	}

	return nil
}

func (s *Sudoku) GetLine(lineNumber int) (line *Line) {
	line = &Line{}

	if lineNumber < 0 || lineNumber >= s.Lines*s.Lines {
		return
	}

	for i := 0; i < s.Columns*s.SectorColumnCount; i++ {
		tile := s.GetTileByCoord(lineNumber, i)
		line.Tiles = append(line.Tiles, tile)

	}

	return
}

func (s *Sudoku) GetColumn(colNumber int) (column *Column) {
	column = &Column{}

	if colNumber < 0 || colNumber >= s.Columns*s.Columns {
		return
	}

	for i := 0; i < s.Lines*s.SectorLineCount; i++ {
		tile := s.GetTileByCoord(i, colNumber)
		column.Tiles = append(column.Tiles, tile)

	}

	return
}

func (s *Sudoku) fillSector(sector *Sector, size int) {
	var number = 1

	for i := 0; i < s.SectorLineCount; i++ {
		for j := 0; j < s.SectorColumnCount; j++ {
			sector.Tiles = append(sector.Tiles, Tile{
				Line:   sector.StartLine + i,
				Column: sector.StartColumn + j,
				Sector: sector,
			})

			number++
		}
	}
}

func (s *Sudoku) generateAllNumbers(actualLine, actualCol, linesCount, colsCount int, triedNumbers ...int) bool {
	if triedNumbers == nil {
		triedNumbers = []int{}
	}
	// log.Print(fmt.Sprintf("Trying to search LINE: [%v] COL: [%v]", actualLine, actualCol))

	valid, number := s.generateNumber(actualLine, actualCol, triedNumbers)
	triedNumbers = append(triedNumbers, number)

	// log.Print(fmt.Sprintf("Valid number [%v]? %v", number, valid))

	if valid {
		s.GetTileByCoord(actualLine, actualCol).Value = number

		actualCol++

		// increment to next tile
		if actualCol == colsCount {
			actualCol = 0
			actualLine++

			if actualLine == linesCount {
				return true
			}
		}

		if !s.generateAllNumbers(actualLine, actualCol, linesCount, colsCount) {
			// decrement to back tile
			actualCol--
			if actualCol < 0 {
				actualCol = colsCount - 1
				actualLine--
			}

			s.GetTileByCoord(actualLine, actualCol).Value = 0

			return s.generateAllNumbers(actualLine, actualCol, linesCount, colsCount, triedNumbers...)
		}
	} else {
		return false
	}

	return true
}

func (s *Sudoku) generateNumber(line, col int, triedBefore []int) (valid bool, number int) {
	numbers := []int{}
	size := s.SectorLineCount * s.SectorColumnCount

	for i := 1; i <= size; i++ {
		if !utils.Has(triedBefore, i) {
			numbers = append(numbers, i)
		}
	}

	utils.Shuffle(numbers)

	for i := range numbers {
		number = numbers[i]
		if invalidNumber := s.verifyIfNumberIsAlreadyPlacedOnTrack(number, line, col); !invalidNumber {
			valid = true
			break
		}
	}

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

func (s *Sudoku) verifyIfNumberIsAlreadyPlacedOnTrack(value, line, col int) bool {
	// Get numbers from sector
	sectorTiles := s.GetSectorByTileCoord(line, col).Tiles

	for i := range sectorTiles {
		if sectorTiles[i].Value == value {
			return true
		}
	}

	// Get numbers from line
	lineTiles := s.GetLine(line).Tiles
	for i := range lineTiles {
		if lineTiles[i].Value == value {
			return true
		}
	}

	// Get numbers from collumn
	columnTiles := s.GetColumn(col).Tiles
	for i := range columnTiles {
		if columnTiles[i].Value == value {
			return true
		}
	}

	return false
}
