package forms

import (
	. "github.com/wuyangwang/my-ui"
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
	invalid    bool
	chosen     *string // the last choice, kept across frames
	filterRef  *func(SelectOption, string) bool
	emptyRef   *string
}

// NewComboBox creates a drop-down whose trigger is a text field, with
// the native preset. The text is the app's; the choice is reported by
// Chosen.
func NewComboBox(c *ui.Context, text *string, options []SelectOption) *ComboBox {
	b := &ComboBox{c: c, text: text, options: options}
	p := ui.ComboboxBase(c, text)
	b.parts = p
	b.Element = p.Input
	// A field fills the room it is given; Width(w) takes it over when
	// the app wants a fixed size. The floor keeps it from collapsing.
	b.Element = b.Element.FillWidth().MinWidth(c.Theme().Space(30))
	// The choice MyGo reports lives for one frame. Keep it on the input
	// so that Chosen() stays readable in the frames after the choice.
	b.chosen = ui.Local(b.Element, "combo-chosen", func() string { return "" })
	b.emptyRef = ui.Local(b.Element, "combo-empty", func() string { return "" })
	b.filterRef = ui.Local(b.Element, "combo-filter", func() func(SelectOption, string) bool { return nil })

	// The popup holds the options the text matches.
	p.Popup(func(panel ui.Element) {
		t := c.Theme()
		panel.Background(t.Background).Border(1, t.Border).Radius(t.Radius).
			Padding(t.Space(1)).Shadow(0, 2, 8, 0, t.TextMuted.Alpha(0.15))
		filter := strings.TrimSpace(*text)
		group := ""
		matched := 0
		panel.Children(func() {
			for _, o := range options {
				o := o
				if filter != "" && !b.matches(o, filter) {
					continue
				}
				matched++
				// The name of a group shows above its options.
				if o.Group != "" && o.Group != group {
					group = o.Group
					ui.Text(c, group).FontSize(TextSize(t, TextXs)).TextColor(t.TextMuted).
						Padding(t.Space(1), t.Space(2))
				}
				item := p.Item(o.Value)
				item.Children(func() {
					ui.Text(c, o.Text())
					if o.Value == *b.chosen {
						ui.Icon(c, icons.Check()).Size(IconSizeSm, IconSizeSm).TextColor(t.Accent)
					}
				})
			}
			if matched == 0 && b.emptyRef != nil && *b.emptyRef != "" {
				ui.Text(c, *b.emptyRef).TextColor(t.TextMuted).
					FontSize(TextSize(t, TextSm)).Padding(t.Space(2))
			}
		})
	})
	// MyGo reports the choice for the frame that follows Popup: take it
	// now, while it is available, and keep it. A choice takes the place
	// of what was typed — the label is written back into the text the
	// app holds, so the input echoes what was chosen, as MyGo's own
	// Combobox does.
	if v, ok := p.Chosen(); ok {
		*b.chosen = v
		for _, o := range options {
			if o.Value == v {
				*text = o.Text()
				break
			}
		}
	}
	return b
}

// matches tells whether an option matches what was typed. A filter set
// with Filter replaces the default case-insensitive label match.
func (b *ComboBox) matches(o SelectOption, query string) bool {
	if b.filterRef != nil && *b.filterRef != nil {
		return (*b.filterRef)(o, query)
	}
	return strings.Contains(strings.ToLower(o.Text()), strings.ToLower(query))
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

// Chosen returns the value of the last choice, readable in any frame —
// not just the one the choice was made in. MyGo's own Chosen() is only
// valid during the build and expires with the frame; the combo box copies
// it into a persistent slot in the constructor and reads only that here.
func (b *ComboBox) Chosen() string {
	if b.chosen != nil {
		return *b.chosen
	}
	return ""
}

// Changed reports whether a choice was made in this frame.
func (b *ComboBox) Changed() bool { return b.Element.Changed() }

// Open reports whether the popup shows.
func (b *ComboBox) Open() bool { return b.parts.Open() }

// SetOpen opens or closes the popup, and returns the combo box for
// chaining: an app can open it from its own button or close it on submit.
func (b *ComboBox) SetOpen(on bool) *ComboBox {
	b.parts.SetOpen(on)
	return b
}

// Disabled takes the combo box out of use and returns it, for chaining.
func (b *ComboBox) Disabled(on bool) *ComboBox {
	b.Element = b.Element.Disabled(on)
	return b
}

// Invalid marks the combo box as not holding a good value and returns it,
// for chaining. The trigger takes the danger border.
func (b *ComboBox) Invalid(on bool) *ComboBox {
	b.invalid = on
	b.Element = b.Element.Invalid(on)
	b.apply()
	return b
}

// EmptyText sets the message the popup shows when nothing matches, and
// returns the combo box for chaining.
func (b *ComboBox) EmptyText(s string) *ComboBox {
	if b.emptyRef != nil {
		*b.emptyRef = s
	}
	return b
}

// Filter sets the matcher used to narrow the options by what was typed,
// and returns the combo box for chaining. The default matches the label
// case-insensitively.
func (b *ComboBox) Filter(fn func(o SelectOption, query string) bool) *ComboBox {
	if b.filterRef != nil {
		*b.filterRef = fn
	}
	return b
}

// Control returns the element a field describes: the text field that is
// the trigger, so a Field can focus it and name its error.
func (b *ComboBox) Control() ui.Element { return b.Element }

// apply paints the look of the variant and the size on the field.
func (b *ComboBox) apply() {
	t := b.c.Theme()
	l := FrameLookOf(b.variant, t)
	if !l.Native {
		v, h := FramePadding(t, b.size)
		border := l.Border
		if b.invalid {
			border = t.Danger
		}
		b.Element = b.Element.Background(l.Bg).Border(l.Width, border).Radius(t.Radius).Padding(v, h)
	}
	switch b.shape {
	case ShapeRound, ShapeCircle:
		b.Element = b.Element.Radius(1000)
	default:
		b.Element = b.Element.Radius(t.Radius)
	}
	b.Element = b.Element.FontSize(FrameFontSize(t, b.size))
}
