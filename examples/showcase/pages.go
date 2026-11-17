package main

import (
	"slices"
	"strings"

	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/icons"
)

// page is one entry of the showcase: the key that names it, the label
// and the icon the sidebar shows, the group it belongs to, the line that
// describes it under the title, and the builder of its page.
type page struct {
	key   string
	label string
	icon  *ui.SVG
	group string
	desc  string
	build func(s *showcase, c *ui.Context, t *ui.Theme)
}

// pages are the showcase's pages, in the order the sidebar shows them.
//
// A page adds itself here: nothing else has to know of it, and the
// sidebar and the routing follow. The groups are the folds of the
// sidebar: one of them is open at a time, and a search opens them all.
var pages = []page{
	{key: "button", label: "Button", icon: icons.MousePointerClick(), group: "控件", desc: "同一套变体与尺寸词表，悬停、按下与禁用各有其态。", build: (*showcase).buttonPage},
	{key: "badge", label: "Badge", icon: icons.Tag(), group: "控件", desc: "小标签，四种变体，按语义着色。", build: (*showcase).badgePage},
	{key: "icon", label: "Icon", icon: icons.Shapes(), group: "控件", desc: "Lucide 图标按需链接，没用到的不进二进制。", build: (*showcase).iconPage},
	{key: "input", label: "Input", icon: icons.Type(), group: "表单", desc: "前缀后缀、清除按钮与校验态都长在输入框里。", build: (*showcase).inputPage},
	{key: "textarea", label: "TextArea", icon: icons.AlignLeft(), group: "表单", desc: "多行输入，与 Input 同一套设计语言。", build: (*showcase).textAreaPage},
	{key: "select", label: "Select", icon: icons.List(), group: "表单", desc: "选中即收起的下拉，占位、清除与分组都在。", build: (*showcase).selectPage},
	{key: "combobox", label: "ComboBox", icon: icons.Search(), group: "表单", desc: "可搜索的下拉，触发器本身就是输入框。", build: (*showcase).comboBoxPage},
	{key: "multiselect", label: "MultiSelect", icon: icons.Tags(), group: "表单", desc: "选中项以标签呈现，可以逐个移除。", build: (*showcase).multiSelectPage},
	{key: "field", label: "Field", icon: icons.Braces(), group: "表单", desc: "字段层：标签、说明、错误与校验规则。", build: (*showcase).formPage},
	{key: "checkbox", label: "Checkbox", icon: icons.SquareCheck(), group: "表单", desc: "方框勾选，可带说明与禁用态。", build: (*showcase).checkboxPage},
	{key: "radio", label: "Radio", icon: icons.CircleDot(), group: "表单", desc: "一组里选一个，圆环中央亮起圆点。", build: (*showcase).radioPage},
	{key: "switch", label: "Switch", icon: icons.ToggleRight(), group: "表单", desc: "开与关的胶囊，随主题换色。", build: (*showcase).switchPage},
	{key: "slider", label: "Slider", icon: icons.Gauge(), group: "表单", desc: "拖动的刻度线，走到哪算哪。", build: (*showcase).sliderPage},
	{key: "alert", label: "Alert", icon: icons.TriangleAlert(), group: "反馈", desc: "页面里的消息，读屏会念出它。", build: (*showcase).alertPage},
	{key: "toast", label: "Toast", icon: icons.Bell(), group: "反馈", desc: "右下角短暂的通知，过一会儿自己消失。", build: (*showcase).toastPage},
	{key: "overlays", label: "浮层", icon: icons.Layers(), group: "浮层", desc: "提示、浮层与对话框，都开在该在的位置。", build: (*showcase).overlaysPage},
	{key: "drawer", label: "Drawer", icon: icons.PanelRight(), group: "浮层", desc: "从窗口边缘滑入的面板。", build: (*showcase).drawerPage},
	{key: "card", label: "Card", icon: icons.LayoutPanelTop(), group: "版式", desc: "页面里的面板，标题、说明与内容各就各位。", build: (*showcase).cardPage},
	{key: "tabs", label: "Tabs", icon: icons.PanelsTopLeft(), group: "版式", desc: "一排标签，选中的那块单独成形。", build: (*showcase).tabsPage},
	{key: "scaffold", label: "Scaffold", icon: icons.LayoutDashboard(), group: "版式", desc: "固定头尾、内容滚动的页面骨架。", build: (*showcase).scaffoldPage},
	{key: "theme", label: "外观", icon: icons.Palette(), group: "主题", desc: "明暗、色板与子树预览，下一帧生效。", build: (*showcase).themePage},
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
