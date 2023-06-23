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

	// Testing banner
	reader, err := fs.ReadFile(os.DirFS("./banners"), "standard.txt")
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

	// Text to print
	withNewLine := asciiart.NewLineCheck(args[1])
	if withNewLine {
		asciiart.SepAtNL(args[1])
		return

	}
	// Test string preparation, formating and printing
	s := args[1]
	asciiart.Indexer(s)
	asciiart.LineFmt(asciiart.Indx)
	asciiart.Printer(asciiart.Words)

}
