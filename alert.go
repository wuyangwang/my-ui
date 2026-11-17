package myui

import (
	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/icons"
)

// Alert is a message that stands in the page: a card with a title, a
// description and an icon.
//
//	myui.NewAlert(c, "注意").Description("这条信息需要留意").Warning()
//
// The alert is styled by the library, as shadcn's: a bordered card with
// the icon at its left, the title and the description beside it. The
// variant says what kind of message it is: Default, Destructive,
// Success, Warning or Info. It is a live region, so a screen reader
// announces it as it comes.
type Alert struct {
	ui.Element
	c       *ui.Context
	title   string
	icon    *ui.SVG
	variant Variant

	texts     ui.Element // the title and the description
	iconSlot  ui.Element // the slot of the icon at the left
	descSlot  ui.Element // the slot of the description
	action    ui.Element // the slot of the action at the right
	closed    *bool      // the app's flag, set true when closed
	inherited bool
}

// NewAlert creates an alert with a title.
func NewAlert(c *ui.Context, title string) *Alert {
	a := &Alert{c: c, title: title}
	th := c.Theme()
	alert := ui.Box(c).Row().AlignItems(ui.Start).Gap(th.Space(3)).
		Background(th.Surface).Border(1, th.Border).Radius(th.Radius).
		Padding(th.Space(3), th.Space(4)).
		FontSize(TextSize(th, TextSm)).
		Role(ui.RoleAlert).
		Transition(stateTransition())
	alert.Children(func() {
		a.iconSlot = ui.Row(c).AlignItems(ui.Start)
		a.texts = ui.Column(c).Grow(1).Gap(0).Children(func() {
			ui.Text(c, title).Bold()
			a.descSlot = ui.Box(c)
		})
		a.action = ui.Row(c).AlignItems(ui.Center)
	})
	a.Element = alert
	return a
}

// Description adds the text under the title of the alert.
func (a *Alert) Description(s string) *Alert {
	if s == "" {
		return a
	}
	th := a.c.Theme()
	a.descSlot.Children(func() {
		ui.Text(a.c, s).
			TextColor(a.descriptionColor()).
			Margin(th.Space(1), 0, 0, 0)
	})
	return a
}

// Icon puts an icon at the left of the alert and returns it, for
// chaining. The icon takes the color of the alert's ink, as shadcn's
// `[&>svg]:text-current` does, so a destructive alert's icon is red.
func (a *Alert) Icon(svg *ui.SVG) *Alert {
	a.icon = svg
	a.iconSlot.Children(func() {
		ic := NewIcon(a.c, svg).Size(SizeSm)
		if ink, ok := a.variantInk(); ok {
			ic.TextColor(ink)
		}
	})
	return a
}

// Variant applies the look of v and returns the alert, for chaining: the
// ink of the alert says what kind of message it is. The title, the
// description and the icon all take it.
func (a *Alert) Variant(v Variant) *Alert {
	a.variant = v
	if ink, ok := a.variantInk(); ok {
		a.texts.TextColor(ink)
		a.iconSlot.TextColor(ink)
	}
	return a
}

// variantInk is the ink the variant paints, if it paints one.
func (a *Alert) variantInk() (ui.Color, bool) {
	th := a.c.Theme()
	switch a.variant {
	case Destructive:
		return th.Danger, true
	case Success, Info, Warning:
		role, _ := semanticOf(a.variant)
		return semanticColor(th, role), true
	}
	return ui.Color{}, false
}

// Default is Variant(myui.Default), the built-in look.
func (a *Alert) Default() *Alert { return a.Variant(Default) }

// Native is Variant(myui.Native), the old name of Default.
func (a *Alert) Native() *Alert { return a.Variant(Native) }

// Destructive is Variant(myui.Destructive): a message of an error.
func (a *Alert) Destructive() *Alert { return a.Variant(Destructive) }

// Success is Variant(myui.Success).
func (a *Alert) Success() *Alert { return a.Variant(Success) }

// Info is Variant(myui.Info).
func (a *Alert) Info() *Alert { return a.Variant(Info) }

// Warning is Variant(myui.Warning).
func (a *Alert) Warning() *Alert { return a.Variant(Warning) }

// Action builds a control at the right of the alert — shadcn's
// AlertAction — a link or a small button the reader can act on.
func (a *Alert) Action(build func()) *Alert {
	if a.action.Valid() && build != nil {
		a.action.Children(build)
	}
	return a
}

// ShowIcon puts the icon of the variant at the left when none was given
// with Icon, and returns the alert, for chaining. The icon follows the
// variant's ink.
func (a *Alert) ShowIcon(on bool) *Alert {
	if !on || a.icon != nil {
		return a
	}
	svg := icons.Info()
	switch a.variant {
	case Destructive:
		svg = icons.CircleX()
	case Success:
		svg = icons.CircleCheck()
	case Warning:
		svg = icons.TriangleAlert()
	}
	return a.Icon(svg)
}

// Closable puts a close button at the right of the alert, and returns it,
// for chaining. Clicking it sets *closed to true; the app hides the alert
// when it does.
func (a *Alert) Closable(closed *bool) *Alert {
	a.closed = closed
	if !a.action.Valid() {
		return a
	}
	a.action.Children(func() {
		b := NewButton(a.c, "").Ghost().Xs()
		b.Icon(icons.X()).Label(Message(a.c, "close"))
		if b.Clicked() && a.closed != nil {
			*a.closed = true
		}
	})
	return a
}

// Role sets the accessibility role of the alert — RoleAlert (the
// default) or RoleStatus for a non-urgent message — and returns it.
func (a *Alert) Role(r ui.Role) *Alert {
	a.Element = a.Element.Role(r)
	return a
}

// descriptionColor is the color of the description: the muted ink, or
// the ink of the variant at nine tenths, as shadcn's destructive alert
// does.
func (a *Alert) descriptionColor() ui.Color {
	th := a.c.Theme()
	switch a.variant {
	case Destructive:
		return th.Danger.Alpha(0.9)
	case Success, Info, Warning:
		role, _ := semanticOf(a.variant)
		return semanticColor(th, role).Alpha(0.9)
	}
	return th.TextMuted
}
