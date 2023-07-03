package main

import (
	"asciiart/asciiart"
	"fmt"
	"io/fs"
	"os"
)

func main() {

	// TODO: CMD line args handling
	args := os.Args
	if len(args) != 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("EX: go run . something standard")
		return
	}
	bannerFile := args[2] + ".txt"

	// Testing banner
	reader, err := fs.ReadFile(os.DirFS("./banners"), bannerFile)
	if err != nil {
		fmt.Println("Error Opening File")
	}

	// Formate banner file
	asciiart.BannerFmt(reader)

	// Handle "" && "\n"
	if args[1] == "" {
		return
	} else if args[1] == "\\n" {
		fmt.Println()
		return
	}

	// Test string preparation, formating and printing
	s := args[1]
	asciiart.AsciiPrep(s)

}
