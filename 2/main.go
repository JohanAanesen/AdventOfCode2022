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

	opponentPlays := make([]string, 0)
	myPlays := make([]string, 0)

	for r.Scan() {
		line := r.Text()
		var opponent string
		var me string

		fmt.Sscanf(line, "%s %s\n", &opponent, &me)

		opponentPlays = append(opponentPlays, opponent)
		myPlays = append(myPlays, me)
	}

	totScore := 0

	for i, opponent := range opponentPlays {
		switch opponent {
		case "A":
			totScore += calculateRock(myPlays[i])
		case "B":
			totScore += calculatePaper(myPlays[i])
		case "C":
			totScore += calculateScissor(myPlays[i])
		}
	}

	fmt.Println("Part 1: ", totScore)

	totScore2 := 0

	for i, opponent := range opponentPlays {
		switch opponent {
		case "A":
			totScore2 += calculateRock2(myPlays[i])
		case "B":
			totScore2 += calculatePaper2(myPlays[i])
		case "C":
			totScore2 += calculateScissor2(myPlays[i])
		}
	}

	fmt.Println("Part 2: ", totScore2)
}

func calculateRock(me string) int {

	switch me {
	case "X":
		return 1 + 3
	case "Y":
		return 2 + 6
	case "Z":
		return 3 + 0
	}

	return 0
}

func calculatePaper(me string) int {

	switch me {
	case "X":
		return 1 + 0
	case "Y":
		return 2 + 3
	case "Z":
		return 3 + 6
	}

	return 0
}

func calculateScissor(me string) int {

	switch me {
	case "X":
		return 1 + 6
	case "Y":
		return 2 + 0
	case "Z":
		return 3 + 3
	}

	return 0
}

func calculateRock2(me string) int {

	switch me {
	case "X":
		return 3 + 0
	case "Y":
		return 1 + 3
	case "Z":
		return 2 + 6
	}

	return 0
}

func calculatePaper2(me string) int {

	switch me {
	case "X":
		return 1 + 0
	case "Y":
		return 2 + 3
	case "Z":
		return 3 + 6
	}

	return 0
}

func calculateScissor2(me string) int {

	switch me {
	case "X":
		return 2 + 0
	case "Y":
		return 3 + 3
	case "Z":
		return 1 + 6
	}

	return 0
}
