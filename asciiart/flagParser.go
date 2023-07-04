package asciiart

import "flag"

var Justify string
var Color string
var Output bool

func init() {
	const (
		defaultJustify = "left"
		usageJustify   = "Justify text"

		defaultColor   = "white"
		usageColor     = "Color text"

		defaultOutput  = false
		usageOutput    = "Output text"
	)
	
	flag.StringVar(&Justify, "justify", defaultJustify, usageJustify)
	flag.StringVar(&Justify, "j", defaultJustify, usageJustify+" shorthand")

	flag.StringVar(&Color, "color", defaultColor, usageColor)
	flag.StringVar(&Color, "c", defaultColor, usageColor+" shorthand")

	flag.BoolVar(&Output, "output", defaultOutput, usageOutput)
	flag.BoolVar(&Output, "o", defaultOutput, usageOutput+" shorthand")
}
