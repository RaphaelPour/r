package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	BuildDate    string
	BuildVersion string
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "--version" {
		fmt.Println("BuildVersion:", BuildVersion)
		fmt.Println("BuildDate:", BuildDate)
		return
	}

	if len(os.Args) != 3 {
		fmt.Println("usage: r <path> <new_filename>")
		return
	}

	path := filepath.Dir(os.Args[1])
	newPath := filepath.Join(path, os.Args[2])

	if err := os.Rename(os.Args[1], newPath); err != nil {
		fmt.Printf("error renaming %s to %s: %s\n", os.Args[1], newPath, err)
		return
	}

	fmt.Printf("%s -> %s\n", os.Args[1], newPath)
}
