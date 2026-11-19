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

// NewMenubar builds one menu button per menu, in a row. An optional key
// scopes the row: pass a different one for each Menubar under the same
// parent, or the two share an identity and the later one overwrites the
// earlier. With no key the row uses its position in the parent, which is
// enough for a single menubar.
func NewMenubar(c *ui.Context, menus []MenubarMenu, key ...string) *Menubar {
	if len(key) > 0 && key[0] != "" {
		c = c.Key(key[0])
	}
	t := c.Theme()
	row := ui.Row(c).AlignItems(ui.Center).Gap(t.Space(1))
	row.Children(func() {
		for _, menu := range menus {
			open := menu.Open
			if open == nil {
				// A fresh pointer each frame loses the menu's state;
				// keep it on the row, keyed by the menu's label.
				open = ui.Local(row, "open "+menu.Label, func() bool { return false })
			}
			// Key each menu by its label so a menu keeps its identity
			// even if the row's menus are reordered.
			parts := ui.MenuBase(c.Key("menubar "+menu.Label), open)
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
