package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		io.WriteString(os.Stderr, "Please give one or more floats\n")
		os.Exit(1)
	}

	args := os.Args
	min, _ := strconv.ParseFloat(args[1], 64)
	max, _ := strconv.ParseFloat(args[1], 64)

	for _, arg := range args[1:] {
		n, _ := strconv.ParseFloat(arg, 64)

		if min > n {
			min = n
		}

		if max < n {
			max = n
		}
	}

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
