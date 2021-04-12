package color

import (
	"fmt"
	"testing"
)

func TestColor(t *testing.T) {
	tests := []struct {
		name   string
		colors *Colors
		want   string
	}{
		{
			name:   "Foreground red",
			colors: &Colors{Foreground: New(ColorRed)},
			want:   "\x034",
		},
		{
			name: "Foreground red, background black",
			colors: &Colors{
				Foreground: New(ColorRed),
				Background: New(ColorBlack),
			},
			want: "\x034,1",
		},
		{
			name:   "Background black (should not color)",
			colors: &Colors{Background: New(ColorBlack)},
			want:   "",
		},
		{
			name:   "No colors (should not color)",
			colors: &Colors{},
			want:   "",
		},
	}

	for _, tt := range tests {
		got := Color("Hello, world", tt.colors)

		if got != fmt.Sprintf("%sHello, world\x03", tt.want) {
			t.Errorf("Color() got = %v, want %v", got, tt.want)
			t.Fail()
		}
	}
}

func TestColors(t *testing.T) {
	colorFunctions := map[rune]func(string) string{
		0:  White,
		1:  Black,
		2:  Blue,
		3:  Green,
		4:  Red,
		5:  Brown,
		6:  Magenta,
		7:  Orange,
		8:  Yellow,
		9:  LightGreen,
		10: Cyan,
		11: LightCyan,
		12: LightBlue,
		13: Pink,
		14: Grey,
		15: LightGrey,
	}

	for colorCode, colorFunc := range colorFunctions {
		got := colorFunc("Hello, world")
		want := fmt.Sprintf("\x03%dHello, world\x03", colorCode)

		if got != want {
			t.Fail()
		}
	}
}
