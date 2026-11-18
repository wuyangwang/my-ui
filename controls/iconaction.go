package controls

import (
	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

// checkGlyph clamps a checkbox's tick to the shared icon scale: it
// follows the box's size but never drops below Xs or grows past Md.
func checkGlyph(box float32) float32 {
	s := box - 4
	if s < IconSizeXs {
		return IconSizeXs
	}
	if s > IconSizeSm {
		return IconSizeSm
	}
	return s
}

// iconActionHit is the shared hit area of an icon action: 20 DIPs,
// with a 14 DIP glyph inside. Clear and close affordances in every
// component use it, so the same icon has the same size and the same
// states everywhere.
const IconActionHit = float32(20)

// IconActionOptions configures NewIconAction.
type IconActionOptions struct {
	// SVG is the glyph to draw.
	SVG *ui.SVG
	// Label names the action for assistive technology.
	Label string
	// OnClick runs when the icon is pressed.
	OnClick func()
	// Danger paints the hover and press in the theme's danger color,
	// as a close icon does.
	Danger bool
	// Size overrides the 20 DIP hit area; IconSize overrides the 14
	// DIP glyph.
	Size, IconSize float32
}

// NewIconAction builds an icon that is itself the control, as shadcn's
// ghost icon button and Element Plus's input clear icon are: there is
// no button face, only the glyph changing color under the pointer.
// The element carries a button role and a tab stop for assistive
// technology, and its colors fade through StateTransition.
func NewIconAction(c *ui.Context, o IconActionOptions) ui.Element {
	hit, glyph := o.Size, o.IconSize
	if hit == 0 {
		hit = IconActionHit
	}
	if glyph == 0 {
		glyph = IconSizeSm
	}
	t := c.Theme()
	a := ui.Box(c).Size(hit, hit).AlignItems(ui.Center).Justify(ui.Center).
		Role(ui.RoleButton).TabStop(true).Label(o.Label).
		Cursor(ui.CursorPointer).Transition(StateTransition())

	// The glyph's ink is the whole state: muted at rest, the text's
	// own ink on hover/press, and the danger shade for a close icon.
	ink := t.TextMuted
	if a.Pressed() || a.Hovered() {
		if o.Danger {
			hover, pressed, _ := Shades(t.Danger, t.Dark)
			ink = hover
			if a.Pressed() {
				ink = pressed
			}
		} else {
			ink = t.Text
		}
	}
	a.Children(func() {
		ui.Icon(c, o.SVG).Size(glyph, glyph).TextColor(ink)
	})
	if a.Clicked() && o.OnClick != nil {
		o.OnClick()
	}
	return a
}
