package forms

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

// ImageSelectItem is one option of an ImageSelect. Value is the app's
// value; Image is the picture drawn in the tile. ImageFn, when set, makes
// the picture lazily and takes the place of Image.
type ImageSelectItem[T comparable] struct {
	// Value is the app's value written to the selected pointer.
	Value T
	// Label is the caption under the picture.
	Label string
	// Icon, when set, is drawn beside the caption.
	Icon func() *ui.SVG
	// Image is the picture drawn in the tile.
	Image ui.ImageSource
	// ImageFn makes the picture lazily and takes the place of Image.
	ImageFn func() ui.ImageSource
}

// ImageSelectProps configures NewImageSelect.
type ImageSelectProps[T comparable] struct {
	// Options are the tiles, in order.
	Options []ImageSelectItem[T]
	// MinTileWidth is the least width of a tile; zero uses the theme's
	// Space(40) (160 DIPs). The row wraps when the tiles no longer fit.
	MinTileWidth float32
}

// ImageSelect is a radio group of pictures with labels. A click chooses a
// tile and writes the app's pointer; the arrows move the choice while the
// group has the focus. The chosen tile takes a ring in the theme's accent,
// which darkens on hover, as a theme picker does.
type ImageSelect[T comparable] struct {
	ui.Element
	selected *T
}

// NewImageSelect builds the tiles. selected is the app's pointer to the
// current value; a click writes the value of the tile. A nil pointer is
// allowed, in which case the choice is kept in a ui.Local of the root.
func NewImageSelect[T comparable](c *ui.Context, selected *T, p ImageSelectProps[T]) *ImageSelect[T] {
	t := c.Theme()
	s := &ImageSelect[T]{selected: selected}
	root := ui.Row(c).Gap(t.Space(3)).Wrap()
	if selected == nil {
		selected = ui.Local(root, "image select value", func() T {
			var zero T
			if len(p.Options) > 0 {
				return p.Options[0].Value
			}
			return zero
		})
	}
	s.selected = selected
	tileW := p.MinTileWidth
	if tileW == 0 {
		tileW = t.Space(40)
	}
	root.Children(func() {
		for i, opt := range p.Options {
			opt, i := opt, i
			chosen := *selected == opt.Value
			tile := ui.Column(c).Width(tileW).
				Background(t.Surface).Border(2, t.Border).
				Radius(Radius(t, RadiusLg)).Clip().
				Cursor(ui.CursorPointer).TabStop(true).
				Transition(StateTransition())
			if chosen {
				ring := t.Accent
				if tile.Hovered() || tile.Pressed() {
					ring = t.Text
				}
				tile = tile.Border(2, ring)
			} else if tile.Hovered() {
				tile = tile.Border(2, t.TextMuted)
			}
			tile.Children(func() {
				src := opt.Image
				if opt.ImageFn != nil {
					src = opt.ImageFn()
				}
				if src != nil {
					ui.Image(c, src).FillWidth().Fit(ui.Cover).
						AspectRatio(16.0 / 9.0)
				}
				label := ui.Row(c).FillWidth().AlignItems(ui.Center).
					Gap(t.Space(2)).Padding(t.Space(2), t.Space(3))
				label.Children(func() {
					ink := t.TextMuted
					if chosen {
						ink = t.Accent
					}
					if opt.Icon != nil {
						ui.Icon(c, opt.Icon()).Size(IconSizeSm, IconSizeSm).TextColor(ink)
					}
					ui.Text(c, opt.Label).FontSize(TextSize(t, TextSm)).TextColor(ink)
				})
			})
			if tile.Clicked() {
				*selected = opt.Value
			}
			// The arrows move the choice without a click, as a radio
			// group does. Shortcut is mygo's, so no key timer is needed.
			if tile.Shortcut(0, ui.KeyRight) || tile.Shortcut(0, ui.KeyDown) {
				*selected = p.Options[(i+1)%len(p.Options)].Value
			}
			if tile.Shortcut(0, ui.KeyLeft) || tile.Shortcut(0, ui.KeyUp) {
				*selected = p.Options[(i+len(p.Options)-1)%len(p.Options)].Value
			}
		}
	})
	s.Element = root
	return s
}

// Value reports the currently chosen value.
func (s *ImageSelect[T]) Value() T {
	if s.selected == nil {
		var zero T
		return zero
	}
	return *s.selected
}
