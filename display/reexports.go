package display

import "github.com/egoist/mygo/ui"

// Only MyGo widgets my-ui has no counterpart for are re-exported.
// Badge, Progress and Tabs have themed my-ui versions, so their MyGo
// equivalents are deliberately left out.
var (
	// NewText is ui.Text.
	NewText = ui.Text
	// NewTextf is ui.Textf.
	NewTextf = ui.Textf
	// NewDivider is ui.Divider.
	NewDivider = ui.Divider
	// NewAvatar is ui.Avatar.
	NewAvatar = ui.Avatar
	// NewImage is ui.Image.
	NewImage = ui.Image
	// NewVideo is ui.Video.
	NewVideo = ui.Video
	// NewWebView is ui.WebView.
	NewWebView = ui.WebView
	// NewRichText is ui.RichText.
	NewRichText = ui.RichText
	// NewLink is ui.Link.
	NewLink = ui.Link
	// NewTable is ui.Table.
	NewTable = ui.Table
	// NewTree is ui.Tree.
	NewTree = ui.Tree
	// NewTreeBase is ui.TreeBase.
	NewTreeBase = ui.TreeBase
	// NewTreeItem is ui.TreeItem.
	NewTreeItem = ui.TreeItem
	// NewAccordion is ui.Accordion.
	NewAccordion = ui.Accordion
	// NewAccordionBase is ui.AccordionBase.
	NewAccordionBase = ui.AccordionBase
	// NewAccordionItem is ui.AccordionItem.
	NewAccordionItem = ui.AccordionItem
	// NewCollapsible is ui.Collapsible.
	NewCollapsible = ui.Collapsible
	// NewCollapsibleBase is ui.CollapsibleBase.
	NewCollapsibleBase = ui.CollapsibleBase
	// NewBreadcrumbs is ui.Breadcrumbs.
	NewBreadcrumbs = ui.Breadcrumbs
	// NewCalendar is ui.Calendar.
	NewCalendar = ui.Calendar
	// NewGrid is ui.Grid.
	NewGrid = ui.Grid
	// NewGridView is ui.GridView.
	NewGridView = ui.GridView
	// NewScroll is ui.Scroll.
	NewScroll = ui.Scroll
	// NewScrollBoth is ui.ScrollBoth.
	NewScrollBoth = ui.ScrollBoth
	// NewScrollHorizontal is ui.ScrollHorizontal.
	NewScrollHorizontal = ui.ScrollHorizontal
	// NewSidebar is ui.Sidebar.
	NewSidebar = ui.Sidebar
	// NewSidebarItem is ui.SidebarItem.
	NewSidebarItem = ui.SidebarItem
	// NewSidebarSection is ui.SidebarSection.
	NewSidebarSection = ui.SidebarSection
	// NewSplit is ui.Split.
	NewSplit = ui.Split
	// NewSplitPercent is ui.SplitPercent.
	NewSplitPercent = ui.SplitPercent
	// NewSplitVertical is ui.SplitVertical.
	NewSplitVertical = ui.SplitVertical
	// NewSplitVerticalPercent is ui.SplitVerticalPercent.
	NewSplitVerticalPercent = ui.SplitVerticalPercent
	// NewSpacer is ui.Spacer.
	NewSpacer = ui.Spacer
	// NewThemed is ui.Themed, which paints a subtree in its own theme.
	NewThemed = ui.Themed
)
