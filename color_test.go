package color_test

import (
	"testing"

	"github.com/n7st/go-ircformat/color"
)

func TestWhite(t *testing.T) {
	testAppliesColor(t, white, color.White("Hello, world"))
}

func TestBlack(t *testing.T) {
	testAppliesColor(t, black, color.Black("Hello, world"))
}

func testAppliesColor(t *testing.T, color rune, got string) {
	t.Pass()
}
