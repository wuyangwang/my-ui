package forms

import (
	"strings"

	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/icons"
)

// ComboOptions configures a ComboBox at construction. The selected value
// and the typed query are separate: Value is the app-owned choice, Text is
// the (usually internal) text of the trigger.
type ComboOptions struct {
	// Value is the app-owned selected value. It is written when an option
	// is chosen and cleared to "" when the clear button is pressed.
	Value *string
	// Text is the trigger's text and the query that filters the options.
	// When nil the combo keeps its own text in a per-frame slot.
	Text *string
	// Options are the choices shown in the popup.
	Options []SelectOption
	// Clearable shows a clear icon in place of the chevron while a value
	// is selected.
	Clearable bool
	// EmptyText is shown when nothing matches the query. When empty the
	// library's localized "no results" string is used.
	EmptyText string
	// Placeholder is shown while the trigger's text is empty.
	Placeholder string
	// Label names the combo for assistive technology.
	Label string
	// Filter replaces the default case-insensitive label match.
	Filter func(SelectOption, string) bool
	// OnChanged is called with the new value on a choice or a clear. It may
	// also be attached later with the fluent OnChanged method.
	OnChanged func(value string)
}

// ComboBox is a drop-down whose trigger is a text field: what is typed
// narrows the options down, and a choice sets the app's selected value.
// The typed query and the selected value are different state — Value in
// ComboOptions is the choice; Text is the query.
//
//	forms.NewComboBox(c, forms.ComboOptions{
//		Value:     &app.lang,
//		Options:   []forms.SelectOption{{Label: "English", Value: "en"}},
//		Clearable: true,
//	}).Outline().Placeholder("搜索语言…")
type ComboBox struct {
	ui.Element // the text input, the control a Field describes
	c          *ui.Context
	parts      ui.ComboboxParts
	root       ui.Element // the frame around the input and the trailing glyph
	glyph      ui.Element // the chevron, or the clear button
	value      *string
	text       *string
	options    []SelectOption
	clearable  bool
	size       Size
	variant    Variant
	shape      Shape
	width      float32
	invalid    bool
	disabled   bool
	chosen     *string
	pending    *pendingChoice
	emptyRef   *string
	filterRef  *func(SelectOption, string) bool
	onChanged  *func(string)
}

// pendingChoice is a choice waiting for a registered callback. It lives
// in a per-frame slot, so a callback attached after construction — even
// on a later frame — still receives the choice.
type pendingChoice struct {
	v  string
	ok bool
}

// NewComboBox creates a searchable drop-down configured by o. If o.Value
// is nil the selected value is kept in a private slot, so Chosen() stays
// readable; the app should normally pass its own pointer.
func NewComboBox(c *ui.Context, o ComboOptions) *ComboBox {
	b := &ComboBox{
		c: c, options: o.Options, clearable: o.Clearable,
		value: o.Value,
	}

	// The frame has to exist before the input, because the input's text
	// pointer — when it is internal — lives on that frame.
	root := ui.Row(c).AlignItems(ui.Center)
	b.root = root
	if b.value == nil {
		b.value = ui.Local(root, "combo-value", func() string { return "" })
	}
	text := o.Text
	if text == nil {
		text = ui.Local(root, "combo-text", func() string { return "" })
	}
	b.text = text

	root.Children(func() {
		b.parts = ui.ComboboxBase(c, text)
		b.parts.Input = b.parts.Input.Grow(1)
		b.glyph = ui.Row(c).AlignItems(ui.Center)
	})
	b.Element = b.parts.Input

	b.chosen = ui.Local(root, "combo-chosen", func() string { return "" })
	b.pending = ui.Local(root, "combo-pending", func() pendingChoice { return pendingChoice{} })
	b.emptyRef = ui.Local(root, "combo-empty", func() string { return comboEmptyText(c, o.EmptyText) })
	b.filterRef = ui.Local(root, "combo-filter", func() func(SelectOption, string) bool { return o.Filter })
	b.onChanged = ui.Local(root, "combo-onchanged", func() func(string) { return o.OnChanged })
	if o.EmptyText != "" {
		*b.emptyRef = o.EmptyText
	}

	// The popup holds the options the query matches.
	b.parts.Popup(func(panel ui.Element) {
		t := c.Theme()
		panel.Background(t.Background).Border(1, t.Border).Radius(Radius(t, RadiusMd)).
			Padding(t.Space(1)).Shadow(0, 2, 8, 0, t.TextMuted.Alpha(0.15))
		query := strings.TrimSpace(*text)
		group := ""
		matched := 0
		panel.Children(func() {
			for _, o := range b.options {
				o := o
				if query != "" && !b.matches(o, query) {
					continue
				}
				matched++
				if o.Group != "" && o.Group != group {
					group = o.Group
					ui.Text(c, group).FontSize(TextSize(t, TextXs)).TextColor(t.TextMuted).
						Padding(t.Space(1), t.Space(2))
				}
				item := b.parts.Item(o.Value)
				item.Children(func() {
					ui.Text(c, o.Text())
					if o.Value == *b.value {
						ui.Icon(c, icons.Check()).Size(IconSizeSm, IconSizeSm).TextColor(t.Accent)
					}
				})
				if o.Disabled {
					item.Disabled(true).TextColor(t.TextMuted)
				}
			}
			if matched == 0 && *b.emptyRef != "" {
				ui.Text(c, *b.emptyRef).TextColor(t.TextMuted).
					FontSize(TextSize(t, TextSm)).Padding(t.Space(2))
			}
		})
	})

	// The trailing glyph: a clear button replaces the chevron while a
	// value is selected, as in Select. One glyph at a time, in one slot.
	b.glyph.Children(func() {
		if !b.clearable || *b.value == "" || b.disabled {
			b.chevron()
			return
		}
		ctrlpkg.NewIconAction(c, ctrlpkg.IconActionOptions{
			SVG:      icons.X(),
			Label:    Message(c, "clear"),
			Size:     16,
			IconSize: IconSizeXs,
			OnClick:  b.clear,
		})
	})

	// MyGo reports the choice for the frame that follows Popup: take it
	// now, while it is available, and keep it. The choice becomes the
	// app's selected value and the query becomes the option's label.
	if v, ok := b.parts.Chosen(); ok {
		*b.chosen = v
		*b.value = v
		for _, o := range b.options {
			if o.Value == v {
				*text = o.Text()
				break
			}
		}
		b.pending.v, b.pending.ok = v, true
	}
	b.fire()

	if o.Placeholder != "" {
		b.Placeholder(o.Placeholder)
	}
	if o.Label != "" {
		b.Label(o.Label)
	}
	b.root = b.root.Transition(StateTransition())
	b.apply()
	return b
}

// chevron draws the arrow that opens the popup. It turns to point up
// while the popup is open, through mygo's animation so Reduce Motion is
// respected.
func (b *ComboBox) chevron() {
	c := b.c
	turn := float32(0)
	if b.parts.Open() {
		turn = 180
	}
	arrow := ui.Icon(c, icons.ChevronDown()).Size(IconSizeSm, IconSizeSm).
		TextColor(c.Theme().TextMuted)
	arrow = arrow.Rotate(arrow.AnimateTo("combo-arrow", turn, ui.Animation{Duration: MotionNormal}))
}

// clear empties both the selected value and the query, and reports the
// empty value through the registered callback.
func (b *ComboBox) clear() {
	*b.value = ""
	*b.text = ""
	*b.chosen = ""
	b.pending.v, b.pending.ok = "", true
	b.fire()
}

// fire calls the registered callback with a pending choice, if any. A
// callback registered after the choice — even on a later frame — receives
// it, because the pending slot lives until it is consumed.
func (b *ComboBox) fire() {
	if b.pending == nil || !b.pending.ok || b.onChanged == nil || *b.onChanged == nil {
		return
	}
	fn := *b.onChanged
	b.pending.ok = false
	fn(b.pending.v)
}

// comboEmptyText is the library's default text when nothing matches. It
// goes through the root package's i18n so an app can translate it; until
// a catalog entry exists, Message returns the key itself and the English
// fallback is shown rather than a bare key.
func comboEmptyText(c *ui.Context, supplied string) string {
	if supplied != "" {
		return supplied
	}
	const key = "noResults"
	if s := Message(c, key); s != key {
		return s
	}
	return "No results found"
}

// matches tells whether an option matches the query. A filter set with
// Filter or in ComboOptions replaces the default label match.
func (b *ComboBox) matches(o SelectOption, query string) bool {
	if b.filterRef != nil && *b.filterRef != nil {
		return (*b.filterRef)(o, query)
	}
	return strings.Contains(strings.ToLower(o.Text()), strings.ToLower(query))
}

// Variant applies the look of v and returns the combo box, for chaining.
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

// Width pins the frame's width in DIPs and returns the combo box. Zero
// fills the parent; a positive width keeps a floor so the field cannot
// collapse.
func (b *ComboBox) Width(w float32) *ComboBox {
	b.width = w
	b.apply()
	return b
}

// Placeholder shows s while the trigger's text is empty, and returns the
// combo box for chaining.
func (b *ComboBox) Placeholder(s string) *ComboBox {
	if s != "" {
		b.Element = b.Element.Placeholder(s)
	}
	return b
}

// Label names the combo box for assistive technology and returns it, for
// chaining.
func (b *ComboBox) Label(s string) *ComboBox {
	if s != "" {
		b.Element = b.Element.Label(s)
	}
	return b
}

// OnChanged registers fn as the callback for a choice or a clear, and
// returns the combo box for chaining. It may be called at any point — a
// choice that happened before the callback was attached is kept in a
// per-frame slot and delivered here.
func (b *ComboBox) OnChanged(fn func(value string)) *ComboBox {
	if b.onChanged != nil {
		*b.onChanged = fn
	}
	b.fire()
	return b
}

// Chosen returns the value of the last choice, readable in any frame. It
// is emptied by a clear. MyGo's own Chosen() expires with the frame; the
// combo copies it into a persistent slot in the constructor.
func (b *ComboBox) Chosen() string {
	if b.chosen != nil {
		return *b.chosen
	}
	return ""
}

// Selected returns the app-owned selected value, readable in any frame.
// It is the value in ComboOptions.Value, not the typed query.
func (b *ComboBox) Selected() string {
	if b.value != nil {
		return *b.value
	}
	return ""
}

// Text returns the trigger's text — the query the options are filtered
// by. It is the app's Text pointer, or the combo's internal one.
func (b *ComboBox) Text() string {
	if b.text != nil {
		return *b.text
	}
	return ""
}

// Changed reports whether the trigger changed in this frame.
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
	b.disabled = on
	b.Element = b.Element.Disabled(on)
	return b
}

// Invalid marks the combo box as not holding a good value and returns it,
// for chaining. The frame takes the danger border.
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

// Control returns the element a field describes: the text input that is
// the trigger, so a Field can focus it and name its error.
func (b *ComboBox) Control() ui.Element { return b.Element }

// apply paints the shared look of the variant, size and shape on the
// frame; colors and metrics come from c.Theme() on this frame.
func (b *ComboBox) apply() {
	t := b.c.Theme()
	l := FrameLookOf(b.variant, t)
	border := l.Border
	if b.invalid {
		border = t.Danger
	}
	if !l.Native {
		v, h := FramePadding(t, b.size)
		b.root = b.root.Background(l.Bg).Border(l.Width, border).Radius(Radius(t, RadiusMd)).Padding(v, h)
		// The frame owns the border; the input has none of its own.
		b.Element = b.Element.Background(ui.Transparent).Border(0, ui.Transparent).Padding(0)
	}
	switch b.shape {
	case ShapeRound, ShapeCircle:
		b.root = b.root.Radius(1000)
	default:
		b.root = b.root.Radius(Radius(t, RadiusMd))
	}
	b.Element = b.Element.FontSize(FrameFontSize(t, b.size))
	if b.width > 0 {
		b.root = b.root.Width(b.width)
	} else {
		b.root = b.root.FillWidth().MinWidth(t.Space(30))
	}
}
