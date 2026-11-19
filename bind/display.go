package bind

import (
	"github.com/egoist/mygo/ui"
	displaypkg "github.com/wuyangwang/my-ui/display"
)

// EmptyState builds a centered explanation of an empty list or page.
func (u UI) EmptyState(o displaypkg.EmptyStateOptions) ui.Element {
	return displaypkg.NewEmptyState(u.c, o)
}

// Tag builds a small label for metadata, configured by o.
func (u UI) Tag(o displaypkg.TagOptions) *displaypkg.Tag {
	return displaypkg.NewTag(u.c, o)
}

// Accordion builds a group of collapsible sections.
func (u UI) Accordion(fn func()) *displaypkg.Accordion {
	return displaypkg.NewAccordion(u.c, fn)
}

// AccordionItem builds one titled section of an accordion, editing *open.
func (u UI) AccordionItem(title string, open *bool, fn func()) *displaypkg.AccordionItem {
	return displaypkg.NewAccordionItem(u.c, title, open, fn)
}

// Collapsible builds a single labelled section editing *open.
func (u UI) Collapsible(label string, open *bool, fn func()) *displaypkg.Collapsible {
	return displaypkg.NewCollapsible(u.c, label, open, fn)
}

// Carousel builds a carousel showing item index of count.
func (u UI) Carousel(index *int, count int, build func(i int)) *displaypkg.Carousel {
	return displaypkg.NewCarousel(u.c, index, count, build)
}

// Table builds a table of n rows over state, with columns and cell.
func (u UI) Table(state *ui.ListState, columns []ui.TableColumn, n int, cell func(row, col int)) *displaypkg.Table {
	return displaypkg.NewTable(u.c, state, columns, n, cell)
}

// CardPanel builds the body of a card, without a title.
func (u UI) CardPanel() *displaypkg.Card {
	return displaypkg.NewCardPanel(u.c)
}

// Spin builds a loading spinner.
func (u UI) Spin() *displaypkg.Spin {
	return displaypkg.NewSpin(u.c)
}

// Breadcrumb builds a path of items with chosen selected.
func (u UI) Breadcrumb(items []string, chosen *int) *displaypkg.Breadcrumb {
	return displaypkg.NewBreadcrumb(u.c, items, chosen)
}

// SortableList builds a reorderable list of n items over state.
func (u UI) SortableList(state *ui.ListState, n int, item func(i int)) *displaypkg.SortableList {
	return displaypkg.NewSortableList(u.c, state, n, item)
}

// BadgeOn builds an anchor with a badge pinned to its corner, configured
// by o.
func (u UI) BadgeOn(anchor func() ui.Element, o displaypkg.BadgeOptions) *displaypkg.Badge {
	return displaypkg.NewBadgeOn(u.c, anchor, o)
}

// Skeleton builds lines of placeholder text while content loads.
func (u UI) Skeleton(lines int) *displaypkg.Skeleton {
	return displaypkg.NewSkeleton(u.c, lines)
}

// Alert builds a live-region message with title.
func (u UI) Alert(title string) *displaypkg.Alert {
	return displaypkg.NewAlert(u.c, title)
}

// Result builds a centered outcome page, configured by o.
func (u UI) Result(o displaypkg.ResultOptions) ui.Element {
	return displaypkg.NewResult(u.c, o)
}

// Pagination builds page controls for *current of total pages.
func (u UI) Pagination(current *int, total int) *displaypkg.Pagination {
	return displaypkg.NewPagination(u.c, current, total)
}
