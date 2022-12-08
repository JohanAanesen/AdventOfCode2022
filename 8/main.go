package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type tree struct {
	height  int
	visible bool
}

func main() {
	b, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)

	trees := make(map[int][]tree)

	rowCount := 0

	for r.Scan() {
		line := r.Text()

		var temp []string
		temp = strings.Split(line, "")

		for i, v := range temp {
			num, _ := strconv.Atoi(v)
			if i == 0 || i == 98 || rowCount == 0 || rowCount == 98 {
				trees[rowCount] = append(trees[rowCount], tree{height: num, visible: true})
			} else {
				trees[rowCount] = append(trees[rowCount], tree{height: num, visible: false})
			}

		}

		rowCount++

	}

	visible := 0

	for y := 0; y < 99; y++ {
		for x := 0; x < 99; x++ {
			if y == 0 || y == 98 || x == 0 || x == 98 {
				continue
			}

			if checkXLeft(x, trees[y]) {
				trees[y][x].visible = true
				continue
			} else if checkXRight(x, trees[y]) {
				trees[y][x].visible = true
				continue
			} else if checkYUp(x, y, trees) {
				trees[y][x].visible = true
				continue
			} else if checkYDown(x, y, trees) {
				trees[y][x].visible = true
				continue
			}

		}
	}

	for y := 0; y < 99; y++ {
		for x := 0; x < 99; x++ {
			if trees[y][x].visible {
				visible++
			}
		}
	}

	fmt.Println("Part 1: ", visible)

	highScore := 0

	for y := 1; y < 98; y++ {
		for x := 1; x < 98; x++ {

			//right
			rightCounter := 0
			for i := x + 1; i < len(trees[y]); i++ {
				if trees[y][i].height < trees[y][x].height {
					rightCounter++
				} else if trees[y][i].height >= trees[y][x].height {
					rightCounter++
					break
				}
			}

			//left
			leftCounter := 0
			for i := x - 1; i >= 0; i-- {
				if trees[y][i].height < trees[y][x].height {
					leftCounter++
				} else if trees[y][i].height >= trees[y][x].height {
					leftCounter++
					break
				}
			}

			//up
			upCounter := 0
			for i := y - 1; i >= 0; i-- {
				if trees[i][x].height < trees[y][x].height {
					upCounter++
				} else if trees[i][x].height >= trees[y][x].height {
					upCounter++
					break
				}
			}

			//down
			downCounter := 0
			for i := y + 1; i < len(trees); i++ {
				if trees[i][x].height < trees[y][x].height {
					downCounter++
				} else if trees[i][x].height >= trees[y][x].height {
					downCounter++
					break
				}
			}

			score := rightCounter * leftCounter * upCounter * downCounter
			if score > highScore {
				highScore = score
			}
		}
	}

	fmt.Println("Part 2: ", highScore)

}

func checkXLeft(position int, row []tree) bool {

	for i := 0; i < position; i++ {
		if row[i].height >= row[position].height {
			return false
		}
	}

	return true
}

func checkXRight(position int, row []tree) bool {

	for i := position + 1; i < len(row); i++ {
		if row[i].height >= row[position].height {
			return false
		}
	}

	return true
}

func checkYUp(x int, y int, cols map[int][]tree) bool {

	for i := 0; i < y; i++ {
		if cols[i][x].height >= cols[y][x].height {
			return false
		}
	}

	return true
}

func checkYDown(x int, y int, cols map[int][]tree) bool {

	for i := y + 1; i < len(cols); i++ {
		if cols[i][x].height >= cols[y][x].height {
			return false
		}
	}

	return true
}
