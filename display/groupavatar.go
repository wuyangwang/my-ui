package display

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

// GroupGrid says how many columns a group avatar's members take.
type GroupGrid int

const (
	// GroupGridTwo lays out up to four members in two columns.
	GroupGridTwo GroupGrid = iota
	// GroupGridThree lays out up to nine members in three columns.
	GroupGridThree
	// GroupGridAuto takes three columns past four members, two otherwise.
	GroupGridAuto
)

// GroupCorner shapes the group's outline. A square group uses the shared
// radius token; a circle is fully round.
type GroupCorner int

const (
	// GroupCornerSquare rounds the corners with Radius(t, RadiusMd).
	GroupCornerSquare GroupCorner = iota
	// GroupCornerCircle makes the group round, with smaller members.
	GroupCornerCircle
)

// GroupAvatarItem is one member of a group avatar.
type GroupAvatarItem struct {
	// Avatar is the member's name or emoji, as in Avatar.
	Avatar string
	// Image shows instead of the name, covering the member's avatar.
	Image ui.ImageSource
	// Background fills the member's avatar.
	Background ui.Color
	// Fallback is what the name shows without an image.
	Fallback AvatarFallback
}

// GroupAvatarOptions configures NewGroupAvatar.
type GroupAvatarOptions struct {
	// Avatars are the members, in order; those past the grid are left out.
	Avatars []GroupAvatarItem
	// Size is the group's width and height in DIPs; 32 when zero.
	Size float32
	// Grid chooses the number of columns.
	Grid GroupGrid
	// Corner shapes the group.
	Corner GroupCorner
	// Shape shapes each member's avatar.
	Shape Shape
}

// GroupAvatar shows a group, as a group chat in a list, by the avatars of
// its members in a grid on a theme-tinted fill; a single member shows
// large. The fill is the shadcn Surface slot and the outline's rounding is
// Radius(t, RadiusMd) (or a full circle), both read from c.Theme() on this
// frame. Member avatars reuse NewAvatar and so follow the same tokens.
func NewGroupAvatar(c *ui.Context, o GroupAvatarOptions) ui.Element {
	t := c.Theme()
	size := o.Size
	if size <= 0 {
		size = 32
	}
	cols := 2
	if o.Grid == GroupGridThree || (o.Grid == GroupGridAuto && len(o.Avatars) > 4) {
		cols = 3
	}
	circle := o.Corner == GroupCornerCircle
	share, gaps := float32(0.75), float32(4)
	if circle {
		share, gaps = 0.65, 6
	}
	member := float32(int(size / float32(cols) * share))
	gap := float32(int((size - member*float32(cols)) / gaps))
	total := member*float32(cols) + gap
	pitch := (total-float32(cols-1)*gap)/float32(cols) + gap
	items := o.Avatars[:min(len(o.Avatars), cols*cols)]
	radius := Radius(t, RadiusMd)
	if circle {
		radius = size / 2
	}
	root := ui.Box(c).Size(size, size).Shrink(0).Clip().
		Background(t.Surface).Radius(radius)
	root.Children(func() {
		if len(items) == 1 {
			one := size * 0.8
			groupMember(c.Key(0), items[0], one, o.Shape).
				Absolute().Left((size - one) / 2).Top((size - one) / 2)
			return
		}
		rows := (len(items) + cols - 1) / cols
		top := (size - (float32(rows-1)*pitch + member)) / 2
		left := (size - total) / 2
		for i, item := range items {
			groupMember(c.Key(i), item, member, o.Shape).Absolute().
				Left(left + float32(i%cols)*pitch).Top(top + float32(i/cols)*pitch)
		}
	})
	return root
}

func groupMember(c *ui.Context, item GroupAvatarItem, size float32, shape Shape) ui.Element {
	return NewAvatar(c, item.Avatar, AvatarOptions{
		Size:       size,
		Shape:      shape,
		Image:      item.Image,
		Background: item.Background,
		Fallback:   item.Fallback,
	}).Element
}
