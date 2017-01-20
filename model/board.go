package model

const (
	rows= 4
	cols= 4
)

type Board struct {
	Matrix [][]int
}

func NewBoard() *Board {
	board:= Board{Matrix: make([][]int, rows)}
	for i,_ := range board.Matrix {
		board.Matrix[i]= make([]int, cols)
	}
	
	return &board
}