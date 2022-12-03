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

	rucksacks := make([]string, 0)

	for r.Scan() {
		line := r.Text()
		var sack string

		fmt.Sscanf(line, "%s\n", &sack)

		rucksacks = append(rucksacks, sack)
	}

	characters := make([]byte, 0)

	for _, rucksack := range rucksacks {
		split := []byte(rucksack)

		firstHalf := split[0:(len(split) / 2)]
		secondHalf := split[(len(split) / 2):]
		for _, c := range firstHalf {
			if contains(secondHalf, byte(c)) {
				characters = append(characters, c)
				break
			}
		}
	}

	tally := 0

	for _, v := range characters {
		tally += priority(v)
	}

	fmt.Println("Part 1: ", tally)

	part2Chars := make([]byte, 0)

	for i := 0; i < len(rucksacks); i += 3 {

		a := []byte(rucksacks[i])
		b := []byte(rucksacks[i+1])
		c := []byte(rucksacks[i+2])

		var simChar byte

		for j := 0; j < len(a); j++ {
			for k := 0; k < len(b); k++ {
				for l := 0; l < len(c); l++ {
					if a[j] == b[k] && a[j] == c[l] && b[k] == c[l] {
						simChar = a[j]
					}
				}
			}
		}

		part2Chars = append(part2Chars, simChar)

	}

	tally2 := 0
	for _, v := range part2Chars {
		tally2 += priority(v)
	}
	fmt.Println("Part 2: ", tally2)
}

func priority(v byte) int {
	if v >= 97 {
		return int(v) - 96
	} else {
		return int(v) - 64 + 26
	}
}

func contains(s []byte, e byte) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
