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

	var x int // horizontal
	var y int // depth

	var val int
	var cmd string

	for {
		// pattern to scan
		_, err := fmt.Fscanf(file, "%s %d\n", &cmd, &val)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		switch {
		case cmd == "forward":
			x += val
		case cmd == "down":
			y += val
		case cmd == "up":
			y -= val
		}
		//fmt.Println("Command: ", cmd, " Value: ", val)
	}
	fmt.Println("Multiplication result: ", x*y)
}
