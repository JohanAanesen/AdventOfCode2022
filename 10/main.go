package main

import (
	"bufio"
	"fmt"
	"os"
)

var x = 1
var sycle = 1
var crtXCount = 0
var crtYCount = 0
var crtString = ""

func main() {

	b, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)

	sycles := make(map[int]int)

	for r.Scan() {
		line := r.Text()

		var val int

		if line == "noop" {
			calculateCrt()
			sycles[sycle] = x
			sycle++
			continue
		}

		fmt.Sscanf(line, "addx %d\n", &val)

		calculateCrt()
		sycles[sycle] = x
		sycle++

		calculateCrt()
		sycles[sycle] = x
		sycle++
		x += val

	}

	//20th, 60th, 100th, 140th, 180th, and 220th
	v20 := sycles[20] * 20
	v60 := sycles[60] * 60
	v100 := sycles[100] * 100
	v140 := sycles[140] * 140
	v180 := sycles[180] * 180
	v220 := sycles[220] * 220

	fmt.Println("Part 1: ", v20+v60+v100+v140+v180+v220)
	fmt.Println("-----------------------------------------------------")
	fmt.Println("Part 2:")
	fmt.Println("-----------------------------------------------------")
	fmt.Println(crtString[0:39])
	fmt.Println(crtString[40:79])
	fmt.Println(crtString[80:119])
	fmt.Println(crtString[120:159])
	fmt.Println(crtString[160:199])
	fmt.Println(crtString[200:239])
	fmt.Println("-----------------------------------------------------")
}

func calculateCrt() {
	if x == crtXCount || x+1 == crtXCount || x-1 == crtXCount {
		crtString += "#"
	} else {
		crtString += "."
	}
	crtXCount++
	if int(sycle)%40 == 0 {
		crtYCount++
		crtXCount = 0
	}
}
