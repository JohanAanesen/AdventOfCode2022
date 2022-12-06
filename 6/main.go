package main

import (
	"bufio"
	"fmt"
	"os"
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

	var input = make([]string, 0)

	for r.Scan() {
		line := r.Text()

		input = strings.Split(line, "")

	}

	for i := 0; i < len(input)-4; i++ {
		if checkUnique(input[i:i+4], 4) {
			fmt.Println("Part 1: ", i+4)
			break
		}
	}

	for i := 0; i < len(input)-14; i++ {
		if checkUnique(input[i:i+14], 14) {
			fmt.Println("Part 2: ", i+14)
			break
		}
	}

}

func checkUnique(input []string, len int) bool {

	for i := 0; i < len; i++ {
		for j := 0; j < len; j++ {
			if input[i] == input[j] && i != j {
				return false
			}
		}
	}

	return true
}
