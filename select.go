package myui

import (
	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/icons"
)

// SelectOption is one option of a Select: the label the trigger and the
// popup show, and the value the app holds. An option without a label
// shows its value. The group of an option names the group it sits in:
// the popup shows the name above the options of the group.
//
//	{Label: "绠€浣撲腑鏂?, Value: "zh-Hans"}
//	{Value: "zh-Hant"}
//	{Group: "鏈€杩?, Label: "鑻辫", Value: "en"}
type SelectOption struct {
	Label    string
	Value    string
	Group    string
	Disabled bool
}

// label is the text the option shows.
func (o SelectOption) label() string {
	if o.Label != "" {
		return o.Label
	}
	return o.Value
}

// Select is a drop-down choosing one of its options:
//
//	myui.NewSelect(c, &size, []myui.SelectOption{{Value: "灏?}, {Value: "澶?}}).Outline()
//	myui.NewSelect(c, &lang, []myui.SelectOption{
//		{Label: "绠€浣撲腑鏂?, Value: "zh-Hans"},
//		{Label: "English", Value: "en"},
//	}).Placeholder("閫夋嫨璇█").OnChanged(func(v string) { app.lang = v })
//
// The trigger looks like an input: the same frame, the same border, the
// same sizes, with an arrow at its end that opens the popup. The popup
// below holds the options; the one under the pointer is highlighted and
// the chosen one is marked with a check.
//
// The app holds the **value** of the choice; the trigger shows its
// label.
//
// The chain applies the presets from the current theme every frame, so
// it follows palette and appearance changes. Call it every frame, ahead
// of anything else the app styles on it, as the view runs.
type Select struct {
	ui.Element      // the trigger
	parts           ui.SelectParts[string]
	c               *ui.Context
	value           *string
	options         []SelectOption
	valueSlot       ui.Element // the slot of the choice
	placeholderSlot ui.Element // the slot of the placeholder
	clearSlot       ui.Element // the slot of the clear button
	variant         Variant
	size            Size
	disabled        bool
	cleared         bool
}

// NewSelect creates a drop-down choosing one of options into *selected,
// with the native preset. The value of the choice is what *selected
// holds; the label of the option is what the trigger shows.
func NewSelect(c *ui.Context, selected *string, options []SelectOption) *Select {
	s := &Select{c: c, value: selected, options: options}
	s.parts = ui.SelectBase(c, selected)

	// The trigger holds the choice, the placeholder, the button that
	// empties the choice and the arrow that opens the popup.
	s.parts.Trigger.Children(func() {
		s.valueSlot = ui.Row(c).Grow(1).AlignItems(ui.Center)
		s.placeholderSlot = ui.Row(c).AlignItems(ui.Center)
		s.clearSlot = ui.Row(c).AlignItems(ui.Center)
		ui.Icon(c, icons.ChevronDown()).Margin(0, 0, 0, c.Theme().Space(1.5))
	})

	// The choice the trigger shows; an empty choice leaves room for the
	// placeholder.
	s.valueSlot.Children(func() {
		if o, ok := s.chosen(); ok {
			ui.Text(c, o.label())
		}
	})

	// The button that empties the choice, while there is one.
	s.clearSlot.Children(func() {
		if _, ok := s.chosen(); !ok {
			return
		}
		t := c.Theme()
		d := float32(16)
		// The button keeps the focus in the field while it is pressed, and
		// it is a stop of the Tab key, as a clear button is in Ant Design.
		b := ui.ButtonBase(c).Size(d, d).KeepFocus().Label("娓呴櫎")
		b.Children(func() {
			ic := ui.Icon(c, icons.X()).Size(d-5, d-5).TextColor(t.TextMuted)
			if b.Hovered() {
				ic = ic.TextColor(t.Text)
			}
		})
		if b.Clicked() {
			*selected = ""
			s.cleared = true
		}
	})

	// The popup holds the options.
	//
	// The popup's content is the machinery's: the options it builds are
	// the popup's children, and what a view adds to the panel beside
	// them is dropped. A filter can therefore not live in the popup 鈥?	// the trigger of a select is a button, and a button takes no text.
	// A select that can be searched is another component, a ComboBox.
	s.parts.Popup(func(panel ui.Element) {
		t := c.Theme()
		panel.Background(t.Background).Border(1, t.Border).Radius(t.Radius).
			Padding(t.Space(1)).Shadow(0, 2, 8, 0, t.TextMuted.Alpha(0.15))

		group := ""
		for _, o := range options {
			o := o
			// The name of a group shows above its options.
			if o.Group != "" && o.Group != group {
				group = o.Group
				ui.Text(c, group).FontSize(TextSize(t, TextXs)).TextColor(t.TextMuted).
					Padding(t.Space(1), t.Space(2))
			}
			item := s.parts.Item(o.Value).Padding(t.Space(1.5), t.Space(2)).Radius(t.Radius / 2)
			if item.Highlighted() {
				item.Background(t.Accent.Alpha(0.10))
			}
			item.Children(func() {
				ui.Text(c, o.label()).Grow(1)
				if o.Value == *selected {
					ui.Icon(c, icons.Check()).TextColor(t.Accent)
				}
			})
			if o.Disabled {
				item.Disabled(true).TextColor(t.TextMuted)
			}
		}
	})

	s.Element = s.parts.Trigger
	// The trigger takes the room of its content, but keeps a width of
	// its own so that it does not collapse to the width of a word. It
	// does not fill the row it is in: a drop-down that is as wide as the
	// page is not a drop-down.
	s.Element = s.Element.AlignSelf(ui.Start).MinWidth(c.Theme().Space(30))
	return s
}

// Variant applies the look of v and returns the select, for chaining.
// The zero variant is Native, which keeps the built-in look.
func (s *Select) Variant(v Variant) *Select {
	s.variant = v
	s.apply()
	return s
}

// Primary is Variant(myui.Primary).
func (s *Select) Primary() *Select { return s.Variant(Primary) }

// Secondary is Variant(myui.Secondary), a muted fill.
func (s *Select) Secondary() *Select { return s.Variant(Secondary) }

// Default is Variant(myui.Default), the look of an input field.
func (s *Select) Default() *Select { return s.Variant(Default) }

// Native is Variant(myui.Native), the old name of Default.
func (s *Select) Native() *Select { return s.Variant(Native) }

// Outline is Variant(myui.Outline).
func (s *Select) Outline() *Select { return s.Variant(Outline) }

// Ghost is Variant(myui.Ghost).
func (s *Select) Ghost() *Select { return s.Variant(Ghost) }

// Destructive is Variant(myui.Destructive).
func (s *Select) Destructive() *Select { return s.Variant(Destructive) }

// Link is Variant(myui.Link).
func (s *Select) Link() *Select { return s.Variant(Link) }

// Success is Variant(myui.Success).
func (s *Select) Success() *Select { return s.Variant(Success) }

// Info is Variant(myui.Info).
func (s *Select) Info() *Select { return s.Variant(Info) }

// Warning is Variant(myui.Warning).
func (s *Select) Warning() *Select { return s.Variant(Warning) }

// Size applies the sizing of s and returns the select, for chaining.
func (s *Select) Size(size Size) *Select {
	s.size = size
	s.apply()
	return s
}

// Sm sizes the select down.
func (s *Select) Sm() *Select { return s.Size(SizeSm) }

// Xs sizes the select to the smallest size.
func (s *Select) Xs() *Select { return s.Size(SizeXs) }

// Md gives the built-in sizing.
func (s *Select) Md() *Select { return s.Size(SizeMd) }

// Lg sizes the select up.
func (s *Select) Lg() *Select { return s.Size(SizeLg) }

// Control returns the element a field describes: the trigger, which is
// the one that takes the focus.
func (s *Select) Control() ui.Element { return s.Element }

// Width sets the width of the trigger in DIPs and returns the select,
// for chaining. Without it the trigger takes the room of its content,
// with a width of its own as the floor.
func (s *Select) Width(w float32) *Select {
	s.Element = s.Element.Width(w)
	return s
}

// Placeholder shows s while nothing is chosen and returns the select,
// for chaining.
func (s *Select) Placeholder(text string) *Select {
	c := s.c
	if _, ok := s.chosen(); !ok {
		s.placeholderSlot.Children(func() {
			ui.Text(c, text).TextColor(c.Theme().TextMuted)
		})
	}
	return s
}

// Label names the select for assistive technology and returns it, for
// chaining.
func (s *Select) Label(text string) *Select {
	s.Element = s.Element.Label(text)
	return s
}

// Disabled takes the select out of use and returns it, for chaining.
func (s *Select) Disabled(on bool) *Select {
	s.disabled = on
	s.Element = s.Element.Disabled(on)
	return s
}

// OnChanged calls fn with the new value when it changed in this frame,
// and returns the select, for chaining.
func (s *Select) OnChanged(fn func(value string)) *Select {
	if fn != nil && s.Changed() {
		fn(*s.value)
	}
	return s
}

// OnCleared calls fn when the clear button emptied the choice, and
// returns the select, for chaining.
func (s *Select) OnCleared(fn func()) *Select {
	if fn != nil && s.cleared {
		fn()
	}
	return s
}

// Value returns the value of the choice.
func (s *Select) Value() string { return *s.value }

// Open reports whether the popup shows.
func (s *Select) Open() bool { return s.parts.Open() }

// Changed reports whether the choice changed in this frame.
func (s *Select) Changed() bool { return s.Element.Changed() }

// chosen is the option the select holds.
func (s *Select) chosen() (SelectOption, bool) {
	for _, o := range s.options {
		if o.Value == *s.value {
			return o, true
		}
	}
	return SelectOption{}, false
}

// apply paints the look of the variant and the size on the trigger.
//
// The trigger of a select has no built-in look of its own 鈥?MyGo's
// SelectBase hands over a bare button 鈥?so the native variant takes the
// look of an input field: a face and a border, or the trigger would be
// a shell with no room in it.
func (s *Select) apply() {
	t := s.c.Theme()
	l := selectLookOf(s.variant, t)
	s.Element = s.Element.Background(l.bg).Border(l.width, l.border).Radius(t.Radius)
	v, h := fieldPadding(t, s.size)
	s.Element = s.Element.Padding(v, h).FontSize(fieldFontSize(t, s.size))
}

// selectLookOf resolves one variant's look for a select's trigger.
//
// The trigger of a select has no built-in look 鈥?MyGo's SelectBase hands
// over a bare button 鈥?so the field look is its default one.
func selectLookOf(v Variant, t *ui.Theme) fieldLook {
	l := fieldLookOf(v, t)
	if l.native {
		return fieldLook{bg: t.Surface, border: t.Border, width: 1}
	}
	return l
}
