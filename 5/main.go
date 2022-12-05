package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	b, err := os.Open("input2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)

	stacks := [][]string{
		{"Q", "S", "W", "C", "Z", "V", "F", "T"},
		{"Q", "R", "B"},
		{"B", "Z", "T", "Q", "P", "M", "S"},
		{"D", "V", "F", "R", "Q", "H"},
		{"J", "G", "L", "D", "B", "S", "T", "P"},
		{"W", "R", "T", "Z"},
		{"H", "Q", "M", "N", "S", "F", "R", "J"},
		{"R", "N", "F", "H", "W"},
		{"J", "Z", "T", "Q", "P", "R", "B"},
	}
	/*stacks := [][]string{
		{"Z", "N"},
		{"M", "C", "D"},
		{"P"},
	}*/
	/*
		[T]             [P]     [J]
		[F]     [S]     [T]     [R]     [B]
		[V]     [M] [H] [S]     [F]     [R]
		[Z]     [P] [Q] [B]     [S] [W] [P]
		[C]     [Q] [R] [D] [Z] [N] [H] [Q]
		[W] [B] [T] [F] [L] [T] [M] [F] [T]
		[S] [R] [Z] [V] [G] [R] [Q] [N] [Z]
		[Q] [Q] [B] [D] [J] [W] [H] [R] [J]
		 1   2   3   4   5   6   7   8   9
	*/

	for r.Scan() {
		line := r.Text()

		var count, from, to int

		fmt.Sscanf(line, "move %d from %d to %d\n", &count, &from, &to)

		from = from - 1
		to = to - 1 //0 indexing

		stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-count:len(stacks[from])]...)

		//slice = slice[:len(slice)-1]
		stacks[from] = stacks[from][:len(stacks[from])-count]

		/* PART 1
		for i := 0; i < count; i++ {
			stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-1]) //add value to stack
			stacks[from] = stacks[from][:len(stacks[from])-1]                  //remove from previous
		}
		*/
	}

	part := ""

	for _, v := range stacks {
		part += v[len(v)-1]
	}

	fmt.Println("Answer: ", part)
}
