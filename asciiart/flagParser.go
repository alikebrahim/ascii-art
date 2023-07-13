package asciiart

import "flag"

var Align string
var Color string
var Output bool

func init() {
	const (
		defaultAlign = ""
		usageAlign   = "Align text"

		defaultColor   = ""
		usageColor     = "Color text"

		defaultOutput  = false
		usageOutput    = "Output text"
	)
	
	flag.StringVar(&Align, "align", defaultAlign, usageAlign)
	flag.StringVar(&Align, "a", defaultAlign, usageAlign+" shorthand")

	flag.StringVar(&Color, "color", defaultColor, usageColor)
	flag.StringVar(&Color, "c", defaultColor, usageColor+" shorthand")

	flag.BoolVar(&Output, "output", defaultOutput, usageOutput)
	flag.BoolVar(&Output, "o", defaultOutput, usageOutput+" shorthand")
}
