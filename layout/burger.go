package layout

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
	"github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/icons"
)

// BurgerItem is one row of a Burger menu.
type BurgerItem struct {
	// Key is the value written to Selected when the row is chosen.
	Key string
	// Label is the text drawn in the row.
	Label string
	// Icon, when set, is drawn ahead of the label.
	Icon *ui.SVG
	// Divider draws a separator instead of a clickable row.
	Divider bool
}

// BurgerProps configures NewBurger.
type BurgerProps struct {
	// Items are the rows of the menu, in order.
	Items []BurgerItem
	// Selected is the app's pointer to the active key. A click on a row
	// writes it. It may be nil when only OnSelect is wanted.
	Selected *string
	// OnSelect, when set, also runs when a row is chosen.
	OnSelect func(key string)
	// Fullscreen covers the window with large rows instead of dropping a
	// panel under the header.
	Fullscreen bool
	// HeaderHeight is the gap under the header for the dropped panel;
	// zero uses the theme's Space(16).
	HeaderHeight float32
	// Footer builds a row at the foot of the menu, as a sign-in button.
	Footer func()
}

// Burger is a header's menu button that opens a panel of rows. It is
// my-ui's desktop stand-in for a mobile burger: mygo v0.10.0's ui/mobile
// package has TopBar, Drawer and Sheet but no burger/menu button, so this
// component builds the button and its panel from the public ui API. The
// fullscreen mode uses ui.DialogBase, which is a modal panel; the dropped
// mode uses ui.PopoverBase, which lets a press outside close the menu.
type Burger struct {
	ui.Element
	open *bool
}

// NewBurger builds the trigger first and the panel after it, so a click on
// the trigger lands in the same frame. open is the app's flag; the menu
// writes false when a row is chosen or the user presses outside.
func NewBurger(c *ui.Context, open *bool, p BurgerProps) *Burger {
	b := &Burger{open: open}
	if open == nil {
		return b
	}
	t := c.Theme()
	glyph := icons.Menu()
	if *open {
		glyph = icons.X()
	}
	trigger := controls.NewIconAction(c, controls.IconActionOptions{
		SVG: glyph, Label: Message(c, "menu"), Size: t.Space(8),
	})
	if trigger.Clicked() {
		*open = !*open
	}
	b.Element = trigger
	if !*open {
		return b
	}
	choose := func(key string) {
		if p.Selected != nil {
			*p.Selected = key
		}
		if p.OnSelect != nil {
			p.OnSelect(key)
		}
		*open = false
	}
	if p.Fullscreen {
		panel := ui.DialogBase(c, open, func(backdrop, panel ui.Element) {
			backdrop.Background(t.Background.Alpha(0.98))
			panel.Background(t.Background).Fill()
			panel.Children(func() {
				fullRows(c, open, p, choose)
			})
		})
		b.Element = panel
		return b
	}
	gap := p.HeaderHeight
	if gap == 0 {
		gap = t.Space(16)
	}
	panel := ui.PopoverBase(c, trigger, open, func(panel ui.Element) {
		panel.Background(t.Background).Border(1, t.Border).
			Radius(Radius(t, RadiusLg)).Padding(t.Space(1)).
			Shadow(0, 8, 24, -8, t.Text.Alpha(0.18)).
			Width(t.Space(64)).Gap(t.Space(0.5)).
			Margin(gap, 0, 0, 0).Transition(StateTransition())
		OverlayEnter(c, panel, "burger enter")
		panel.Children(func() {
			rows(c, p, choose, false)
		})
	})
	b.Element = panel
	return b
}

// rows builds the menu's rows inside the panel.
func rows(c *ui.Context, p BurgerProps, choose func(string), large bool) {
	t := c.Theme()
	for _, it := range p.Items {
		if it.Divider {
			ui.Box(c).FillWidth().Height(1).Background(t.Border).
				Margin(t.Space(1), t.Space(2))
			continue
		}
		it := it
		row := ui.Row(c).FillWidth().AlignItems(ui.Center).
			Gap(t.Space(3)).Padding(t.Space(2), t.Space(3)).
			Radius(Radius(t, RadiusMd)).Cursor(ui.CursorPointer).
			Transition(StateTransition())
		active := p.Selected != nil && *p.Selected == it.Key
		if active {
			row = row.Background(t.Accent.Alpha(0.10)).TextColor(t.Accent)
		} else if row.Hovered() {
			row = row.Background(t.SurfaceHover)
		}
		row.Children(func() {
			if it.Icon != nil {
				ink := t.TextMuted
				if active {
					ink = t.Accent
				}
				ui.Icon(c, it.Icon).Size(IconSizeMd, IconSizeMd).TextColor(ink)
			}
			font := TextSize(t, TextSm)
			if large {
				font = TextSize(t, TextLg)
			}
			ui.Text(c, it.Label).FontSize(font)
		})
		if row.Clicked() {
			choose(it.Key)
		}
	}
	if p.Footer != nil {
		ui.Box(c).FillWidth().Height(1).Background(t.Border).
			Margin(t.Space(1), t.Space(2))
		foot := ui.Row(c).FillWidth().Padding(t.Space(1), t.Space(3))
		foot.Children(p.Footer)
	}
}

// fullRows builds the large rows of the fullscreen menu, with a close
// affordance in the corner.
func fullRows(c *ui.Context, open *bool, p BurgerProps, choose func(string)) {
	t := c.Theme()
	head := ui.Row(c).FillWidth().Height(t.Space(14)).
		AlignItems(ui.Center).Justify(ui.End).Padding(0, t.Space(6))
	head.Children(func() {
		close := controls.NewIconAction(c, controls.IconActionOptions{
			SVG: icons.X(), Label: Message(c, "close"), Size: t.Space(10),
		})
		if close.Clicked() && open != nil {
			*open = false
		}
	})
	list := ui.Column(c).Grow(1).Justify(ui.Center).
		Gap(t.Space(2)).Padding(t.Space(6))
	list.Children(func() { rows(c, p, choose, true) })
}
