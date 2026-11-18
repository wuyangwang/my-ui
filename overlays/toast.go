package overlays

import (
	. "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
	"time"

	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/icons"
)

// toastFade is how long a toast takes to fade as its time runs out.
const toastFade = 200 * time.Millisecond

// ToastViewport holds the toasts of the window: where the toasts the app
// shows come out.
//
//	myui.NewToastViewport(c)
//
// The viewport is styled by the library, as shadcn's: the toasts stack
// in the bottom right corner of the window, as cards. It takes the place
// of the theme's own look of the toasts, so call it once in the view,
// wherever.
type ToastViewport struct {
	ui.Element
}

// ToastPosition is a corner or edge of the window the toasts stack at.
type ToastPosition uint8

const (
	ToastBottomRight ToastPosition = iota
	ToastBottomLeft
	ToastTopRight
	ToastTopLeft
	ToastBottomCenter
	ToastTopCenter
)

// ViewportOption configures a ToastViewport.
type ViewportOption func(*viewportConfig)

type viewportConfig struct {
	position ToastPosition
	fadeIn   bool
}

// ViewportPosition sets the corner the toasts stack at.
func ViewportPosition(p ToastPosition) ViewportOption {
	return func(cfg *viewportConfig) { cfg.position = p }
}

// ViewportFadeIn fades a toast in when it appears. On by default.
func ViewportFadeIn(on bool) ViewportOption {
	return func(cfg *viewportConfig) { cfg.fadeIn = on }
}

// NewToastViewport builds the window's toasts with the library's look.
// Call it once in the view, wherever. Without options the toasts stack in
// the bottom right corner, as in shadcn's.
func NewToastViewport(c *ui.Context, opts ...ViewportOption) *ToastViewport {
	cfg := &viewportConfig{fadeIn: true}
	for _, o := range opts {
		if o != nil {
			o(cfg)
		}
	}
	v := &ToastViewport{}
	v.Element = ui.ToastViewportBase(c, func(viewport ui.Element, toasts []ui.Toast) {
		th := c.Theme()
		// The viewport column runs top-to-bottom; justify picks the
		// vertical edge and AlignItems the horizontal one.
		switch cfg.position {
		case ToastTopRight, ToastTopLeft, ToastTopCenter:
			viewport.Justify(ui.Start)
		default:
			viewport.Justify(ui.End)
		}
		switch cfg.position {
		case ToastBottomLeft, ToastTopLeft:
			viewport.AlignItems(ui.Start)
		case ToastBottomCenter, ToastTopCenter:
			viewport.AlignItems(ui.Center)
		default:
			viewport.AlignItems(ui.End)
		}
		viewport.Padding(th.Space(4)).Gap(th.Space(2))
		for _, t := range toasts {
			toastCard(c, t, cfg.fadeIn)
		}
	})
	return v
}

// toastCard builds one toast as a card: the icon of its kind, the title
// and the description, the action and the close button.
func toastCard(c *ui.Context, t ui.Toast, fadeIn bool) {
	th := c.Theme()
	parts := ui.ToastBase(c, t)
	card := parts.Root.Row().AlignItems(ui.Center).Gap(th.Space(3)).
		Background(th.Surface).Border(1, th.Border).Radius(Radius(th, RadiusMd)).
		Padding(th.Space(3), th.Space(4)).
		Shadow(0, 8, 24, 0, ui.RGBA(0, 0, 0, 0.18)).
		MaxWidth(th.Space(80)).
		Transition(StateTransition())
	// The toast fades in as it appears, then fades as its time runs out.
	opacity := float32(1)
	if fadeIn {
		opacity = card.AnimateTo("toast-in", 1, ui.Animation{Duration: toastFade, NonMotion: true})
	}
	if left := parts.Left(); left > 0 && left < toastFade {
		opacity *= float32(left) / float32(toastFade)
		c.AnimationFrame()
	}
	if opacity < 1 {
		card.Opacity(opacity)
	}
	card.Children(func() {
		if svg := toastIcon(t.Type); svg != nil {
			ctrlpkg.NewIcon(c, svg).Variant(toastVariant(t.Type)).Size(SizeSm)
		}
		ui.Column(c).Grow(1).Gap(th.Space(1)).Children(func() {
			ui.Text(c, t.Title).Bold()
			if t.Description != "" {
				ui.Text(c, t.Description).TextColor(th.TextMuted)
			}
		})
		if t.Action != "" {
			action := parts.ActionButton().TextColor(th.Accent).FontWeight(600)
			action.Children(func() { ui.Text(c, t.Action) })
		}
		close := parts.CloseButton().TextColor(th.TextMuted)
		close.Children(func() { ctrlpkg.NewIcon(c, icons.X()).Size(SizeSm) })
		close.Label(Message(c, "close"))
	})
}

// Toast is a message that shows near the corner of the window, for a
// few seconds, as the outcome of what the user just did.
//
//	myui.NewToast(c, "已保存").Success().Show()
//
// The window needs a viewport for the toasts: myui.NewToastViewport(c),
// once in the view.
type Toast struct {
	c *ui.Context
	t ui.Toast
}

// NewToast creates a toast with a title, to show with Show.
func NewToast(c *ui.Context, title string) *Toast {
	return &Toast{c: c, t: ui.Toast{Title: title}}
}

// Description adds the text under the title of the toast.
func (t *Toast) Description(s string) *Toast {
	t.t.Description = s
	return t
}

// Variant tells what kind of toast it is, which picks its icon and its
// color: Success, Info, Warning or Destructive, which is an error. It is
// the shared word for the toast's kind, as Type is on other components.
func (t *Toast) Variant(v Variant) *Toast { return t.Type(v) }

// Type tells what kind of toast it is: the same as Variant.
func (t *Toast) Type(v Variant) *Toast {
	if kind := toastType(v); kind != "" {
		t.t.Type = kind
	}
	return t
}

// Success is Type(myui.Success).
func (t *Toast) Success() *Toast { return t.Type(Success) }

// Info is Type(myui.Info).
func (t *Toast) Info() *Toast { return t.Type(Info) }

// Warning is Type(myui.Warning).
func (t *Toast) Warning() *Toast { return t.Type(Warning) }

// Destructive is Type(myui.Destructive): a toast of an error.
func (t *Toast) Destructive() *Toast { return t.Type(Destructive) }

// Timeout sets how long the toast shows: four seconds by default, until
// it is closed when negative.
func (t *Toast) Timeout(d time.Duration) *Toast {
	t.t.Timeout = d
	return t
}

// Action puts a button in the toast that runs fn and closes the toast,
// as Undo after deleting.
func (t *Toast) Action(label string, fn func()) *Toast {
	t.t.Action = label
	t.t.OnAction = fn
	return t
}

// ID names the toast: showing one with the ID of a toast showing changes
// that toast, which shows anew.
func (t *Toast) ID(s string) *Toast {
	t.t.ID = s
	return t
}

// OnClose runs as the toast closes: its time is over, or the user
// closed it.
func (t *Toast) OnClose(fn func()) *Toast {
	t.t.OnClose = fn
	return t
}

// Show shows the toast and returns its ID, to close it with
// Context.CloseToast.
func (t *Toast) Show() string {
	return t.c.AddToast(t.t)
}

// toastType tells what kind of toast a variant means, in the words the
// look of the toast uses.
func toastType(v Variant) string {
	switch v {
	case Success:
		return "success"
	case Info:
		return "info"
	case Warning:
		return "warning"
	case Destructive:
		return "error"
	}
	return ""
}

// toastVariant is the variant a kind of toast takes its color from.
func toastVariant(kind string) Variant {
	switch kind {
	case "success":
		return Success
	case "info":
		return Info
	case "warning":
		return Warning
	case "error":
		return Destructive
	}
	return Default
}

// toastIcon is the icon of a kind of toast.
func toastIcon(kind string) *ui.SVG {
	switch kind {
	case "success":
		return icons.CircleCheck()
	case "info":
		return icons.Info()
	case "warning":
		return icons.TriangleAlert()
	case "error":
		return icons.CircleX()
	}
	return nil
}
