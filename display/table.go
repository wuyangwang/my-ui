package display

import (
	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

// Table is mygo's Table with shadcn's container: a hairline border,
// the theme's radius and background, and the shared text scale.
// Sorting, selection and column resizing stay with mygo's ListState.
type Table struct {
	ui.Element
	c       *ui.Context
	state   *ui.ListState
	columns []ui.TableColumn
	n       int
	cell    func(row, col int)
	variant Variant
	size    Size
	width   float32
}

// NewTable builds a table of n rows. cell fills one cell per call;
// mygo's Table owns the rest of the structure.
func NewTable(c *ui.Context, state *ui.ListState, columns []ui.TableColumn, n int, cell func(row, col int)) *Table {
	t := &Table{c: c, state: state, columns: columns, n: n, cell: cell}
	t.build()
	return t
}

// Variant paints the table's face with the shared look.
func (t *Table) Variant(v Variant) *Table {
	t.variant = v
	t.build()
	return t
}

// Size picks the text scale.
func (t *Table) Size(sz Size) *Table {
	t.size = sz
	t.build()
	return t
}

// Width pins the table's width in DIPs; zero fills the parent.
func (t *Table) Width(w float32) *Table {
	t.width = w
	t.build()
	return t
}

func (t *Table) build() {
	th := t.c.Theme()
	el := ui.Table(t.c, t.state, t.columns, t.n, t.cell).
		Transition(StateTransition()).Border(1, th.Border).
		Radius(th.Radius).Background(th.Background)
	el = el.FontSize(t.fontSize(th))
	if t.width > 0 {
		el = el.Width(t.width)
	} else {
		el = el.FillWidth()
	}
	t.Element = el
}

func (t *Table) fontSize(th *ui.Theme) float32 {
	switch t.size {
	case SizeXs:
		return TextSize(th, TextXs)
	case SizeLg:
		return TextSize(th, TextLg)
	default:
		return TextSize(th, TextSm)
	}
}
