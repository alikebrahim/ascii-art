package asciiart

import (
	"bytes"
	"fmt"
	"strings"
)

func PrinterRight(Text [][]byte, w int) {
	adj := len(Text[0])
	if (w/2)-(adj/2) < 0 {
		fmt.Println("Terminal size too small to print text") // to avoid error when text width is larger than Terminal width
		return
	}
	for _, line := range Text {
		fmt.Printf("%s", strings.Repeat(" ", w-adj))
		for _, char := range line {
			fmt.Printf("%s", string(char))
		}
	}
}

func AsciiPrepCR(s string, w int) {
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
	if Align == "right" {

		PrinterRight(Text, w)
	} else if Align == "center" {
		PrinterCenter(Text, w)
	}
}

func PrinterCenter(Text [][]byte, w int) {
	adj := len(Text[0])
	if (w/2)-(adj/2) < 0 {
		fmt.Println("Terminal size too small to print text") // to avoid error when text width is larger than Terminal width
		return
	}
	for _, line := range Text {
		fmt.Printf("%s", strings.Repeat(" ", (w/2)-(adj/2)))
		for _, char := range line {
			fmt.Printf("%s", string(char))
		}
	}
}

func AsciiPrepJustify(s string, w int) {
	words := StringToBytes(s)
	var edtwords []byte
	for _, line := range words {
		edtwords = append(edtwords, line...)
	}
	splitwords, wordsNum := bytes.Split(edtwords, []byte(" ")), len(bytes.Split(edtwords, []byte(" ")))
	llen := Count(s, w)
	Indx = nil
	Text = nil
	w_llen := w - llen
	wordsn := wordsNum - 1

	for i, slice := range splitwords {
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

		if i != wordsn {

			for j := 0; j < ((w - llen) / wordsn); j++ {
				Indx = append(Indx, 3200)
			}
		}

	}
	LineFmt(Indx)
	if w_llen < 0 {
		fmt.Println("Error: Stdout screen to small to display ASCII ART correctly aligned")
		return
	}
	Printer(Text)


}

func Count(s string, w int) (llen int) {
	words := StringToBytes(s)

	var edtwords []byte
	for _, line := range words {
		edtwords = append(edtwords, line...)
	}
	splitwords := bytes.Split(edtwords, []byte(" "))

	for i, slice := range splitwords {

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

	}
	LineFmt(Indx)
	return len(Text[0])
}
