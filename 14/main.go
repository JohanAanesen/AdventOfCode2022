package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)

	var rocks [1000][1000]int

	for r.Scan() {
		line := r.Text()

		vals := strings.Split(line, " -> ")
		prevx := 0
		prevy := 0
		for i := 0; i < len(vals); i++ {
			s := strings.Split(vals[i], ",")
			x, _ := strconv.Atoi(s[0])
			y, _ := strconv.Atoi(s[1])

			//fmt.Println(x, y, prevx, prevy)

			if prevx == 0 && prevy == 0 {
				prevx = x
				prevy = y
				continue
			}

			if x == prevx {
				if y < prevy {
					for j := y; j <= prevy; j++ {
						rocks[j][x] = 1
					}
				} else {
					for j := prevy; j <= y; j++ {
						rocks[j][x] = 1
					}
				}
			} else if y == prevy {
				if x < prevx {
					for j := x; j <= prevx; j++ {
						rocks[y][j] = 1
					}
				} else {
					for j := prevx; j <= x; j++ {
						rocks[y][j] = 1
					}
				}
			}

			prevx = x
			prevy = y

		}
	}

	highestY := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if rocks[i][j] == 1 {
				highestY = i
			}
		}
	}

	fmt.Println(highestY)

	//part2
	for i := 0; i < 1000; i++ {
		rocks[highestY+2][i] = 1
	}

	counter := 0

	for {
		stop := false
		x, y := 500, 0
		for {

			//part1
			/*
				if(y == 999){
					stop = true
					break
				}
			*/

			if rocks[y+1][x] == 0 {
				y++
				continue
			} else if rocks[y+1][x-1] == 0 {
				y++
				x--
				continue
			} else if rocks[y+1][x+1] == 0 {
				y++
				x++
				continue
			} else {
				rocks[y][x] = 1

				counter++
				//part2
				if x == 500 && y == 0 {
					stop = true
				}
				break
			}
		}
		if stop {
			break
		}

	}

	fmt.Println("Part 2: ", counter)

}
