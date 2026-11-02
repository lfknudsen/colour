package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		printTable()
		os.Exit(0)
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid number.")
	}
	text := fmt.Sprintf("Colour code: %03d", num)
	fmt.Print(format(text, num))
}

func format(text string, num int) string {
	return fmt.Sprintf("\033[%03dm%s\033[0m", num, text)
}

func printTable() {
	counter := 1
	for i := 0; i <= 127; i++ {
		text := fmt.Sprintf("%03d", i)
		fmt.Print(format(text, i))
		fmt.Print("  ")
		if counter%16 == 0 {
			fmt.Println()
		}
		counter++
	}
}
