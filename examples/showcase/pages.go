package main

import (
	"slices"

	"github.com/egoist/mygo/ui"
)

// page is one entry of the showcase: the key that names it, the label
// the sidebar shows, the group it belongs to, and the builder of its
// page.
type page struct {
	key   string
	label string
	group string
	build func(s *showcase, c *ui.Context, t *ui.Theme)
}

// pages are the showcase's pages, in the order the sidebar shows them.
//
// A page adds itself here: nothing else has to know of it, and the
// sidebar and the routing follow.
var pages = []page{
	{key: "button", label: "Button", group: "组件", build: (*showcase).buttonPage},
	{key: "input", label: "Input", group: "组件", build: (*showcase).inputPage},
	{key: "textarea", label: "TextArea", group: "组件", build: (*showcase).textAreaPage},
	{key: "select", label: "Select", group: "组件", build: (*showcase).selectPage},
	{key: "combobox", label: "ComboBox", group: "组件", build: (*showcase).comboBoxPage},
	{key: "multiselect", label: "MultiSelect", group: "组件", build: (*showcase).multiSelectPage},
	{key: "overlays", label: "弹层", group: "组件", build: (*showcase).overlaysPage},
	{key: "form", label: "表单", group: "组件", build: (*showcase).formPage},
	{key: "badge", label: "Badge", group: "组件", build: (*showcase).badgePage},
	{key: "icon", label: "Icon", group: "组件", build: (*showcase).iconPage},
	{key: "scaffold", label: "Scaffold", group: "组件", build: (*showcase).scaffoldPage},
	{key: "theme", label: "外观", group: "主题", build: (*showcase).themePage},
}

// groups are the groups the pages belong to, in the order they first
// appear.
func groups() []string {
	var out []string
	for _, p := range pages {
		if !slices.Contains(out, p.group) {
			out = append(out, p.group)
		}
	}
	return out
}

// pageOf is the page with that key; the first page stands in for a key
// that names none.
func pageOf(key string) page {
	for _, p := range pages {
		if p.key == key {
			return p
		}
	}
	return pages[0]
}

// pageLabel is the label of the page with that key.
func pageLabel(key string) string {
	return pageOf(key).label
}
