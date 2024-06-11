package main

import (
	"fmt"
)

type Data struct {
	X, Y float64
	O bool
}

func main() {
	var inputQueue []Data
	var inputCount int

	fmt.Print("先攻: ")

	for {
		var x, y float64

		for {
			fmt.Print("1〜3の範囲で座標を入力してください（例: 2 3）: ")
			_, err := fmt.Scan(&x, &y)

			if err != nil || (x < 1 || x > 3) || (y < 1 || y > 3) {
				fmt.Println("1〜3の範囲で入力してください")
				continue
			}

			if duplicateCheck(inputQueue, x, y) {
				fmt.Println("座標が重複しています")
				continue
			}

			break
		}

		data := Data{X: x, Y: y, O: inputCount%2 == 0}
		inputQueue = append(inputQueue, data)

		inputCount++

		if len(inputQueue) > 6 {
			inputQueue = inputQueue[len(inputQueue)-6:]
		}

		fmt.Printf("現在の入力: %+v\n", inputQueue)
		
		drawMap(generateMap(inputQueue))
		
		if judgement(generateMap(inputQueue), 2 - int(inputCount%2)) {
			switch data.O {
				case true:
					fmt.Println("先攻の勝利です")
				default:
					fmt.Println("後攻の勝利です ")
			}
			break
		}

		if len(inputQueue) > 5 {
			switch data.O {
				case true:
					fmt.Printf("\n手持ちの駒を使い切ったため、次に打つ後攻に駒が返却されました\n")
				default:
					fmt.Printf("\n手持ちの駒を使い切ったため、次に打つ先攻に駒が返却されました\n")
			}

			inputQueue = inputQueue[len(inputQueue)-5:]
			drawMap(generateMap(inputQueue))
		}

		switch data.O {
			case true:
				fmt.Print("後攻: ")
			default:
				fmt.Print("先攻: ")
		}
	}
}

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
					fmt.Print("o")
				case 2:
					fmt.Print("x")
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
		if i[0] == turn && i[1] == turn && i[2] == turn {
			result = true
		}
	}

	return result
}
