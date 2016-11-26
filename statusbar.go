package tui

import (
	"image"

	termbox "github.com/nsf/termbox-go"
)

// StatusBar is a widget to display status information.
type StatusBar struct {
	size image.Point

	text     Text
	permText Text

	fg, bg termbox.Attribute
}

// NewStatusBar returns a new StatusBar.
func NewStatusBar(text string) *StatusBar {
	return &StatusBar{
		text:     Text(text),
		permText: "",
	}
}

// Draw draws the spacer.
func (b *StatusBar) Draw(p *Painter) {
	s := b.Size()
	p.SetBrush(b.fg, b.bg)
	p.FillRect(0, 0, s.X, 1)
	p.DrawText(0, 0, b.text)
	p.DrawText(s.X-len(b.permText), 0, b.permText)
	p.SetBrush(termbox.ColorDefault, termbox.ColorDefault)
}

// Size returns the size of the spacer.
func (b *StatusBar) Size() image.Point {
	return b.size
}

// SizeHint returns the recommended size for the spacer.
func (b *StatusBar) SizeHint() image.Point {
	return image.Point{0, 1}
}

// SizePolicy returns the default layout behavior.
func (b *StatusBar) SizePolicy() (SizePolicy, SizePolicy) {
	return Expanding, Minimum
}

// Resize updates the size of the spacer.
func (b *StatusBar) Resize(size image.Point) {
	b.size = size
}

func (b *StatusBar) OnEvent(_ termbox.Event) {
}

func (b *StatusBar) SetBrush(fg, bg termbox.Attribute) {
	b.fg = fg
	b.bg = bg
}

func (b *StatusBar) SetText(text string) {
	b.text = Text(text)
}

func (b *StatusBar) SetPermanentText(text string) {
	b.permText = Text(text)
}
