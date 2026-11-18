package forms

import (
	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/icons"
)

// SearchBar is a text field with a search icon and a clear affordance.
// It is the shared Field frame, so it takes the same variants, sizes,
// states and default full width as Input.
type SearchBar struct {
	Frame[*SearchBar]
}

// NewSearchBar creates a search field editing *query.
func NewSearchBar(c *ui.Context, query *string) *SearchBar {
	s := &SearchBar{}
	s.Frame = *NewFrame[*SearchBar](c, query, func(c *ui.Context, value *string, host ui.Element) ui.Element {
		return ui.TextInput(c, value)
	})
	s.Self = s
	s.PrefixIcon(icons.Search())
	s.Clearable()
	return s
}

// Placeholder shows v while the query is empty.
func (s *SearchBar) Placeholder(v string) *SearchBar {
	s.Frame.Placeholder(v)
	return s
}

// Width pins the bar's width in DIPs; zero fills its parent.
func (s *SearchBar) Width(w float32) *SearchBar {
	s.Frame.Width(w)
	return s
}

// Label names the field for assistive technology.
func (s *SearchBar) Label(v string) *SearchBar {
	s.Frame.Label(v)
	return s
}
