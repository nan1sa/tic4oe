package main

import (
	"fmt"
)

func duplicateCheck(inputQueue []Data, x float64, y float64) bool {
	result := false

	for _, i := range inputQueue {
		if x == i.X && y == i.Y {
			result = true
		}
	}

	return result
}

func generateMap(inputQueue []Data) [][]int {
	result := [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}

	if len(inputQueue) > 4 {
		i := inputQueue[0]

		iX := int(i.X) - 1
		iY := int(i.Y) - 1

		if i.O {
			piece := 3
			result[iX][iY] = piece
		} else {
			piece := 4
			result[iX][iY] = piece
		}

		inputQueue = inputQueue[1:]
	}

	for _, i := range inputQueue {
		iX := int(i.X) - 1
		iY := int(i.Y) - 1

		if i.O {
			piece := 1
			result[iX][iY] = piece
		} else {
			piece := 2
			result[iX][iY] = piece
		}
	}

	return result
}

func drawMap(gameMap [][]int) {
	fmt.Println("")

	for _, i := range gameMap {
		for _, j := range i {
			switch j {
				case 1:
					fmt.Printf("\x1b[31m%s\x1b[0m", "o")
				case 2:
					fmt.Printf("\x1b[34m%s\x1b[0m", "x")
				case 3:
					fmt.Printf("%s", "o")
				case 4:
					fmt.Printf("%s", "x")
				default:
					fmt.Print("-")
			}
		}
		fmt.Println("")
	}

	fmt.Println("")
}

func judgement(gameMap [][]int, turn int) bool {
	judgementList := [][]int{{gameMap[0][0], gameMap[0][1], gameMap[0][2]}, {gameMap[1][0], gameMap[1][1], gameMap[1][2]}, {gameMap[2][0], gameMap[2][1], gameMap[2][2]}, {gameMap[0][0], gameMap[1][0], gameMap[2][0]}, {gameMap[0][1], gameMap[1][1], gameMap[2][1]}, {gameMap[0][0], gameMap[1][0], gameMap[2][0]}, {gameMap[0][0], gameMap[1][1], gameMap[2][2]}, {gameMap[0][2], gameMap[1][1], gameMap[2][0]}}
	result := false

	for _, i := range judgementList {
		if (i[0] == turn || i[0] == turn + 2) && (i[1] == turn || i[1] == turn + 2) && (i[2] == turn || i[2] == turn + 2) {
			result = true
		}
	}

	return result
}

