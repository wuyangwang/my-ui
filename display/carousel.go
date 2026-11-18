package display

import (
	"strconv"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/icons"
)

// Carousel shows one slide at a time, with arrow buttons and a row of
// dots. The current index is the app's, as in a pager.
type Carousel struct {
	ui.Element
	c       *ui.Context
	index   *int
	count   int
	buildIt func(i int)
	arrows  bool
	dots    bool
}

// NewCarousel builds count slides; build draws the one at index i.
// A nil index is treated as a read-only zero.
func NewCarousel(c *ui.Context, index *int, count int, build func(i int)) *Carousel {
	carousel := &Carousel{c: c, index: index, count: count, buildIt: build,
		arrows: true, dots: true}
	carousel.build()
	return carousel
}

// Arrows shows or hides the previous/next controls.
func (c *Carousel) Arrows(on bool) *Carousel {
	c.arrows = on
	c.build()
	return c
}

// Dots shows or hides the indicator row.
func (c *Carousel) Dots(on bool) *Carousel {
	c.dots = on
	c.build()
	return c
}

func (c *Carousel) build() {
	count := c.count
	if count < 1 {
		count = 1
	}
	cur := 0
	if c.index != nil && *c.index > 0 {
		cur = *c.index
	}
	if cur >= count {
		cur = count - 1
	}
	if c.index != nil {
		*c.index = cur
	}
	goTo := func(i int) {
		if i < 0 || i >= count || c.index == nil {
			return
		}
		*c.index = i
	}
	t := c.c.Theme()
	c.Element = ui.Column(c.c).Gap(t.Space(3)).Children(func() {
		ui.Row(c.c).AlignItems(ui.Center).Gap(t.Space(2)).Children(func() {
			if c.arrows {
				prev := ctrlpkg.NewIconAction(c.c, ctrlpkg.IconActionOptions{
					SVG: icons.ChevronLeft(), Label: Message(c.c, "previous"),
					IconSize: IconSizeSm, OnClick: func() { goTo(cur - 1) },
				})
				if cur == 0 {
					prev = prev.Disabled(true)
				}
			}
			if c.buildIt != nil {
				ui.Box(c.c).Grow(1).Children(func() { c.buildIt(cur) })
			}
			if c.arrows {
				next := ctrlpkg.NewIconAction(c.c, ctrlpkg.IconActionOptions{
					SVG: icons.ChevronRight(), Label: Message(c.c, "next"),
					IconSize: IconSizeSm, OnClick: func() { goTo(cur + 1) },
				})
				if cur == count-1 {
					next = next.Disabled(true)
				}
			}
		})
		if c.dots {
			ui.Row(c.c).AlignItems(ui.Center).Justify(ui.Center).
				Gap(t.Space(1)).Children(func() {
				for i := 0; i < count; i++ {
					label := strconv.Itoa(i + 1)
					dot := ctrlpkg.NewButton(c.c.Key("dot "+label), "").
						Xs().Width(t.Space(2)).Height(t.Space(2)).Radius(1000)
					dot.Label("第 " + label + " 张")
					if i == cur {
						dot = dot.Background(t.Accent)
					} else {
						dot = dot.Background(t.SurfaceHover)
					}
					if dot.Clicked() {
						goTo(i)
					}
				}
			})
		}
	})
}
