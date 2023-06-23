package main

import (
	"asciiart/asciiart"
	"fmt"
	"io/fs"
	"os"
)

func main() {

	// TODO: CMD line args handling

	// Testing banner
	reader, err := fs.ReadFile(os.DirFS("./banners"), "standard.txt")
	if err != nil {
		fmt.Println("Error Opening File")
	}

	// Formate banner file
	asciiart.BannerFmt(reader)

	// Test string
	s := "Hello$%^Ali123456789"

	// Test string preparation, formating and printing
	asciiart.Indexer(s)
	fmtTxt := asciiart.LineFmt(asciiart.Indx)
	asciiart.Printer(fmtTxt)

}
