package asciiart

import (
	"fmt"
	"strings"
)


func PrinterJustifyRight(Text [][]byte, w int) {
	adj:= len(Text[0]) 
	for _, line := range Text {
		fmt.Printf("%s", strings.Repeat(" ", w-adj))
		for _, char := range line {
			fmt.Printf("%s", string(char))
		}
	}
}

func AsciiPrepJustify(s string, w int) {
	words := StringToBytes(s)
	for i, slice := range words {
		if slice[0] == 10 && i == 0 {
			Text = append(Text, []byte("\n"))
			Indx = nil
			continue
		}
		if slice[0] == 10 {
			if words[i-1][0] == 10 {
				Text = append(Text, []byte("\n"))
				Indx = nil
				continue
			} else {
				Indx = nil
				continue
			}

		}
		for _, char := range slice {
			Indexer(char)
		}
		LineFmt(Indx)

	}
	if Justify == "right" {

		PrinterJustifyRight(Text, w)
	} else if Justify == "center" {
		PrinterJustifyCenter(Text, w)	
	}
}

func PrinterJustifyCenter(Text [][]byte, w int) {
	adj:= len(Text[0]) -1
	for _, line := range Text {
		fmt.Printf("%s", strings.Repeat(" ", (w/2)-(adj/2)))
		for _, char := range line {
			fmt.Printf("%s", string(char))
		}
	}
}
