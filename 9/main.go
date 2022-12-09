package main

import (
	"bufio"
	"fmt"
	"os"
)

type command struct {
	direction string
	amount    int
}

type rope struct {
	x int
	y int
}

var gridSize = 10000
var start = 5000

var grid [10000][10000]int
var grid2 [10000][10000]int

func main() {
	b, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)

	var commands = make([]command, 0)

	for r.Scan() {
		line := r.Text()
		var input command
		fmt.Sscanf(line, "%s %d\n", &input.direction, &input.amount)

		commands = append(commands, input)

	}

	head := rope{start, start}
	tailOne := rope{start, start}
	tailTwo := rope{start, start}
	tailThree := rope{start, start}
	tailFour := rope{start, start}
	tailFive := rope{start, start}
	tailSix := rope{start, start}
	tailSeven := rope{start, start}
	tailEight := rope{start, start}
	tailNine := rope{start, start}

	grid[start][start] = 1 //start
	grid2[start][start] = 1

	for _, v := range commands {
		if v.direction == "R" {
			for i := 0; i < v.amount; i++ {
				head.x++
				moveTail(&head, &tailOne)
				moveTail(&tailOne, &tailTwo)
				moveTail(&tailTwo, &tailThree)
				moveTail(&tailThree, &tailFour)
				moveTail(&tailFour, &tailFive)
				moveTail(&tailFive, &tailSix)
				moveTail(&tailSix, &tailSeven)
				moveTail(&tailSeven, &tailEight)
				moveTail(&tailEight, &tailNine)
				grid[tailOne.x][tailOne.y] = 1
				grid2[tailNine.x][tailNine.y] = 1
			}
		} else if v.direction == "L" {
			for i := 0; i < v.amount; i++ {
				head.x--
				moveTail(&head, &tailOne)
				moveTail(&tailOne, &tailTwo)
				moveTail(&tailTwo, &tailThree)
				moveTail(&tailThree, &tailFour)
				moveTail(&tailFour, &tailFive)
				moveTail(&tailFive, &tailSix)
				moveTail(&tailSix, &tailSeven)
				moveTail(&tailSeven, &tailEight)
				moveTail(&tailEight, &tailNine)
				grid[tailOne.x][tailOne.y] = 1
				grid2[tailNine.x][tailNine.y] = 1
			}
		} else if v.direction == "U" {
			for i := 0; i < v.amount; i++ {
				head.y--
				moveTail(&head, &tailOne)
				moveTail(&tailOne, &tailTwo)
				moveTail(&tailTwo, &tailThree)
				moveTail(&tailThree, &tailFour)
				moveTail(&tailFour, &tailFive)
				moveTail(&tailFive, &tailSix)
				moveTail(&tailSix, &tailSeven)
				moveTail(&tailSeven, &tailEight)
				moveTail(&tailEight, &tailNine)
				grid[tailOne.x][tailOne.y] = 1
				grid2[tailNine.x][tailNine.y] = 1
			}
		} else if v.direction == "D" {
			for i := 0; i < v.amount; i++ {
				head.y++
				moveTail(&head, &tailOne)
				moveTail(&tailOne, &tailTwo)
				moveTail(&tailTwo, &tailThree)
				moveTail(&tailThree, &tailFour)
				moveTail(&tailFour, &tailFive)
				moveTail(&tailFive, &tailSix)
				moveTail(&tailSix, &tailSeven)
				moveTail(&tailSeven, &tailEight)
				moveTail(&tailEight, &tailNine)
				grid[tailOne.x][tailOne.y] = 1
				grid2[tailNine.x][tailNine.y] = 1
			}
		}
	}

	sum1 := 0
	sum2 := 0

	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			sum1 += grid[i][j]
			sum2 += grid2[i][j]
		}
	}
	fmt.Println("Part 1: ", sum1)
	fmt.Println("Part 2: ", sum2)
}

func moveTail(head *rope, tail *rope) {

	deltaX := head.x - tail.x
	deltaY := head.y - tail.y

	if abs(deltaX) > 1 || abs(deltaY) > 1 {

		tail.x += sign(deltaX)
		tail.y += sign(deltaY)
	}

	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sign(x int) int {
	if x < 0 {
		return -1
	}
	if x == 0 {
		return 0
	}
	return 1
}
