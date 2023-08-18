package logger

import (
	"fmt"
	"time"
)

type Logger interface {
	StrlogV(text string, data string, s time.Time)
	StrlogE(text string, data string, s time.Time)
	StrlogR(text string, data string, s time.Time)
}

var x, r, g, bg, rb, gr, u, clr, yellow, red, prp = "\u001b[30;1m", "\033[39m", "\033[32m", "\u001b[34;1m", "\u001b[0m", "\u001b[30;1m", "\u001b[4m", "\033[36m", "\033[33m", "\u001B[31m", "\033[35m"

func StrlogV(text string, data string, s time.Time) {
	e := time.Since(s)
	fmt.Printf("[%s%sms%s] [%s✓%s]%s: %s%s%s\n", bg, e.String()[:3], rb, g, r, text, gr, data, rb+r)
}

func StrlogE(text string, data string, s time.Time) {
	e := time.Since(s)
	fmt.Printf("[%s%sms%s] [%sX%s]%s: %s%s%s\n", bg, e.String()[:3], rb, red, r, text, gr, data, rb+r)
}

func StrlogR(text string, data string, s time.Time) {
	e := time.Since(s)
	fmt.Printf("[%s%sms%s] [%s-%s]%s: %s%s%s\n", bg, e.String()[:3], rb, yellow, r, text, gr, data, rb+r)

}
