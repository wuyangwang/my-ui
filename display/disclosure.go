package display

import (
	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

// Accordion groups collapsible items on one card-like surface.
type Accordion struct {
	ui.Element
}

// NewAccordion builds an accordion; items are built by the app inside
// fn with NewAccordionItem.
func NewAccordion(c *ui.Context, fn func()) *Accordion {
	t := c.Theme()
	return &Accordion{Element: ui.Accordion(c, fn).
		Background(t.Background).Border(1, t.Border).
		Radius(Radius(t, RadiusXl)).Padding(t.Space(2)).
		Transition(StateTransition())}
}

// AccordionItem is one titled section of an accordion.
type AccordionItem struct {
	ui.Element
}

// NewAccordionItem builds a titled section bound to *open.
func NewAccordionItem(c *ui.Context, title string, open *bool, fn func()) *AccordionItem {
	t := c.Theme()
	return &AccordionItem{Element: ui.AccordionItem(c, title, open, fn).
		Padding(t.Space(2), t.Space(3)).
		Radius(Radius(t, RadiusMd)).Transition(StateTransition())}
}

// Collapsible is a single section opened by its label.
type Collapsible struct {
	ui.Element
}

// NewCollapsible builds a collapsible section bound to *open.
func NewCollapsible(c *ui.Context, label string, open *bool, fn func()) *Collapsible {
	t := c.Theme()
	return &Collapsible{Element: ui.Collapsible(c, label, open, fn).
		Background(t.Surface).Border(1, t.Border).
		Radius(Radius(t, RadiusMd)).Padding(t.Space(2), t.Space(3)).
		Transition(StateTransition())}
}
