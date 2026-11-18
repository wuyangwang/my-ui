package overlays

import (
	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

// MenuItem is one row of a dropdown or context menu. A Separator item
// draws a hairline instead, and OnClick runs when the row is chosen.
type MenuItem struct {
	Label     string
	Icon      *ui.SVG
	Disabled  bool
	Separator bool
	OnClick   func()
}

// buildMenuPanel gives a mygo menu panel shadcn's look: the theme's
// background, a hairline border, the theme's radius and a soft shadow.
func buildMenuPanel(c *ui.Context, t *ui.Theme, panel ui.Element) {
	panel.Padding(t.Space(1)).Radius(Radius(t, RadiusMd)).
		Background(t.Background).Border(1, t.Border).
		Shadow(0, 4, 16, -4, t.Text.Alpha(0.14))
}

// buildMenuItems fills a mygo MenuBase/ContextMenuBase panel with the
// app's rows. Highlighted rows take a tint of the accent and its ink,
// like shadcn's menu item; mygo still owns opening, keyboard nav and
// Clicked().
func buildMenuItems(c *ui.Context, t *ui.Theme, parts ui.MenuParts, items []MenuItem, open *bool) {
	for _, item := range items {
		if item.Separator {
			ui.Box(c).Height(1).FillWidth().Background(t.Border).
				Margin(t.Space(1), 0)
			continue
		}
		// shadcn's menu item is px-2 py-1: keep the rows compact, mygo's
		// own item already carries some inset.
		row := parts.Item().Padding(t.Space(1), t.Space(2)).Radius(Radius(t, RadiusSm))
		switch {
		case item.Disabled:
			row = row.Disabled(true).TextColor(t.TextMuted)
		case row.Highlighted():
			row = row.Background(t.Accent.Alpha(0.08)).TextColor(t.Accent)
		default:
			row = row.TextColor(t.Text)
		}
		row.Children(func() {
			if item.Icon != nil {
				ui.Icon(c, item.Icon).Size(IconSizeSm, IconSizeSm).
					Margin(0, t.Space(1.5), 0, 0)
			}
			ui.Text(c, item.Label).FontSize(TextSize(t, TextSm))
		})
		if row.Clicked() {
			if item.OnClick != nil {
				item.OnClick()
			}
			if open != nil {
				*open = false
			}
		}
	}
}
