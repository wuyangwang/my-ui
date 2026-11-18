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
}

// NewTag builds a tag. A nil Closed makes the close icon read-only.
func NewTag(c *ui.Context, o TagOptions) *Tag {
	t := &Tag{c: c, o: o}
	t.build()
	return t
}

// Variant applies the shared semantic color.
func (t *Tag) Variant(v Variant) *Tag {
	t.variant = v
	t.build()
	return t
}

// Size picks the text scale.
func (t *Tag) Size(sz Size) *Tag {
	t.size = sz
	t.build()
	return t
}

func (t *Tag) build() {
	th := t.c.Theme()
	fill, border, ink := th.Surface, th.Border, th.Text
	if role, ok := SemanticOf(t.variant); ok {
		col := SemanticColor(th, role)
		fill, border, ink = col.Alpha(0.12), ui.Transparent, col
	} else if t.variant == Destructive {
		fill, border, ink = th.Danger.Alpha(0.12), ui.Transparent, th.Danger
	}
	row := ui.Row(t.c).AlignItems(ui.Center).Gap(th.Space(1)).
		Padding(th.Space(0.5), th.Space(2)).Radius(th.Radius/2).
		Background(fill).Border(1, border)
	row.Children(func() {
		if t.o.Icon != nil {
			ui.Icon(t.c, t.o.Icon).Size(IconSizeXs, IconSizeXs).TextColor(ink)
		}
		if t.o.Label != "" {
			ui.Text(t.c, t.o.Label).FontSize(t.fontSize(th)).TextColor(ink)
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
	t.Element = row
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
