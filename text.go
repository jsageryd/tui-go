package tui

import (
	"strings"
	"unicode/utf8"

	runewidth "github.com/mattn/go-runewidth"
)

type Text string

func (t Text) cols() int {
	width := 0
	for _, line := range strings.Split(string(t), "\n") {
		if w := runewidth.StringWidth(string(line)); w > width {
			width = w
		}
	}
	return width
}

func (t Text) rows() int {
	return strings.Count(string(t), "\n") + 1
}

// trimRightLen returns t with n runes trimmed off
func trimRightLen(t Text, n int) Text {
	if n <= 0 {
		return t
	}
	c := utf8.RuneCountInString(string(t))
	runeCount := 0
	var i int
	for i, _ = range t {
		if runeCount >= c-n {
			break
		}
		runeCount++
	}
	return t[:i]
}
