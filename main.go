package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: colour <number>")
		os.Exit(1)
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid number.")
	}
	fmt.Printf("\033[%03dmColour code: %03d\033[0m\n", num, num)
}
