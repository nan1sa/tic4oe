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

	for {
		var x, y float64

		for {
			fmt.Print("1〜3の範囲で座標を入力してください（例: 2 3）: ")

			_, err := fmt.Scan(&x, &y)
			if err != nil || (x < 1 || x > 3) || (y < 1 || y > 3) {
				fmt.Println("1〜3の範囲で入力してください")
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
	}
}
