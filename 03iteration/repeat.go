package _3iteration

import (
	"fmt"
	"strings"
)

func Repeat(letterToRepeat string, times int) string {
	var repeat string
	for i := 0; i < times; i++ {
		repeat += letterToRepeat
	}
	return repeat
}

func Builder() string {
	var b strings.Builder
	for i := 1; i <= 3; i++ {
		fmt.Fprintf(&b, "%d...", i)
	}
	b.WriteString("repeat")
	return b.String()
}
