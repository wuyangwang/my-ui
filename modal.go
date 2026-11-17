package myui

import "github.com/egoist/mygo/ui"

// Modal shows a dialog over a dimmed backdrop, while the app's flag is
// true.
//
//	myui.NewModal(c, &app.open, trigger, func(panel ui.Element) {
//		ui.Text(c, "确定删除这个项目吗？")
//	}).Title("删除项目").Description("此操作不能撤销").
//		Footer(func() { myui.NewButton(c, "取消").Ghost() })
//
// The trigger is built by the app, and a click on it opens the dialog; a
// click on the backdrop closes the dialog, and Escape closes it too.
//
// The backdrop and the panel are styled by the library, as shadcn's: the
// scrim dims the window, the panel is a card in the middle of it, with a
// close button in its corner. Title, Description and Footer add the
// structured parts of a shadcn dialog; the build callback fills the body
// between them.
type Modal struct {
	ui.Element
	open *bool
	*dialogParts
}

// NewModal shows a dialog over the window while open is true. The flag
// is the app's: it opens and closes the dialog.
//
// The trigger is built where the modal is, so it sits beside the dialog
// it opens; a click on it opens the dialog in the same frame.
func NewModal(c *ui.Context, open *bool, trigger func() ui.Element, build func(panel ui.Element)) *Modal {
	m := &Modal{open: open}
	if trigger != nil {
		t := trigger()
		// A click on the trigger opens the dialog.
		if t.Clicked() && m.open != nil {
			*m.open = true
		}
	}
	modal := ui.DialogBase(c, open, func(backdrop ui.Element, panel ui.Element) {
		th := c.Theme()
		backdrop.Background(ui.RGBA(0, 0, 0, 0.5))
		panel.Background(th.Background).Border(1, th.Border).Radius(th.Radius).
			Padding(th.Space(6)).Shadow(0, 8, 24, 0, ui.RGBA(0, 0, 0, 0.2)).
			MaxWidth(th.Space(128)).
			Transition(stateTransition())
		overlayEnter(c, panel, "modal-enter")
		m.dialogParts = buildDialogSlots(c, build, panel, open)
	})
	m.Element = modal
	return m
}

// Open reports whether the dialog shows.
func (m *Modal) Open() bool { return m.open != nil && *m.open }

// SetOpen opens or closes the dialog and returns it, for chaining.
func (m *Modal) SetOpen(open bool) *Modal {
	if m.open != nil {
		*m.open = open
	}
	return m
}

// Dismissed reports whether the dialog was dismissed in this frame — by
// Escape, the backdrop or the close button — as mygo's DialogBase does.
func (m *Modal) Dismissed() bool { return m.Element.Dismissed() }

// OnDismiss calls fn when the dialog was dismissed in this frame, and
// returns it for chaining.
func (m *Modal) OnDismiss(fn func()) *Modal {
	if fn != nil && m.Dismissed() {
		fn()
	}
	return m
}

// Title sets the heading of the dialog, for chaining.
func (m *Modal) Title(s string) *Modal {
	if m.dialogParts != nil {
		m.dialogParts.Title(s)
	}
	return m
}

// Description sets the muted line under the title, for chaining.
func (m *Modal) Description(s string) *Modal {
	if m.dialogParts != nil {
		m.dialogParts.Description(s)
	}
	return m
}

// Footer builds the right-aligned action row of the dialog, for chaining.
func (m *Modal) Footer(build func()) *Modal {
	if m.dialogParts != nil {
		m.dialogParts.Footer(build)
	}
	return m
}

// ShowClose shows or hides the corner close button, for chaining. On by
// default, as in shadcn's dialog.
func (m *Modal) ShowClose(on bool) *Modal {
	if m.dialogParts != nil {
		m.dialogParts.ShowClose(on)
	}
	return m
}
