package overlays

import "github.com/egoist/mygo/ui"

// ContextMenu is mygo's ContextMenuBase with shadcn's panel and item
// look. The target is built by the app before this call.
type ContextMenu struct {
	ui.Element
}

// NewContextMenu attaches a menu to target. open is the app's flag for
// whether the menu is showing.
func NewContextMenu(c *ui.Context, open *bool, target ui.Element, items []MenuItem) *ContextMenu {
	t := c.Theme()
	parts := ui.ContextMenuBase(c, target, open)
	parts.Popup(func(panel ui.Element) {
		buildMenuPanel(c, t, panel)
		buildMenuItems(c, t, parts, items, open)
	})
	return &ContextMenu{Element: parts.Trigger}
}
