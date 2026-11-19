package display

import (
	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/icons"
)

// TagVariant is the surface a Tag is painted on. It is independent of the
// color: the variant says how much of the color shows, the color says
// which hue it is.
type TagVariant uint8

const (
	// TagFilled tints the surface with the color (the zero value).
	TagFilled TagVariant = iota
	// TagOutlined draws a border in the color over a faint face.
	TagOutlined
	// TagBorderless keeps only the colored text.
	TagBorderless
	// TagSolid fills the whole tag with the color.
	TagSolid
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
	// Color is an arbitrary color of the tag. A zero Color falls back to
	// the semantic color of Variant, or to the theme's neutral ink.
	Color ui.Color
	// Round makes the tag's corners fully round.
	Round bool
	// Size picks the text scale (SizeXs/Sm/Md/Lg); zero is SizeMd.
	Size Size
}

// Tag is a small label for metadata, like Element Plus's tag and
// shadcn's badge: a tinted face, the shared text scale and an optional
// close icon that is itself the control. Its surface is picked by
// TagVariant and its hue by Color or by the shared semantic Variant.
type Tag struct {
	ui.Element
	c          *ui.Context
	o          TagOptions
	variant    Variant
	tagVariant TagVariant
	color      ui.Color
	round      bool
	row        ui.Element
	icon       ui.Element
	label      ui.Element
}

// NewTag builds a tag. A nil Closed makes the close icon read-only in the
// sense that no app flag is written, but OnClose still runs.
func NewTag(c *ui.Context, o TagOptions) *Tag {
	t := &Tag{c: c, o: o, color: o.Color, round: o.Round}
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

// TagVariant picks the surface: Filled, Outlined, Borderless or Solid. It
// is independent of the color chosen by Color or Variant.
func (t *Tag) TagVariant(v TagVariant) *Tag {
	t.tagVariant = v
	t.applyLook()
	return t
}

// Color sets an arbitrary hue for the tag. A zero Color clears it and
// falls back to Variant's semantic color (or the theme's neutral ink).
func (t *Tag) Color(c ui.Color) *Tag {
	t.color = c
	t.applyLook()
	return t
}

// Round makes the tag's corners fully round when on, and restores the
// theme's radius when off.
func (t *Tag) Round(on bool) *Tag {
	t.round = on
	t.applyLook()
	return t
}

// Size picks the text scale. As Variant, it only restyles this frame's
// elements, so the close button's click is read once.
func (t *Tag) Size(sz Size) *Tag {
	t.o.Size = sz
	t.applyLook()
	return t
}

func (t *Tag) build() {
	th := t.c.Theme()
	t.row = ui.Row(t.c).AlignItems(ui.Center).Gap(th.Space(1)).
		Padding(th.Space(0.5), th.Space(2)).Radius(th.Radius / 2).
		Transition(StateTransition())
	// The closed state lives on the element so a tag that was closed
	// stays hidden across frames, even if its app's flag is not read.
	closed := ui.Local(t.row, "tag-closed", func() bool { return false })
	t.row.Hidden(*closed)
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
					// Hide at once and still run OnClose: hiding is the
					// visual, OnClose is the app's chance to clean up.
					*closed = true
					t.row.Hidden(true)
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

// tagInk resolves the hue of the tag from this frame's theme: an explicit
// Color wins, then the semantic Variant, then no hue at all.
func (t *Tag) tagInk() (ui.Color, bool) {
	if t.color != (ui.Color{}) {
		return t.color, true
	}
	if role, ok := SemanticOf(t.variant); ok {
		return SemanticColor(t.c.Theme(), role), true
	}
	if t.variant == Destructive {
		return t.c.Theme().Danger, true
	}
	if t.variant == Primary || t.variant == Link {
		return t.c.Theme().Accent, true
	}
	return ui.Color{}, false
}

// applyLook paints the shared colors, scale and radius on the elements of
// this frame. It never rebuilds the row, so the close control keeps its
// click.
func (t *Tag) applyLook() {
	th := t.c.Theme()
	fill, border, ink := th.Surface, th.Border, th.Text
	base, hasHue := t.tagInk()
	if hasHue {
		switch t.tagVariant {
		case TagOutlined:
			fill, border, ink = base.Alpha(0.06), base.Alpha(0.45), base
		case TagBorderless:
			fill, border, ink = ui.Transparent, ui.Transparent, base
		case TagSolid:
			fill, border, ink = base, base, OnAccent(base)
		default: // TagFilled
			fill, border, ink = base.Alpha(0.12), ui.Transparent, base
		}
	}
	radius := th.Radius / 2
	if t.round {
		radius = 999
	}
	t.row = t.row.Background(fill).Border(1, border).Radius(radius)
	if t.icon.Valid() {
		t.icon = t.icon.TextColor(ink)
	}
	if t.label.Valid() {
		t.label = t.label.TextColor(ink).FontSize(t.fontSize(th))
	}
	t.Element = t.row
}

func (t *Tag) fontSize(th *ui.Theme) float32 {
	switch t.o.Size {
	case SizeXs:
		return TextSize(th, TextXs)
	case SizeLg:
		return TextSize(th, TextLg)
	default:
		return TextSize(th, TextSm)
	}
}
