package display

import (
	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/icons"
)

// Breadcrumb shows where the current page sits in a hierarchy. The
// app owns the chosen index; clicking an earlier crumb writes it back.
type Breadcrumb struct {
	ui.Element
	c    *ui.Context
	size Size
}

// NewBreadcrumb builds a crumb for each item, separated by a small
// chevron. chosen may be nil to make the trail read-only.
func NewBreadcrumb(c *ui.Context, items []string, chosen *int) *Breadcrumb {
	b := &Breadcrumb{c: c}
	t := c.Theme()
	cur := -1
	if chosen != nil {
		cur = *chosen
	}
	row := ui.Row(c).AlignItems(ui.Center).Gap(t.Space(1.5))
	row.Children(func() {
		for i, item := range items {
			if i > 0 {
				ui.Icon(c, icons.ChevronRight()).
					Size(IconSizeXs, IconSizeXs).TextColor(t.TextMuted)
			}
			if i == cur || (cur < 0 && i == len(items)-1) {
				ui.Text(c, item).FontSize(b.fontSize(t)).Bold().TextColor(t.Text)
				continue
			}
			crumb := ctrlpkg.NewButton(c, item).Ghost().Xs().
				FontSize(b.fontSize(t)).TextColor(t.TextMuted)
			if crumb.Clicked() && chosen != nil {
				*chosen = i
			}
		}
	})
	b.Element = row
	return b
}

// Size picks the crumb text size.
func (b *Breadcrumb) Size(sz Size) *Breadcrumb {
	b.size = sz
	return b
}

func (b *Breadcrumb) fontSize(t *ui.Theme) float32 {
	switch b.size {
	case SizeXs:
		return TextSize(t, TextXs)
	case SizeLg:
		return TextSize(t, TextLg)
	default:
		return TextSize(t, TextSm)
	}
}
