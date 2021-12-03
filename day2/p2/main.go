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
	var z int // aim

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
			y += z * val
		case cmd == "down":
			z += val
		case cmd == "up":
			z -= val
		}
		//fmt.Println("Horizontal : ", x, " Depth: ", y, " Aim: ", z)
		//fmt.Println("Command: ", cmd, " Value: ", val)
	}
	fmt.Println("Multiplication result: ", x*y)
}
