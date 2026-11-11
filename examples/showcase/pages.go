package main

import (
	"slices"
	"strings"

	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/icons"
)

// page is one entry of the showcase: the key that names it, the label
// and the icon the sidebar shows, the group it belongs to, and the
// builder of its page.
type page struct {
	key   string
	label string
	icon  *ui.SVG
	group string
	build func(s *showcase, c *ui.Context, t *ui.Theme)
}

// pages are the showcase's pages, in the order the sidebar shows them.
//
// A page adds itself here: nothing else has to know of it, and the
// sidebar and the routing follow. The groups are the folds of the
// sidebar: one of them is open at a time, and a search opens them all.
var pages = []page{
	{key: "button", label: "Button", icon: icons.MousePointerClick(), group: "控件", build: (*showcase).buttonPage},
	{key: "badge", label: "Badge", icon: icons.Tag(), group: "控件", build: (*showcase).badgePage},
	{key: "icon", label: "Icon", icon: icons.Shapes(), group: "控件", build: (*showcase).iconPage},
	{key: "input", label: "Input", icon: icons.Type(), group: "表单", build: (*showcase).inputPage},
	{key: "textarea", label: "TextArea", icon: icons.AlignLeft(), group: "表单", build: (*showcase).textAreaPage},
	{key: "select", label: "Select", icon: icons.List(), group: "表单", build: (*showcase).selectPage},
	{key: "combobox", label: "ComboBox", icon: icons.Search(), group: "表单", build: (*showcase).comboBoxPage},
	{key: "multiselect", label: "MultiSelect", icon: icons.Tags(), group: "表单", build: (*showcase).multiSelectPage},
	{key: "field", label: "Field", icon: icons.Braces(), group: "表单", build: (*showcase).formPage},
	{key: "checkbox", label: "Checkbox", icon: icons.SquareCheck(), group: "表单", build: (*showcase).checkboxPage},
	{key: "radio", label: "Radio", icon: icons.CircleDot(), group: "表单", build: (*showcase).radioPage},
	{key: "switch", label: "Switch", icon: icons.ToggleRight(), group: "表单", build: (*showcase).switchPage},
	{key: "slider", label: "Slider", icon: icons.Gauge(), group: "表单", build: (*showcase).sliderPage},
	{key: "alert", label: "Alert", icon: icons.TriangleAlert(), group: "反馈", build: (*showcase).alertPage},
	{key: "toast", label: "Toast", icon: icons.Bell(), group: "反馈", build: (*showcase).toastPage},
	{key: "overlays", label: "浮层", icon: icons.Layers(), group: "浮层", build: (*showcase).overlaysPage},
	{key: "drawer", label: "Drawer", icon: icons.PanelRight(), group: "浮层", build: (*showcase).drawerPage},
	{key: "card", label: "Card", icon: icons.LayoutPanelTop(), group: "版式", build: (*showcase).cardPage},
	{key: "tabs", label: "Tabs", icon: icons.PanelsTopLeft(), group: "版式", build: (*showcase).tabsPage},
	{key: "scaffold", label: "Scaffold", icon: icons.LayoutDashboard(), group: "版式", build: (*showcase).scaffoldPage},
	{key: "theme", label: "外观", icon: icons.Palette(), group: "主题", build: (*showcase).themePage},
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
