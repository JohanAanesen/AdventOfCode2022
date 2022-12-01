package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	b, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)

	dwarfCounter := 0
	dwarfSum := make([]int, 0)
	dwarfSum = append(dwarfSum, 0) //start an index

	for r.Scan() {
		line := r.Text()
		var value int

		if line == "" {
			dwarfCounter++
			dwarfSum = append(dwarfSum, 0)
		} else {
			fmt.Sscanf(line, "%d\n", &value)

			dwarfSum[dwarfCounter] += value
		}

	}

	sort.Ints(dwarfSum)

	fmt.Println("Part1: ", dwarfSum[len(dwarfSum)-1])

	top3 := 0

	top3 += dwarfSum[len(dwarfSum)-1]
	top3 += dwarfSum[len(dwarfSum)-2]
	top3 += dwarfSum[len(dwarfSum)-3]

	fmt.Println("Part2: ", top3)
}
