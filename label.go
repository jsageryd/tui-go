package tui

import (
	"image"
	"strings"

	termbox "github.com/nsf/termbox-go"
)

// Label is a widget to display read-only text.
type Label struct {
	text Text

	size image.Point
}

// NewLabel returns a new Label.
func NewLabel(text string) *Label {
	return &Label{
		text: Text(text),
	}
}

// Draw draws the label.
func (l *Label) Draw(p *Painter) {
	lines := strings.Split(string(l.text), "\n")
	for i, line := range lines {
		p.DrawText(0, i, Text(line))
	}
}

// Size returns the size of the label.
func (l *Label) Size() image.Point {
	return l.size
}

// SizeHint returns the recommended size for the label.
func (l *Label) SizeHint() image.Point {
	return image.Point{
		X: l.text.cols(),
		Y: l.text.rows(),
	}
}

// SizePolicy returns the default layout behavior.
func (l *Label) SizePolicy() (SizePolicy, SizePolicy) {
	return Minimum, Minimum
}

// Resize updates the size of the label.
func (l *Label) Resize(_ image.Point) {
	l.size = l.SizeHint()
}

func (l *Label) OnEvent(_ termbox.Event) {
}

func (l *Label) SetText(text string) {
	l.text = Text(text)
}
