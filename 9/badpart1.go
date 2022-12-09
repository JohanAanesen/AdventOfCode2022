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

var grid [10000][10000]int
var Hx, Hy, Tx, Ty int

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

	for i := 0; i < 10000; i++ {
		for j := 0; j < 10000; j++ {
			grid[i][j] = 0
		}
	}

	//grid := [100][100]int{}
	start := 5000

	Hx, Hy, Tx, Ty = start, start, start, start

	grid[start][start] = 1 //start

	for _, v := range commands {
		if v.direction == "R" {
			for i := 0; i < v.amount; i++ {
				moveRight()
			}
		} else if v.direction == "L" {
			for i := 0; i < v.amount; i++ {
				moveLeft()
			}
		} else if v.direction == "U" {
			for i := 0; i < v.amount; i++ {
				moveUp()
			}
		} else if v.direction == "D" {
			for i := 0; i < v.amount; i++ {
				moveDown()
			}
		}
	}

	sum := 0

	for i := 0; i < 10000; i++ {
		for j := 0; j < 10000; j++ {
			sum += grid[i][j]
		}
	}
	fmt.Println("Part 1: ", sum)

}

func moveRight() {
	Hx++

	if Hy == Ty && abs(Hx-Tx) > 1 { //right
		Tx++
		grid[Ty][Tx] = 1
	} else if Hy > Ty && abs(Hx-Tx) > 1 { //right down
		Ty++
		Tx++
		grid[Ty][Tx] = 1
	} else if Hy < Ty && abs(Hx-Tx) > 1 { //right up
		Ty--
		Tx++
		grid[Ty][Tx] = 1
	}
}

func moveLeft() {
	Hx--

	if Hy == Ty && abs(Hx-Tx) > 1 { //left
		Tx--
		grid[Ty][Tx] = 1
	} else if Hy > Ty && abs(Hx-Tx) > 1 { //left up
		Ty++
		Tx--
		grid[Ty][Tx] = 1
	} else if Hy < Ty && abs(Hx-Tx) > 1 {
		Ty--
		Tx--
		grid[Ty][Tx] = 1
	}
}

func moveUp() {
	Hy--

	if Hx == Tx && abs(Hy-Ty) > 1 {
		Ty--
		grid[Ty][Tx] = 1
	} else if Hx > Tx && abs(Hy-Ty) > 1 {
		Ty--
		Tx++
		grid[Ty][Tx] = 1
	} else if Hx < Tx && abs(Hy-Ty) > 1 {
		Ty--
		Tx--
		grid[Ty][Tx] = 1
	}
}

func moveDown() {
	Hy++

	if Hx == Tx && abs(Hy-Ty) > 1 {
		Ty++
		grid[Ty][Tx] = 1
	} else if Hx > Tx && abs(Hy-Ty) > 1 {
		Ty++
		Tx++
		grid[Ty][Tx] = 1
	} else if Hx < Tx && abs(Hy-Ty) > 1 {
		Ty++
		Tx--
		grid[Ty][Tx] = 1
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
