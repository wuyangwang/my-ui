package forms

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/icons"
)

// SelectOption is one option of a Select: the label the trigger and the
// popup show, and the value the app holds. An option without a label
// shows its value. Icon, when set, shows before the label in both the
// trigger (once chosen) and the popup. The group of an option names the
// group it sits in: the popup shows the name above the options of the group.
//
//	{Label: "绠€浣撲腑鏂?, Value: "zh-Hans"}
//	{Value: "zh-Hant"}
//	{Group: "鏈€杩?, Label: "鑻辫", Value: "en"}
type SelectOption struct {
	Label    string
	Value    string
	Group    string
	Icon     *ui.SVG
	Disabled bool
}

// label is the text the option shows.
func (o SelectOption) Text() string {
	if o.Label != "" {
		return o.Label
	}
	return o.Value
}

// Select is a drop-down choosing one of its options:
//
//	forms.NewSelect(c, &size, []forms.SelectOption{{Value: "灏?}, {Value: "澶?}}).Outline()
//	forms.NewSelect(c, &lang, []forms.SelectOption{
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
	clearSlot       ui.Element // the slot of the clear icon
	arrowSlot       ui.Element // the slot of the chevron, hidden while clearing
	clearable       bool
	variant         Variant
	size            Size
	disabled        bool
	cleared         bool
	invalid         bool
	emptyTextRef    *string
	shape           Shape
}

// NewSelect creates a drop-down choosing one of options into *selected,
// with the native preset and a clear button.
func NewSelect(c *ui.Context, selected *string, options []SelectOption) *Select {
	return NewSelectClear(c, selected, options, true)
}

// NewSelectClear creates a drop-down like NewSelect, with clearable
// deciding whether the button that empties the choice is built. It must be
// settled here: the trigger's children are built while the select is
// created, so changing it afterwards has no effect.
func NewSelectClear(c *ui.Context, selected *string, options []SelectOption, clearable bool) *Select {
	s := &Select{c: c, value: selected, options: options, clearable: clearable}
	s.parts = ui.SelectBase(c, selected)
	s.emptyTextRef = ui.Local(s.parts.Trigger, "select-empty", func() string { return "" })

	// The trigger holds the choice, the placeholder, the button that
	// empties the choice and the arrow that opens the popup.
	s.parts.Trigger.Children(func() {
		s.valueSlot = ui.Row(c).Grow(1).AlignItems(ui.Center).Gap(c.Theme().Space(1))
		s.placeholderSlot = ui.Row(c).AlignItems(ui.Center)
		s.clearSlot = ui.Row(c).AlignItems(ui.Center)
		// A clearable select with a choice shows the clear icon in place
		// of the chevron, as Element Plus and shadcn do: one glyph at a
		// time, in the same slot.
		s.arrowSlot = ui.Row(c).AlignItems(ui.Center)
		s.arrowSlot.Children(func() {
			if _, ok := s.chosen(); ok && s.clearable {
				return
			}
			// The chevron turns to point up while the popup is open, as
			// Lobe's does. AnimateTo keeps Reduce Motion with the desktop.
			turn := float32(0)
			if s.parts.Open() {
				turn = 180
			}
			arrow := ui.Icon(c, icons.ChevronDown()).Size(IconSizeSm, IconSizeSm).
				TextColor(c.Theme().TextMuted)
			arrow = arrow.Rotate(arrow.AnimateTo("select-arrow", turn, ui.Animation{Duration: MotionNormal}))
		})
	})

	// The choice the trigger shows; an empty choice leaves room for the
	// placeholder.
	s.valueSlot.Children(func() {
		if o, ok := s.chosen(); ok {
			if o.Icon != nil {
				ui.Icon(c, o.Icon).Size(IconSizeXs, IconSizeXs)
			}
			ui.Text(c, o.Text())
		}
	})

	// The clear icon takes the chevron's place while there is a choice.
	s.clearSlot.Children(func() {
		if _, ok := s.chosen(); !ok || !s.clearable {
			return
		}
		ctrlpkg.NewIconAction(c, ctrlpkg.IconActionOptions{
			SVG:      icons.X(),
			Label:    Message(c, "clear"),
			Size:     16,
			IconSize: IconSizeXs,
			OnClick: func() {
				*selected = ""
				s.cleared = true
			},
		})
	})

	s.Element = s.parts.Trigger
	// A field fills the room it is given; Width(w) takes it over when
	// the app wants a fixed size. The floor keeps it from collapsing.
	s.Element = s.Element.FillWidth().MinWidth(c.Theme().Space(30))

	// The popup never gets narrower than its trigger. Bounds() reads the
	// last committed frame, so remember the widest one seen on the host.
	minWidth := ui.Local(s.parts.Trigger, "select-popup-min-width", func() float32 { return 0 })
	if w := s.parts.Trigger.Bounds().W; w > *minWidth {
		*minWidth = w
	}

	// The popup holds the options.
	//
	// The popup's content is the machinery's: the options it builds are
	// the popup's children, and what a view adds to the panel beside
	// them is dropped. A filter can therefore not live in the popup 鈥?	// the trigger of a select is a button, and a button takes no text.
	// A select that can be searched is another component, a ComboBox.
	s.parts.Popup(func(panel ui.Element) {
		t := c.Theme()
		panel.Background(t.Background).Border(1, t.Border).Radius(Radius(t, RadiusMd)).
			Padding(t.Space(1)).MinWidth(*minWidth).
			Shadow(0, 2, 8, 0, t.TextMuted.Alpha(0.15))

		// A long list scrolls inside the popup instead of running off the
		// window, so the last option stays reachable.
		ui.Scroll(c).FillWidth().MaxHeight(t.Space(80)).Children(func() {

			group := ""
			for _, o := range options {
				o := o
				// The name of a group shows above its options.
				if o.Group != "" && o.Group != group {
					group = o.Group
					ui.Text(c, group).FontSize(TextSize(t, TextXs)).TextColor(t.TextMuted).
						Padding(t.Space(1), t.Space(2))
				}
				item := s.parts.Item(o.Value).Padding(t.Space(1.5), t.Space(2)).Radius(Radius(t, RadiusSm)).
					AlignItems(ui.Center).Gap(t.Space(1))
				if item.Highlighted() {
					item.Background(t.Accent.Alpha(0.10))
				}
				item.Children(func() {
					if o.Icon != nil {
						ui.Icon(c, o.Icon).Size(IconSizeXs, IconSizeXs)
					}
					ui.Text(c, o.Text()).Grow(1)
					if o.Value == *selected {
						ui.Icon(c, icons.Check()).Size(IconSizeSm, IconSizeSm).TextColor(t.Accent)
					}
				})
				if o.Disabled {
					item.Disabled(true).TextColor(t.TextMuted)
				}
			}
			if len(options) == 0 && s.emptyTextRef != nil && *s.emptyTextRef != "" {
				ui.Text(c, *s.emptyTextRef).TextColor(t.TextMuted).
					FontSize(TextSize(t, TextSm)).Padding(t.Space(1), t.Space(2))
			}
		})
	})

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

// Shape sets the shape of the trigger and returns the select, for
// chaining: a pill or a circle, or the theme's own corners.
func (s *Select) Shape(shape Shape) *Select {
	s.shape = shape
	s.apply()
	return s
}

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

// Invalid marks the select as not holding a good value and returns it,
// for chaining. The trigger takes the danger border of the theme, as a
// field's control does when the field has an error.
func (s *Select) Invalid(on bool) *Select {
	s.invalid = on
	s.Element = s.Element.Invalid(on)
	s.apply()
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

// SetOpen opens or closes the popup, and returns the select for chaining:
// an app can open it from its own button or close it on submit.
func (s *Select) SetOpen(on bool) *Select {
	s.parts.SetOpen(on)
	return s
}

// Highlight moves the highlight to the option with value, as a search
// field jumps to a match, and returns the select for chaining.
func (s *Select) Highlight(value string) *Select {
	s.parts.Highlight(value)
	return s
}

// EmptyText sets the message the popup shows when there are no options,
// and returns the select for chaining.
func (s *Select) EmptyText(text string) *Select {
	// The popup is rebuilt each frame from the options; the text lives
	// on the trigger so the next frame's build reads it.
	if s.emptyTextRef != nil {
		*s.emptyTextRef = text
	}
	return s
}

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
	border := l.Border
	if s.invalid {
		border = t.Danger
	}
	radius := Radius(t, RadiusMd)
	if s.shape == ShapeRound || s.shape == ShapeCircle {
		radius = 1000
	}
	s.Element = s.Element.Background(l.Bg).Border(l.Width, border).Radius(radius)
	v, h := FramePadding(t, s.size)
	s.Element = s.Element.Padding(v, h).FontSize(FrameFontSize(t, s.size))
}

// selectLookOf resolves one variant's look for a select's trigger.
//
// The trigger of a select has no built-in look 鈥?MyGo's SelectBase hands
// over a bare button 鈥?so the field look is its default one.
func selectLookOf(v Variant, t *ui.Theme) FrameLook {
	l := FrameLookOf(v, t)
	if l.Native {
		return FrameLook{Bg: t.Surface, Border: t.Border, Width: 1}
	}
	return l
}
