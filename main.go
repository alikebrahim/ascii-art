package main

import (
	"asciiart/asciiart"
	"flag"
	"fmt"
	"io/fs"
	"os"
)

func main() {

	// TODO: CMD line args handling
	args := os.Args
	flag.Parse()

	fmt.Println(len(args))
	if len(args) != 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("EX: go run . something standard")
		return
	}

	// fmt.Println(flag.NArg(), flag.NFlag())
	// fmt.Println("*Justify: ", asciiart.Justify, " *Color: ", asciiart.Color, "*Output: ", asciiart.Output)
	// os.Exit(1)

	// if flag.NFlag() == 0 && len(args) != 3 {
	// 	fmt.Println("Entered @ 0")
	// 	fmt.Println("Usage: go run . [STRING] [BANNER]")
	// 	fmt.Println("EX: go run . something standard")
	// 	return
	// } else if flag.NFlag() != 0 && len(args) != 5 {
	// 	fmt.Println("Entered @ > 0")
	// 	fmt.Println("Usage: go run . [STRING] [BANNER]")
	// 	fmt.Println("EX: go run . something standard")
	// 	return
	// }

	bannerFile := args[2] + ".txt"
	// var bannerFile string
	// if flag.NFlag() == 0 {
	// 	bannerFile = args[2] + ".txt"
	// } else {
	// 	bannerFile = args[4] + ".txt"
	// }

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
	var s string
	if flag.NFlag() == 0 {
		s = args[1]
	} else {
		s = args[3]
	}

	asciiart.AsciiPrep(s)

}
