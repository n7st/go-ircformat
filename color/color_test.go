package color

import (
	"fmt"
	"testing"
)

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
