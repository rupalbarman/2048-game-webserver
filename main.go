package main 

import (
	"github.com/rupalbarman/2048-game-webserver/model"
	"github.com/rupalbarman/2048-game-webserver/view"
	//"github.com/rupalbarman/2048-game-webserver/controller"

	"fmt"
	"os"
)

func main() {

	webMode()
	//driverMode()
	//consoleMode()

}

func webMode() {
	b:= model.NewBoard()
	view.DisplayBoardConsole(b)

	//	Initiate the Board with some random 2's to start the game with.
	model.GenerateTile(b, 2, len(b.Matrix), len(b.Matrix[0]))
	model.GenerateTile(b, 2, len(b.Matrix), len(b.Matrix[0]))
	model.GenerateTile(b, 2, len(b.Matrix), len(b.Matrix[0]))
	model.GenerateTile(b, 2, len(b.Matrix), len(b.Matrix[0]))
	
	//controller.Run(b)
	run(b)
}

func consoleMode() {

	b:= model.NewBoard()
	view.DisplayBoardConsole(b)

	helpText := func() {
		fmt.Println("Press W A S D. To exit, press E")
	}

	//	Initiate the Board with some random 2's to start the game with.
	model.GenerateTile(b, 2, len(b.Matrix), len(b.Matrix[0]))
	model.GenerateTile(b, 2, len(b.Matrix), len(b.Matrix[0]))
	model.GenerateTile(b, 2, len(b.Matrix), len(b.Matrix[0]))
	model.GenerateTile(b, 2, len(b.Matrix), len(b.Matrix[0]))

	view.DisplayBoardConsole(b)
	helpText()

	var key []byte = make([]byte, 1)

	//	The main game loop. After each action, a new 2 is generated..
	//	at a position randomely decided.
	for {
		os.Stdin.Read(key)
		switch {
		case string(key) == "w":
			//UP
			model.SwipeVertical(b, true)
			model.GenerateTile(b, 2, len(b.Matrix), len(b.Matrix[0]))
			view.DisplayBoardConsole(b)
		case string(key) == "a":
			//LEFT
			model.SwipeHorizontal(b, true)
			model.GenerateTile(b, 2, len(b.Matrix), len(b.Matrix[0]))
			view.DisplayBoardConsole(b)
		case string(key) == "s":
			//DOWN
			model.SwipeVertical(b, false)
			model.GenerateTile(b, 2, len(b.Matrix), len(b.Matrix[0]))
			view.DisplayBoardConsole(b)
		case string(key) == "d":
			//Right
			model.SwipeHorizontal(b, false)
			model.GenerateTile(b, 2, len(b.Matrix), len(b.Matrix[0]))
			view.DisplayBoardConsole(b)
		case string(key) == "e":
			return
		}
	}
}

func driverMode() {
	fmt.Println("Driver Mode")

	b:= model.NewBoard()
	view.DisplayBoardConsole(b)

	b.Matrix[0][0] = 2
	b.Matrix[0][1] = 4
	b.Matrix[0][3] = 2
	b.Matrix[0][2] = 2
	b.Matrix[0][3] = 2
	b.Matrix[2][3] = 2
	b.Matrix[1][3] = 2
	b.Matrix[3][1] = 2
	fmt.Println("Initial arrangement")
	//INIT
	view.DisplayBoardConsole(b)
	//LEFT
	model.SwipeHorizontal(b, true)
	view.DisplayBoardConsole(b)
	//UP
	model.SwipeVertical(b, true)
	view.DisplayBoardConsole(b)
	//RIGHT
	model.SwipeHorizontal(b, false)
	view.DisplayBoardConsole(b)
	//DOWN
	model.SwipeVertical(b, false)
	view.DisplayBoardConsole(b)
}