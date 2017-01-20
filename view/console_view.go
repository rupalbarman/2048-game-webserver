package view

import (
	"fmt"
	"model"
)

func DisplayBoardConsole( b *model.Board) {
	for i,_:= range b.Matrix {
		for j,_:= range b.Matrix[i] {
			if b.Matrix[i][j] != 0 {
				fmt.Printf("|\t%d\t", b.Matrix[i][j])
			} else {
				fmt.Printf("|\t \t")
			}
		}

		fmt.Printf("\n")
	}

	fmt.Printf("\n")
}