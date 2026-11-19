package ai

import (
	"strings"

	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
	formspkg "github.com/wuyangwang/my-ui/forms"
	"github.com/wuyangwang/my-ui/icons"
)

// ModelOption is one model the picker offers. Group heads a section, Icon
// replaces the default model glyph, and Context, Speed and Price show as
// small capability chips. Disabled models are shown but cannot be chosen.
type ModelOption struct {
	ID          string
	Name        string
	Group       string
	Description string
	Icon        *ui.SVG
	Context     string
	Speed       string
	Price       string
	Disabled    bool
}

// Text is the label shown in the popup and echoed into the search field
// when the model is chosen.
func (o ModelOption) Text() string {
	if o.Name != "" {
		return o.Name
	}
	return o.ID
}

// ModelPicker is a searchable drop-down of models. Its trigger is a text
// field whose content the app owns; the chosen ID is reported for one frame
// by mygo and kept on the input for later frames.
//
//	var query string
//	picker := ai.NewModelPicker(c, &query, models).Outline()
//	if picker.Changed() {
//		app.model = picker.Chosen()
//	}
type ModelPicker struct {
	ui.Element
	parts     ui.ComboboxParts
	c         *ui.Context
	text      *string
	options   []ModelOption
	variant   Variant
	size      Size
	shape     Shape
	invalid   bool
	chosen    *string
	emptyRef  *string
	filterRef *func(ModelOption, string) bool
}

// NewModelPicker builds a model drop-down over text and options. text is the
// app's search field: the picker writes the chosen model's name into it and
// reads what the user typed to narrow the list.
func NewModelPicker(c *ui.Context, text *string, options []ModelOption) *ModelPicker {
	p := &ModelPicker{c: c, text: text, options: options}
	parts := ui.ComboboxBase(c, text)
	p.parts = parts
	p.Element = parts.Input.FillWidth().MinWidth(c.Theme().Space(30))
	p.chosen = ui.Local(p.Element, "model-chosen", func() string { return "" })
	p.emptyRef = ui.Local(p.Element, "model-empty", func() string { return "" })
	p.filterRef = ui.Local(p.Element, "model-filter", func() func(ModelOption, string) bool { return nil })

	parts.Popup(func(panel ui.Element) {
		t := c.Theme()
		panel.Background(t.Background).Border(1, t.Border).Radius(Radius(t, RadiusMd)).
			Padding(t.Space(1)).Shadow(0, 2, 8, 0, t.TextMuted.Alpha(0.15))
		filter := strings.TrimSpace(*text)
		group := ""
		matched := 0
		panel.Children(func() {
			for _, o := range options {
				if filter != "" && !p.matches(o, filter) {
					continue
				}
				matched++
				if o.Group != "" && o.Group != group {
					group = o.Group
					ui.Text(c, group).FontSize(TextSize(t, TextXs)).TextColor(t.TextMuted).
						Padding(t.Space(1), t.Space(2))
				}
				item := parts.Item(o.ID)
				item.Disabled(o.Disabled)
				item.Children(func() {
					ui.Row(c).AlignItems(ui.Center).Gap(t.Space(2)).Children(func() {
						glyph := o.Icon
						if glyph == nil {
							glyph = icons.CpuAi()
						}
						ui.Icon(c, glyph).Size(IconSizeSm, IconSizeSm).TextColor(t.Accent)
						ui.Column(c).Gap(0).Children(func() {
							ui.Text(c, o.Text())
							if o.Description != "" {
								ui.Text(c, o.Description).FontSize(TextSize(t, TextXs)).TextColor(t.TextMuted)
							}
						})
						if o.ID == *p.chosen {
							ui.Icon(c, icons.Check()).Size(IconSizeSm, IconSizeSm).TextColor(t.Accent)
						}
					})
					if o.Context != "" || o.Speed != "" || o.Price != "" {
						ui.Row(c).Gap(t.Space(1)).Padding(0, t.Space(4)).Children(func() {
							for _, chip := range []string{o.Context, o.Speed, o.Price} {
								if chip != "" {
									modelChip(c, t, chip)
								}
							}
						})
					}
				})
			}
			if matched == 0 && *p.emptyRef != "" {
				ui.Text(c, *p.emptyRef).TextColor(t.TextMuted).
					FontSize(TextSize(t, TextSm)).Padding(t.Space(2))
			}
		})
	})
	// MyGo reports a choice for one frame only. Copy it onto the input now
	// and echo the model's name into the text the app holds.
	if v, ok := parts.Chosen(); ok {
		*p.chosen = v
		for _, o := range options {
			if o.ID == v {
				*text = o.Text()
				break
			}
		}
	}
	return p
}

// modelChip draws one capability as a muted pill.
func modelChip(c *ui.Context, t *ui.Theme, s string) {
	ui.Row(c).Padding(t.Space(0.5), t.Space(1)).Background(t.SurfaceHover).
		Radius(Radius(t, RadiusSm)).Children(func() {
		ui.Text(c, s).FontSize(TextSize(t, TextXs)).TextColor(t.TextMuted)
	})
}

// matches narrows the options by what was typed. A filter set with Filter
// replaces the default case-insensitive match on the model's name.
func (p *ModelPicker) matches(o ModelOption, query string) bool {
	if p.filterRef != nil && *p.filterRef != nil {
		return (*p.filterRef)(o, query)
	}
	return strings.Contains(strings.ToLower(o.Text()), strings.ToLower(query))
}

// Variant applies the look of v and returns the picker, for chaining.
func (p *ModelPicker) Variant(v Variant) *ModelPicker {
	p.variant = v
	p.apply()
	return p
}

// Primary is Variant(Primary).
func (p *ModelPicker) Primary() *ModelPicker { return p.Variant(Primary) }

// Secondary is Variant(Secondary), a muted fill.
func (p *ModelPicker) Secondary() *ModelPicker { return p.Variant(Secondary) }

// Default is Variant(Default).
func (p *ModelPicker) Default() *ModelPicker { return p.Variant(Default) }

// Native is Variant(Native), the old name of Default.
func (p *ModelPicker) Native() *ModelPicker { return p.Variant(Native) }

// Outline is Variant(Outline).
func (p *ModelPicker) Outline() *ModelPicker { return p.Variant(Outline) }

// Ghost is Variant(Ghost).
func (p *ModelPicker) Ghost() *ModelPicker { return p.Variant(Ghost) }

// Destructive is Variant(Destructive).
func (p *ModelPicker) Destructive() *ModelPicker { return p.Variant(Destructive) }

// Link is Variant(Link).
func (p *ModelPicker) Link() *ModelPicker { return p.Variant(Link) }

// Success is Variant(Success).
func (p *ModelPicker) Success() *ModelPicker { return p.Variant(Success) }

// Info is Variant(Info).
func (p *ModelPicker) Info() *ModelPicker { return p.Variant(Info) }

// Warning is Variant(Warning).
func (p *ModelPicker) Warning() *ModelPicker { return p.Variant(Warning) }

// Size applies the sizing of s and returns the picker, for chaining.
func (p *ModelPicker) Size(s Size) *ModelPicker {
	p.size = s
	p.apply()
	return p
}

// Xs sizes the picker down to the smallest shared size.
func (p *ModelPicker) Xs() *ModelPicker { return p.Size(SizeXs) }

// Sm sizes the picker down.
func (p *ModelPicker) Sm() *ModelPicker { return p.Size(SizeSm) }

// Md gives the built-in sizing.
func (p *ModelPicker) Md() *ModelPicker { return p.Size(SizeMd) }

// Lg sizes the picker up.
func (p *ModelPicker) Lg() *ModelPicker { return p.Size(SizeLg) }

// Shape sets the shape of the trigger and returns the picker, for chaining.
func (p *ModelPicker) Shape(s Shape) *ModelPicker {
	p.shape = s
	p.apply()
	return p
}

// Width pins the trigger to w DIPs and returns the picker, for chaining.
func (p *ModelPicker) Width(w float32) *ModelPicker {
	p.Element = p.Element.Width(w)
	return p
}

// Placeholder shows s in the empty field and returns the picker, for
// chaining.
func (p *ModelPicker) Placeholder(s string) *ModelPicker {
	p.Element = p.Element.Placeholder(s)
	return p
}

// Label names the field for assistive technology and returns the picker.
func (p *ModelPicker) Label(s string) *ModelPicker {
	p.Element = p.Element.Label(s)
	return p
}

// OnChanged calls fn with the ID of a model chosen in this frame.
func (p *ModelPicker) OnChanged(fn func(id string)) *ModelPicker {
	if fn != nil {
		if v, ok := p.parts.Chosen(); ok {
			fn(v)
		}
	}
	return p
}

// Chosen returns the ID of the last chosen model, readable in any frame.
func (p *ModelPicker) Chosen() string {
	if p.chosen != nil {
		return *p.chosen
	}
	return ""
}

// Selected returns the last chosen option, or nil when none was chosen.
func (p *ModelPicker) Selected() *ModelOption {
	id := p.Chosen()
	if id == "" {
		return nil
	}
	for i := range p.options {
		if p.options[i].ID == id {
			return &p.options[i]
		}
	}
	return nil
}

// Changed reports whether a model was chosen in this frame.
func (p *ModelPicker) Changed() bool { return p.Element.Changed() }

// Open reports whether the popup shows.
func (p *ModelPicker) Open() bool { return p.parts.Open() }

// SetOpen opens or closes the popup, and returns the picker for chaining.
func (p *ModelPicker) SetOpen(on bool) *ModelPicker {
	p.parts.SetOpen(on)
	return p
}

// Disabled takes the picker out of use and returns it, for chaining.
func (p *ModelPicker) Disabled(on bool) *ModelPicker {
	p.Element = p.Element.Disabled(on)
	return p
}

// Invalid marks the picker as not holding a good value and returns it.
func (p *ModelPicker) Invalid(on bool) *ModelPicker {
	p.invalid = on
	p.Element = p.Element.Invalid(on)
	p.apply()
	return p
}

// EmptyText sets what the popup shows when nothing matches and returns the
// picker, for chaining.
func (p *ModelPicker) EmptyText(s string) *ModelPicker {
	if p.emptyRef != nil {
		*p.emptyRef = s
	}
	return p
}

// Filter sets the matcher used to narrow the models by what was typed, and
// returns the picker, for chaining. The default matches the name
// case-insensitively.
func (p *ModelPicker) Filter(fn func(o ModelOption, query string) bool) *ModelPicker {
	if p.filterRef != nil {
		*p.filterRef = fn
	}
	return p
}

// Control returns the element a field describes: the trigger input.
func (p *ModelPicker) Control() ui.Element { return p.Element }

// apply paints the trigger's variant and size on this frame's theme.
func (p *ModelPicker) apply() {
	t := p.c.Theme()
	l := formspkg.FrameLookOf(p.variant, t)
	if !l.Native {
		v, h := formspkg.FramePadding(t, p.size)
		border := l.Border
		if p.invalid {
			border = t.Danger
		}
		p.Element = p.Element.Background(l.Bg).Border(l.Width, border).
			Radius(Radius(t, RadiusMd)).Padding(v, h)
	}
	switch p.shape {
	case ShapeRound, ShapeCircle:
		p.Element = p.Element.Radius(1000)
	default:
		p.Element = p.Element.Radius(Radius(t, RadiusMd))
	}
	p.Element = p.Element.FontSize(formspkg.FrameFontSize(t, p.size))
}
