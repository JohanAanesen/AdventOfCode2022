package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	b, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)

	//pair := make([]string, 0)
	contains := 0
	overlaps := 0

	for r.Scan() {
		line := r.Text()
		var oneLower int
		var oneHigher int
		var twoLower int
		var twoHigher int

		fmt.Sscanf(line, "%d-%d,%d-%d\n", &oneLower, &oneHigher, &twoLower, &twoHigher)

		if checkContains(oneLower, oneHigher, twoLower, twoHigher) {
			contains++
		}

		if checkOverlap(oneLower, oneHigher, twoLower, twoHigher) {
			overlaps++
		}
	}

	fmt.Println("Part 1: ", contains)
	fmt.Println("Part 2: ", overlaps)

}

func checkContains(a int, b int, c int, d int) bool {

	if a >= c && a <= d && b >= c && b <= d {
		return true
	} else if c >= a && c <= b && d >= a && d <= b {
		return true
	}

	return false
}

func checkOverlap(a int, b int, c int, d int) bool {

	if (a >= c && a <= d) || (b >= c && b <= d) {
		return true
	} else if (c >= a && c <= b) || (d >= a && d <= b) {
		return true
	}

	return false
}
