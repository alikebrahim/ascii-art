package main

import (
	"asciiart/asciiart"
	"bytes"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	flag.Parse()
	//fmt.Printf("Justify: %v ** Color: %v ** Output: %v\n", asciiart.Justify, asciiart.Color, asciiart.Output)

	// CMD line args handling
	args := os.Args[1:]
	lArgs := len(args)
	if lArgs < 2 {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("EX: go run . something standard")
		return
	}

	bannerFile := args[lArgs-1] + ".txt"

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
	s := args[lArgs-2]

	if lArgs > 2 {
		if asciiart.Align != ""{
			if len(args) != 3 {
				fmt.Println("Usage: go run . [FLAG] [OPTION] [STRING] [BANNER]")
				fmt.Println("EX: go run . -a center something standard")
				return
			} 

	}


		// Get terminal width
		cmd := exec.Command("tput", "cols")
		cmd.Stdin = os.Stdin
		out, err := cmd.Output()
		if err != nil {
			fmt.Println("Error getting terminal width")
			return
		}
		n := bytes.Split(out, []byte("\n"))[0]
		width, err := strconv.Atoi(string(n))
		if err != nil {
			fmt.Println("Error getting terminal width")
			return
		}
		if asciiart.Align == "justify" {
			asciiart.AsciiPrepJustify(s, width)
			return
		} else if asciiart.Align == "right" || asciiart.Align == "center" {

			asciiart.AsciiPrepCR(s, width)
			return
		} else if asciiart.Align == "left" {
			asciiart.AsciiPrep(s)
			return
		} else {
			fmt.Println("Usage: go run .  [OPTION] [STRING] [BANNER]")
			fmt.Println("Example: go run . --align=right  something  standard")
			return
		}
	}

	asciiart.AsciiPrep(s)

}
