package main

import (
	"fmt"
	"log"

	"github.com/izacgaldino23/daily-sudoku-server/sudoku"
)

func main() {
	sudoku := sudoku.GenerateSudoku(3, 3)

	printSudoku(sudoku)
}

func printSudoku(s sudoku.Sudoku) {
	log.Printf("Sectors count: [%v]", len(s.Sectors))
	log.Printf("Sector length: [%v]", len(s.Sectors[0].Tiles))

	for i := 0; i < s.Lines*s.Lines; i++ {
		for j := 0; j < s.Columns*s.Columns; j++ {
			fmt.Print(s.GetTileByCoord(i, j).Value, " ")

			if (j+1)%s.Columns == 0 && j+1 != s.Columns*s.Columns {
				fmt.Print("| ")
			}
		}
		fmt.Println("")
		if (i+1)%s.Lines == 0 && i+1 != s.Lines*s.Lines {
			fmt.Println("- - - + - - - + - - -")
		}
	}
}
