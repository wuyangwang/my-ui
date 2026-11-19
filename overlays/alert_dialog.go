// Package overlays provides dialogs, sheets, menus and other panels that
// float above a window.
package overlays

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
)

// DialogAction is one button of an AlertDialog. Each action owns its own
// behavior, so the app never maps a returned index back to a label.
type DialogAction struct {
	// Label is the text on the button.
	Label string
	// Kind picks the button's preset: Outline for a quiet action and
	// Destructive for the dangerous one, for example. The zero value
	// keeps the old default — Outline, except the last action, which is
	// Destructive and focused.
	Kind Variant
	// OnClick runs when the button is pressed, on the main thread in the
	// frame of the click, before the alert closes.
	OnClick func()
}

// AlertDialogOptions configures an AlertDialog at construction time. The
// open flag is the app's. Every field is read again each frame from the
// options the view passes.
type AlertDialogOptions struct {
	// Open is the app's flag; the alert belongs to the app.
	Open *bool
	// Title is the bold heading and the panel's accessible name.
	Title string
	// Message is the muted line under the title and the panel's
	// accessible description.
	Message string
	// Actions are the buttons, in the order they show. Each carries its
	// own label and behavior.
	Actions []DialogAction
	// OnDismiss runs when Escape closes the alert. It is kept across
	// frames, so attaching it late does not lose the dismissal.
	OnDismiss func()
}

// alertState is what an AlertDialog keeps in its Local between frames.
type alertState struct {
	open    openState
	title   string
	message string
}

// AlertDialog shows a short, blocking alert over the window while the
// app's flag is true.
//
//	overlays.NewAlertDialog(c, overlays.AlertDialogOptions{
//		Open:    &app.confirmOpen,
//		Title:   "删除项目？",
//		Message: "此操作不能撤销。",
//		Actions: []overlays.DialogAction{
//			{Label: "取消"},
//			{Label: "删除", Kind: myui.Destructive, OnClick: app.delete},
//		},
//		OnDismiss: func() { app.result = "cancelled" },
//	})
//
// mygo's own ui.AlertDialog returns an int and draws its buttons in the
// overlay, so there is no public handle to restyle. This component uses
// ui.DialogBase instead: the panel is an Element the library can paint
// with shadcn's look, and each action carries its own click behavior.
type AlertDialog struct {
	ui.Element
	st   *alertState
	open *bool
}

// NewAlertDialog shows an alert over the window while o.Open is true.
func NewAlertDialog(c *ui.Context, o AlertDialogOptions) *AlertDialog {
	if c == nil {
		return nil
	}
	st := stateFor(c, o.Open, "alert-dialog", func() alertState {
		return alertState{}
	})
	if o.Title != "" {
		st.title = o.Title
	}
	if o.Message != "" {
		st.message = o.Message
	}
	if o.OnDismiss != nil {
		st.open.onDismiss = o.OnDismiss
	}
	st.open.flush()

	open := st.open.openPtr(o.Open)
	panel := ui.DialogBase(c, open, func(backdrop, panel ui.Element) {
		t := c.Theme()
		backdrop.Background(ui.RGBA(0, 0, 0, 0.5))
		panel.Background(t.Background).Border(1, t.Border).
			Radius(Radius(t, RadiusXl)).Padding(t.Space(6)).Gap(t.Space(3)).
			MaxWidth(t.Space(128)).
			Shadow(0, 10, 30, 0, ui.RGBA(0, 0, 0, 0.3)).
			Transition(StateTransition()).
			Role(ui.RoleAlertDialog).Label(st.title).Description(st.message)
		OverlayEnter(c, panel, "alert-dialog-enter")
		panel.Children(func() {
			ui.Column(c).Gap(t.Space(2)).Children(func() {
				if st.title != "" {
					ui.Text(c, st.title).FontSize(TextSize(t, TextLg)).Bold()
				}
				if st.message != "" {
					ui.Text(c, st.message).FontSize(TextSize(t, TextSm)).
						TextColor(t.TextMuted)
				}
				ui.Row(c).Wrap().Justify(ui.End).Gap(t.Space(2)).
					Margin(t.Space(2), 0, 0, 0).Children(func() {
					for i, action := range o.Actions {
						kind := action.Kind
						if kind == Default {
							if i == len(o.Actions)-1 {
								kind = Destructive
							} else {
								kind = Outline
							}
						}
						b := ctrlpkg.NewButton(c, action.Label).Variant(kind)
						if i == len(o.Actions)-1 {
							b.Element = b.Element.AutoFocus()
						}
						if b.Clicked() {
							*open = false
							if action.OnClick != nil {
								action.OnClick()
							}
						}
					}
				})
			})
		})
	})
	if panel.Dismissed() {
		st.open.recordDismiss()
	}
	return &AlertDialog{Element: panel, st: st, open: open}
}

// Open reports whether the alert shows.
func (a *AlertDialog) Open() bool { return a.open != nil && *a.open }

// Dismissed reports whether Escape dismissed the alert in this frame.
func (a *AlertDialog) Dismissed() bool { return a.Element.Dismissed() }

// OnDismiss registers fn for dismissals and returns the alert. A dismissal
// that happened before this call is delivered at once.
func (a *AlertDialog) OnDismiss(fn func()) *AlertDialog {
	a.st.open.onDismiss = fn
	a.st.open.flush()
	return a
}
