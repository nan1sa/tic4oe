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

