package sudoku

import (
	"testing"

	"github.com/izacgaldino23/daily-sudoku-server/utils"
	"github.com/stretchr/testify/assert"
)

const (
	MAGIC_NUMBER  = 45
	LINES         = 3
	COLUMNS       = 3
	SECTOR_LENGTH = 3
	ELEMENTS_SIZE = LINES * COLUMNS
)

func TestSudokuGeneration(t *testing.T) {

	var sudoku = GenerateSudoku(LINES, COLUMNS, SECTOR_LENGTH, SECTOR_LENGTH)

	t.Run("TestSector", func(t *testing.T) {

		t.Run("TestSectorLength", func(t *testing.T) {
			assert.Equal(t, ELEMENTS_SIZE, len(sudoku.Sectors))
		})

		for _, sector := range sudoku.Sectors {
			var numbers = utils.Map(sector.Tiles, func(v Tile) int {
				return v.Value
			})

			t.Run("TestSumOfSectorsNumbers", func(t *testing.T) {
				assert.Equal(t, MAGIC_NUMBER, sumElements(numbers...))
			})

			t.Run("TestHasAllNumber", func(t *testing.T) {
				for i := 1; i <= ELEMENTS_SIZE; i++ {
					assert.True(t, utils.Has(numbers, i))
				}
			})
		}
	})

	t.Run("TestLines", func(t *testing.T) {
		for i := 0; i < ELEMENTS_SIZE; i++ {
			line := sudoku.GetLine(i)

			var numbers = utils.Map(line.Tiles, func(v *Tile) int {
				return v.Value
			})

			t.Run("TestLineElementsCount", func(t *testing.T) {
				assert.Equal(t, ELEMENTS_SIZE, len(line.Tiles))
			})

			t.Run("TestLineElementsSum", func(t *testing.T) {
				assert.Equal(t, MAGIC_NUMBER, sumElements(numbers...))
			})

			t.Run("TestHasAllNumber", func(t *testing.T) {
				for j := 1; j <= ELEMENTS_SIZE; j++ {
					assert.True(t, utils.Has(numbers, j))
				}
			})
		}
	})

	t.Run("TestColumns", func(t *testing.T) {
		for i := 0; i < ELEMENTS_SIZE; i++ {
			column := sudoku.GetColumn(i)

			var numbers = utils.Map(column.Tiles, func(v *Tile) int {
				return v.Value
			})

			t.Run("TestColumnElementsCount", func(t *testing.T) {
				assert.Equal(t, ELEMENTS_SIZE, len(column.Tiles))
			})

			t.Run("TestColumnElementsSum", func(t *testing.T) {
				assert.Equal(t, MAGIC_NUMBER, sumElements(numbers...))
			})

			t.Run("TestHasAllNumber", func(t *testing.T) {
				for j := 1; j <= ELEMENTS_SIZE; j++ {
					assert.True(t, utils.Has(numbers, j))
				}
			})
		}
	})

}

func sumElements(numbers ...int) int {
	var total = 0

	for _, v := range numbers {
		total += v
	}

	return total
}
