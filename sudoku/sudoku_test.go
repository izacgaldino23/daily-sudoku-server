package sudoku

import (
	"testing"

	arrayFuncs "github.com/izacgaldino23/array-funcs"
	"github.com/stretchr/testify/assert"
)

const (
	MAGIC_NUMBER = 45
)

func TestSudokuGeneration(t *testing.T) {

	var sudoku = GenerateSudoku()

	t.Run("TestSector", func(t *testing.T) {
		var sectorLines = 3
		var sectorColumns = 3
		var sectorsLength = 9

		t.Run("TestSectorLength", func(t *testing.T) {
			assert.Equal(t, sectorsLength, len(sudoku))
		})

		t.Run("TestSumOfSectorsNumbers", func(t *testing.T) {
			for _, sector := range sudoku {
				var tilesArray = arrayFuncs.AnyToArrayKind(sector.Tiles)
				var numbers = tilesArray.Map(func(v *Tile, i int) {

				})
				sumElements(sector.Tiles)
			}
		})
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
