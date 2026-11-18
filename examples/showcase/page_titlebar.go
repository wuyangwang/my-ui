package main

import (
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	"github.com/wuyangwang/my-ui/icons"
)

// titleBarPage demonstrates the custom title bar: both looks, wired to
// state the demo owns.
func (s *showcase) titleBarPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "用途", "TitleBar 自绘窗口标题栏：左侧品牌、右侧窗口按钮，颜色随明暗主题每帧更新（系统标题栏不会跟随应用明暗切换）。真实窗口需配 mygo.WindowOptions{TitleBarStyle: mygo.TitleBarHidden, HideWindowButtons: true}，回调里转调 win.Minimize / win.ToggleMaximize / win.Close", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(3)).Children(func() {
				ui.Text(c, "Windows 风格（Look: TitleBarWindows）").
					FontSize(myui.TextSize(t, myui.TextXs)).TextColor(t.TextMuted)
				s.titleBarDemo(c, t, myui.TitleBarWindows)
			})
		})
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(3)).Children(func() {
				ui.Text(c, "macOS 风格（Look: TitleBarMac，红绿灯在左）").
					FontSize(myui.TextSize(t, myui.TextXs)).TextColor(t.TextMuted)
				s.titleBarDemo(c, t, myui.TitleBarMac)
			})
		})
	})
	s.sectionNote(c, t, "窗口接线", "示例应用自身的窗口就用它：回调转调窗口动作，最大化状态每帧从 win.IsMaximized() 读回", func() {
		s.card(c, t, func() {
			ui.Column(c).Gap(t.Space(2)).Children(func() {
				ui.Text(c, "win := mygo.NewWindow(mygo.WindowOptions{").
					FontSize(myui.TextSize(t, myui.TextXs))
				ui.Text(c, "  TitleBarStyle: mygo.TitleBarHidden, HideWindowButtons: true,").
					FontSize(myui.TextSize(t, myui.TextXs))
				ui.Text(c, "  Content: ui.View(view)})").
					FontSize(myui.TextSize(t, myui.TextXs))
				ui.Text(c, "myui.NewTitleBar(c, myui.TitleBarOptions{").
					FontSize(myui.TextSize(t, myui.TextXs))
				ui.Text(c, "  Name: \"MyGo App\", Maximized: &maximized,").
					FontSize(myui.TextSize(t, myui.TextXs))
				ui.Text(c, "  OnMinimize: win.Minimize, OnMaximize: win.ToggleMaximize, OnClose: win.Close})").
					FontSize(myui.TextSize(t, myui.TextXs))
				ui.Text(c, "Look: TitleBarAuto 按构建平台取 Windows/macOS 风格；系统控件仍在时（c.TitleBar 有安全区）自绘按钮让位。").
					FontSize(myui.TextSize(t, myui.TextXs)).TextColor(t.TextMuted)
			})
		})
	})
}

// titleBarDemo lays one title bar into a fake window frame, its
// callbacks feeding a line of feedback under it.
func (s *showcase) titleBarDemo(c *ui.Context, t *ui.Theme, look myui.TitleBarLook) {
	// The demo's window: a frame with a title bar and a body.
	ui.Column(c).FillWidth().Background(t.Background).Border(1, t.Border).
		Radius(myui.Radius(t, myui.RadiusMd)).Children(func() {
		myui.NewTitleBar(c, myui.TitleBarOptions{
			Look:       look,
			Icon:       icons.Component(),
			Name:       "my-ui",
			OnMinimize: func() { s.titleBarEvent = "最小化了窗口" },
			OnMaximize: func() { s.titleBarEvent = "切换了最大化" },
			OnClose:    func() { s.titleBarEvent = "关闭了窗口" },
		})
		ui.Box(c).FillWidth().Height(1).Background(t.Border)
		ui.Row(c).Height(t.Space(16)).PaddingX(t.Space(3)).AlignItems(ui.Center).
			Children(func() {
				ui.Text(c, "窗口主体（拖拽上面那条试试；本页演示里的按钮只改这行字）").
					FontSize(myui.TextSize(t, myui.TextXs)).TextColor(t.TextMuted)
			})
	})
	if s.titleBarEvent != "" {
		ui.Text(c, "刚发生："+s.titleBarEvent).
			FontSize(myui.TextSize(t, myui.TextXs)).TextColor(t.Accent)
	}
}
