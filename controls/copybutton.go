package controls

import (
	"time"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/icons"
)

// copyFeedback is how long the check stays after a copy. It is the
// duration of a framework animation, not a timer of the component.
const copyFeedback = 2 * time.Second

// copyState is the animated value that says the check is showing.
const copyState = "myui.copy.copied"

// CopyButton copies text to the clipboard and shows a check for a short
// while after. The icon is a NewIconAction, so it shares the hover and
// press of every other pure icon affordance; the lingering check rides
// on a framework animation and needs no timer or goroutine.
type CopyButton struct {
	ui.Element
	c          *ui.Context
	content    string
	copyLabel  string
	copiedText string
	variant    ActionIconVariant
	size       Size
	action     ui.Element
	copied     bool
}

// NewCopyButton builds a button that copies content when clicked. Its
// accessible name is "Copy" and becomes "Copied" while the check shows;
// Labels replaces both strings for an app's own language.
func NewCopyButton(c *ui.Context, content string) *CopyButton {
	b := &CopyButton{c: c, content: content, copyLabel: "Copy", copiedText: "Copied"}
	b.Element = ui.Box(c).Cursor(ui.CursorPointer).Transition(StateTransition())
	// The animation lives on the box: Jump sets it on a click, AnimateTo
	// walks it back to rest and keeps the check up until it gets there.
	copied := b.Element.AnimateTo(copyState, 0,
		ui.Animation{Duration: copyFeedback, NonMotion: true}) > 0
	b.copied = copied
	b.Element.Children(func() {
		svg, label := icons.Copy(), b.copyLabel
		if copied {
			svg, label = icons.Check(), b.copiedText
		}
		box, glyph := actionIconMetrics(b.size)
		b.action = NewIconAction(c, IconActionOptions{
			SVG:      svg,
			Label:    label,
			Size:     box,
			IconSize: glyph,
			OnClick: func() {
				c.WriteClipboard(content)
				b.Element.Jump(copyState, 1)
				b.copied = true
			},
		})
	})
	b.apply()
	return b
}

// Variant applies the surface of v and returns the button, for chaining.
func (b *CopyButton) Variant(v ActionIconVariant) *CopyButton {
	b.variant = v
	b.apply()
	return b
}

// Borderless is Variant(ActionIconBorderless).
func (b *CopyButton) Borderless() *CopyButton { return b.Variant(ActionIconBorderless) }

// Filled is Variant(ActionIconFilled).
func (b *CopyButton) Filled() *CopyButton { return b.Variant(ActionIconFilled) }

// Outlined is Variant(ActionIconOutlined).
func (b *CopyButton) Outlined() *CopyButton { return b.Variant(ActionIconOutlined) }

// Size applies the shared size s and returns the button, for chaining.
func (b *CopyButton) Size(s Size) *CopyButton {
	b.size = s
	b.apply()
	return b
}

// Xs sizes the button down to the smallest shared size.
func (b *CopyButton) Xs() *CopyButton { return b.Size(SizeXs) }

// Sm sizes the button down.
func (b *CopyButton) Sm() *CopyButton { return b.Size(SizeSm) }

// Md gives the built-in sizing.
func (b *CopyButton) Md() *CopyButton { return b.Size(SizeMd) }

// Lg sizes the button up.
func (b *CopyButton) Lg() *CopyButton { return b.Size(SizeLg) }

// Labels replaces the accessible names of the copy and copied states,
// and returns the button, for chaining.
func (b *CopyButton) Labels(copy, copied string) *CopyButton {
	b.copyLabel, b.copiedText = copy, copied
	if b.action.Valid() {
		name := copy
		if b.copied {
			name = copied
		}
		b.action = b.action.Label(name)
	}
	return b
}

// Copied reports whether the check is showing.
func (b *CopyButton) Copied() bool { return b.copied }

// apply paints the surface on this frame's theme.
func (b *CopyButton) apply() {
	t := b.c.Theme()
	box, _ := actionIconMetrics(b.size)
	bg, border := ui.Transparent, ui.Transparent
	width := float32(0)
	switch b.variant {
	case ActionIconFilled:
		bg = t.Surface
	case ActionIconOutlined:
		bg, border, width = t.Background, t.Border, 1
	}
	b.Element = b.Element.Size(box, box).Radius(Radius(t, RadiusMd)).
		Background(bg).Border(width, border)
	if b.action.Valid() {
		b.action = b.action.Size(box, box)
	}
}
