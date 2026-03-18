package main

import (
	"fmt"
	"os"
	"strconv"

	dispatcher "smarter-screen"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Fprintf(os.Stderr, "Usage: sort <width> <height> <length> <mass>\n")
		os.Exit(1)
	}

	width, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid width: %s\n", os.Args[1])
		os.Exit(1)
	}

	height, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid height: %s\n", os.Args[2])
		os.Exit(1)
	}

	length, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid length: %s\n", os.Args[3])
		os.Exit(1)
	}

	mass, err := strconv.ParseFloat(os.Args[4], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid mass: %s\n", os.Args[4])
		os.Exit(1)
	}

	fmt.Println(dispatcher.Sort(width, height, length, mass))
}
