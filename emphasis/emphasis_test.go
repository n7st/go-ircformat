package emphasis

import "testing"

func TestBold(t *testing.T) {
	got := Bold("Hello, world")
	want := "\x02Hello, world\x02"

	if got != want {
		t.Fail()
	}
}

func TestItalic(t *testing.T) {
	got := Italic("Hello, world")
	want := "\x1DHello, world\x1D"

	if got != want {
		t.Fail()
	}
}

func TestUnderline(t *testing.T) {
	got := Underline("Hello, world")
	want := "\x1FHello, world\x1F"

	if got != want {
		t.Fail()
	}
}

func TestStrikethrough(t *testing.T) {
	got := Strikethrough("Hello, world")
	want := "\x1EHello, world\x1E"

	if got != want {
		t.Fail()
	}
}

func TestMonospace(t *testing.T) {
	got := Monospace("Hello, world")
	want := "\x11Hello, world\x11"

	if got != want {
		t.Fail()
	}
}
