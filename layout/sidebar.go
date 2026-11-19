package layout

import "github.com/egoist/mygo/ui"

// NewSidebar builds mygo's sidebar with my-ui's surface. It keeps mygo's
// ui.Sidebar signature, so a call written for the old re-export compiles.
// The selected id is the app's pointer, written by mygo as the user picks
// an item, and the surface, hairline and gap follow c.Theme() each frame.
//
//	layout.NewSidebar(c, &app.section, func() {
//		ui.SidebarItem(c, "drafts", icons.File(), "草稿")
//	})
func NewSidebar(c *ui.Context, selected *string, fn func()) ui.Element {
	t := c.Theme()
	return ui.Sidebar(c, selected, fn).
		Background(t.Surface).
		BorderWidth(0, 1, 0, 0).BorderColor(t.Border).
		Gap(t.Space(1)).Padding(t.Space(2), 0, t.Space(2), t.Space(2))
}
