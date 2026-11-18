package display

import (
	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

// EmptyStateOptions configures NewEmptyState.
type EmptyStateOptions struct {
	// Icon is the glyph above the title; nil leaves the space out.
	Icon *ui.SVG
	// Title is the short line that says what is missing.
	Title string
	// Description says what the app can do next.
	Description string
	// Action builds the control under the description.
	Action func()
}

// NewEmptyState builds a centered, icon-first explanation of an empty
// list or page. The icon draws in the theme's accent, at the library's
// shared large size.
func NewEmptyState(c *ui.Context, o EmptyStateOptions) ui.Element {
	t := c.Theme()
	return ui.Column(c).FillWidth().Padding(t.Space(10), t.Space(6)).
		AlignItems(ui.Center).Gap(t.Space(3)).Children(func() {
		if o.Icon != nil {
			disc := ui.Row(c).Size(t.Space(14), t.Space(14)).
				AlignItems(ui.Center).Justify(ui.Center).
				Radius(t.Space(7)).Background(t.Accent.Alpha(0.12))
			disc.Children(func() {
				ui.Icon(c, o.Icon).Size(IconSizeLg, IconSizeLg).TextColor(t.Accent)
			})
		}
		if o.Title != "" {
			ui.Text(c, o.Title).FontSize(TextSize(t, TextLg)).Bold()
		}
		if o.Description != "" {
			ui.Text(c, o.Description).FontSize(TextSize(t, TextSm)).
				TextColor(t.TextMuted)
		}
		if o.Action != nil {
			ui.Row(c).Gap(t.Space(2)).Children(o.Action)
		}
	})
}
