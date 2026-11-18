package display

import (
	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	"github.com/wuyangwang/my-ui/icons"
)

// Spin is a loading indicator: a Lucide circle turning under the
// framework's own motion rules.
type Spin struct {
	ui.Element
	c       *ui.Context
	size    Size
	variant Variant
}

// NewSpin creates a spinner at the text's own size and the accent's
// color.
func NewSpin(c *ui.Context) *Spin {
	s := &Spin{c: c}
	s.build()
	return s
}

// Size picks the spinner's DIP size.
func (s *Spin) Size(sz Size) *Spin {
	s.size = sz
	s.build()
	return s
}

// Variant colors the spinner with the shared semantic colors.
func (s *Spin) Variant(v Variant) *Spin {
	s.variant = v
	s.build()
	return s
}

func (s *Spin) build() {
	d := IconSizeMd
	switch s.size {
	case SizeXs:
		d = IconSizeXs
	case SizeSm:
		d = IconSizeSm
	case SizeLg:
		d = IconSizeLg
	}
	icon := ctrlpkg.NewIcon(s.c, icons.LoaderCircle())
	icon.Element = icon.Element.Size(d, d)
	icon.Variant(s.variant)
	s.Element = icon.Spin().Element
}
