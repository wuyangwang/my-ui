package myui

import "github.com/egoist/mygo/ui"

// FieldGroup gathers fields that belong together, as shadcn's FieldGroup.
func NewFieldGroup(c *ui.Context, build func()) ui.Element {
	col := ui.Column(c).Gap(c.Theme().Space(4))
	if build != nil {
		col.Children(build)
	}
	return col
}

// FieldSet is a named group of fields, as shadcn's FieldSet/FieldLegend.
type FieldSet struct {
	ui.Element
}

// NewFieldSet creates a group with a legend and the fields build puts in
// it.
func NewFieldSet(c *ui.Context, legend string, build func()) *FieldSet {
	col := ui.Column(c).Gap(c.Theme().Space(4))
	col.Children(func() {
		ui.Text(c, legend).FontSize(TextSize(c.Theme(), TextMd)).Bold()
		ui.Box(c).Height(1).Background(c.Theme().Border).Margin(0, 0, c.Theme().Space(1), 0)
		if build != nil {
			build()
		}
	})
	return &FieldSet{Element: col}
}

// FieldSeparator is a hairline that divides fields, as shadcn's
// FieldSeparator.
func FieldSeparator(c *ui.Context) ui.Element {
	return ui.Box(c).Height(1).FillWidth().Background(c.Theme().Border).
		Margin(c.Theme().Space(2), 0)
}
