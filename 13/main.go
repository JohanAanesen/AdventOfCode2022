package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type pair struct {
	pair1 string
	pair2 string
}

func main() {

	b, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)

	var lines []string

	for r.Scan() {
		line := r.Text()
		var row string
		fmt.Sscanf(line, "%s\n", &row)

		lines = append(lines, row)

	}

	pairs := parseInput(lines)

	counter := 0

	for i, pair := range pairs {
		if res := comparePairs(pair[0], pair[1]); res == 0 || res == 1 {
			counter += i + 1
		}
	}
	fmt.Println("Part 1: ", counter)

	lines = append(lines, []string{"[[2]]", "[[6]]"}...)
	pairs2 := parseInput(lines)

	allPackets := make([]any, 0, len(pairs2)*2)
	for _, p := range pairs2 {
		allPackets = append(allPackets, p...)
	}
	sort.Slice(allPackets, func(i int, j int) bool {
		a := allPackets[i]
		b := allPackets[j]
		return comparePairs(a, b) == 1
	})
	counter2 := 1
	for i, packet := range allPackets {
		asString := fmt.Sprintf("%v", packet)
		if asString == "[[[2]]]" || asString == "[[[6]]]" {
			counter2 *= i + 1
		}
	}

	fmt.Println("Part 2: ", counter2)
}

func parseInput(lines []string) [][]any {
	pairs := make([][]any, 0)
	pair := make([]any, 0, 2)
	for _, line := range lines {
		if line == "" {
			continue
		}
		expr, _ := parseExpression(line)
		pair = append(pair, expr)
		if len(pair) == 2 {
			pairs = append(pairs, pair)
			pair = make([]any, 0, 2)
		}
	}

	return pairs
}

func comparePairs(a any, b any) int {
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)
	if aIsInt && bIsInt {
		if aInt > bInt {
			return -1
		} else if bInt > aInt {
			return 1
		}
		return 0
	}
	aList, aIsList := a.([]any)
	bList, bIsList := b.([]any)
	if !aIsList {
		aList = []any{aInt}
	}
	if !bIsList {
		bList = []any{bInt}
	}

	maxLen := len(aList)
	if len(bList) > maxLen {
		maxLen = len(bList)
	}

	for i := 0; i < maxLen; i++ {
		if i >= len(aList) {
			return 1
		}
		if i >= len(bList) {
			return -1
		}
		if sub := comparePairs(aList[i], bList[i]); sub != 0 {
			return sub
		}
	}
	return 0
}

func parseExpression(expr string) (any, int) {
	chars := []rune(expr)
	idx := 0
	out := make([]any, 0)
	nChars := make([]rune, 0)
	for idx < len(chars) {
		char := chars[idx]
		switch char {
		case '[':
			x, i := parseExpression(string(chars[idx+1:]))
			out = append(out, x)
			idx += i + 1
		case ']':
			if len(nChars) > 0 {
				n, _ := strconv.Atoi(string(nChars))
				out = append(out, n)
				nChars = make([]rune, 0)
			}
			idx++
			return out, idx
		case ',':
			if len(nChars) > 0 {
				n, _ := strconv.Atoi(string(nChars))
				out = append(out, n)
				nChars = make([]rune, 0)
			}
			idx++
		default:
			nChars = append(nChars, char)
			idx++
		}
	}
	return out, idx
}
