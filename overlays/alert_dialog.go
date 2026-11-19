// Package overlays provides dialogs, sheets, menus and other panels that
// float above a window.
package overlays

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
	ctrlpkg "github.com/wuyangwang/my-ui/controls"
)

// AlertDialogOptions is the content of an alert dialog: a title, a
// message, the button labels in the order they show, and a callback for
// the result. The app owns the open flag, passed to NewAlertDialog
// separately, and the result value, which OnResult reports.
type AlertDialogOptions struct {
	// Title is the bold heading of the alert and the name assistive
	// technology reads for its panel.
	Title string
	// Message is the muted line under the title, also used as the
	// panel's accessible description.
	Message string
	// Buttons holds the labels of the action buttons, from left to
	// right. The last one is the default: it gets the destructive
	// preset and the focus, so Enter clicks it.
	Buttons []string
	// OnResult receives the index of the button clicked in this frame,
	// or ui.AlertDismissed when Escape closed the alert. It is not
	// called when the alert has no result yet. It runs on the main
	// thread, in the frame of the click.
	OnResult func(i int)
}

// NewAlertDialog shows an alert over the window while open is true. It
// follows mygo's ui.AlertDialog contract — the open flag is the app's, a
// button click reports its index and closes the alert, Escape reports
// ui.AlertDismissed, and nothing is reported while the alert is idle —
// but gives the panel shadcn's look instead of mygo's built-in one.
//
// The reason for rebuilding it is that ui.AlertDialog returns an int, not
// an Element, and draws its own buttons: there is nothing public to
// restyle. This constructor uses ui.DialogBase, whose panel is an Element,
// and paints the buttons with myui's presets. The last button is
// destructive by default, as in the spec; the others are outlined.
//
//	open := false
//	overlays.NewAlertDialog(c, &open, overlays.AlertDialogOptions{
//		Title:   "删除项目？",
//		Message: "此操作不能撤销。",
//		Buttons: []string{"取消", "删除"},
//		OnResult: func(i int) {
//			if i == 1 {
//				app.delete()
//			}
//		},
//	})
//
// Opening the alert is the app's job: set *open to true in a trigger's
// click handler. An alert that must be answered can set *open back to
// true inside OnResult when i is ui.AlertDismissed.
//
// The panel, its colors and its metrics are rebuilt from c.Theme() every
// frame, so theme and appearance changes take effect on the next frame.
// There are no timers or goroutines: the panel's fade-in uses the
// framework's own animation, and state changes use StateTransition.
func NewAlertDialog(c *ui.Context, open *bool, o AlertDialogOptions) ui.Element {
	if c == nil || open == nil {
		return ui.Element{}
	}
	result := -1
	panel := ui.DialogBase(c, open, func(backdrop, panel ui.Element) {
		t := c.Theme()
		backdrop.Background(ui.RGBA(0, 0, 0, 0.5))
		panel.Background(t.Background).Border(1, t.Border).
			Radius(Radius(t, RadiusXl)).Padding(t.Space(6)).Gap(t.Space(3)).
			MaxWidth(t.Space(128)).
			Shadow(0, 10, 30, 0, ui.RGBA(0, 0, 0, 0.3)).
			Transition(StateTransition()).
			Role(ui.RoleAlertDialog).Label(o.Title).Description(o.Message)
		OverlayEnter(c, panel, "alert-dialog-enter")
		panel.Children(func() {
			ui.Column(c).Gap(t.Space(2)).Children(func() {
				if o.Title != "" {
					ui.Text(c, o.Title).FontSize(TextSize(t, TextLg)).Bold()
				}
				if o.Message != "" {
					ui.Text(c, o.Message).FontSize(TextSize(t, TextSm)).
						TextColor(t.TextMuted)
				}
				ui.Row(c).Wrap().Justify(ui.End).Gap(t.Space(2)).
					Margin(t.Space(2), 0, 0, 0).Children(func() {
					for i, label := range o.Buttons {
						i := i
						b := ctrlpkg.NewButton(c, label)
						if i == len(o.Buttons)-1 {
							b.Destructive()
							b.Element = b.Element.AutoFocus()
						} else {
							b.Outline()
						}
						if b.Clicked() {
							result = i
						}
					}
				})
			})
		})
	})
	if panel.Dismissed() && result == -1 {
		result = ui.AlertDismissed
	}
	if result != -1 {
		*open = false
		if o.OnResult != nil {
			o.OnResult(result)
		}
	}
	return panel
}
