package myui

import (
	"strings"

	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/icons"
)

// ComboBox is a drop-down whose trigger is a text field: what is typed
// narrows the options down, and a choice takes the place of the text.
//
//	myui.NewComboBox(c, &query, []myui.SelectOption{
//		{Label: "简体中文", Value: "zh-Hans"},
//		{Label: "English", Value: "en"},
//	}).Outline().Placeholder("搜索语言…")
//
// The text is the app's: the app holds what the user typed, and the
// choice is reported by Chosen and OnChanged.
type ComboBox struct {
	ui.Element // the text field, which is the trigger
	parts      ui.ComboboxParts
	c          *ui.Context
	text       *string
	options    []SelectOption
	size       Size
	variant    Variant
	shape      Shape
}

// NewComboBox creates a drop-down whose trigger is a text field, with
// the native preset. The text is the app's; the choice is reported by
// Chosen.
func NewComboBox(c *ui.Context, text *string, options []SelectOption) *ComboBox {
	b := &ComboBox{c: c, text: text, options: options}
	p := ui.ComboboxBase(c, text)
	b.parts = p
	b.Element = p.Input
	// The field takes the room of its content, but keeps a width of its
	// own so that it does not collapse to the width of a word, and it
	// does not fill the row it is in.
	b.Element = b.Element.AlignSelf(ui.Start).MinWidth(c.Theme().Space(30))

	// The popup holds the options the text matches.
	p.Popup(func(panel ui.Element) {
		t := c.Theme()
		panel.Background(t.Background).Border(1, t.Border).Radius(t.Radius).
			Padding(t.Space(1)).Shadow(0, 2, 8, 0, t.TextMuted.Alpha(0.15))
		filter := strings.TrimSpace(*text)
		group := ""
		panel.Children(func() {
			for _, o := range options {
				o := o
				if filter != "" && !strings.Contains(o.label(), filter) {
					continue
				}
				// The name of a group shows above its options.
				if o.Group != "" && o.Group != group {
					group = o.Group
					ui.Text(c, group).FontSize(TextSize(t, TextXs)).TextColor(t.TextMuted).
						Padding(t.Space(1), t.Space(2))
				}
				item := p.Item(o.Value)
				item.Children(func() {
					ui.Text(c, o.label())
					if o.Value == b.Chosen() {
						ui.Icon(c, icons.Check()).TextColor(t.Accent)
					}
				})
			}
		})
	})
	return b
}

// Variant applies the look of v and returns the combo box, for chaining.
// The zero variant is Native, which keeps the built-in look of the text
// field.
func (b *ComboBox) Variant(v Variant) *ComboBox {
	b.variant = v
	b.apply()
	return b
}

// Primary is Variant(myui.Primary).
func (b *ComboBox) Primary() *ComboBox { return b.Variant(Primary) }

// Secondary is Variant(myui.Secondary), a muted fill.
func (b *ComboBox) Secondary() *ComboBox { return b.Variant(Secondary) }

// Default is Variant(myui.Default), the built-in look.
func (b *ComboBox) Default() *ComboBox { return b.Variant(Default) }

// Native is Variant(myui.Native), the old name of Default.
func (b *ComboBox) Native() *ComboBox { return b.Variant(Native) }

// Outline is Variant(myui.Outline).
func (b *ComboBox) Outline() *ComboBox { return b.Variant(Outline) }

// Ghost is Variant(myui.Ghost).
func (b *ComboBox) Ghost() *ComboBox { return b.Variant(Ghost) }

// Destructive is Variant(myui.Destructive).
func (b *ComboBox) Destructive() *ComboBox { return b.Variant(Destructive) }

// Link is Variant(myui.Link).
func (b *ComboBox) Link() *ComboBox { return b.Variant(Link) }

// Success is Variant(myui.Success).
func (b *ComboBox) Success() *ComboBox { return b.Variant(Success) }

// Info is Variant(myui.Info).
func (b *ComboBox) Info() *ComboBox { return b.Variant(Info) }

// Warning is Variant(myui.Warning).
func (b *ComboBox) Warning() *ComboBox { return b.Variant(Warning) }

// Size applies the sizing of s and returns the combo box, for chaining.
func (b *ComboBox) Size(s Size) *ComboBox {
	b.size = s
	b.apply()
	return b
}

// Sm sizes the combo box down.
func (b *ComboBox) Sm() *ComboBox { return b.Size(SizeSm) }

// Md gives the built-in sizing.
func (b *ComboBox) Md() *ComboBox { return b.Size(SizeMd) }

// Lg sizes the combo box up.
func (b *ComboBox) Lg() *ComboBox { return b.Size(SizeLg) }

// Xs sizes the combo box to the smallest size.
func (b *ComboBox) Xs() *ComboBox { return b.Size(SizeXs) }

// Shape sets the shape of the combo box and returns it, for chaining.
func (b *ComboBox) Shape(s Shape) *ComboBox {
	b.shape = s
	b.apply()
	return b
}

// Width sets the width of the combo box in DIPs and returns it, for
// chaining.
func (b *ComboBox) Width(w float32) *ComboBox {
	b.Element = b.Element.Width(w)
	return b
}

// Placeholder shows s in the empty field and returns the combo box, for
// chaining.
func (b *ComboBox) Placeholder(s string) *ComboBox {
	b.Element = b.Element.Placeholder(s)
	return b
}

// Label names the combo box for assistive technology and returns it,
// for chaining.
func (b *ComboBox) Label(s string) *ComboBox {
	b.Element = b.Element.Label(s)
	return b
}

// OnChanged calls fn with the value of the new choice when one was
// made, and returns the combo box, for chaining.
//
// MyGo's combobox reports the choice for the frame that follows it:
// the field holds it then, and not after.
func (b *ComboBox) OnChanged(fn func(value string)) *ComboBox {
	if fn != nil {
		if v, ok := b.parts.Chosen(); ok {
			fn(v)
		}
	}
	return b
}

// Chosen returns the value of the choice the field holds.
func (b *ComboBox) Chosen() string {
	if v, ok := b.parts.Chosen(); ok {
		return v
	}
	return ""
}

// Changed reports whether a choice was made in this frame.
func (b *ComboBox) Changed() bool { return b.Element.Changed() }

// Open reports whether the popup shows.
func (b *ComboBox) Open() bool { return b.parts.Open() }

// apply paints the look of the variant and the size on the field.
func (b *ComboBox) apply() {
	t := b.c.Theme()
	l := fieldLookOf(b.variant, t)
	if !l.native {
		v, h := fieldPadding(t, b.size)
		b.Element = b.Element.Background(l.bg).Border(l.width, l.border).Radius(t.Radius).Padding(v, h)
	}
	switch b.shape {
	case ShapeRound, ShapeCircle:
		b.Element = b.Element.Radius(1000)
	default:
		b.Element = b.Element.Radius(t.Radius)
	}
	b.Element = b.Element.FontSize(fieldFontSize(t, b.size))
}
