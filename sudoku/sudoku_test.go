package sudoku

import (
	"testing"

	"github.com/izacgaldino23/daily-sudoku-server/utils"
	"github.com/stretchr/testify/assert"
)

const (
	MAGIC_NUMBER = 45
)

func TestSudokuGeneration(t *testing.T) {

	var sudoku = GenerateSudoku()

	t.Run("TestSector", func(t *testing.T) {
		// var sectorLines = 3
		// var sectorColumns = 3
		var sectorsLength = 9

		t.Run("TestSectorLength", func(t *testing.T) {
			assert.Equal(t, sectorsLength, len(sudoku.Sectors))
		})

		for _, sector := range sudoku.Sectors {
			var numbers = utils.Map(sector.Tiles, func(v Tile) int {
				return v.Value
			})

			t.Run("TestSumOfSectorsNumbers", func(t *testing.T) {
				assert.Equal(t, MAGIC_NUMBER, sumElements(numbers...))
			})

			t.Run("TestHasAllNumber", func(t *testing.T) {
				for i := 1; i <= 9; i++ {
					assert.True(t, utils.Has(numbers, i))
				}
			})
		}
	})

	t.Run("TestLines", func(t *testing.T) {

	})

	t.Run("TestColumns", func(t *testing.T) {

	})

}

func sumElements(numbers ...int) int {
	var total = 0

	for _, v := range numbers {
		total += v
	}

	return total
}
