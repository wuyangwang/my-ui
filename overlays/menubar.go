package overlays

import (
	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

// MenubarMenu is one pull-down of a Menubar. Open may be nil, in which
// case the menu keeps its own state inside mygo's MenuBase.
type MenubarMenu struct {
	Label string
	Open  *bool
	Items []MenuItem
}

// Menubar is a row of menus that share DropdownMenu's panel and item
// look, as shadcn's Menubar does.
type Menubar struct {
	ui.Element
}

// NewMenubar builds one menu button per menu, in a row.
func NewMenubar(c *ui.Context, menus []MenubarMenu) *Menubar {
	t := c.Theme()
	row := ui.Row(c.Key("menubar")).AlignItems(ui.Center).Gap(t.Space(1))
	row.Children(func() {
		for _, menu := range menus {
			open := menu.Open
			if open == nil {
				// A fresh pointer each frame loses the menu's state;
				// keep it on the row, keyed by the menu's label.
				open = ui.Local(row, "open "+menu.Label, func() bool { return false })
			}
			parts := ui.MenuBase(c, open)
			parts.Trigger.Padding(t.Space(2), t.Space(3)).
				Radius(Radius(t, RadiusSm)).Children(func() {
				ui.Text(c, menu.Label).FontSize(TextSize(t, TextSm))
			})
			parts.Popup(func(panel ui.Element) {
				buildMenuPanel(c, t, panel)
				buildMenuItems(c, t, parts, menu.Items, open)
			})
		}
	})
	return &Menubar{Element: row}
}
