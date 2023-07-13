package asciiart

import "testing"

func TestAsciiPrepJustify(t *testing.T) {
	type args struct {
		s string
		w int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AsciiPrepJustify(tt.args.s, tt.args.w)
		})
	}
}
