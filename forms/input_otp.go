package forms

import (
	"unicode"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

// OTPOptions configures an InputOTP.
type OTPOptions struct {
	// Length is how many characters the code has: 6 when zero or negative.
	Length int
	// Alphanumeric takes letters as well as digits.
	Alphanumeric bool
	// Variant picks the cells' surface, as an Input's.
	Variant Variant
	// Size sets the cells' side (24, 32 or 40 DIPs) and their font size.
	Size Size
	// OnComplete runs with the code once every cell holds a character. It
	// runs once when the code becomes complete, not on every frame after.
	OnComplete func(code string)
}

// InputOTP edits a one-time code in a row of cells, one character each.
// The value pointer is the app's; the hidden input behind the cells takes
// the typing, so a paste of the whole code lands in the right cells, and
// the app's string is kept to digits (or letters too with Alphanumeric).
type InputOTP struct {
	ui.Element
	cells []ui.Element
}

// NewOTP builds a row of cells editing *value. Characters that are not
// digits (or letters with Alphanumeric) are dropped, and no more than
// Length of them are kept. When the code is complete, OnComplete runs
// once, on the frame the last character arrives.
func NewOTP(c *ui.Context, value *string, o OTPOptions) *InputOTP {
	t := c.Theme()
	n := o.Length
	if n <= 0 {
		n = 6
	}
	valid := func(r rune) bool {
		return unicode.IsDigit(r) || (o.Alphanumeric && unicode.IsLetter(r))
	}
	clean := make([]rune, 0, n)
	for _, r := range *value {
		if len(clean) < n && valid(r) {
			clean = append(clean, r)
		}
	}
	if string(clean) != *value {
		*value = string(clean)
	}

	root := ui.Row(c).Gap(t.Space(2)).AlignItems(ui.Center).AlignSelf(ui.Start).Shrink(0)
	// Remember the last complete code on the host, so the callback fires
	// once when a code fills and not again on the frames that follow.
	last := ui.Local(root, "otp-last-complete", func() string {
		if len(clean) == n {
			return *value
		}
		return ""
	})
	if len(clean) == n && *value != *last {
		*last = *value
		if o.OnComplete != nil {
			o.OnComplete(*value)
		}
	}

	side := otpCellSide(t, o.Size)
	look := otpCellLook(o.Variant, t)
	runes := []rune(*value)
	otp := &InputOTP{Element: root}
	root.Children(func() {
		input := ui.TextInputBase(c, value).MaxLength(n).Absolute().Left(0).Top(0).Right(0).Bottom(0).
			Opacity(0).FocusRing(false)
		if len(runes) == 0 {
			input = input.Focus()
		}
		focused := input.Focused()
		for i := range n {
			cell := ui.Row(c).Width(side).Height(side).AlignItems(ui.Center).Justify(ui.Center).
				Radius(Radius(t, RadiusMd)).Transition(StateTransition())
			border := look.Border
			if focused && (i == len(runes) || (i == n-1 && len(runes) == n)) {
				border = t.Accent
			}
			cell.Background(look.Bg).Border(look.Width, border)
			if look.HasText {
				cell.TextColor(look.Text)
			}
			cell.Children(func() {
				if i < len(runes) {
					ui.Text(c, string(runes[i])).FontFeatures("tnum").FontSize(FrameFontSize(t, o.Size))
				}
			})
			otp.cells = append(otp.cells, cell)
		}
	})
	return otp
}

// NewInputOTP is the old constructor: it builds a code of digits cells
// with the native look. An optional key scopes the host, as before. New
// code uses NewOTP, which carries the look, the size and the callback.
func NewInputOTP(c *ui.Context, value *string, digits int, key ...string) *InputOTP {
	if len(key) > 0 && key[0] != "" {
		c = c.Key(key[0])
	}
	return NewOTP(c, value, OTPOptions{Length: digits})
}

// Width sets the width of each cell in DIPs and returns the input, for
// chaining. Without it a cell is as wide as its size tier (24, 32 or 40
// DIPs); an explicit width overrides that on the next frame.
func (o *InputOTP) Width(w float32) *InputOTP {
	for i := range o.cells {
		o.cells[i] = o.cells[i].Width(w)
	}
	return o
}

// otpCellSide is the side of a cell on size s: 24, 32 or 40 DIPs at the
// theme's spacing, so a looser theme makes every tier roomier in kind.
func otpCellSide(t *ui.Theme, s Size) float32 {
	switch s {
	case SizeXs, SizeSm:
		return t.Space(6)
	case SizeLg:
		return t.Space(10)
	default:
		return t.Space(8)
	}
}

// otpCellLook resolves the surface of one cell. The native variant, as on
// Select's trigger, keeps a bordered face: with no paint the boxes would
// have no shape of their own.
func otpCellLook(v Variant, t *ui.Theme) FrameLook {
	l := FrameLookOf(v, t)
	if l.Native {
		return FrameLook{Bg: t.Surface, Border: t.Border, Width: 1}
	}
	return l
}
