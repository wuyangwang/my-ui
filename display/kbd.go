package display

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

// Kbd is a key cap: the key a shortcut needs.
//
//	myui.NewKbd(c, "Ctrl", "K")
//	myui.NewKbd(c, "Cmd", "Shift", "P").Separator("+")
//
// The key cap is styled by the library, as shadcn's: a small box with a
// border, in the monospaced face. Separator("+") renders "+" as a plain
// joiner instead of a key cap. Size and Variant use the shared vocabularies.
type Kbd struct {
	ui.Element
	c         *ui.Context
	keys      []string
	separator string
	size      Size
	variant   Variant
}

// NewKbd creates the key caps of a shortcut, one for each key.
func NewKbd(c *ui.Context, keys ...string) *Kbd {
	k := &Kbd{c: c, keys: keys}
	k.build()
	return k
}

// build paints the caps. A preset builds again so a size or variant set
// after the constructor takes effect.
func (k *Kbd) build() {
	t := k.c.Theme()
	row := ui.Row(k.c).Gap(t.Space(1)).AlignItems(ui.Center)
	row.Children(func() {
		for _, key := range k.keys {
			if k.separator != "" && key == k.separator {
				ui.Text(k.c, key).TextColor(t.TextMuted)
				continue
			}
			text := ui.Text(k.c, key).Font("monospace").
				TextColor(k.ink()).Background(k.fill()).
				Border(1, t.Border).Radius(Radius(t, RadiusSm)).
				Padding(k.padV(), k.padH())
			text.FontSize(k.fontSize())
		}
	})
	k.Element = row
}

// ink and fill are the colors of a cap on its variant.
func (k *Kbd) ink() ui.Color {
	t := k.c.Theme()
	switch k.variant {
	case Primary:
		return t.AccentText
	case Secondary, Outline:
		return t.Text
	}
	return t.TextMuted
}

func (k *Kbd) fill() ui.Color {
	t := k.c.Theme()
	switch k.variant {
	case Primary:
		return t.Accent
	case Secondary:
		return t.Surface
	}
	return t.SurfaceHover
}

func (k *Kbd) fontSize() float32 {
	switch k.size {
	case SizeXs:
		return TextSize(k.c.Theme(), TextXs)
	case SizeLg:
		return TextSize(k.c.Theme(), TextSm)
	}
	return TextSize(k.c.Theme(), TextXs)
}

func (k *Kbd) padV() float32 {
	if k.size == SizeLg {
		return k.c.Theme().Space(1)
	}
	return k.c.Theme().Space(0.5)
}

func (k *Kbd) padH() float32 {
	if k.size == SizeLg {
		return k.c.Theme().Space(2)
	}
	return k.c.Theme().Space(1.5)
}

// Separator treats the given token as a plain joiner between caps instead
// of a key cap of its own, and returns the kbd for chaining.
func (k *Kbd) Separator(s string) *Kbd {
	k.separator = s
	k.build()
	return k
}

// Size applies the shared size s and returns the kbd for chaining.
func (k *Kbd) Size(s Size) *Kbd {
	k.size = s
	k.build()
	return k
}

// Xs sizes the caps down.
func (k *Kbd) Xs() *Kbd { return k.Size(SizeXs) }

// Lg sizes the caps up.
func (k *Kbd) Lg() *Kbd { return k.Size(SizeLg) }

// Variant applies the shared variant v and returns the kbd for chaining.
func (k *Kbd) Variant(v Variant) *Kbd {
	k.variant = v
	k.build()
	return k
}
