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

	var curNum int
	var prevNum int
	var countLTP int

	for {
		// pattern to scan
		_, err := fmt.Fscanf(file, "%d\n", &curNum)

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}

		if curNum > prevNum {
			//fmt.Printf("prevNum: %d - curNum: %d (Incremented)\n", prevNum, curNum)
			countLTP++
		}
		prevNum = curNum
	}

	// print out the result minus first check
	fmt.Println("Count: ", countLTP-1)
}
