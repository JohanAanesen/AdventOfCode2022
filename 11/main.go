package main

import "fmt"

var monkeys = [][]int{
	{54, 98, 50, 94, 69, 62, 53, 85},
	{71, 55, 82},
	{77, 73, 86, 72, 87},
	{97, 91},
	{78, 97, 51, 85, 66, 63, 62},
	{88},
	{87, 57, 63, 86, 87, 53},
	{73, 59, 82, 65},
}

var monkey0Counter uint64 = 0
var monkey1Counter uint64 = 0
var monkey2Counter uint64 = 0
var monkey3Counter uint64 = 0
var monkey4Counter uint64 = 0
var monkey5Counter uint64 = 0
var monkey6Counter uint64 = 0
var monkey7Counter uint64 = 0

func main() {
	for i := 0; i < 10000; i++ {
		monkey0()
		monkey1()
		monkey2()
		monkey3()
		monkey4()
		monkey5()
		monkey6()
		monkey7()
	}

	fmt.Println(monkey0Counter,
		monkey1Counter,
		monkey2Counter,
		monkey3Counter,
		monkey4Counter,
		monkey5Counter,
		monkey6Counter,
		monkey7Counter)

	fmt.Println(161943 * 156056)
}

func monkey0() {
	/*
	   Operation: new = old * 13
	  Test: divisible by 3
	    If true: throw to monkey 2
	    If false: throw to monkey 1
	*/

	for i := 0; i < len(monkeys[0]); i++ {
		monkey0Counter++

		newWorry := monkeys[0][i] * 13
		newWorry = newWorry % (2 * 11 * 7 * 5 * 17 * 19 * 13 * 3)
		if newWorry%3 == 0 {
			monkeys[2] = append(monkeys[2], newWorry)
		} else {
			monkeys[1] = append(monkeys[1], newWorry)
		}
	}
	monkeys[0] = nil

}
func monkey1() {
	for i := 0; i < len(monkeys[1]); i++ {
		monkey1Counter++

		newWorry := monkeys[1][i] + 2
		newWorry = newWorry % (2 * 11 * 7 * 5 * 17 * 19 * 13 * 3)
		if newWorry%13 == 0 {
			monkeys[7] = append(monkeys[7], newWorry)
		} else {
			monkeys[2] = append(monkeys[2], newWorry)
		}
	}
	monkeys[1] = nil

}
func monkey2() {
	for i := 0; i < len(monkeys[2]); i++ {
		monkey2Counter++

		newWorry := monkeys[2][i] + 8
		newWorry = newWorry % (2 * 11 * 7 * 5 * 17 * 19 * 13 * 3)
		if newWorry%19 == 0 {
			monkeys[4] = append(monkeys[4], newWorry)
		} else {
			monkeys[7] = append(monkeys[7], newWorry)
		}
	}
	monkeys[2] = nil
}
func monkey3() {
	for i := 0; i < len(monkeys[3]); i++ {
		monkey3Counter++

		newWorry := monkeys[3][i] + 1
		newWorry = newWorry % (2 * 11 * 7 * 5 * 17 * 19 * 13 * 3)
		if newWorry%17 == 0 {
			monkeys[6] = append(monkeys[6], newWorry)
		} else {
			monkeys[5] = append(monkeys[5], newWorry)
		}
	}
	monkeys[3] = nil
}
func monkey4() {
	for i := 0; i < len(monkeys[4]); i++ {
		monkey4Counter++

		newWorry := monkeys[4][i] * 17
		newWorry = newWorry % (2 * 11 * 7 * 5 * 17 * 19 * 13 * 3)
		if newWorry%5 == 0 {
			monkeys[6] = append(monkeys[6], newWorry)
		} else {
			monkeys[3] = append(monkeys[3], newWorry)
		}
	}
	monkeys[4] = nil
}
func monkey5() {
	for i := 0; i < len(monkeys[5]); i++ {
		monkey5Counter++

		newWorry := monkeys[5][i] + 3
		newWorry = newWorry % (2 * 11 * 7 * 5 * 17 * 19 * 13 * 3)
		if newWorry%7 == 0 {
			monkeys[1] = append(monkeys[1], newWorry)
		} else {
			monkeys[0] = append(monkeys[0], newWorry)
		}
	}
	monkeys[5] = nil
}
func monkey6() {
	for i := 0; i < len(monkeys[6]); i++ {
		monkey6Counter++

		newWorry := monkeys[6][i] * monkeys[6][i]
		newWorry = newWorry % (2 * 11 * 7 * 5 * 17 * 19 * 13 * 3)
		if newWorry%11 == 0 {
			monkeys[5] = append(monkeys[5], newWorry)
		} else {
			monkeys[0] = append(monkeys[0], newWorry)
		}
	}
	monkeys[6] = nil

}
func monkey7() {
	for i := 0; i < len(monkeys[7]); i++ {
		monkey7Counter++

		newWorry := monkeys[7][i] + 6
		newWorry = newWorry % (2 * 11 * 7 * 5 * 17 * 19 * 13 * 3)
		if newWorry%2 == 0 {
			monkeys[4] = append(monkeys[4], newWorry)
		} else {
			monkeys[3] = append(monkeys[3], newWorry)
		}
	}
	monkeys[7] = nil

}
