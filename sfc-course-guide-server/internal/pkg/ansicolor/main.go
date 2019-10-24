package ansicolor

import (
	"fmt"
)

// Color represents a ANSI escape code supported color
// range [30, 37] and [90, 97]
type Color uint8

// Colors
const (
	Black = Color(iota + 30)
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	BrightBlack = Color(iota + 90 - 8)
	BrightRed
	BrightGreen
	BrightYellow
	BrightBlue
	BrightMagenta
	BrightCyan
	BrightWhite
)

// ToColor return the string s that can be printed in frontground color c
func ToColor(c Color, s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", c, s)
}

// ToBlack return the string s that can be printed in frontground color Black
func ToBlack(s string) string {
	return ToColor(Black, s)
}

// ToRed return the string s that can be printed in frontground color Red
func ToRed(s string) string {
	return ToColor(Red, s)
}

// ToGreen return the string s that can be printed in frontground color Green
func ToGreen(s string) string {
	return ToColor(Green, s)
}

// ToYellow return the string s that can be printed in frontground color Yellow
func ToYellow(s string) string {
	return ToColor(Yellow, s)
}

// ToBlue return the string s that can be printed in frontground color Blue
func ToBlue(s string) string {
	return ToColor(Blue, s)
}

// ToMagenta return the string s that can be printed in frontground color Magenta
func ToMagenta(s string) string {
	return ToColor(Magenta, s)
}

// ToCyan return the string s that can be printed in frontground color Cyan
func ToCyan(s string) string {
	return ToColor(Cyan, s)
}

// ToWhite return the string s that can be printed in frontground color White
func ToWhite(s string) string {
	return ToColor(White, s)
}

// ToBrightBlack return the string s that can be printed in frontground color BrightBlack
func ToBrightBlack(s string) string {
	return ToColor(BrightBlack, s)
}

// ToBrightRed return the string s that can be printed in frontground color BrightRed
func ToBrightRed(s string) string {
	return ToColor(BrightRed, s)
}

// ToBrightGreen return the string s that can be printed in frontground color BrightGreen
func ToBrightGreen(s string) string {
	return ToColor(BrightGreen, s)
}

// ToBrightYellow return the string s that can be printed in frontground color BrightYellow
func ToBrightYellow(s string) string {
	return ToColor(BrightYellow, s)
}

// ToBrightBlue return the string s that can be printed in frontground color BrightBlue
func ToBrightBlue(s string) string {
	return ToColor(BrightBlue, s)
}

// ToBrightMagenta return the string s that can be printed in frontground color BrightMagenta
func ToBrightMagenta(s string) string {
	return ToColor(BrightMagenta, s)
}

// ToBrightCyan return the string s that can be printed in frontground color BrightCyan
func ToBrightCyan(s string) string {
	return ToColor(BrightCyan, s)
}

// ToBrightWhite return the string s that can be printed in frontground color BrightWhite
func ToBrightWhite(s string) string {
	return ToColor(BrightWhite, s)
}
