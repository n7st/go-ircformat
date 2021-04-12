// Package color applies standard IRC color strings to text.
package color

import "fmt"

const (
	white      = 0
	black      = 1
	blue       = 2
	green      = 3
	red        = 4
	brown      = 5
	magenta    = 6
	orange     = 7
	yellow     = 8
	lightGreen = 9
	cyan       = 10
	lightCyan  = 11
	lightBlue  = 12
	pink       = 13
	grey       = 14
	lightGrey  = 15
)

// White makes text white.
func White(input string) string {
	return applyColor(input, white)
}

// Black makes text black.
func Black(input string) string {
	return applyColor(input, black)
}

// Blue makes text blue.
func Blue(input string) string {
	return applyColor(input, blue)
}

// Green makes text green.
func Green(input string) string {
	return applyColor(input, green)
}

// Red makes text red.
func Red(input string) string {
	return applyColor(input, red)
}

// Brown makes text brown.
func Brown(input string) string {
	return applyColor(input, brown)
}

// Magenta makes text magenta.
func Magenta(input string) string {
	return applyColor(input, magenta)
}

// Orange makes text orange.
func Orange(input string) string {
	return applyColor(input, orange)
}

// Yellow makes text yellow.
func Yellow(input string) string {
	return applyColor(input, yellow)
}

// LightGreen makes text light green.
func LightGreen(input string) string {
	return applyColor(input, lightGreen)
}

// Cyan makes text cyan.
func Cyan(input string) string {
	return applyColor(input, cyan)
}

// LightCyan makes text light cyan.
func LightCyan(input string) string {
	return applyColor(input, lightCyan)
}

// LightBlue makes text light blue.
func LightBlue(input string) string {
	return applyColor(input, lightBlue)
}

// Pink makes text pink.
func Pink(input string) string {
	return applyColor(input, pink)
}

// Grey makes text grey.
func Grey(input string) string {
	return applyColor(input, grey)
}

// LightGrey makes text light grey.
func LightGrey(input string) string {
	return applyColor(input, lightGrey)
}

func applyColor(input string, color rune) string {
	return fmt.Sprintf("\x03%i%s\x03")
}
