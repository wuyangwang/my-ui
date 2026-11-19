package display

import "github.com/egoist/mygo/ui"

// Only widgets that my-ui neither wraps nor plans to wrap are
// re-exported unchanged. Names that now have my-ui versions (Text,
// Divider, Avatar, Image, Grid, Scroll, Sidebar, Split, …) were removed
// from this file when their styled components landed.
var (
	// NewVideo is ui.Video.
	NewVideo = ui.Video
	// NewWebView is ui.WebView.
	NewWebView = ui.WebView
	// NewRichText is ui.RichText, the read-only renderer.
	NewRichText = ui.RichText
	// NewLink is ui.Link.
	NewLink = ui.Link
	// NewTree is ui.Tree.
	NewTree = ui.Tree
	// NewTreeBase is ui.TreeBase.
	NewTreeBase = ui.TreeBase
	// NewTreeItem is ui.TreeItem.
	NewTreeItem = ui.TreeItem
	// NewThemed is ui.Themed, which paints a subtree in its own theme.
	NewThemed = ui.Themed
)
