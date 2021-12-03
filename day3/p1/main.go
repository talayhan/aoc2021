package main

import (
	"fmt"
	"io"
	"os"
)

func bitCheck(val int, mask int, res *int) {
	if (val & mask) == mask {
		*res++
	} else {
		*res--
	}
}

func setZeroOrOne(val *int) {
	if *val > 0 {
		*val = 1
	} else {
		*val = 0
	}
}

func notEquals(val int) int {
	if val == 0 {
		return 1
	} else {
		return 0
	}
}

func main() {
	file, err := os.Open("../input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var val int
	// most common bits

	var bitCount = 12
	var bits [12]int

	var gammaRate int
	var epsilonRate int

	for {
		// pattern to scan
		_, err := fmt.Fscanf(file, "%b\n", &val)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}

		for i := 0; i < bitCount; i++ {
			bitCheck(val, 1<<(bitCount-1-i), &bits[i])
		}
	}

	for i := 0; i < bitCount; i++ {
		setZeroOrOne(&bits[i])
		gammaRate += bits[i] << (bitCount - 1 - i)
		epsilonRate += notEquals(bits[i]) << (bitCount - 1 - i)
	}

	fmt.Println("Gamma Rate in dec: ", gammaRate)
	fmt.Println("Epsilon Rate in dec: ", epsilonRate)
	fmt.Println("Power consumption : ", gammaRate*epsilonRate)
}
