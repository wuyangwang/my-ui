package forms

import (
	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

// AutoComplete is mygo's Autocomplete with my-ui's shared field look.
// Filtering, suggestion choice and the keyboard stay with mygo;
// value is the app's pointer.
type AutoComplete struct {
	ui.Element
	c       *ui.Context
	variant Variant
	size    Size
	width   float32
}

// NewAutoComplete builds an input with suggestions drawn from the
// app's slice.
func NewAutoComplete(c *ui.Context, value *string, suggestions []string) *AutoComplete {
	a := &AutoComplete{c: c}
	a.Element = ui.Autocomplete(c, value, suggestions).Transition(StateTransition())
	a.apply()
	return a
}

// Variant applies the shared look of v.
func (a *AutoComplete) Variant(v Variant) *AutoComplete {
	a.variant = v
	a.apply()
	return a
}

// Size picks the text and padding scale.
func (a *AutoComplete) Size(sz Size) *AutoComplete {
	a.size = sz
	a.apply()
	return a
}

// Placeholder shows v while the value is empty.
func (a *AutoComplete) Placeholder(v string) *AutoComplete {
	a.Element = a.Element.Placeholder(v)
	return a
}

// Width pins the field's width in DIPs; zero fills the parent.
func (a *AutoComplete) Width(w float32) *AutoComplete {
	a.width = w
	a.apply()
	return a
}

func (a *AutoComplete) apply() {
	t := a.c.Theme()
	look := FrameLookOf(a.variant, t)
	if !look.Native {
		a.Element = a.Element.Background(look.Bg).Border(look.Width, look.Border)
		if look.HasText {
			a.Element = a.Element.TextColor(look.Text)
		}
	}
	v, h := FramePadding(t, a.size)
	a.Element = a.Element.Radius(t.Radius).Padding(v, h).FontSize(FrameFontSize(t, a.size))
	if a.width > 0 {
		a.Element = a.Element.Width(a.width)
	} else {
		a.Element = a.Element.FillWidth()
	}
}
