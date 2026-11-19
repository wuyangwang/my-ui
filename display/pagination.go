package display

import (
	"strconv"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/icons"
)

// Pagination pages through a count with a page pointer the app owns.
type Pagination struct {
	ui.Element
	c       *ui.Context
	current *int
	total   int
	sizeRef *Size // button size, staged for the next frame's build
	changed func(int)
	pending int // a page chosen before OnChanged was called this frame
}

// NewPagination builds a pager for total pages, reading and writing
// *current. A nil pointer makes the pager read-only.
func NewPagination(c *ui.Context, current *int, total int) *Pagination {
	// The host is rebuilt at the same spot every frame, so the Local it
	// carries survives the per-frame rebuild of the page buttons. Size
	// stages the choice there; writing a field on p would be lost as soon
	// as the view runs again.
	host := ui.Row(c).AlignItems(ui.Center).Gap(c.Theme().Space(1))
	p := &Pagination{c: c, current: current, total: total}
	p.sizeRef = ui.Local(host, "pagination-size", func() Size { return SizeMd })
	p.build(host)
	return p
}

// OnChanged runs fn when a page button changes *current.
func (p *Pagination) OnChanged(fn func(page int)) *Pagination {
	p.changed = fn
	// The click may have been read while the pager was built, before
	// this callback was attached in the same frame.
	if fn != nil && p.pending > 0 {
		fn(p.pending)
		p.pending = 0
	}
	return p
}

// Size picks the pager's button size. The choice is staged on the host
// element and read by the next frame's build: it survives the per-frame
// rebuild of the buttons, and staging it here builds no control that
// could consume this frame's click.
func (p *Pagination) Size(sz Size) *Pagination {
	if p.sizeRef != nil {
		*p.sizeRef = sz
	}
	return p
}

func (p *Pagination) build(host ui.Element) {
	t := p.c.Theme()
	cur := 1
	if p.current != nil && *p.current > 0 {
		cur = *p.current
	}
	buttonSize := SizeMd
	if p.sizeRef != nil {
		buttonSize = *p.sizeRef
	}
	w := sz(buttonSize)
	// Every control is built inside the host's children, so its click
	// lands in the same frame and its parent is the pager.
	host.Children(func() {
		// Previous: an icon that is itself the control, disabled at the
		// first page. Pressing it writes the app's pointer.
		prev := ctrlpkg.NewIconAction(p.c, ctrlpkg.IconActionOptions{
			SVG: icons.ChevronLeft(), Label: Message(p.c, "previous"),
			Size: ctrlpkg.IconActionHit, IconSize: IconSizeSm,
			OnClick: func() { p.goTo(cur - 1) },
		})
		if cur <= 1 {
			prev = prev.Disabled(true)
		}

		for _, page := range pageWindow(cur, p.total) {
			if page < 0 {
				ui.Text(p.c, "…").TextColor(t.TextMuted).Padding(0, t.Space(1))
				continue
			}
			label := strconv.Itoa(page)
			b := ctrlpkg.NewButton(p.c, label)
			b.Element = b.Element.Width(w).Height(w).Justify(ui.Center)
			if page == cur {
				b = b.Primary()
			} else {
				b = b.Ghost()
			}
			if b.Clicked() {
				p.goTo(page)
			}
		}

		next := ctrlpkg.NewIconAction(p.c, ctrlpkg.IconActionOptions{
			SVG: icons.ChevronRight(), Label: Message(p.c, "next"),
			Size: ctrlpkg.IconActionHit, IconSize: IconSizeSm,
			OnClick: func() { p.goTo(cur + 1) },
		})
		if cur >= p.total {
			next = next.Disabled(true)
		}
	})
	p.Element = host
}

func (p *Pagination) goTo(page int) {
	if page < 1 || page > p.total || p.current == nil {
		return
	}
	*p.current = page
	p.pending = page
	if p.changed != nil {
		p.changed(page)
		p.pending = 0
	}
}

// pageWindow gives up to seven page numbers with -1 for an ellipsis:
// the first and last page, and two pages around the current one.
func pageWindow(cur, total int) []int {
	if total <= 7 {
		out := make([]int, total)
		for i := range out {
			out[i] = i + 1
		}
		return out
	}
	keep := map[int]bool{1: true, total: true}
	for i := cur - 1; i <= cur+1; i++ {
		if i >= 1 && i <= total {
			keep[i] = true
		}
	}
	out := []int{}
	last := 0
	for i := 1; i <= total; i++ {
		if !keep[i] {
			continue
		}
		if last > 0 && i-last > 1 {
			out = append(out, -1)
		}
		out = append(out, i)
		last = i
	}
	return out
}

func sz(s Size) float32 {
	switch s {
	case SizeXs:
		return 20
	case SizeSm:
		return 24
	case SizeLg:
		return 36
	default:
		return 28
	}
}
