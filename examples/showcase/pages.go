package main

import (
	"slices"
	"strings"

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
// sidebar and the routing follow. The groups are the folds of the
// sidebar: one of them is open at a time, and a search opens them all.
var pages = []page{
	{key: "button", label: "Button", group: "控件", build: (*showcase).buttonPage},
	{key: "badge", label: "Badge", group: "控件", build: (*showcase).badgePage},
	{key: "icon", label: "Icon", group: "控件", build: (*showcase).iconPage},
	{key: "input", label: "Input", group: "表单", build: (*showcase).inputPage},
	{key: "textarea", label: "TextArea", group: "表单", build: (*showcase).textAreaPage},
	{key: "select", label: "Select", group: "表单", build: (*showcase).selectPage},
	{key: "combobox", label: "ComboBox", group: "表单", build: (*showcase).comboBoxPage},
	{key: "multiselect", label: "MultiSelect", group: "表单", build: (*showcase).multiSelectPage},
	{key: "field", label: "Field", group: "表单", build: (*showcase).formPage},
	{key: "checkbox", label: "Checkbox", group: "表单", build: (*showcase).checkboxPage},
	{key: "radio", label: "Radio", group: "表单", build: (*showcase).radioPage},
	{key: "switch", label: "Switch", group: "表单", build: (*showcase).switchPage},
	{key: "slider", label: "Slider", group: "表单", build: (*showcase).sliderPage},
	{key: "alert", label: "Alert", group: "反馈", build: (*showcase).alertPage},
	{key: "toast", label: "Toast", group: "反馈", build: (*showcase).toastPage},
	{key: "overlays", label: "弹层", group: "浮层", build: (*showcase).overlaysPage},
	{key: "drawer", label: "Drawer", group: "浮层", build: (*showcase).drawerPage},
	{key: "scaffold", label: "Scaffold", group: "布局", build: (*showcase).scaffoldPage},
	{key: "card", label: "Card", group: "布局", build: (*showcase).cardPage},
	{key: "tabs", label: "Tabs", group: "布局", build: (*showcase).tabsPage},
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

// pagesOf are the pages of the group g, in the order they show.
func pagesOf(g string) []page {
	var out []page
	for _, p := range pages {
		if p.group == g {
			out = append(out, p)
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

// filterPages are the pages whose label or key holds the query: what a
// search in the sidebar shows.
func filterPages(ps []page, query string) []page {
	if query == "" {
		return ps
	}
	q := strings.ToLower(query)
	var out []page
	for _, p := range ps {
		if strings.Contains(strings.ToLower(p.label), q) ||
			strings.Contains(strings.ToLower(p.key), q) {
			out = append(out, p)
		}
	}
	return out
}
