package main

import (
	"fmt"
)

func helper(in []string) {

	if len(in) < 2 {
		fmt.Println("Usage: help <command>")
		return
	}

	in = in[1:] //Remove the first element from the array

	switch in[0] {
	case "help":
		fmt.Println("Help would show you all you need to know")
	case "build":
		helpBuild()
	}
}

func helpBuild() {
	fmt.Println("With build you can compile all parts of the project")
	fmt.Println("It will also make sure you have all necessary dependencies.")
	fmt.Println("Available Components:")
	fmt.Println("\tgame")
	fmt.Println("\tandroid-app")
	fmt.Println("\twebsite")
	fmt.Println("Available options")
	fmt.Println("\t--remove-dependencies")
}

// This is a very hacky clear... Sorry
func clear() {
	fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
}
