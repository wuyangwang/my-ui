package overlays

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/icons"
)

// Parts are the slots a shadcn-style overlay panel is built from:
// a header with a title and a description, a body the app fills, and a
// footer for the actions. The slots are built into the panel during the
// overlay's build callback and filled by the fluent methods, which the
// app calls in the same frame, right after the constructor — so the
// handles are always from the current frame.
type Parts struct {
	c                          *ui.Context
	title, desc, footer, close ui.Element
}

// buildSlots makes the slots inside the panel. build fills the body with
// whatever the app gave the overlay constructor; close is the corner
// button, which sets *open to false when clicked. It returns the parts.
func BuildSlots(c *ui.Context, build func(panel ui.Element), panel ui.Element, open *bool) *Parts {
	d := &Parts{c: c}
	th := c.Theme()
	panel.Children(func() {
		d.title = ui.Row(c)
		d.desc = ui.Row(c)
		body := ui.Column(c)
		body.Children(func() {
			if build != nil {
				build(panel)
			}
		})
		d.footer = ui.Row(c).Justify(ui.End).Gap(th.Space(2))
		// The corner close button, as shadcn's dialog and sheet.
		close := ctrlpkg.NewButton(c, "").Ghost().Xs()
		close.Icon(icons.X()).Label(Message(c, "close"))
		close.Element = close.Element.Absolute().Top(th.Space(4)).Right(th.Space(4))
		if close.Clicked() && open != nil {
			*open = false
		}
		d.close = close.Element
	})
	return d
}

// Title puts the bold heading in the panel and returns the parts.
func (d *Parts) Title(s string) *Parts {
	if d.title.Valid() {
		d.title.Children(func() {
			ui.Text(d.c, s).FontSize(TextSize(d.c.Theme(), TextLg)).Bold()
		})
	}
	return d
}

// Description puts the muted line under the title and returns the parts.
func (d *Parts) Description(s string) *Parts {
	if d.desc.Valid() {
		d.desc.Children(func() {
			ui.Text(d.c, s).FontSize(TextSize(d.c.Theme(), TextSm)).
				TextColor(d.c.Theme().TextMuted)
		})
	}
	return d
}

// Footer builds the right-aligned action row at the foot of the panel.
func (d *Parts) Footer(build func()) *Parts {
	if d.footer.Valid() && build != nil {
		d.footer.Children(build)
	}
	return d
}

// ShowClose shows or hides the corner close button.
func (d *Parts) ShowClose(on bool) *Parts {
	if d.close.Valid() {
		d.close.Hidden(!on)
	}
	return d
}
