package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\n\n")
	fmt.Println("=============================================================")
	fmt.Println("=                 StellarVerge Build Utility                =")
	fmt.Println("=============================================================")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("\tbuild <component> [options]")
	fmt.Println("\thelp <command>")

	fmt.Println("\texit")
	fmt.Print("\n> ")

	for scanner.Scan() {
		parseInput(scanner.Text())
		fmt.Print("\n> ")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func parseInput(s string) {
	in := strings.Split(s, " ")
	switch in[0] {
	case "exit":
		os.Exit(0)
	case "help":
		helper(in)
	case "clear":
		clear()
	case "build":
		builder(in)
	default:
		fmt.Println("Unknown Command")
	}
}
