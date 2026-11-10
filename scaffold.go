package myui

import "github.com/egoist/mygo/ui"

// ScaffoldOptions configures Scaffold.
type ScaffoldOptions struct {
	// UseTitleBar reserves the window control area (c.TitleBar) at the
	// top so content stays clear of traffic lights / window buttons.
	UseTitleBar bool
	// Top builds an optional header below the title bar area.
	Top func()
	// Bottom builds an optional bar fixed at the bottom of the scaffold.
	Bottom func()
	// BottomInset is extra bottom padding for scroll content. When a
	// Bottom bar is provided and BottomInset is 0, a default inset is
	// applied so the last content row is not obscured.
	BottomInset float32
	// Content builds the scrollable body.
	Content func()
}

// Scaffold composes a page layout: optional title bar safe area and
// header, a scrollable content region, and an optional bottom bar. When
// UseTitleBar is set, the top area is padded using c.TitleBar() and made
// draggable. BottomInset (or a default when Bottom is set) keeps the end
// of scroll content visible above the bottom bar.
func Scaffold(c *ui.Context, o ScaffoldOptions) {
	t := c.Theme()
	bar := c.TitleBar()
	topPad := float32(0)
	leftPad := float32(0)
	rightPad := float32(0)
	if o.UseTitleBar {
		topPad = bar.Height
		leftPad = bar.Left
		rightPad = bar.Right
	}

	bottomInset := o.BottomInset
	if o.Bottom != nil && bottomInset == 0 {
		bottomInset = Space(t, Spacing2Xl) + Space(t, SpacingLg)
	}

	ui.Column(c).Fill().Children(func() {
		if o.UseTitleBar || o.Top != nil {
			h := float32(32)
			if topPad > h {
				h = topPad
			}
			row := ui.Row(c).Height(h).Padding(topPad, rightPad, 0, leftPad).AlignItems(ui.Center)
			if o.UseTitleBar {
				row.DragWindow()
			}
			row.Children(func() {
				if o.Top != nil {
					o.Top()
				}
			})
		}

		ui.Column(c).Grow(1).Children(func() {
			scroll := ui.Scroll(c).Fill().Padding(0, 0, bottomInset, 0)
			if o.Bottom != nil {
				scroll.ScrollbarInsets(0, 0, bottomInset, 0)
			}
			scroll.Children(func() {
				if o.Content != nil {
					o.Content()
				}
			})

			if o.Bottom != nil {
				ui.Row(c).Children(o.Bottom)
			}
		})
	})
}
