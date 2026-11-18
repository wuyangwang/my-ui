package overlays

import "github.com/egoist/mygo/ui"

// DropdownMenu is mygo's MenuBase with shadcn's panel and item look.
// The open flag belongs to the app, as in ui.MenuBase.
type DropdownMenu struct {
	ui.Element
}

// NewDropdownMenu builds a menu button around trigger and a panel of
// items. trigger builds the *content* of mygo's menu button (text,
// icon or a Row) — put an interactive Button inside it and that button
// would take the pointer before the menu button sees it.
func NewDropdownMenu(c *ui.Context, open *bool, trigger func() ui.Element, items []MenuItem) *DropdownMenu {
	t := c.Theme()
	parts := ui.MenuBase(c, open)
	parts.Trigger.Padding(t.Space(1), t.Space(2)).Radius(t.Radius).
		Children(func() {
			if trigger != nil {
				trigger()
			}
		})
	parts.Popup(func(panel ui.Element) {
		buildMenuPanel(c, t, panel)
		buildMenuItems(c, t, parts, items, open)
	})
	return &DropdownMenu{Element: parts.Trigger}
}
