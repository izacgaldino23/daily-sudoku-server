package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/izacgaldino23/daily-sudoku-server/sudoku"
)

func main() {
	sudoku := sudoku.GenerateSudoku(2, 3, 3, 2)

	printSudoku(sudoku)
}

func printSudoku(s sudoku.Sudoku) {
	log.Printf("Sectors count: [%v]", len(s.Sectors))
	log.Printf("Sector length: [%v]", len(s.Sectors[0].Tiles))

	for i := 0; i < s.Lines*s.SectorLineCount; i++ {
		for j := 0; j < s.Columns*s.SectorColumnCount; j++ {
			fmt.Print(s.GetTileByCoord(i, j).Value, " ")

			if (j+1)%s.SectorColumnCount == 0 && j+1 != s.Columns*s.SectorColumnCount {
				fmt.Print("| ")
			}
		}
		fmt.Println("")
		if (i+1)%s.SectorLineCount == 0 && i+1 != s.Lines*s.SectorLineCount {
			line := []string{}
			for k := 0; k < s.Columns; k++ {
				line = append(line, strings.Repeat("- ", s.SectorColumnCount))
			}
			fmt.Println(strings.Join(line, "+ "))
		}
	}
}
