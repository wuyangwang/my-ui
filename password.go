package myui

import (
	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/icons"
)

// PasswordInput is a text field that hides what it holds, with a button
// that shows it:
//
//	myui.NewPasswordInput(c, &password).Outline().Placeholder("密码")
//
// MyGo's text input hides what it holds for good: there is no public way
// to show it again. The button therefore shows the text in a field of
// its own — the same value, another field — and the two take turns.
// The value is the app's, so it survives the change; the state of the
// editor, the caret and the undo, does not.
//
// The button takes the focus back to the field after it is pressed, so
// that typing goes on where it left off.
type PasswordInput struct {
	field[*PasswordInput]
	shown    *bool
	focus    *bool
	eyeOn    *ui.SVG // the icon of the button while the text shows
	eyeOff   *ui.SVG // the icon of the button while the text is hidden
	eyeColor *ui.Color
}

// NewPasswordInput creates a password field editing *value with the
// native preset.
func NewPasswordInput(c *ui.Context, value *string) *PasswordInput {
	p := &PasswordInput{}
	p.field = *newField[*PasswordInput](c, value, func(c *ui.Context, value *string, host ui.Element) ui.Element {
		// The state of the button: whether the text shows. The host
		// keeps it from frame to frame.
		p.shown = ui.Local(host, "password-shown", func() bool { return false })
		p.focus = ui.Local(host, "password-focus", func() bool { return false })
		if *p.shown {
			c.Key("password-shown")
			return ui.TextInput(c, value)
		}
		c.Key("password-hidden")
		return ui.TextInput(c, value).Password()
	})
	p.self = p

	// The focus goes back to the text after the button was pressed. It
	// runs before the button is built, so that the press of the button
	// is read in the frame after the one that focuses the text.
	if *p.focus {
		*p.focus = false
		p.inner = p.inner.Focus()
	}

	// The button that shows and hides, at the end of the field.
	p.Suffix(func() {
		t := c.Theme()
		b := ui.ButtonBase(c).Size(18, 18).KeepFocus().Label("显示密码")
		// The icon is drawn when the frame is painted, so that the app's
		// own icon and color are read then.
		b.Draw(func(painter *ui.Painter, r ui.Rect) {
			icon := p.eyeIcon()
			color := p.eyeStateColor(t)
			if b.Hovered() {
				color = t.Text
			}
			painter.Icon(icon, ui.Rect{X: r.X + 2, Y: r.Y + 2, W: r.W - 4, H: r.H - 4}, color)
		})
		if b.Clicked() {
			*p.shown = !*p.shown
			// The text is in another field now: the focus goes back to
			// it in the frame that follows.
			*p.focus = true
		}
	})
	return p
}

// EyeIcons sets the icons of the button: the one it shows while the
// text shows, and the one it shows while the text is hidden. It returns
// the field, for chaining. The icons of the set are the default.
func (p *PasswordInput) EyeIcons(shown, hidden *ui.SVG) *PasswordInput {
	p.eyeOn = shown
	p.eyeOff = hidden
	return p
}

// EyeColor sets the color of the button's icon, which is the muted
// tone of the theme by default. It returns the field, for chaining.
func (p *PasswordInput) EyeColor(c ui.Color) *PasswordInput {
	p.eyeColor = &c
	return p
}

// eyeIcon is the icon of the button in the state the field is in.
func (p *PasswordInput) eyeIcon() *ui.SVG {
	if *p.shown {
		if p.eyeOn != nil {
			return p.eyeOn
		}
		return icons.EyeOff()
	}
	if p.eyeOff != nil {
		return p.eyeOff
	}
	return icons.Eye()
}

// eyeStateColor is the color of the button's icon.
func (p *PasswordInput) eyeStateColor(t *ui.Theme) ui.Color {
	if p.eyeColor != nil {
		return *p.eyeColor
	}
	return t.TextMuted
}
