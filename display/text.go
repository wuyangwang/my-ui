package display

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

// TextType names the semantic color of a text.
type TextType int

const (
	// TextDefault takes the theme's own ink.
	TextDefault TextType = iota
	// TextSecondary uses the muted ink.
	TextSecondary
	// TextSuccess uses the success color.
	TextSuccess
	// TextWarning uses the warning color.
	TextWarning
	// TextDanger uses the danger color.
	TextDanger
	// TextInfo uses the info color.
	TextInfo
)

// TextOptions configures NewText. The zero value is body text in the
// theme's own ink.
type TextOptions struct {
	// Type colors the text by role.
	Type TextType
	// Variant is the shared vocabulary for the same roles; it takes effect
	// when Type is TextDefault.
	Variant Variant
	// Heading sets a heading's size from 1 to 5; 0 is body text.
	Heading int
	// Size applies the shared size vocabulary when FontSize is zero.
	Size Size
	// FontSize overrides the size the heading or the theme gives.
	FontSize float32
	// Strong, Italic, Underline and Delete set the weight, slant and lines.
	Strong, Italic, Underline, Delete bool
	// Code sets the text in the monospaced face.
	Code bool
	// Mark highlights the text in the theme's accent.
	Mark bool
	// Disabled dims the text whatever its type.
	Disabled bool
	// Rows shows at most that many lines; 0 shows them all.
	Rows int
}

// Text is a styled run of text. It embeds ui.Element, so the mygo
// text methods remain available for anything the options do not cover.
// Fluent methods paint the same element in place; they never build a
// second text in the frame.
type Text struct {
	ui.Element
	c *ui.Context
	o TextOptions
	e ui.Element
}

// NewText shows s in my-ui's type styles. The options are optional, so
// NewText(c, s) is plain body text.
//
//	display.NewText(c, "Hello").Strong()
//	display.NewText(c, "Saved", display.TextOptions{Type: TextSuccess})
func NewText(c *ui.Context, s string, opts ...TextOptions) *Text {
	x := &Text{c: c, e: ui.Text(c, s)}
	x.Element = x.e
	if len(opts) > 0 {
		x.o = opts[0]
	}
	x.apply()
	return x
}

var textHeadingScale = [6]float32{1, 2.5, 2, 1.5, 1.25, 1}

func (x *Text) apply() {
	t := x.c.Theme()
	ink := t.Text
	role := x.o.Type
	if role == TextDefault {
		switch x.o.Variant {
		case Secondary:
			role = TextSecondary
		case Success:
			role = TextSuccess
		case Warning:
			role = TextWarning
		case Destructive:
			role = TextDanger
		case Info, Link, Primary:
			role = TextInfo
		}
	}
	switch role {
	case TextSecondary:
		ink = t.TextMuted
	case TextSuccess:
		ink = SemanticColor(t, SemanticSuccess)
	case TextWarning:
		ink = SemanticColor(t, SemanticWarning)
	case TextDanger:
		ink = t.Danger
	case TextInfo:
		ink = SemanticColor(t, SemanticInfo)
	}
	if x.o.Heading >= 1 && x.o.Heading <= 5 {
		x.e.FontSize(TextSize(t, textHeadingScale[x.o.Heading])).FontWeight(700).LineHeight(1.25)
	}
	size := x.o.FontSize
	if size == 0 {
		switch x.o.Size {
		case SizeXs:
			size = TextSize(t, TextXs)
		case SizeSm:
			size = TextSize(t, TextSm)
		case SizeLg:
			size = TextSize(t, TextLg)
		}
	}
	if size > 0 {
		x.e.FontSize(size)
	}
	if x.o.Strong || (x.o.Heading >= 1 && x.o.Heading <= 5) {
		x.e.FontWeight(700)
	}
	if x.o.Italic {
		x.e.Italic()
	}
	if x.o.Underline {
		x.e.Underline()
	}
	if x.o.Delete {
		x.e.Strikethrough()
	}
	if x.o.Code {
		x.e.Font("monospace")
	}
	if x.o.Mark {
		x.e.Background(t.Accent)
		ink = t.AccentText
	}
	if x.o.Disabled {
		ink = t.TextMuted
		x.e.Cursor(ui.CursorNotAllowed)
	}
	switch {
	case x.o.Rows == 1:
		x.e.SingleLine()
	case x.o.Rows > 1:
		x.e.MaxLines(x.o.Rows)
	}
	x.e.TextColor(ink)
}

// Size applies the shared size s and returns the text for chaining.
func (x *Text) Size(s Size) *Text {
	x.o.Size = s
	x.apply()
	return x
}

// Variant applies a semantic variant and returns the text for chaining.
func (x *Text) Variant(v Variant) *Text {
	x.o.Variant = v
	x.apply()
	return x
}

// Secondary uses the muted ink.
func (x *Text) Secondary() *Text { return x.Variant(Secondary) }

// Success uses the success color.
func (x *Text) Success() *Text { return x.Variant(Success) }

// Warning uses the warning color.
func (x *Text) Warning() *Text { return x.Variant(Warning) }

// Danger uses the danger color.
func (x *Text) Danger() *Text { return x.Variant(Destructive) }

// Info uses the info color.
func (x *Text) Info() *Text { return x.Variant(Info) }

// Heading sets the heading level n (1 through 5).
func (x *Text) Heading(n int) *Text {
	x.o.Heading = n
	x.apply()
	return x
}

// Strong makes the text bold.
func (x *Text) Strong() *Text {
	x.o.Strong = true
	x.apply()
	return x
}

// Italic slants the text.
func (x *Text) Italic() *Text {
	x.o.Italic = true
	x.apply()
	return x
}

// Code sets the monospaced face.
func (x *Text) Code() *Text {
	x.o.Code = true
	x.apply()
	return x
}

// Mark highlights the text in the accent color.
func (x *Text) Mark() *Text {
	x.o.Mark = true
	x.apply()
	return x
}

// Disabled dims the text and returns it for chaining.
func (x *Text) Disabled(off bool) *Text {
	x.o.Disabled = off
	x.apply()
	return x
}

// Rows limits the text to n lines with an ellipsis.
func (x *Text) Rows(n int) *Text {
	x.o.Rows = n
	x.apply()
	return x
}
