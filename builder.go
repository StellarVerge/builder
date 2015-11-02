package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
)

type flagsStruct struct {
	preserveDependencies bool
	saveLog              bool
}

var flags flagsStruct

func builder(in []string) {
	if len(in) < 3 {
		fmt.Println("Usage: build <component> <destinationDir> [options]")
		return
	}

	component := in[1]
	destDir := in[2]
	in = in[3:] // Remove the first 2 arguments, keep the rest as they are the options

	tmpDir, err := ioutil.TempDir("", "sv")
	if err != nil {
		fmt.Println("FATAL: Could not create temporary directory")
		os.Exit(1)
	}

	// Check which flags are present
	for _, flag := range in {
		if flag == "--preserve-dependencies" {
			flags.preserveDependencies = true
		}
		if flag == "--save-log" {
			flags.saveLog = true
		}
	}

	// Build specific components
	switch component {
	case "game":
		buildGame(tmpDir, destDir)
	case "android-app":
		buildAndroid(tmpDir)
	}
}

func buildGame(tmpDir, destDir string) {
	buildDir := path.Join(tmpDir, "game")

	print("Game build started...")
	print("Using tmp directory", buildDir)

	// Clone the game repository
	print("Cloning latest version of the game repo")
	cloneCmd := exec.Command("git", "clone", "https://github.com/StellarVerge/game.git", buildDir)
	if err := cloneCmd.Run(); err != nil {
		print("Error cloning the game repo:", err)
		return
	}

	// Build the game
	print("Building game")
	buildCmd := exec.Command(path.Join(buildDir, "make"))
	if err := buildCmd.Run(); err != nil {
		print(err)
	}

	//Moving game to destination folder
	print("Moving game into destination directory", destDir)
	moveCmd := exec.Command("mv", path.Join(buildDir, "export"), destDir)
	if err := moveCmd.Run(); err != nil {
		print(err)
	}

	print("Build Finished")

}

func buildAndroid(dir string) {
	// TODO
}

// Implement own version of print so that things can be logged or just send to stdout
func print(s ...interface{}) {
	if flags.saveLog {
		log.Println(s)
	} else {
		fmt.Println(s)
	}
}
