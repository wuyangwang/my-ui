package overlays

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/icons"
)

// panelParts is the structured content of a shadcn-style overlay panel:
// a header with a title and a description, a body the app fills, and a
// footer for the actions. The library builds all of it inside
// panel.Children, so no content ever lands in the window's corner.
type panelParts struct {
	Title, Description string
	Body, Footer       func()
	ShowClose          bool
}

// buildPanelParts fills panel with the structured parts. close runs when
// the corner close button is pressed: it sets the app's flag and records
// the change for OnOpenChange.
func buildPanelParts(c *ui.Context, panel ui.Element, p panelParts, close func()) {
	th := c.Theme()
	panel.Children(func() {
		if p.Title != "" {
			ui.Text(c, p.Title).FontSize(TextSize(th, TextLg)).Bold()
		}
		if p.Description != "" {
			ui.Text(c, p.Description).FontSize(TextSize(th, TextSm)).
				TextColor(th.TextMuted)
		}
		if p.Body != nil {
			ui.Column(c).Children(p.Body)
		}
		if p.Footer != nil {
			ui.Row(c).Justify(ui.End).Gap(th.Space(2)).Children(p.Footer)
		}
		if p.ShowClose {
			ctrlpkg.NewIconAction(c, ctrlpkg.IconActionOptions{
				SVG:      icons.X(),
				Label:    Message(c, "close"),
				IconSize: IconSizeSm,
				OnClick:  close,
			}).Absolute().Top(th.Space(4)).Right(th.Space(4))
		}
	})
}
