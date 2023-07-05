package main

import (
	"asciiart/asciiart"
	"flag"
	"fmt"
	"io/fs"
	"os"
)

func main() {

	// CMD line args handling
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
		return
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
	var s string
	if flag.NFlag() == 0 {
		s = args[1]
	} else {
		s = args[3]
	}

	asciiart.AsciiPrep(s)

}
