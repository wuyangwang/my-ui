package controls

import (
	"strconv"
	"strings"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/icons"
)

// HotkeyInput records a keyboard shortcut into a bound string. Focus it
// and press the keys: a shortcut is written as names joined by "+", as
// "mod+shift+k", with mod standing for the primary modifier. A
// combination without a modifier, or one of the conflicts, is refused.
type HotkeyInput struct {
	ui.Element
	c          *ui.Context
	value      *string
	control    ui.Element
	frame      ui.Element
	resetSlot  ui.Element
	clearSlot  ui.Element
	rec        *hotkeyRecord
	focused    bool
	cfg        *hotkeyConfig
	resetValue string
	allowClear bool
	conflicts  []string
	onConflict func(string)
	invalid    bool
	conflict   bool
	resetDone  bool
	clearDone  bool
}

// hotkeyRecord is the reason the last key press was refused. It lives
// in the frame's state, so the message stays until another key arrives.
type hotkeyRecord struct{ invalid, conflict bool }

// hotkeyConfig holds the options the fluent methods set. It lives in the
// frame's state so a value set after the constructor is read by the next
// frame's build, as every my-ui field does.
type hotkeyConfig struct {
	placeholder, invalid, busy string
	variant                    Variant
	disabled                   bool
}

// NewHotkeyInput builds a recorder bound to *value. The returned
// element is the whole field; Control() is the focusable inner control.
func NewHotkeyInput(c *ui.Context, value *string) *HotkeyInput {
	h := &HotkeyInput{c: c, value: value}
	col := ui.Column(c).Gap(c.Theme().Space(2)).FillWidth()
	h.Element = col
	col.Children(func() {
		h.frame = ui.Row(c).MinHeight(32).Padding(c.Theme().Space(1), c.Theme().Space(2.5)).
			Gap(c.Theme().Space(2)).AlignItems(ui.Center).
			Transition(StateTransition())
		h.cfg = ui.Local(h.frame, "hotkey-config", func() hotkeyConfig {
			return hotkeyConfig{
				placeholder: "Press keys to record shortcut",
				invalid:     "A shortcut needs a modifier and a key",
				busy:        "This shortcut is already in use",
			}
		})
		h.rec = ui.Local(h.frame, "hotkey-record", func() hotkeyRecord { return hotkeyRecord{} })
		h.frame.Children(func() {
			h.control = ui.Box(c).Role(ui.RoleTextField).TabStop(true).
				Grow(1).MinWidth(0).AlignSelf(ui.Stretch).
				AlignItems(ui.Center).Justify(ui.Start).
				Cursor(ui.CursorText).FocusRing(false)
			if h.cfg.disabled {
				h.control = h.control.Disabled(true).Cursor(ui.CursorNotAllowed)
			}
			h.control = h.control.HandleInput(func(ev ui.InputEvent) bool {
				if ev.Kind != ui.InputKeyDown || h.cfg.disabled {
					return false
				}
				// Tab moves the focus away; Shift-Tab is the same.
				if ev.Key == ui.KeyTab && ev.Mods&^ui.Shift == 0 {
					return false
				}
				keys, ok := hotkeyName(ev.Mods, ev.Key)
				if !ok {
					return true
				}
				h.rec.invalid, h.rec.conflict = ev.Mods&(ui.Ctrl|ui.Alt|ui.Shift|ui.Super) == 0, false
				switch {
				case h.rec.invalid:
				case containsKey(h.conflicts, keys) && keys != h.resetValue:
					h.rec.conflict = true
					if h.onConflict != nil {
						h.onConflict(keys)
					}
				default:
					*value = keys
				}
				return true
			})
			h.control.Children(func() {
				if *value == "" {
					ui.Text(c, h.cfg.placeholder).TextColor(c.Theme().TextMuted).SingleLine()
					return
				}
				hotkeyCaps(c, *value)
			})
			h.focused = h.control.Focused()
			// Each slot is filled by its own fluent method; an empty
			// slot takes no room.
			row := ui.Row(c).Gap(c.Theme().Space(0.5)).
				Margin(0, -c.Theme().Space(1.5), 0, 0)
			row.Children(func() {
				h.resetSlot = ui.Row(c)
				h.clearSlot = ui.Row(c)
			})
		})
		h.invalid, h.conflict = h.rec.invalid, h.rec.conflict
		h.apply()
		if h.rec.conflict {
			ui.Text(c, h.cfg.busy).TextColor(c.Theme().Danger).FontSize(TextSize(c.Theme(), TextXs))
		} else if h.rec.invalid {
			ui.Text(c, h.cfg.invalid).TextColor(c.Theme().Danger).FontSize(TextSize(c.Theme(), TextXs))
		}
	})
	return h
}

// Reset sets the default shortcut and returns the field, for chaining.
// A reset button appears beside a recorded shortcut while its value
// differs from the default.
func (h *HotkeyInput) Reset(s string) *HotkeyInput {
	h.resetValue = s
	if h.resetDone || h.focused || *h.value == "" || *h.value == s || h.cfg.disabled {
		return h
	}
	h.resetDone = true
	h.resetSlot.Children(func() {
		NewIconAction(h.c.Key("hotkey-reset"), IconActionOptions{
			SVG: icons.Undo2(), Label: "Reset shortcut",
			OnClick: func() {
				*h.value, h.rec.invalid, h.rec.conflict = h.resetValue, false, false
			},
		})
	})
	return h
}

// AllowClear adds a button that empties the shortcut and returns the
// field, for chaining.
func (h *HotkeyInput) AllowClear(on bool) *HotkeyInput {
	h.allowClear = on
	if h.clearDone || !on || h.focused || *h.value == "" || h.cfg.disabled {
		return h
	}
	h.clearDone = true
	h.clearSlot.Children(func() {
		NewIconAction(h.c.Key("hotkey-clear"), IconActionOptions{
			SVG: icons.X(), Label: "Clear shortcut",
			OnClick: func() {
				*h.value, h.rec.invalid, h.rec.conflict = "", false, false
			},
		})
	})
	return h
}

// Placeholder sets the text shown while no shortcut is recorded and
// returns the field, for chaining.
func (h *HotkeyInput) Placeholder(s string) *HotkeyInput {
	h.cfg.placeholder = s
	return h
}

// Conflicts sets the shortcuts the recorder refuses and returns the
// field, for chaining.
func (h *HotkeyInput) Conflicts(keys ...string) *HotkeyInput {
	h.conflicts = keys
	return h
}

// OnConflict calls fn with the refused shortcut and returns the field,
// for chaining.
func (h *HotkeyInput) OnConflict(fn func(string)) *HotkeyInput {
	h.onConflict = fn
	return h
}

// Messages replaces the two refusal messages and returns the field, for
// chaining.
func (h *HotkeyInput) Messages(invalid, conflict string) *HotkeyInput {
	h.cfg.invalid, h.cfg.busy = invalid, conflict
	return h
}

// Variant applies the shared look of v to the field's frame and returns
// the field, for chaining.
func (h *HotkeyInput) Variant(v Variant) *HotkeyInput {
	h.cfg.variant = v
	h.apply()
	return h
}

// Disabled takes the recorder out of use and returns it, for chaining.
func (h *HotkeyInput) Disabled(on bool) *HotkeyInput {
	h.cfg.disabled = on
	h.apply()
	return h
}

// Control returns the focusable control inside the field.
func (h *HotkeyInput) Control() ui.Element { return h.control }

// Value returns the recorded shortcut.
func (h *HotkeyInput) Value() string { return *h.value }

// Changed reports whether the recorded value changed in this frame.
func (h *HotkeyInput) Changed() bool { return h.control.Changed() }

// apply paints the frame on this frame's theme.
func (h *HotkeyInput) apply() {
	t := h.c.Theme()
	bg, border := t.Background, t.Border
	switch h.cfg.variant {
	case Ghost:
		bg = ui.Transparent
	case Secondary:
		bg = t.Surface
	}
	switch {
	case h.invalid || h.conflict:
		border = t.Danger
	case h.focused:
		border = t.Accent
	}
	h.frame.Background(bg).Border(1, border).Radius(Radius(t, RadiusMd))
	if h.cfg.disabled {
		h.frame.Opacity(0.5)
	}
}

// itoa is the number as a string.
func itoa(n int) string { return strconv.Itoa(n) }

// containsKey reports whether keys holds s.
func containsKey(keys []string, s string) bool {
	for _, k := range keys {
		if k == s {
			return true
		}
	}
	return false
}

// hotkeyName writes the key and its modifiers as a recorder stores them.
func hotkeyName(mods ui.Modifiers, key ui.Key) (string, bool) {
	name, ok := namedKey(key)
	if !ok {
		return "", false
	}
	var parts []string
	if mods&ui.Ctrl != 0 {
		parts = append(parts, "mod")
	}
	if mods&ui.Super != 0 {
		parts = append(parts, "meta")
	}
	if mods&ui.Alt != 0 {
		parts = append(parts, "alt")
	}
	if mods&ui.Shift != 0 {
		parts = append(parts, "shift")
	}
	return strings.Join(append(parts, name), "+"), true
}

// namedKey is the name a key is written under.
func namedKey(key ui.Key) (string, bool) {
	switch {
	case key >= ui.KeyA && key <= ui.KeyZ:
		return string(rune('a' + key - ui.KeyA)), true
	case key >= ui.Key0 && key <= ui.Key9:
		return string(rune('0' + key - ui.Key0)), true
	}
	switch key {
	case ui.KeyEnter:
		return "enter", true
	case ui.KeyEscape:
		return "escape", true
	case ui.KeyBackspace:
		return "backspace", true
	case ui.KeyTab:
		return "tab", true
	case ui.KeySpace:
		return "space", true
	case ui.KeyDelete:
		return "delete", true
	case ui.KeyUp:
		return "up", true
	case ui.KeyDown:
		return "down", true
	case ui.KeyLeft:
		return "left", true
	case ui.KeyRight:
		return "right", true
	}
	if key >= ui.KeyF1 && key <= ui.KeyF12 {
		return "f" + itoa(int(key-ui.KeyF1)+1), true
	}
	return "", false
}

// hotkeyCaps shows a recorded shortcut as a row of key caps.
func hotkeyCaps(c *ui.Context, keys string) {
	t := c.Theme()
	for i, k := range strings.Split(keys, "+") {
		ui.Box(c.Key(i)).Padding(t.Space(0.5), t.Space(1.5)).
			Radius(Radius(t, RadiusSm)).Background(t.Surface).Border(1, t.Border).
			Transition(StateTransition()).Children(func() {
			ui.Text(c, capLabel(k)).FontSize(TextSize(t, TextXs)).
				TextColor(t.Text).SingleLine()
		})
	}
}

// capLabel is how a key name reads on its cap.
func capLabel(k string) string {
	switch k {
	case "mod":
		return "Mod"
	case "meta":
		return "Meta"
	default:
		r := []rune(k)
		if len(r) == 0 {
			return k
		}
		r[0] = []rune(strings.ToUpper(string(r[0])))[0]
		return string(r)
	}
}
