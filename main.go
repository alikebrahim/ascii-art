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
	"strings"
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

	var bannerFile string

	if lArgs == 2 {
		if args[lArgs-1] != "standard" || args[lArgs-1] != "thinkertoy" || args[lArgs-1] != "shadow" || args[lArgs-1] != "banner-3D" {
			bannerFile = "standard.txt"
		}
	} else {
		bannerFile = args[lArgs-1] + ".txt"
	}

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
	if lArgs == 2 {
		s = args[lArgs-1]
	} else {

		s = args[lArgs-2]
	}

	// Flag format issue needs to be fixed: only --flag format should be accepted
	if asciiart.Align != "" {

		if len(args[0]) < 8 || args[0][:8] != "--align=" {
			fmt.Println("Usage: go run . [FLAG] [STRING] [BANNER]")
			fmt.Println("EX: go run . --align=right something standard")
			return
		} else if args[0][:8] != "--align=" {
			fmt.Println("Usage: go run . [FLAG] [STRING] [BANNER]")
			fmt.Println("EX: go run . --align=right something standard")
			return
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
			splitS := strings.Split(s, "\\n")
			// fmt.Println(splitS, len(splitS))
			// fmt.Printf("%T\n",splitS[0])
			// return

			for _, item := range splitS {

				asciiart.AsciiPrepJustify(item, width)
			}
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
