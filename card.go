package myui

import (
	"github.com/egoist/mygo/ui"
)

// Card is a panel of the page: a surface with a title, a description
// and what the app puts in it.
//
//	myui.NewCard(c, "概览").Description("今天的情况").
//		Content(func() { ... })
//
// The card is styled by the library, as shadcn's: a surface with a
// border, rounded, with room in it.
type Card struct {
	ui.Element
	c    *ui.Context
	desc ui.Element
}

// NewCard creates a card with a title.
func NewCard(c *ui.Context, title string) *Card {
	card := &Card{c: c}
	t := c.Theme()
	col := ui.Column(c).Background(t.Surface).Border(1, t.Border).Radius(t.Radius).
		Padding(t.Space(4)).Gap(0)
	col.Children(func() {
		ui.Text(c, title).FontSize(TextSize(t, TextLg)).Bold()
		card.desc = ui.Box(c)
	})
	card.Element = col
	return card
}

// Description adds the text under the title of the card.
func (card *Card) Description(s string) *Card {
	if s == "" {
		return card
	}
	t := card.c.Theme()
	card.desc.Children(func() {
		ui.Text(card.c, s).TextColor(t.TextMuted).Margin(t.Space(1), 0, 0, 0)
	})
	return card
}

// Content puts what the card shows in it, and returns the card, for
// chaining.
func (card *Card) Content(build func()) *Card {
	card.Element.Children(func() {
		if build != nil {
			build()
		}
	})
	return card
}
