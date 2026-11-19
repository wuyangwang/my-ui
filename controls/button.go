package controls

import (
	. "github.com/wuyangwang/my-ui"
	"time"

	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/icons"
)

// NewButtonGroup gathers buttons into one row, as shadcn's ButtonGroup:
// the buttons sit together, sharing edges.
func NewButtonGroup(c *ui.Context, build func()) ui.Element {
	row := ui.Row(c).Gap(0)
	if build != nil {
		row.Children(build)
	}
	return row
}

// Button is a ui button with a variant's preset look:
//
//	controls.NewButton(c, "Save").Primary()
//	controls.NewButton(c, "Remove").Destructive().Sm()
//
// The chain applies the preset from the build's theme, so it keeps
// following palette and appearance changes; call it every frame, ahead
// of anything else the app styles on the button, as the view runs. What
// a variant does not set keeps the built-in button's look, and every
// method of ui.Element is on hand: Clicked, Disabled (which dims the
// button as the theme dims disabled ones), Children and the rest.
//
// The text of the button is a child of it, as the built-in button's
// own text is, so an icon can sit beside it:
//
//	controls.NewButton(c, "Save").Icon(icons.Save())
//	controls.NewButton(c, "").Icon(icons.Save()).Label("Save")
type Button struct {
	ui.Element
	c       *ui.Context
	icon    ui.Element // the slot of the icon before the text
	text    ui.Element // the text of the button
	end     ui.Element // the slot of the icon after the text
	size    Size
	shape   Shape
	busy    bool
	pressed *bool   // non-nil for a toggle button
	iconSVG *ui.SVG // the icon to restore when the busy state clears
}

// NewButton creates a button with the native preset: controls.NewButton(c,
// "Save"). Clicked() reports the click on it as on ui.Button.
func NewButton(c *ui.Context, label string) *Button {
	b := &Button{c: c}
	btn := ui.Button(c, "")
	btn.Children(func() {
		b.icon = ui.Row(c).AlignItems(ui.Center)
		if label != "" {
			b.text = ui.Text(c, label).SingleLine()
		}
		b.end = ui.Row(c).AlignItems(ui.Center)
	})
	// The room between the icon and the text rides on the icon, so that
	// a button without an icon has no room in it.
	b.Element = btn.Gap(0).Transition(StateTransition())
	// Apply the default scale even without a Variant: a bare button is
	// still a shadcn-sized button, not the framework's built-in one.
	b.applySizing()
	// Size to the content as a button does: a column stretches its
	// children across, which a button never asks for.
	b.AlignSelf(ui.Start)
	return b
}

// look is the colors one variant paints on a build's theme: the fill at
// rest, while hovered and while pressed, and the outline and the ink
// beside them. native marks the built-in look, which the variant keeps
// and paints nothing over.
type look struct {
	native  bool
	rest    ui.Color
	hover   ui.Color
	pressed ui.Color
	width   float32
	border  ui.Color
	text    ui.Color
	hasText bool
}

// lookOf resolves one variant's look on a theme. The sizing of a
// variant is not here: it rides the Size scale, which runs beside the
// variants. The theme's triplet slots feed the primary straight, the
// surface's feed ghost and outline; the destructive and the link Shades
// come from Shades, which is where those rules live.
func lookOf(v Variant, t *ui.Theme) look {
	switch v {
	case Default:
		return look{native: true}
	case Secondary:
		return look{
			rest: t.Surface, hover: t.SurfaceHover, pressed: t.SurfacePressed,
			width: 0, border: t.Border,
		}
	case Primary:
		return look{
			rest: t.Accent, hover: t.AccentHover, pressed: t.AccentPressed,
			width: 0, border: t.Border, text: t.AccentText, hasText: true,
		}
	case Outline:
		return look{
			rest: t.Background, hover: t.SurfaceHover, pressed: t.SurfacePressed,
			width: 1, border: t.Border,
		}
	case Ghost:
		return look{
			rest: ui.Transparent, hover: t.SurfaceHover, pressed: t.SurfacePressed,
			width: 0, border: t.Border,
		}
	case Destructive:
		hover, pressed, ink := Shades(t.Danger, t.Dark)
		return look{
			rest: t.Danger, hover: hover, pressed: pressed,
			width: 0, border: t.Border, text: ink, hasText: true,
		}
	case Link:
		a := t.Accent
		return look{
			rest: ui.Transparent, hover: a.Alpha(0.08), pressed: a.Alpha(0.16),
			width: 0, border: t.Border, text: a, hasText: true,
		}
	case Success, Info, Warning:
		role, _ := SemanticOf(v)
		base := SemanticColor(t, role)
		hover, pressed, ink := Shades(base, t.Dark)
		return look{
			rest: base, hover: hover, pressed: pressed,
			width: 0, border: t.Border, text: ink, hasText: true,
		}
	}
	// An unknown variant reads as the built-in look.
	return look{native: true}
}

// Variant applies the look of v and returns the button, for chaining.
// The zero variant is Native.
func (b *Button) Variant(v Variant) *Button {
	b.applySizing()
	if l := lookOf(v, b.c.Theme()); !l.native {
		b.Background(l.rest).Border(l.width, l.border)
		if l.hasText {
			b.TextColor(l.text)
		}
		// Hover and pressed paint over the rest, each frame, where the
		// built-in button's own state colors do not reach a preset's
		// rebuilt look. MyGo's transitions fade the changes as the
		// pointer comes and goes, and a disabled button keeps its rest
		// colors as the theme dims it.
		if !b.IsDisabled() {
			switch {
			case b.Pressed():
				b.Background(l.pressed)
			case b.Hovered():
				b.Background(l.hover)
			}
		}
	}
	return b
}

// Primary is Variant(myui.Primary).
func (b *Button) Primary() *Button { return b.Variant(Primary) }

// Secondary is Variant(myui.Secondary), the shadcn secondary look: a
// muted fill of the theme's surface, without a border.
func (b *Button) Secondary() *Button { return b.Variant(Secondary) }

// Default is Variant(myui.Default), the built-in look.
func (b *Button) Default() *Button { return b.Variant(Default) }

// Native is Variant(myui.Native), the old name of Default.
func (b *Button) Native() *Button { return b.Variant(Native) }

// Outline is Variant(myui.Outline).
func (b *Button) Outline() *Button { return b.Variant(Outline) }

// Ghost is Variant(myui.Ghost).
func (b *Button) Ghost() *Button { return b.Variant(Ghost) }

// Destructive is Variant(myui.Destructive).
func (b *Button) Destructive() *Button { return b.Variant(Destructive) }

// Link is Variant(myui.Link).
func (b *Button) Link() *Button { return b.Variant(Link) }

// Success is Variant(myui.Success).
func (b *Button) Success() *Button { return b.Variant(Success) }

// Info is Variant(myui.Info).
func (b *Button) Info() *Button { return b.Variant(Info) }

// Warning is Variant(myui.Warning).
func (b *Button) Warning() *Button { return b.Variant(Warning) }

// Size applies the sizing of s and returns the button, for chaining.
func (b *Button) Size(s Size) *Button {
	b.size = s
	b.applySizing()
	return b
}

// Sm sizes the button down.
func (b *Button) Sm() *Button { return b.Size(SizeSm) }

// Xs sizes the button to the smallest size.
func (b *Button) Xs() *Button { return b.Size(SizeXs) }

// Md gives the sizing of the built-in button.
func (b *Button) Md() *Button { return b.Size(SizeMd) }

// Lg sizes the button up.
func (b *Button) Lg() *Button { return b.Size(SizeLg) }

// Icon puts an icon before the text of the button and returns the
// button, for chaining. The icon takes the color of the text and the
// size of it. A button with no text takes the icon alone, in the middle
// of it.
func (b *Button) Icon(svg *ui.SVG) *Button {
	c := b.c
	b.iconSVG = svg
	b.icon.Children(func() {
		icon := ui.Icon(c, svg)
		// The room between the icon and the text rides on the icon, and
		// a button with no text has no room to keep.
		if b.text.Valid() {
			icon.Margin(0, c.Theme().Space(1.5), 0, 0)
		}
	})
	return b
}

// EndIcon puts an icon after the text of the button and returns the
// button, for chaining. A button with no text takes the icon alone, in
// the middle of it.
func (b *Button) EndIcon(svg *ui.SVG) *Button {
	c := b.c
	b.end.Children(func() {
		icon := ui.Icon(c, svg)
		if b.text.Valid() {
			icon.Margin(0, 0, 0, c.Theme().Space(1.5))
		}
	})
	return b
}

// Busy sets or clears the busy state and returns the button, for
// chaining. While busy a spinner takes the place of the icon before the
// text, and the button takes no clicks. The icon is restored when the
// state clears.
func (b *Button) Busy(on bool) *Button {
	c := b.c
	b.busy = on
	b.icon.Children(func() {
		if on {
			spinner := ui.Icon(c, icons.LoaderCircle())
			spinner = spinner.Rotate(spinner.Loop("spin", time.Second, ui.Linear) * 360)
			if b.text.Valid() {
				spinner.Margin(0, c.Theme().Space(1.5), 0, 0)
			}
			return
		}
		if b.iconSVG != nil {
			ui.Icon(c, b.iconSVG)
		}
	})
	b.Element = b.Element.Disabled(on)
	return b
}

// Loading turns the button into a busy one and returns it, for chaining:
// the same as Busy(true).
func (b *Button) Loading() *Button { return b.Busy(true) }

// Toggle makes the button a two-state control reflecting *pressed, as
// shadcn's Toggle: a click flips the flag, and assistive technology sees
// the pressed state. It returns the button for chaining.
func (b *Button) Toggle(pressed *bool) *Button {
	b.pressed = pressed
	b.Element = b.Element.Role(ui.RoleToggleButton)
	if pressed != nil {
		b.Element = b.Element.AccessSelected(*pressed).Checked(*pressed)
		if clicked := b.Clicked(); clicked {
			*pressed = !*pressed
		}
	}
	return b
}

// Label names the button for assistive technology and returns it, for
// chaining. A button that shows only an icon needs a name: the name of
// what the icon stands for.
func (b *Button) Label(s string) *Button {
	b.Element = b.Element.Label(s)
	return b
}

// applySizing sizes the button on the theme of its build. The paddings
// of SizeMd are the built-in button's; the others move along the
// spacing scale the theme's Space says.
func (b *Button) applySizing() {
	t := b.c.Theme()
	switch b.size {
	case SizeXs:
		b.Padding(t.Space(0.5), t.Space(2)).FontSize(TextSize(t, TextXs)).Height(28)
	case SizeSm:
		b.Padding(t.Space(1), t.Space(2.5)).FontSize(TextSize(t, TextSm)).Height(32)
	case SizeLg:
		b.Padding(t.Space(2.5), t.Space(4.5)).FontSize(TextSize(t, TextLg)).Height(40)
	default:
		b.Padding(t.Space(1.5), t.Space(3.5)).FontSize(t.FontSize).Height(36)
	}
	b.applyShape()
}

// Width sets the width of the button in DIPs and returns it, for
// chaining.
func (b *Button) Width(w float32) *Button {
	b.Element = b.Element.Width(w)
	return b
}

// Block makes the button fill the width of its row and returns it, for
// chaining.
func (b *Button) Block() *Button {
	b.Element = b.Element.AlignSelf(ui.Stretch)
	return b
}

// Shape sets the shape of the button and returns it, for chaining: a
// pill or a circle, or the theme's own corners.
func (b *Button) Shape(s Shape) *Button {
	b.shape = s
	b.applyShape()
	return b
}

// applyShape rounds the corners of the button as its shape says.
func (b *Button) applyShape() {
	t := b.c.Theme()
	switch b.shape {
	case ShapeRound, ShapeCircle:
		b.Radius(1000)
	default:
		b.Radius(Radius(t, RadiusMd))
	}
}
