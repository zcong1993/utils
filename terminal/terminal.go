package terminal

import (
	"fmt"
	"github.com/zcong1993/utils/colors"
)

const space = " "

// Success log success message with colors
func Success(str string) {
	println()
	fmt.Printf("%s%s%s", colors.Green("SUCCESS"), space, str)
	println()
}

// Fail log error message with colors
func Fail(str string) {
	println()
	fmt.Printf("%s%s%s", colors.Red("ERROR"), space, str)
	println()
}

// Info log info message with colors
func Info(str string) {
	println()
	fmt.Printf("%s%s%s", colors.Cyan("INFO"), space, str)
	println()
}

// Pad helper.
func Pad() func() {
	println()
	return func() {
		println()
	}
}

// LogPad outputs a log message with padding.
func LogPad(msg string) {
	defer Pad()()
	fmt.Println(msg)
}

// LogErrPad outputs error message with color and pad
func LogErrPad(err error) {
	LogPad(colors.Red("ERROR") + space + colors.Purple(err.Error()))
}

// LogErrPad outputs success message with color and pad
func LogSuccessPad(msg string) {
	LogPad(colors.Green("SUCCESS") + space + msg)
}
