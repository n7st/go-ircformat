// Package color applies standard IRC color strings to text.
package color

import "fmt"

const (
	ColorWhite      = 0
	ColorBlack      = 1
	ColorBlue       = 2
	ColorGreen      = 3
	ColorRed        = 4
	ColorBrown      = 5
	ColorMagenta    = 6
	ColorOrange     = 7
	ColorYellow     = 8
	ColorLightGreen = 9
	ColorCyan       = 10
	ColorLightCyan  = 11
	ColorLightBlue  = 12
	ColorPink       = 13
	ColorGrey       = 14
	ColorLightGrey  = 15
)

// White makes text white.
func White(input string) string {
	return applyColor(input, ColorWhite)
}

// Black makes text black.
func Black(input string) string {
	return applyColor(input, ColorBlack)
}

// Blue makes text blue.
func Blue(input string) string {
	return applyColor(input, ColorBlue)
}

// Green makes text green.
func Green(input string) string {
	return applyColor(input, ColorGreen)
}

// Red makes text red.
func Red(input string) string {
	return applyColor(input, ColorRed)
}

// Brown makes text brown.
func Brown(input string) string {
	return applyColor(input, ColorBrown)
}

// Magenta makes text magenta.
func Magenta(input string) string {
	return applyColor(input, ColorMagenta)
}

// Orange makes text orange.
func Orange(input string) string {
	return applyColor(input, ColorOrange)
}

// Yellow makes text yellow.
func Yellow(input string) string {
	return applyColor(input, ColorYellow)
}

// LightGreen makes text light green.
func LightGreen(input string) string {
	return applyColor(input, ColorLightGreen)
}

// Cyan makes text cyan.
func Cyan(input string) string {
	return applyColor(input, ColorCyan)
}

// LightCyan makes text light cyan.
func LightCyan(input string) string {
	return applyColor(input, ColorLightCyan)
}

// LightBlue makes text light blue.
func LightBlue(input string) string {
	return applyColor(input, ColorLightBlue)
}

// Pink makes text pink.
func Pink(input string) string {
	return applyColor(input, ColorPink)
}

// Grey makes text grey.
func Grey(input string) string {
	return applyColor(input, ColorGrey)
}

// LightGrey makes text light grey.
func LightGrey(input string) string {
	return applyColor(input, ColorLightGrey)
}

func applyColor(input string, color rune) string {
	return fmt.Sprintf("\x03%d%s\x03", color, input)
}
