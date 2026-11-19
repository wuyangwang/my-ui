package overlays

import "github.com/egoist/mygo/ui"

// overlayStateKey scopes the Local that keeps an overlay's state to one
// instance. The app's open flag is the identity: two overlays that share a
// parent never share a callback just because they share a name.
type overlayStateKey struct {
	open *bool
	name string
}

// openState is the part of an overlay's state that bridges the frame
// boundary. The open flag stays in the app; what the library keeps here is
// the registered callbacks and the changes still waiting for them.
//
// A change is recorded when it happens, then delivered by flush. Keeping
// the callback in a Local means a caller that attaches it after the
// constructor — even several frames later — still receives changes that
// were recorded before it arrived.
type openState struct {
	owned          bool // used when the app supplies no open flag
	onChange       func(bool)
	onDismiss      func()
	pendingOpen    []bool
	pendingDismiss bool
}

// openPtr returns the flag to read and write: the app's when it gave one,
// else a stable one the state owns.
func (s *openState) openPtr(external *bool) *bool {
	if external != nil {
		return external
	}
	return &s.owned
}

// record notes an open change for the registered OnOpenChange.
func (s *openState) record(open bool) {
	s.pendingOpen = append(s.pendingOpen, open)
}

// recordDismiss notes a dismissal for the registered OnDismiss.
func (s *openState) recordDismiss() {
	s.pendingDismiss = true
}

// flush delivers changes waiting for a callback and clears them. Without a
// callback it keeps them, so an OnOpenChange or OnDismiss attached on a
// later frame still sees what happened earlier.
func (s *openState) flush() {
	if s.onChange != nil {
		for _, open := range s.pendingOpen {
			s.onChange(open)
		}
		s.pendingOpen = nil
	}
	if s.onDismiss != nil && s.pendingDismiss {
		s.onDismiss()
		s.pendingDismiss = false
	}
}

// stateFor returns the persistent state of one overlay instance. The
// holder is keyed by the app's open flag, so its identity follows the
// instance even when siblings before it come and go.
func stateFor[T any](c *ui.Context, open *bool, name string, init func() T) *T {
	holder := ui.Box(c).Key(overlayStateKey{open: open, name: name}).Hidden(true)
	return ui.Local(holder, name, init)
}
