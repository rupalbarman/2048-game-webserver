package model

import "math/rand"
import "time"

func init() {
	rand.Seed(time.Now().UnixNano())
}

//==========================================================================
// Random 2 generaion

func GenerateTile(b *Board, tileValue, rows, cols int) {
	new:
		x := rand.Intn(rows)
		y := rand.Intn(cols)

		if b.Matrix[x][y] != 0 {
			goto new
		}

		b.Matrix[x][y] = tileValue
}

//==========================================================================
// Shift operations

func ShiftLeft( b *Board) {
	
	count := 0
	for i,_:= range b.Matrix {
		for j,_:= range b.Matrix[i] {
			if b.Matrix[i][j] != 0 {
				b.Matrix[i][count] = b.Matrix[i][j]
				count += 1
			}
		}
		for count < len(b.Matrix[i]) {
			b.Matrix[i][count] = 0
			count += 1
		}
		count = 0
	}
}

func ShiftRight( b *Board) {

	removeZeros := func(slice []int, s int) []int {
		return append(slice[:s], slice[s+1:]...)
	}

	shiftSliceRight := func(slice []int, numOfZeros int) []int {
		slice_temp := make([]int, numOfZeros)
		slice_temp = append(slice_temp, slice[0:]...) //the 3 dots makes it an array,
		//ie. so that we append a fixed size element(array)
		return slice_temp
	}

	countZeros := 0

	for i := 0; i < len(b.Matrix); i++ {
		for j := 0; j < len(b.Matrix[i]); j++ {
			if b.Matrix[i][j] == 0 {
				b.Matrix[i] = removeZeros(b.Matrix[i], j)
				countZeros += 1
				j -= 1 //extra decrement each time to get accurate position and operation (misc)
			}
		}
		b.Matrix[i] = shiftSliceRight(b.Matrix[i], countZeros)
		countZeros = 0
	}
}

func ShiftUp(board_transpose *Board) {
	ShiftLeft(board_transpose)
}

func ShiftDown(board_transpose *Board) {
	ShiftRight(board_transpose)
}

//==========================================================================
// Swipe operations

func SwipeHorizontal(b *Board, left bool) {
	if left {
		ShiftLeft(b)
	} else {
		ShiftRight(b)
	}

	for i := 0; i < len(b.Matrix); i++ {
		for j := 0; j < len(b.Matrix[i]); j++ {
			if b.Matrix[i][j] != 0 && j+1 < len(b.Matrix[i]) {
				if b.Matrix[i][j] == b.Matrix[i][j+1] {
					b.Matrix[i][j] += b.Matrix[i][j+1]
					b.Matrix[i][j+1] = 0
				}
			}

		}
	}
	if left {
		ShiftLeft(b)
	} else {
		ShiftRight(b)
	}
}

func SwipeVertical(b *Board, up bool) {

	board_transpose := NewBoard()
	for i := 0; i < len(b.Matrix); i++ {
		board_transpose.Matrix[i] = make([]int, len(b.Matrix[i]))
	}

	// Transposing it ( swapping rows and colums), so we can operate
	// the swipe and shift operation up using swipe and shift left
	for i := 0; i < len(b.Matrix); i++ {
		for j := 0; j < len(b.Matrix[i]); j++ {
			board_transpose.Matrix[i][j] = b.Matrix[j][i]
		}
	}

	if up {
		ShiftUp(board_transpose)
	} else {
		ShiftDown(board_transpose)
	}

	SwipeHorizontal(board_transpose, true)

	if up {
		ShiftUp(board_transpose)
	} else {
		ShiftDown(board_transpose)
	}

	// Transposing it again to get orginal matrix/ b.Matrix
	for i := 0; i < len(b.Matrix); i++ {
		for j := 0; j < len(b.Matrix[i]); j++ {
			b.Matrix[i][j] = board_transpose.Matrix[j][i]
		}
	}
}

//==========================================================================
