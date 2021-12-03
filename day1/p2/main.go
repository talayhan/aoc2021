package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("../input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var num int
	var first int
	var second int
	var third int

	var curSum int
	var prevSum int
	var countLTP int

	for {
		// pattern to scan
		_, err := fmt.Fscanf(file, "%d\n", &num)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		third = second
		second = first
		first = num

		prevSum = curSum
		curSum = first + second + third
		//fmt.Println("First: ", first, " Second: ", second, "Third: ", third)
		//fmt.Println("prevSum: ", prevSum, " curSum: ", curSum)

		if prevSum < curSum {
			countLTP++
		}
	}

	// print out the result minus first three check
	fmt.Println("Count: ", countLTP-3)
}
