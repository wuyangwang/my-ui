package display

import (
	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/icons"
)

// SortableList is mygo's List with a grip on each row. The app owns the
// order and sets ListState.Reorder, as mygo requires; the icon is only
// the affordance that tells the user rows can be dragged.
type SortableList struct {
	ui.Element
}

// NewSortableList builds n rows; item draws one row's content.
func NewSortableList(c *ui.Context, state *ui.ListState, n int, item func(i int)) *SortableList {
	t := c.Theme()
	el := ui.List(c, state, n, func(i int) {
		ui.Row(c).AlignItems(ui.Center).Gap(t.Space(2)).Children(func() {
			ui.Icon(c, icons.GripVertical()).Size(IconSizeSm, IconSizeSm).
				TextColor(t.TextMuted)
			if item != nil {
				item(i)
			}
		})
	})
	return &SortableList{Element: el}
}
