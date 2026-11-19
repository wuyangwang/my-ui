package overlays

import (
	"runtime"
	"strings"

	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
	"github.com/wuyangwang/my-ui/display"
)

// TooltipSide is the side of the anchor a tip shows on.
type TooltipSide uint8

const (
	TooltipTop TooltipSide = iota
	TooltipBottom
	TooltipLeft
	TooltipRight
)

// TooltipOptions configures a tip at construction time. The same settings
// are also available as fluent methods on the returned Tooltip.
type TooltipOptions struct {
	Side   TooltipSide
	Hotkey string
}

// tooltipAnchors returns the anchor on the target and the one on the tip
// for a side. The framework flips the tip when there is no room.
func tooltipAnchors(s TooltipSide) (at, self ui.Anchor) {
	switch s {
	case TooltipBottom:
		return ui.AnchorBottom, ui.AnchorTop
	case TooltipLeft:
		return ui.AnchorLeft, ui.AnchorRight
	case TooltipRight:
		return ui.AnchorRight, ui.AnchorLeft
	default:
		return ui.AnchorTop, ui.AnchorBottom
	}
}

// hotkeyKeys turns a shortcut written as "mod+shift+k" into the key caps
// display.NewKbd takes. Mod is Command on Apple devices and Ctrl elsewhere;
// single letters are shown upper-case, the way a key cap reads.
func hotkeyKeys(hotkey string) []string {
	var keys []string
	for _, token := range strings.Split(hotkey, "+") {
		switch strings.ToLower(strings.TrimSpace(token)) {
		case "":
		case "mod":
			if runtime.GOOS == "darwin" {
				keys = append(keys, "Cmd")
			} else {
				keys = append(keys, "Ctrl")
			}
		case "cmd", "command", "meta":
			keys = append(keys, "Cmd")
		case "ctrl", "control":
			keys = append(keys, "Ctrl")
		case "shift":
			keys = append(keys, "Shift")
		case "alt", "option":
			keys = append(keys, "Alt")
		case "plus":
			keys = append(keys, "+")
		default:
			runes := []rune(token)
			if len(runes) == 1 && runes[0] >= 'a' && runes[0] <= 'z' {
				keys = append(keys, strings.ToUpper(token))
			} else {
				keys = append(keys, token)
			}
		}
	}
	return keys
}

// Tooltip shows a tip over the anchor, when the pointer rests on it or
// it has the focus.
//
//	overlays.NewTooltip(c, controls.NewButton(c, "帮助").Element, "这是帮助文字").
//		Hotkey("mod+k").Side(overlays.TooltipBottom)
//
// The tip is styled by the library, as shadcn's: the theme turned over, a
// small text, rounded corners and a shadow. Its text becomes the anchor's
// description too, so assistive technology reads it with the anchor.
// Hotkey puts a key cap after the text. Content() builds arbitrary content
// instead of a single text.
type Tooltip struct {
	ui.Element
	c       *ui.Context
	build   func(tip ui.Element)
	side    *TooltipSide
	hotkey  *string
	tooltip ui.Element // the base, rebuilt each frame by the caller
}

// NewTooltip shows a tip with a text over the anchor. Options passed here
// are available to the first frame the tip builds; the fluent methods on
// the result set the same values.
func NewTooltip(c *ui.Context, anchor ui.Element, text string, options ...TooltipOptions) *Tooltip {
	var o TooltipOptions
	if len(options) > 0 {
		o = options[0]
	}
	t := &Tooltip{c: c}
	t.side = ui.Local(anchor, "myui.tooltip.side", func() TooltipSide { return o.Side })
	t.hotkey = ui.Local(anchor, "myui.tooltip.hotkey", func() string { return o.Hotkey })
	t.build = func(tip ui.Element) {
		hotkey := ""
		if t.hotkey != nil {
			hotkey = *t.hotkey
		}
		if hotkey == "" {
			ui.Text(c, text)
			return
		}
		ui.Row(c).Gap(c.Theme().Space(1.5)).AlignItems(ui.Center).Children(func() {
			ui.Text(c, text)
			display.NewKbd(c, hotkeyKeys(hotkey)...).Xs()
		})
	}
	return t.mount(anchor, text)
}

// NewTooltipContent shows a tip built by build over the anchor. The tip
// is built in the overlay, so build must put its content inside the tip
// element with tip.Children(func(){ ... }); content made outside lands at
// the window's corner. Hotkey applies to the text tip of NewTooltip; a
// content tip draws whatever build puts in it.
func NewTooltipContent(c *ui.Context, anchor ui.Element, build func(tip ui.Element), options ...TooltipOptions) *Tooltip {
	var o TooltipOptions
	if len(options) > 0 {
		o = options[0]
	}
	t := &Tooltip{c: c, build: build}
	t.side = ui.Local(anchor, "myui.tooltip.side", func() TooltipSide { return o.Side })
	return t.mount(anchor, "")
}

func (t *Tooltip) mount(anchor ui.Element, description string) *Tooltip {
	c := t.c
	if description != "" {
		anchor = anchor.Description(description)
	}
	t.Element = ui.TooltipBase(c, anchor, func(tip ui.Element) {
		th := c.Theme()
		fill, ink := InverseColors(th)
		side := TooltipTop
		if t.side != nil {
			side = *t.side
		}
		at, self := tooltipAnchors(side)
		tip.Background(fill).TextColor(ink).
			Padding(th.Space(1.5), th.Space(3)).
			Radius(th.Radius).
			FontSize(TextSize(th, TextXs)).
			Shadow(0, 2, 8, 0, ui.RGBA(0, 0, 0, 0.2)).
			MaxWidth(th.Space(60)).
			AttachTo(anchor, at, self).
			Transition(StateTransition())
		OverlayEnter(c, tip, "tooltip-enter")
		if t.build != nil {
			tip.Children(func() { t.build(tip) })
		}
	})
	return t
}

// Side sets the side of the anchor the tip shows on, for chaining. The
// setting is tied to the anchor, so the tip picks it up while it builds.
func (t *Tooltip) Side(s TooltipSide) *Tooltip {
	if t.side != nil {
		*t.side = s
	}
	return t
}

// Hotkey shows a shortcut after the tip's text, for chaining. Keys are
// written as display.Hotkey takes them, such as "mod+k": mod is Cmd on
// Apple devices and Ctrl elsewhere. The key caps are drawn small, after
// the text inside the tip.
func (t *Tooltip) Hotkey(hotkey string) *Tooltip {
	if t.hotkey != nil {
		*t.hotkey = hotkey
	}
	return t
}

// Content replaces the content of the tip with what build puts in it,
// for chaining.
func (t *Tooltip) Content(build func(tip ui.Element)) *Tooltip {
	t.build = build
	return t
}
