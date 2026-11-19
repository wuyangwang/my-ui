package display

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/icons"
)

// GuideCardOptions configures NewGuideCard.
type GuideCardOptions struct {
	// Title and Desc are the card's words, the title bold above the note.
	Title, Desc string
	// Cover is a picture across the top of the card, CoverWidth wide and
	// CoverHeight tall (the card's width when the width is zero).
	Cover                   ui.ImageSource
	CoverWidth, CoverHeight float32
	// Variant picks the surface: Default is a soft fill, Outline a border,
	// Ghost no face.
	Variant Variant
	// Shadow lifts the card.
	Shadow bool
	// NoClose leaves out the close button.
	NoClose bool
	// Closed is the app's flag for the card's visibility. When nil, the
	// component keeps its own state.
	Closed *bool
	// OnClose runs when the close button hides the card.
	OnClose func()
}

// NewGuideCard is a dismissible tip, as the introduction to a feature:
// a cover, a title, a description and a close button that hides it for
// good.
func NewGuideCard(c *ui.Context, o GuideCardOptions) ui.Element {
	t := c.Theme()
	root := ui.Column(c).Radius(Radius(t, RadiusXl)).Clip()
	closed := o.Closed
	if closed == nil {
		closed = ui.Local(root, "guide-closed", func() bool { return false })
	}
	if *closed {
		return root.Hidden(true)
	}
	switch o.Variant {
	case Outline:
		root.Border(1, t.Border).Background(t.Surface)
	case Ghost:
	default:
		root.Gradient(t.SurfaceHover, t.Surface, 180)
	}
	if o.Shadow {
		root.Shadow(0, 1, 3, -1, t.Border.Alpha(0.5))
	}
	return root.Children(func() {
		if o.Cover != nil {
			img := ui.Image(c, o.Cover).Fit(ui.Cover).AlignSelf(ui.Center)
			if o.CoverWidth > 0 {
				img.Width(o.CoverWidth)
			} else {
				img.AlignSelf(ui.Stretch)
			}
			if o.CoverHeight > 0 {
				img.Height(o.CoverHeight)
			}
		}
		ui.Column(c).Padding(t.Space(4)).Gap(t.Space(2)).Children(func() {
			if o.Title != "" {
				ui.Text(c, o.Title).FontSize(TextSize(t, TextLg)).FontWeight(700)
			}
			if o.Desc != "" {
				ui.Text(c, o.Desc).FontSize(TextSize(t, TextSm)).TextColor(t.TextMuted)
			}
		})
		if !o.NoClose {
			ctrlpkg.NewIconAction(c, ctrlpkg.IconActionOptions{
				SVG: icons.X(), Label: "Close", Danger: true,
				OnClick: func() {
					*closed = true
					if o.OnClose != nil {
						o.OnClose()
					}
				},
			}).Absolute().Top(t.Space(2)).Right(t.Space(2))
		}
	})
}
