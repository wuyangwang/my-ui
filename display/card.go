package display

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

// Card is a panel of the page: a heading, a description, content and an
// optional footer, on a surface with a border.
//
//	display.NewCard(c, "概览").Description("今天的情况").
//		Content(func() { ... }).Footer(func() { ... })
//
// The card is styled by the library, as shadcn's: a surface with a
// border, rounded, with room in it. Shadowed() adds shadcn's soft shadow.
// NewCardPanel makes a card with no heading, for free-form content.
type Card struct {
	ui.Element
	c                      *ui.Context
	head, headText, action ui.Element
	desc, content, foot    ui.Element
	shadowed               bool
}

// NewCard creates a card with a title.
func NewCard(c *ui.Context, title string) *Card {
	return newCard(c, title)
}

// NewCardPanel creates a card with no heading, for content that supplies
// its own header, as shadcn's Card does when CardHeader is skipped.
func NewCardPanel(c *ui.Context) *Card {
	return newCard(c, "")
}

func newCard(c *ui.Context, title string) *Card {
	card := &Card{c: c}
	t := c.Theme()
	col := ui.Column(c).Background(t.Surface).Border(1, t.Border).
		FillWidth().Radius(Radius(t, RadiusXl)).Padding(t.Space(5)).Gap(0)
	col.Children(func() {
		card.head = ui.Row(c).AlignItems(ui.Center).Justify(ui.SpaceBetween).
			Gap(t.Space(3))
		card.head.Children(func() {
			left := ui.Column(c).Gap(0).Grow(1)
			left.Children(func() {
				card.headText = ui.Row(c)
				if title != "" {
					card.headText.Children(func() {
						ui.Text(c, title).FontSize(TextSize(t, TextLg)).Bold()
					})
				}
				card.desc = ui.Row(c)
			})
			card.action = ui.Row(c)
		})
		card.content = ui.Column(c)
		card.foot = ui.Row(c)
	})
	card.Element = col
	return card
}

// Title sets the heading of the card, for chaining.
func (card *Card) Title(s string) *Card {
	if card.headText.Valid() {
		card.headText.Children(func() {
			ui.Text(card.c, s).FontSize(TextSize(card.c.Theme(), TextLg)).Bold()
		})
	}
	return card
}

// Description adds the text under the title of the card.
func (card *Card) Description(s string) *Card {
	if s == "" || !card.desc.Valid() {
		return card
	}
	t := card.c.Theme()
	card.desc.Children(func() {
		ui.Text(card.c, s).TextColor(t.TextMuted).Margin(t.Space(1), 0, 0, 0)
	})
	return card
}

// Header builds the heading area of the card, replacing the default
// title row: use it for a custom title, a badge or a CardAction.
func (card *Card) Header(build func()) *Card {
	if card.head.Valid() && build != nil {
		card.head.Children(build)
	}
	return card
}

// Action puts content at the right of the heading, as shadcn's
// CardAction does — a menu button beside the title.
func (card *Card) Action(build func()) *Card {
	if card.action.Valid() && build != nil {
		card.action.Children(build)
	}
	return card
}

// Content puts what the card shows in it, and returns the card for
// chaining.
func (card *Card) Content(build func()) *Card {
	if card.content.Valid() && build != nil {
		card.content.Children(build)
	}
	return card
}

// Footer builds the row at the foot of the card, separated by a hairline,
// for the card's actions.
func (card *Card) Footer(build func()) *Card {
	if !card.foot.Valid() || build == nil {
		return card
	}
	t := card.c.Theme()
	card.foot = card.foot.Margin(t.Space(4), 0, 0, 0).
		Border(0, t.Border).Padding(t.Space(3), 0, 0, 0)
	card.foot.Children(build)
	return card
}

// Shadowed adds shadcn's soft shadow to the card and returns it, for
// chaining.
func (card *Card) Shadowed() *Card {
	if !card.shadowed {
		t := card.c.Theme()
		card.Element = card.Element.
			Shadow(0, 1, 2, 0, t.Text.Alpha(0.06)).
			Shadow(0, 8, 24, -8, t.Text.Alpha(0.10))
		card.shadowed = true
	}
	return card
}
