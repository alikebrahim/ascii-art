package asciiart

import (
	"bytes"
	"fmt"
	"strings"
)

var Chars [][]byte
var Words []byte
var Indx []int

// Create a new [][]byte from banner into Chars
// // case standard & shadow banners
func BannerFmt(r []byte) {
	split := bytes.Split(r, []byte("\n")) // create banner split on new lines

	for _, line := range split {
		if len(line) == 0 {
			continue
		}
		Chars = append(Chars, line)
	}
}

func BannerFmtTT(r []byte) {
	split := bytes.Split(r, []byte("\r")) // create banner split on carriage returns

	for _, line := range split {
		if len(line) == 0 {
			continue
		}
		Chars = append(Chars, line)
	}
	newBannerFMT := ReformatTT(Chars) // copy banner split on carriage return from Chars (type [][]byte) into []byte
	Chars = nil // empty Chars (prepare to repopulate)
	BannerFmt(newBannerFMT) // repopulate Chars with banner split on new lines
}

func ReformatTT(chars [][]byte) []byte {
	var ssB []byte
	for _, line := range chars {
		ssB = append(ssB, line...)
	}
	return ssB
}

// locate first line index for each char in string (word)
func Indexer(s string) {
	for _, char := range s {
		s, _ := Location(char)
		Indx = append(Indx, s)
	}
}

// Locate StartingIndex:EndingIndex for each char
func Location(r rune) (s int, e int) {
	locS := (int(r) - 32) * 8
	locE := locS + 8

	//fmt.Println("r: ", r, " s: ", locS, " e: ", locE)
	return locS, locE
}

// create a new [][]byte of string (word) for printing
// // by appending each line of char byte slice into one []byte
func LineFmt(indx []int) {
	for i := 0; i < 8; i++ {
		var line []byte
		for _, item := range indx {
			slice := Chars[item+i]
			line = append(line, slice...)
		}
		line = append(line, byte('\n'))
		Words = append(Words, line...)
	}
}

func Printer(Text []byte) {
	for _, char := range Text {
		fmt.Printf("%s", string(char))
		// for _, char := range line {
		// }
	}
}

func NewLineCheck(s string) bool {
	for i, char := range s {
		if char == '\\' {
			if s[i+1] == 'n' {
				return true
			}
		}

	}
	return false
}

func SepAtNL(s string) {
	separated := strings.Split(s, "\\n")
	for _, item := range separated {
		if item == "" {
			continue
		}
		Indexer(item)
		LineFmt(Indx)
		Indx = nil
	}
	Printer(Words)
}
