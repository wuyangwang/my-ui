package myui

import (
	"github.com/egoist/mygo/ui"
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
// chaining. The icon takes the color of the alert's ink.
func (a *Alert) Icon(svg *ui.SVG) *Alert {
	a.icon = svg
	a.iconSlot.Children(func() {
		NewIcon(a.c, svg).Size(SizeSm)
	})
	return a
}

// Variant applies the look of v and returns the alert, for chaining: the
// ink of the alert says what kind of message it is.
func (a *Alert) Variant(v Variant) *Alert {
	a.variant = v
	th := a.c.Theme()
	switch v {
	case Destructive:
		a.texts.TextColor(th.Danger)
	case Success, Info, Warning:
		role, _ := semanticOf(v)
		a.texts.TextColor(semanticColor(th, role))
	}
	return a
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
