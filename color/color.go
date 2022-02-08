package color

import (
	"fmt"
)

type Color string

const (
	Black       Color = "\033[0;30m"
	DarkGray    Color = "\033[1;30m"
	Red         Color = "\033[0;31m"
	LightRed    Color = "\033[1;31m"
	Green       Color = "\033[0;32m"
	LightGreen  Color = "\033[1;32m"
	Orange      Color = "\033[0;33m"
	Yellow      Color = "\033[1;33m"
	Blue        Color = "\033[0;34m"
	LightBlue   Color = "\033[1;34m"
	Purple      Color = "\033[0;35m"
	LightPurple Color = "\033[1;35m"
	Cyan        Color = "\033[0;36m"
	LightCyan   Color = "\033[1;36m"
	LightGray   Color = "\033[0;37m"
	White       Color = "\033[1;37m"
	NoColor     Color = "\033[0m"
)

func (c Color) String() string {
	return string(c)
}

func Sprintf(c Color, s string, args ...interface{}) string {
	return fmt.Sprintf(fmt.Sprintf("%s%s%s", c, s, NoColor), args...)
}
