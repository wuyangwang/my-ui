package display

import (
	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/icons"
)

// TagOptions configures NewTag.
type TagOptions struct {
	// Label is the text on the tag.
	Label string
	// Icon is an optional glyph ahead of the label.
	Icon *ui.SVG
	// Closed is the app's flag; a tag with a close icon writes true.
	Closed *bool
	// OnClose runs when the close icon is pressed.
	OnClose func()
}

// Tag is a small label for metadata, like Element Plus's tag and
// shadcn's badge: a tinted face, the shared text scale and an optional
// close icon that is itself the control.
type Tag struct {
	ui.Element
	c       *ui.Context
	o       TagOptions
	variant Variant
	size    Size
	row     ui.Element
	icon    ui.Element
	label   ui.Element
}

// NewTag builds a tag. A nil Closed makes the close icon read-only.
func NewTag(c *ui.Context, o TagOptions) *Tag {
	t := &Tag{c: c, o: o}
	t.build()
	return t
}

// Variant applies the shared semantic color. It restyles the elements
// built this frame; rebuilding the row would fire the close button twice.
func (t *Tag) Variant(v Variant) *Tag {
	t.variant = v
	t.applyLook()
	return t
}

// Size picks the text scale. As Variant, it only restyles this frame's
// elements, so the close button's click is read once.
func (t *Tag) Size(sz Size) *Tag {
	t.size = sz
	t.applyLook()
	return t
}

func (t *Tag) build() {
	th := t.c.Theme()
	t.row = ui.Row(t.c).AlignItems(ui.Center).Gap(th.Space(1)).
		Padding(th.Space(0.5), th.Space(2)).Radius(th.Radius / 2)
	t.row.Children(func() {
		if t.o.Icon != nil {
			t.icon = ui.Icon(t.c, t.o.Icon).Size(IconSizeXs, IconSizeXs)
		}
		if t.o.Label != "" {
			t.label = ui.Text(t.c, t.o.Label)
		}
		if t.o.Closed != nil || t.o.OnClose != nil {
			ctrlpkg.NewIconAction(t.c, ctrlpkg.IconActionOptions{
				SVG: icons.X(), Label: Message(t.c, "close"),
				Size: 16, IconSize: IconSizeXs,
				OnClick: func() {
					if t.o.Closed != nil {
						*t.o.Closed = true
					}
					if t.o.OnClose != nil {
						t.o.OnClose()
					}
				},
			})
		}
	})
	t.Element = t.row
	t.applyLook()
}

// applyLook paints the shared colors and scale on the elements of this
// frame. It never rebuilds the row, so the close control keeps its click.
func (t *Tag) applyLook() {
	th := t.c.Theme()
	fill, border, ink := th.Surface, th.Border, th.Text
	if role, ok := SemanticOf(t.variant); ok {
		col := SemanticColor(th, role)
		fill, border, ink = col.Alpha(0.12), ui.Transparent, col
	} else if t.variant == Destructive {
		fill, border, ink = th.Danger.Alpha(0.12), ui.Transparent, th.Danger
	}
	t.row = t.row.Background(fill).Border(1, border)
	if t.icon.Valid() {
		t.icon = t.icon.TextColor(ink)
	}
	if t.label.Valid() {
		t.label = t.label.TextColor(ink).FontSize(t.fontSize(th))
	}
	t.Element = t.row
}

func (t *Tag) fontSize(th *ui.Theme) float32 {
	switch t.size {
	case SizeXs:
		return TextSize(th, TextXs)
	case SizeLg:
		return TextSize(th, TextLg)
	default:
		return TextSize(th, TextSm)
	}
}
