package display

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/icons"
)

// AlertType is the kind of message an Alert carries. It picks the tone
// (the color of the icon and the tint of a Soft surface) and the default
// icon. It is independent of AlertVariant: any type can wear any surface.
type AlertType uint8

const (
	// AlertInfo is a neutral, informational message (the zero value).
	AlertInfo AlertType = iota
	// AlertSuccess reports that something went well.
	AlertSuccess
	// AlertWarning calls attention to something that needs care.
	AlertWarning
	// AlertError reports a failure.
	AlertError
	// AlertSecondary is a quiet, low-emphasis message.
	AlertSecondary
)

// AlertVariant is the surface an Alert is painted on. It is orthogonal to
// AlertType: the type chooses the hue, the variant chooses how much of it
// shows.
type AlertVariant uint8

const (
	// AlertSoft tints the surface with the type's color and draws a faint
	// ring of the same hue (the zero value).
	AlertSoft AlertVariant = iota
	// AlertOutlined draws a border in the theme's color over a plain face.
	AlertOutlined
	// AlertPlain shows no surface at all, only the content.
	AlertPlain
)

// Alert is a message that stands in the page: an icon, a title, an
// optional description and an optional action. Its two axes are
// independent — Type picks the color and the glyph, Variant picks the
// surface:
//
//	display.NewAlert(c, "注意").Type(AlertWarning).Variant(AlertSoft)
//	display.NewAlert(c, "出错了").Error().NoIcon(true)
//
// It is a live region (RoleAlert by default), so a screen reader
// announces it as it comes. The app's closed flag hides the alert when it
// is set to true. Colors are read from c.Theme() on every frame.
type Alert struct {
	ui.Element
	c           *ui.Context
	title       string
	description string
	typ         AlertType
	variant     AlertVariant
	noIcon      bool
	selectable  bool
	closed      *bool
	role        ui.Role

	root      ui.Element
	texts     ui.Element
	titleText ui.Element
	descSlot  ui.Element
	descText  ui.Element
	action    ui.Element
	icons     [5]ui.Element // one slot per AlertType
	iconEls   [5]ui.Element // the glyph inside each slot
	custom    ui.Element    // the slot filled by Icon
	customOn  bool          // whether Icon supplied a glyph
}

// NewAlert creates an alert with a title, using AlertInfo on an AlertSoft
// surface. The chain runs in the frame the view builds, so call it before
// any app-specific styling.
func NewAlert(c *ui.Context, title string) *Alert {
	a := &Alert{c: c, title: title, typ: AlertInfo, variant: AlertSoft, role: ui.RoleAlert}
	t := c.Theme()

	root := ui.Box(c).Row().AlignItems(ui.Start).Gap(t.Space(3)).AlignSelf(ui.Stretch).
		Padding(t.Space(3), t.Space(4)).FontSize(TextSize(t, TextSm)).
		Transition(StateTransition())
	root.Children(func() {
		// All five glyphs are built once; the active one is shown and the
		// rest are Hidden. That lets Type change the icon without adding
		// children again on a later call in the same frame.
		for k := range a.icons {
			slot := ui.Row(c).AlignItems(ui.Start).Shrink(0)
			k := k
			slot.Children(func() {
				a.iconEls[k] = ui.Icon(c, alertGlyph(AlertType(k))).FontSize(alertGlyphSize(false))
			})
			a.icons[k] = slot
		}
		a.custom = ui.Row(c).AlignItems(ui.Start).Shrink(0).Hidden(true)

		a.texts = ui.Column(c).Grow(1).MinWidth(0)
		a.texts.Children(func() {
			a.titleText = ui.Text(c, title).Bold()
			a.descSlot = ui.Box(c)
		})
		a.action = ui.Row(c).AlignItems(ui.Center).Shrink(0)
	})
	a.root = root
	a.Element = root
	a.apply()
	return a
}

// Type chooses the message's kind: the hue of a Soft surface, the ink of
// the icon and the glyph shown at the left. It is independent of Variant.
func (a *Alert) Type(t AlertType) *Alert {
	a.typ = t
	a.apply()
	return a
}

// Variant chooses the surface: Soft (tinted), Outlined (bordered) or Plain
// (bare). It is independent of Type.
func (a *Alert) Variant(v AlertVariant) *Alert {
	a.variant = v
	a.apply()
	return a
}

// Description adds the text under the title and returns the alert. An
// empty string clears it.
func (a *Alert) Description(s string) *Alert {
	a.description = s
	a.descSlot.Children(func() {
		if s != "" {
			a.descText = ui.Text(a.c, s).TextColor(a.descriptionInk())
		}
	})
	if a.selectable && a.descText.Valid() {
		a.descText.Selectable()
	}
	a.apply()
	return a
}

// Icon replaces the type's glyph with svg and returns the alert. The
// custom glyph takes the type's ink. Call with nil to go back to the
// type's own icon.
func (a *Alert) Icon(svg *ui.SVG) *Alert {
	if svg == nil {
		a.customOn = false
		a.custom.Hidden(true)
		a.apply()
		return a
	}
	a.custom.Children(func() {
		ui.Icon(a.c, svg).FontSize(alertGlyphSize(a.description != "")).TextColor(alertTone(a.c, a.typ))
	})
	a.customOn = true
	a.custom.Hidden(false)
	a.apply()
	return a
}

// NoIcon hides (or shows) the glyph at the left and returns the alert.
func (a *Alert) NoIcon(on bool) *Alert {
	a.noIcon = on
	a.apply()
	return a
}

// Selectable lets people select and copy the title and description, for a
// message whose text they may pass on, such as an error's details.
func (a *Alert) Selectable(on bool) *Alert {
	a.selectable = on
	if a.titleText.Valid() {
		a.titleText.Selectable()
	}
	if a.descText.Valid() {
		a.descText.Selectable()
	}
	return a
}

// Closed shows a close button that writes true to flag, and returns the
// alert. The app hides the alert when flag is true; while it is true the
// alert renders nothing.
func (a *Alert) Closed(flag *bool) *Alert {
	a.closed = flag
	if a.action.Valid() {
		a.action.Children(func() {
			ctrlpkg.NewIconAction(a.c, ctrlpkg.IconActionOptions{
				SVG:      icons.X(),
				Label:    Message(a.c, "close"),
				IconSize: IconSizeSm,
				OnClick: func() {
					if a.closed != nil {
						*a.closed = true
					}
				},
			})
		})
	}
	a.apply()
	return a
}

// Closable is the old name of Closed, kept for code that uses it.
func (a *Alert) Closable(flag *bool) *Alert { return a.Closed(flag) }

// Action builds a control at the right of the alert — shadcn's
// AlertAction — a link or a small button the reader can act on.
func (a *Alert) Action(build func()) *Alert {
	if a.action.Valid() && build != nil {
		a.action.Children(build)
	}
	return a
}

// Role sets the accessibility role — RoleAlert (the default, an urgent
// live region) or RoleStatus for a non-urgent message — and returns it.
func (a *Alert) Role(r ui.Role) *Alert {
	a.role = r
	a.Element = a.Element.Role(r)
	return a
}

// The type presets set Type, not the surface: they are shorthands for the
// hue and glyph. Combine them with a Variant for a different surface.

// Info sets Type to AlertInfo and returns the alert.
func (a *Alert) Info() *Alert { return a.Type(AlertInfo) }

// Success sets Type to AlertSuccess and returns the alert.
func (a *Alert) Success() *Alert { return a.Type(AlertSuccess) }

// Warning sets Type to AlertWarning and returns the alert.
func (a *Alert) Warning() *Alert { return a.Type(AlertWarning) }

// Error sets Type to AlertError and returns the alert.
func (a *Alert) Error() *Alert { return a.Type(AlertError) }

// Destructive is the old name of Error, kept for code that uses it.
func (a *Alert) Destructive() *Alert { return a.Type(AlertError) }

// Secondary sets Type to AlertSecondary and returns the alert.
func (a *Alert) Secondary() *Alert { return a.Type(AlertSecondary) }

// Soft sets Variant to AlertSoft and returns the alert.
func (a *Alert) Soft() *Alert { return a.Variant(AlertSoft) }

// Outlined sets Variant to AlertOutlined and returns the alert.
func (a *Alert) Outlined() *Alert { return a.Variant(AlertOutlined) }

// Plain sets Variant to AlertPlain and returns the alert.
func (a *Alert) Plain() *Alert { return a.Variant(AlertPlain) }

// Default is the built-in look: AlertInfo on AlertSoft.
func (a *Alert) Default() *Alert { return a.Type(AlertInfo).Variant(AlertSoft) }

// Native is the old name of Default, kept for code that uses it.
func (a *Alert) Native() *Alert { return a.Default() }

// apply paints the current type and surface on the elements built this
// frame. It never rebuilds the row, so the close control keeps its click.
func (a *Alert) apply() {
	if !a.root.Valid() {
		return
	}
	t := a.c.Theme()
	accent := alertTone(a.c, a.typ)

	// The surface. Soft tints the theme's face with the type's hue and
	// draws a faint ring of it; Outlined uses the theme's border; Plain
	// paints nothing. Every color comes from this frame's theme.
	face := t.Surface
	switch a.variant {
	case AlertSoft:
		face = t.Surface.MixOklab(accent, 0.06)
		a.root = a.root.Background(face).Border(0, ui.Transparent).
			InsetShadow(0, 0, 0, 1, accent.Alpha(0.16))
	case AlertOutlined:
		a.root = a.root.Background(ui.Transparent).Border(1, t.Border).
			InsetShadow(0, 0, 0, 0, ui.Transparent)
	case AlertPlain:
		a.root = a.root.Background(ui.Transparent).Border(0, ui.Transparent).
			InsetShadow(0, 0, 0, 0, ui.Transparent)
	}
	if a.variant == AlertPlain {
		a.root = a.root.Padding(t.Space(2), 0)
	} else if a.description != "" {
		a.root = a.root.Padding(t.Space(3), t.Space(4))
	} else {
		a.root = a.root.Padding(t.Space(2.5), t.Space(3))
	}

	// Ink: the title is the normal text, the description is muted, and
	// the icon carries the type's hue.
	if a.titleText.Valid() {
		a.titleText = a.titleText.TextColor(t.Text)
	}
	if a.descText.Valid() {
		a.descText = a.descText.TextColor(a.descriptionInk())
	}
	for k, slot := range a.icons {
		visible := !a.noIcon && !a.customOn && AlertType(k) == a.typ
		slot.Hidden(!visible)
		if a.iconEls[k].Valid() {
			a.iconEls[k] = a.iconEls[k].TextColor(accent)
		}
	}
	a.Element = a.root.Role(a.role)
	if a.closed != nil && *a.closed {
		a.Element = a.Element.Hidden(true)
	}
}

// descriptionInk is the muted ink of the description: the theme's muted
// text, tinted a little toward the type on surfaces that carry a hue.
func (a *Alert) descriptionInk() ui.Color {
	t := a.c.Theme()
	if a.variant == AlertPlain {
		return t.TextMuted
	}
	return t.TextMuted.MixOklab(alertTone(a.c, a.typ), 0.12)
}

// alertTone is the hue of a type, read from this frame's theme. Information
// falls back to the library's Info token, since MyGo's Theme has no Info
// slot; Secondary uses the muted text.
func alertTone(c *ui.Context, typ AlertType) ui.Color {
	t := c.Theme()
	switch typ {
	case AlertSuccess:
		return SemanticColor(t, SemanticSuccess)
	case AlertWarning:
		return SemanticColor(t, SemanticWarning)
	case AlertError:
		return t.Danger
	case AlertSecondary:
		return t.TextMuted
	}
	return SemanticColor(t, SemanticInfo)
}

// alertGlyph is the default icon of a type.
func alertGlyph(typ AlertType) *ui.SVG {
	switch typ {
	case AlertSuccess:
		return icons.CircleCheck()
	case AlertWarning:
		return icons.TriangleAlert()
	case AlertError:
		return icons.CircleX()
	case AlertSecondary:
		return icons.CircleHelp()
	}
	return icons.Info()
}

func alertGlyphSize(detailed bool) float32 {
	if detailed {
		return IconSizeMd
	}
	return IconSizeSm
}
