package overlays

import (
	"github.com/egoist/mygo/ui"
	. "github.com/wuyangwang/my-ui"
)

// ModalOptions configures a Modal at construction time. The open flag is
// the app's: it owns the state and opens and closes the dialog. Every
// string and callback is read again each frame from the options the view
// passes, so dynamic values follow the app's state.
type ModalOptions struct {
	// Open is the app's flag. It is required: the dialog is owned by the
	// app, not by the library. Pass the same pointer every frame.
	Open *bool
	// Trigger builds the control that opens the dialog. It runs before the
	// overlay so its click lands in the same frame.
	Trigger func() ui.Element
	// Title is the bold heading and the panel's accessible name.
	Title string
	// Description is the muted line under the title.
	Description string
	// Body fills the space between the header and the footer. The library
	// builds it inside the panel, so it never lands in the window corner.
	Body func()
	// Footer builds the right-aligned action row.
	Footer func()
	// OnOpenChange runs when the dialog opens or closes, including from a
	// trigger click, the close button, the backdrop or Escape. It is kept
	// across frames, so attaching it late does not lose a change.
	OnOpenChange func(open bool)
	// OnDismiss runs when Escape or the backdrop dismisses the dialog.
	OnDismiss func()
	// ShowClose hides the corner close button when false. A nil value
	// keeps it on, as in shadcn's dialog.
	ShowClose *bool
}

// modalState is what a Modal keeps in its Local between frames.
type modalState struct {
	open  openState
	parts panelParts
}

// Modal shows a dialog over a dimmed backdrop while the app's flag is
// true.
//
//	overlays.NewModal(c, overlays.ModalOptions{
//		Open:        &app.deleteOpen,
//		Trigger:     func() ui.Element { return controls.NewButton(c, "删除").Element },
//		Title:       "删除项目",
//		Description: "此操作不能撤销。",
//		Body:        func() { ui.Text(c, "确定删除这个项目吗？") },
//	})
//
// The library builds the trigger first, then the overlay, and puts the
// Body inside the panel. The backdrop and the panel are styled by the
// library as shadcn's: the scrim dims the window, the panel is a card in
// the middle with a close button in its corner.
type Modal struct {
	ui.Element
	st   *modalState
	open *bool
}

// NewModal shows a dialog over the window while o.Open is true.
func NewModal(c *ui.Context, o ModalOptions) *Modal {
	st := stateFor(c, o.Open, "modal", func() modalState {
		return modalState{parts: panelParts{ShowClose: true}}
	})
	if o.Title != "" {
		st.parts.Title = o.Title
	}
	if o.Description != "" {
		st.parts.Description = o.Description
	}
	if o.Body != nil {
		st.parts.Body = o.Body
	}
	if o.Footer != nil {
		st.parts.Footer = o.Footer
	}
	if o.ShowClose != nil {
		st.parts.ShowClose = *o.ShowClose
	}
	if o.OnOpenChange != nil {
		st.open.onChange = o.OnOpenChange
	}
	if o.OnDismiss != nil {
		st.open.onDismiss = o.OnDismiss
	}
	st.open.flush()

	m := &Modal{st: st, open: st.open.openPtr(o.Open)}
	// The trigger is built before the overlay, so a click on it lands in
	// this frame and the dialog opens at once.
	if o.Trigger != nil {
		t := o.Trigger()
		if t.Clicked() && !*m.open {
			*m.open = true
			st.open.record(true)
		}
	}
	modal := ui.DialogBase(c, m.open, func(backdrop, panel ui.Element) {
		th := c.Theme()
		backdrop.Background(ui.RGBA(0, 0, 0, 0.5))
		panel.Background(th.Background).Border(1, th.Border).Radius(th.Radius).
			Padding(th.Space(6)).Shadow(0, 8, 24, 0, ui.RGBA(0, 0, 0, 0.2)).
			MaxWidth(th.Space(128)).
			Transition(StateTransition())
		OverlayEnter(c, panel, "modal-enter")
		if st.parts.Title != "" {
			panel.Label(st.parts.Title)
		}
		buildPanelParts(c, panel, st.parts, func() {
			*m.open = false
			st.open.record(false)
		})
	})
	m.Element = modal
	if modal.Dismissed() {
		st.open.record(false)
		st.open.recordDismiss()
	}
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
// Escape or the backdrop, as mygo's DialogBase does.
func (m *Modal) Dismissed() bool { return m.Element.Dismissed() }

// OnOpenChange registers fn for open changes and returns the modal. A
// change that happened before this call is delivered at once, so a
// callback attached late still fires.
func (m *Modal) OnOpenChange(fn func(open bool)) *Modal {
	m.st.open.onChange = fn
	m.st.open.flush()
	return m
}

// OnDismiss registers fn for dismissals and returns the modal.
func (m *Modal) OnDismiss(fn func()) *Modal {
	m.st.open.onDismiss = fn
	m.st.open.flush()
	return m
}

// Title sets the heading of the dialog, for chaining. It works while the
// dialog is closed: the heading shows when it next opens.
func (m *Modal) Title(s string) *Modal {
	m.st.parts.Title = s
	return m
}

// Description sets the muted line under the title, for chaining.
func (m *Modal) Description(s string) *Modal {
	m.st.parts.Description = s
	return m
}

// Footer builds the right-aligned action row of the dialog, for chaining.
func (m *Modal) Footer(build func()) *Modal {
	m.st.parts.Footer = build
	return m
}

// ShowClose shows or hides the corner close button. On by default.
func (m *Modal) ShowClose(on bool) *Modal {
	m.st.parts.ShowClose = on
	return m
}
