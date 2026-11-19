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
	{key: "titlebar", label: "TitleBar", icon: icons.AppWindow(), group: "版式", desc: "自绘标题栏：左侧品牌，右侧窗口按钮，随明暗换色。", build: (*showcase).titleBarPage},
	{key: "window", label: "Window", icon: icons.ExternalLink(), group: "版式", desc: "一行建出自定义标题栏窗口，随明暗换色。", build: (*showcase).windowPage},
	{key: "skeleton", label: "Skeleton", icon: icons.LoaderCircle(), group: "版式", desc: "数据到达前的占位条，颜色取主题 Surface", build: (*showcase).skeletonPage},
	{key: "empty-state", label: "EmptyState", icon: icons.Inbox(), group: "版式", desc: "没有内容时的图标、说明与下一步", build: (*showcase).emptyStatePage},
	{key: "result", label: "Result", icon: icons.CircleCheck(), group: "版式", desc: "四种结果状态的图标与语义色随主题", build: (*showcase).resultPage},
	{key: "pagination", label: "Pagination", icon: icons.ListOrdered(), group: "版式", desc: "页码按钮写回当前页指针，首尾与省略号自动处理", build: (*showcase).paginationPage},
	{key: "breadcrumb", label: "Breadcrumb", icon: icons.ChevronRight(), group: "版式", desc: "当前项强调墨色，其余可点并写回索引", build: (*showcase).breadcrumbPage},
	{key: "input-number", label: "InputNumber", icon: icons.Hash(), group: "表单", desc: "步进、范围与步长走 mygo 底座，外观用共享字段词表", build: (*showcase).inputNumberPage},
	{key: "search-bar", label: "SearchBar", icon: icons.Search(), group: "表单", desc: "放大镜 + 清除，默认占满、Variant/Size 与 Input 一致", build: (*showcase).searchBarPage},
	{key: "hover-card", label: "HoverCard", icon: icons.MousePointerClick(), group: "浮层", desc: "指针停在触发项上展开，移到面板上保持", build: (*showcase).hoverCardPage},
	{key: "dropdown-menu", label: "DropdownMenu", icon: icons.Menu(), group: "浮层", desc: "mygo MenuBase + shadcn 风格面板与条目", build: (*showcase).dropdownMenuPage},
	{key: "context-menu", label: "ContextMenu", icon: icons.MousePointerClick(), group: "浮层", desc: "mygo ContextMenuBase + shadcn 风格面板与条目", build: (*showcase).contextMenuPage},
	{key: "auto-complete", label: "AutoComplete", icon: icons.List(), group: "表单", desc: "包 mygo Autocomplete，套共享 Variant/Size/Width", build: (*showcase).autoCompletePage},
	{key: "date-picker", label: "DatePicker", icon: icons.CalendarDays(), group: "表单", desc: "包 mygo DateInput，日期指针由 app 持有", build: (*showcase).datePickerPage},
	{key: "carousel", label: "Carousel", icon: icons.Images(), group: "版式", desc: "索引指针由 app 持有，左右图标与圆点翻页", build: (*showcase).carouselPage},
	{key: "tag", label: "Tag", icon: icons.Tag(), group: "版式", desc: "轻量元信息标记，语义弱底与可关闭图标", build: (*showcase).tagPage},
	{key: "spin", label: "Spin", icon: icons.LoaderCircle(), group: "反馈", desc: "旋转图标走框架动效，不带计时器", build: (*showcase).spinPage},
	{key: "table", label: "Table", icon: icons.Table(), group: "版式", desc: "包 mygo Table，排序与选择由 ListState 管", build: (*showcase).tablePage},
	{key: "menubar", label: "Menubar", icon: icons.Menu(), group: "浮层", desc: "横排多个下拉菜单，共用菜单面板外观。", build: (*showcase).menubarPage},
	{key: "command-palette", label: "Command Palette", icon: icons.Search(), group: "浮层", desc: "输入即过滤的命令面板。", build: (*showcase).commandPalettePage},
	{key: "sortable-list", label: "SortableList", icon: icons.ListOrdered(), group: "版式", desc: "拖拽手柄排序，顺序由 app 持有。", build: (*showcase).sortableListPage},
	{key: "input-otp", label: "InputOTP", icon: icons.Hash(), group: "表单", desc: "一格一字符，自动跳到下一格。", build: (*showcase).inputOTPage},
	{key: "upload", label: "Upload", icon: icons.Upload(), group: "表单", desc: "拖入文件区，把路径交给 app。", build: (*showcase).uploadPage},
	{key: "range-slider", label: "RangeSlider", icon: icons.MoveHorizontal(), group: "控件", desc: "双端滑杆，包 mygo RangeSlider + 共享圆角", build: (*showcase).rangeSliderPage},
	{key: "step-slider", label: "StepSlider", icon: icons.Gauge(), group: "控件", desc: "带刻度的滑杆，包 mygo StepSlider", build: (*showcase).stepSliderPage},
	{key: "stepper", label: "Stepper", icon: icons.ChevronsUpDown(), group: "控件", desc: "上下步进器，带 Variant/Size/Width", build: (*showcase).stepperPage},
	{key: "rating", label: "Rating", icon: icons.Star(), group: "控件", desc: "星级评分，包 mygo Rating", build: (*showcase).ratingPage},
	{key: "meter", label: "Meter", icon: icons.Gauge(), group: "控件", desc: "量度计，包 mygo Meter 的状态色", build: (*showcase).meterPage},
	{key: "color-picker", label: "ColorPicker", icon: icons.Palette(), group: "控件", desc: "颜色色块，包 mygo ColorPicker 并加字段边框", build: (*showcase).colorPickerPage},
	{key: "calendar", label: "Calendar", icon: icons.CalendarDays(), group: "表单", desc: "月视图日历，包 mygo Calendar + 卡片外观", build: (*showcase).calendarPage},
	{key: "time-input", label: "TimeInput", icon: icons.Clock(), group: "表单", desc: "时间输入，包 mygo TimeInput + 字段外观", build: (*showcase).timeInputPage},
	{key: "editable-text", label: "EditableText", icon: icons.Type(), group: "表单", desc: "就地重命名文本，包 mygo EditableText", build: (*showcase).editableTextPage},
	{key: "accordion", label: "Accordion", icon: icons.ListTree(), group: "版式", desc: "Accordion 组件，包 mygo 并加 my-ui 外观。", build: (*showcase).accordionPage},
	{key: "collapsible", label: "Collapsible", icon: icons.ChevronsUpDown(), group: "版式", desc: "Collapsible 组件，包 mygo 并加 my-ui 外观。", build: (*showcase).collapsiblePage},
	{key: "decimal-input", label: "DecimalInput", icon: icons.Hash(), group: "表单", desc: "DecimalInput 组件，包 mygo 并加 my-ui 外观。", build: (*showcase).decimalInputPage},
	{key: "split-button", label: "SplitButton", icon: icons.ChevronsUpDown(), group: "控件", desc: "SplitButton 组件，包 mygo 并加 my-ui 外观。", build: (*showcase).splitButtonPage},
	{key: "toggle", label: "Toggle", icon: icons.ToggleRight(), group: "控件", desc: "Toggle 组件，包 mygo 并加 my-ui 外观。", build: (*showcase).togglePage},
	{key: "color-well", label: "ColorWell", icon: icons.Palette(), group: "控件", desc: "ColorWell 组件，包 mygo 并加 my-ui 外观。", build: (*showcase).colorWellPage},
	{key: "alert-dialog", label: "AlertDialog", icon: icons.TriangleAlert(), group: "反馈", desc: "破坏性操作确认，结果走 OnResult。", build: (*showcase).alertDialogPage},
	{key: "list", label: "List", icon: icons.List(), group: "版式", desc: "mygo 虚拟化列表 + my-ui 容器外观。", build: (*showcase).listPage},
	{key: "outline", label: "Outline", icon: icons.ListTree(), group: "版式", desc: "大数据量层级视图，支持展开与选中。", build: (*showcase).outlinePage},
	{key: "markdown", label: "Markdown", icon: icons.FileText(), group: "版式", desc: "完稿与流式 Markdown，主题映射。", build: (*showcase).markdownPage},
	{key: "chart", label: "Chart", icon: icons.BarChart3(), group: "版式", desc: "柱状/折线/面积/饼图预设。", build: (*showcase).chartPage},
	{key: "action-icon", label: "ActionIcon", icon: icons.MousePointerClick(), group: "控件", desc: "ActionIcon，参考 Lobe MyGo 风格。", build: (*showcase).actionIconPage},
	{key: "action-icon-group", label: "ActionIconGroup", icon: icons.Grid2X2(), group: "控件", desc: "ActionIconGroup，参考 Lobe MyGo 风格。", build: (*showcase).actionIconGroupPage},
	{key: "agent-avatar", label: "AgentAvatar", icon: icons.Bot(), group: "版式", desc: "AgentAvatar，参考 Lobe MyGo 风格。", build: (*showcase).agentAvatarPage},
	{key: "avatar", label: "Avatar", icon: icons.CircleUser(), group: "版式", desc: "Avatar，参考 Lobe MyGo 风格。", build: (*showcase).avatarPage},
	{key: "block", label: "Block", icon: icons.Square(), group: "版式", desc: "Block，参考 Lobe MyGo 风格。", build: (*showcase).blockPage},
	{key: "burger", label: "Burger", icon: icons.Menu(), group: "导航", desc: "Burger，参考 Lobe MyGo 风格。", build: (*showcase).burgerPage},
	{key: "color-swatches", label: "ColorSwatches", icon: icons.Palette(), group: "控件", desc: "ColorSwatches，参考 Lobe MyGo 风格。", build: (*showcase).colorSwatchesPage},
	{key: "copy-button", label: "CopyButton", icon: icons.Copy(), group: "控件", desc: "CopyButton，参考 Lobe MyGo 风格。", build: (*showcase).copyButtonPage},
	{key: "divider", label: "Divider", icon: icons.Minus(), group: "版式", desc: "Divider，参考 Lobe MyGo 风格。", build: (*showcase).dividerPage},
	{key: "draggable-panel", label: "DraggablePanel", icon: icons.PanelLeft(), group: "版式", desc: "DraggablePanel，参考 Lobe MyGo 风格。", build: (*showcase).draggablePanelPage},
	{key: "draggable-side-nav", label: "DraggableSideNav", icon: icons.PanelLeft(), group: "版式", desc: "DraggableSideNav，参考 Lobe MyGo 风格。", build: (*showcase).draggableSideNavPage},
	{key: "floating-panel", label: "FloatingPanel", icon: icons.PanelTop(), group: "反馈", desc: "FloatingPanel，参考 Lobe MyGo 风格。", build: (*showcase).floatingPanelPage},
	{key: "floating-sheet", label: "FloatingSheet", icon: icons.PanelBottom(), group: "反馈", desc: "FloatingSheet，参考 Lobe MyGo 风格。", build: (*showcase).floatingSheetPage},
	{key: "form-styled", label: "Form", icon: icons.ClipboardList(), group: "表单", desc: "Form，参考 Lobe MyGo 风格。", build: (*showcase).formStyledPage},
	{key: "grid", label: "Grid", icon: icons.LayoutGrid(), group: "版式", desc: "Grid，参考 Lobe MyGo 风格。", build: (*showcase).gridPage},
	{key: "group-avatar", label: "GroupAvatar", icon: icons.Users(), group: "版式", desc: "GroupAvatar，参考 Lobe MyGo 风格。", build: (*showcase).groupAvatarPage},
	{key: "guide-card", label: "GuideCard", icon: icons.Lightbulb(), group: "版式", desc: "GuideCard，参考 Lobe MyGo 风格。", build: (*showcase).guideCardPage},
	{key: "header", label: "Header", icon: icons.PanelTop(), group: "版式", desc: "Header，参考 Lobe MyGo 风格。", build: (*showcase).headerPage},
	{key: "highlighter", label: "Highlighter", icon: icons.Highlighter(), group: "版式", desc: "Highlighter，参考 Lobe MyGo 风格。", build: (*showcase).highlighterPage},
	{key: "hotkey-input", label: "HotkeyInput", icon: icons.Keyboard(), group: "表单", desc: "HotkeyInput，参考 Lobe MyGo 风格。", build: (*showcase).hotkeyInputPage},
	{key: "image-select", label: "ImageSelect", icon: icons.Images(), group: "表单", desc: "ImageSelect，参考 Lobe MyGo 风格。", build: (*showcase).imageSelectPage},
	{key: "image", label: "Image", icon: icons.Image(), group: "版式", desc: "Image，参考 Lobe MyGo 风格。", build: (*showcase).imagePage},
	{key: "layout", label: "Layout", icon: icons.LayoutGrid(), group: "版式", desc: "Layout，参考 Lobe MyGo 风格。", build: (*showcase).layoutPage},
	{key: "qr-code", label: "QRCode", icon: icons.QrCode(), group: "版式", desc: "QRCode，参考 Lobe MyGo 风格。", build: (*showcase).qrCodePage},
	{key: "scroll-area", label: "ScrollArea", icon: icons.ScrollText(), group: "版式", desc: "ScrollArea，参考 Lobe MyGo 风格。", build: (*showcase).scrollAreaPage},
	{key: "segmented", label: "Segmented", icon: icons.Rows3(), group: "控件", desc: "Segmented，参考 Lobe MyGo 风格。", build: (*showcase).segmentedPage},
	{key: "side-nav", label: "SideNav", icon: icons.PanelLeft(), group: "版式", desc: "SideNav，参考 Lobe MyGo 风格。", build: (*showcase).sideNavPage},
	{key: "snippet", label: "Snippet", icon: icons.Terminal(), group: "版式", desc: "Snippet，参考 Lobe MyGo 风格。", build: (*showcase).snippetPage},
	{key: "text", label: "Text", icon: icons.Type(), group: "版式", desc: "Text，参考 Lobe MyGo 风格。", build: (*showcase).textPage},
	{key: "toggle-group", label: "ToggleGroup", icon: icons.ToggleRight(), group: "控件", desc: "ToggleGroup，参考 Lobe MyGo 风格。", build: (*showcase).toggleGroupPage},
	{key: "steps", label: "Steps", icon: icons.ListOrdered(), group: "导航", desc: "流程步骤条，支持圆点与竖排。", build: (*showcase).stepsPage},
	{key: "chat-list", label: "ChatList", icon: icons.MessageSquare(), group: "AI", desc: "虚拟化消息流，贴底跟随与流式回复。", build: (*showcase).chatListPage},
	{key: "model-picker", label: "ModelPicker", icon: icons.CpuAi(), group: "AI", desc: "分组模型下拉，带图标与能力徽标。", build: (*showcase).modelPickerPage},
	{key: "reasoning", label: "Reasoning", icon: icons.SlidersHorizontal(), group: "AI", desc: "思考强度与温度、top_p 预设。", build: (*showcase).reasoningPage},
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
