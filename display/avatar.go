// Package display holds my-ui's data-display components.
package display

import (
	"hash/fnv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

// AvatarFallback says how a name shows when there is no image.
type AvatarFallback int

const (
	// AvatarMarble paints the name as deterministic marbled colors from the
	// current theme. The same name gets the same colors.
	AvatarMarble AvatarFallback = iota
	// AvatarText shows the first two letters of the name.
	AvatarText
)

// AvatarOptions configures NewAvatar.
type AvatarOptions struct {
	// Size is the width and height in DIPs; 48 when zero.
	Size float32
	// Shape rounds the corners (the theme's look) or makes a circle.
	Shape Shape
	// Image shows instead of the fallback, covering the avatar.
	Image ui.ImageSource
	// Fallback is what a name shows without an image.
	Fallback AvatarFallback
	// Full shows every letter in the text fallback, not just the first two.
	Full bool
	// Background fills the avatar. The letters take a readable ink.
	Background ui.Color
	// Label names the avatar for assistive technology; name when empty.
	Label string
}

// Avatar shows a user or an agent: an image, an emoji, or a name drawn as
// the fallback says.
//
//	display.NewAvatar(c, "Ada Lovelace", display.AvatarOptions{}).Sm()
//	display.NewAvatar(c, "🤯", display.AvatarOptions{Shape: ShapeCircle})
//
// Avatar embeds ui.Element, so the mygo styling methods remain available.
// Every color comes from c.Theme() on the frame the avatar is built, so
// appearance, accent and high-contrast changes take effect at once.
type Avatar struct {
	ui.Element
	c     *ui.Context
	name  string
	o     AvatarOptions
	size  float32
	emoji bool
	glyph ui.Element
	blobs [2]ui.Element
}

// NewAvatar creates an avatar for name with the options o.
func NewAvatar(c *ui.Context, name string, o AvatarOptions) *Avatar {
	a := &Avatar{c: c, name: name, o: o}
	a.size = o.Size
	if a.size <= 0 {
		a.size = 48
	}
	a.emoji = avatarEmoji(name)
	t := c.Theme()
	root := ui.Row(c).Size(a.size, a.size).Shrink(0).
		AlignItems(ui.Center).Justify(ui.Center).Clip().Unselectable()
	bg := o.Background
	if bg == (ui.Color{}) && !a.emoji {
		bg = t.SurfaceHover
	}
	label := o.Label
	if label == "" {
		label = name
	}
	root.Background(bg).Role(ui.RoleImage).Label(label)
	root.Children(func() {
		switch {
		case o.Image != nil:
			a.glyph = ui.Image(c, o.Image).Fill().Fit(ui.Cover)
		case a.emoji:
			a.glyph = ui.Text(c, name).LineHeight(1).SingleLine()
		case o.Fallback == AvatarText:
			text := strings.ToUpper(name)
			if !o.Full {
				r := []rune(text)
				text = string(r[:min(2, len(r))])
			}
			a.glyph = ui.Text(c, text).FontWeight(700).LineHeight(1).
				SingleLine().TextColor(avatarInk(bg, t))
		default:
			a.blobs = avatarMarble(c, name, a.size, t)
		}
	})
	a.applyShape()
	a.Element = root
	return a
}

// applyShape sizes the avatar and its glyphs in place, so a fluent call
// after the constructor does not build a second element in the frame.
func (a *Avatar) applyShape() {
	switch {
	case a.o.Shape == ShapeCircle:
		a.Radius(a.size / 2)
	case a.size < 24:
		a.Radius(a.size * 0.33)
	default:
		a.Radius(max(a.size/6, 2))
	}
	a.Element.Size(a.size, a.size)
	if !a.glyph.Valid() {
		return
	}
	switch {
	case a.o.Image != nil:
	case a.emoji:
		scale := float32(0.8)
		if a.o.Background != (ui.Color{}) {
			scale *= 0.85
		}
		a.glyph.FontSize(a.size * scale)
	case a.o.Fallback == AvatarText:
		a.glyph.FontSize(a.size * 0.5)
	}
}

// Pixels pins the avatar to a size in DIPs and returns it for chaining.
func (a *Avatar) Pixels(px float32) *Avatar {
	a.size = px
	a.o.Size = px
	a.applyShape()
	return a
}

// Size applies the shared size s and returns the avatar for chaining.
func (a *Avatar) Size(s Size) *Avatar {
	switch s {
	case SizeXs:
		return a.Pixels(24)
	case SizeSm:
		return a.Pixels(32)
	case SizeLg:
		return a.Pixels(64)
	default:
		return a.Pixels(48)
	}
}

// Circle makes the avatar round and returns it for chaining.
func (a *Avatar) Circle() *Avatar {
	a.o.Shape = ShapeCircle
	a.applyShape()
	return a
}

// avatarEmoji reports whether s starts with a pictograph.
func avatarEmoji(s string) bool {
	r, _ := utf8.DecodeRuneInString(s)
	return r >= 0x1F000 || unicode.Is(unicode.So, r)
}

// avatarInk returns the ink that reads on bg.
func avatarInk(bg ui.Color, t *ui.Theme) ui.Color {
	if avatarLuma(bg) > 0.4 {
		return t.Text
	}
	return OnAccent(bg)
}

func avatarLuma(c ui.Color) float32 {
	return (0.2126*float32(c.R) + 0.7152*float32(c.G) + 0.0722*float32(c.B)) / 255
}

// avatarMarble paints deterministic soft blobs. The hues come from the
// theme's own palette, so the same name follows light, dark and custom
// colors without a fixed color table in the package.
func avatarMarble(c *ui.Context, seed string, size float32, t *ui.Theme) [2]ui.Element {
	h := fnv.New32a()
	_, _ = h.Write([]byte(seed))
	n := h.Sum32()
	palette := []ui.Color{
		t.Accent,
		SemanticColor(t, SemanticInfo),
		SemanticColor(t, SemanticSuccess),
		SemanticColor(t, SemanticWarning),
		t.TextMuted,
		t.SurfacePressed,
	}
	pick := func(i uint32) ui.Color {
		return palette[(n>>(i*3)+i)%uint32(len(palette))]
	}
	unit := func(i uint32) float32 { return float32((n>>(i*4))&0xf)/15 - 0.5 }
	base, first, second := pick(0), pick(1), pick(2)
	if first == base {
		first = palette[(n+1)%uint32(len(palette))]
	}
	if second == first || second == base {
		second = palette[(n+2)%uint32(len(palette))]
	}
	blur := size * 7 / 80
	ui.Box(c).Absolute().Fill().Background(base)
	b0 := ui.Box(c).Absolute().Size(size*0.9, size*1.1).Radius(size * 0.45).Background(first).
		Left(size * (0.35 + 0.3*unit(1))).Top(size * (-0.2 + 0.3*unit(2))).Rotate(180 * unit(3)).Blur(blur)
	b1 := ui.Box(c).Absolute().Size(size*0.8, size*0.7).Radius(size * 0.35).Background(second).
		Left(size * (-0.15 + 0.3*unit(4))).Top(size * (0.4 + 0.3*unit(5))).Rotate(180 * unit(6)).Blur(blur)
	return [2]ui.Element{b0, b1}
}
