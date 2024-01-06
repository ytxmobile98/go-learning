package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	const ( // should be both >= 256
		rows = 1024
		cols = 1024
	)

	pic := make([][]uint8, rows)
	for row := range pic {
		pic[row] = make([]uint8, cols)
		for col := range pic {
			pic[row][col] = func() uint8 {
				if row == col {
					return 255
				}
				return 0
			}()
		}
	}

	return pic
}

func main() {
	pic.Show(Pic)
}
