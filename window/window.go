// Package window wraps mygo's frameless window with myui's TitleBar:
//
//	window.New(window.Options{
//		Title: "MyGo App",
//		Name:  "MyGo App",
//		Theme: func() myui.Options { return myui.Options{Appearance: myui.AppearanceDark} },
//		View:  appView,
//	})
//
// One call gives a window whose title bar draws in the theme's own
// colors — the system's does not follow the app's light and dark
// switch. The window hides the system's bar and controls
// (mygo.WindowOptions{TitleBarStyle: TitleBarHidden, HideWindowButtons:
// true}), which keeps the shadow, the rounded corners and the resizing
// edges — a Frameless window would drop the top resizing edge on
// Windows — and the bar it draws carries the brand on the left and the
// window controls on the right, with the actions going through the
// window itself.
//
// This package is the one place in my-ui that imports mygo's root
// package: the window is the shell the app builds on, and the actions
// (Minimize, ToggleMaximize, Close) live on mygo.Window, not in ui.
package window

import (
	"fmt"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
	myui "github.com/wuyangwang/my-ui"
	controls "github.com/wuyangwang/my-ui/controls"
)

// The built-in zinc pair, used for the very first frame before the
// theme is applied. Every frame after that sets the window's own
// background to the theme's actual color.
const firstPaintBackground = "light-dark(#ffffff, #09090b)"

// Options configures NewWindow.
type Options struct {
	// Title is the window's title, as the taskbar and the system read it.
	Title string
	// Width and Height size the window, in physical pixels; the zero
	// sizes leave the platform's own.
	Width, Height int
	// MinWidth and MinHeight floor the window's size, in the same
	// pixels; either zero leaves that side unbounded.
	MinWidth, MinHeight int
	// Maximized creates the window maximized.
	Maximized bool
	// Center centers the window on its screen.
	Center bool

	// Logo builds a custom mark ahead of the name, with the frame's
	// context. It takes the place of Icon when both are given.
	Logo func(c *ui.Context)
	// Icon is the logo drawn in a themed square mark ahead of the
	// name: pass icons.Component() and nothing else is needed. When
	// neither Icon nor Logo is given, the title bar draws a default
	// component mark.
	Icon *ui.SVG
	// Name is the name of the app the bar shows beside the logo.
	Name string
	// Look picks the style of the window controls, Windows's or the
	// Mac's; the zero look follows the platform (controls.TitleBarLook).
	Look controls.TitleBarLook

	// Theme returns the theme options the frame applies, every frame:
	// the app's appearance and palette live here, so the bar and the
	// content light together. Nil takes the defaults.
	Theme func() myui.Options

	// BackgroundColor is the window's own color, behind the view and
	// around its rounded corners, in mygo's CSS syntax (for example
	// "light-dark(#ffffff, #09090b)"). Empty follows the theme's
	// Background every frame, so a dark frame never leaves the
	// platform's white showing along the edges.
	BackgroundColor string

	// View builds the content under the bar. It runs in a box that
	// takes the room the bar leaves, so a Column of its own fills the
	// window.
	View func(c *ui.Context)
}

// actions are the window controls' targets. Before the window shows,
// and in a Tester's frame, they are no-ops.
type actions struct {
	minimize, maximize, close func()
	isMaximized               func() bool
	setBg                     func(string) error
	setSource                 func(mygo.ThemeSource)
}

// shell is the window's frame state, kept across frames.
type shell struct {
	o          Options
	act        *actions
	maximized  bool
	lastBg     string
	lastSource mygo.ThemeSource
}

// Open is the short form of New for the common case: a window with
// the icon and the name, and the app's view under the bar. The title,
// the look, the theme and the window's size all take their defaults.
func Open(icon *ui.SVG, name string, view func(c *ui.Context)) *mygo.Window {
	return New(Options{Icon: icon, Name: name, View: view})
}

// withDefaults fills in what an app that passes only an icon and a
// name leaves out: the title is the name.
func withDefaults(o Options) Options {
	if o.Title == "" {
		o.Title = o.Name
	}
	return o
}

// New creates and returns the window. The bar's actions go through the
// window it returns: an app that wants to act on the window (an
// intercepted close, a second window) keeps the *mygo.Window. The
// shell renders without one as well, so a Tester's frame draws it.
func New(o Options) *mygo.Window {
	o = withDefaults(o)
	act := &actions{minimize: func() {}, maximize: func() {}, close: func() {}}
	sh := &shell{o: o, act: act}
	content := ui.View(sh.render)
	bg := o.BackgroundColor
	if bg == "" {
		bg = firstPaintBackground
	}
	win := mygo.NewWindow(mygo.WindowOptions{
		Title:  o.Title,
		Width:  o.Width,
		Height: o.Height,
		// The custom title bar takes the chrome over: the system's bar
		// and buttons hide, the shadow, the rounded corners and the
		// resizing edges stay.
		TitleBarStyle:     mygo.TitleBarHidden,
		HideWindowButtons: true,
		Maximized:         o.Maximized,
		BackgroundColor:   bg,
		Content:           content,
	})
	act.minimize = win.Minimize
	act.maximize = win.ToggleMaximize
	act.close = win.Close
	act.isMaximized = win.IsMaximized
	act.setBg = win.SetBackgroundColor
	act.setSource = mygo.Theme.SetSource
	if o.MinWidth > 0 || o.MinHeight > 0 {
		win.SetMinimumSize(o.MinWidth, o.MinHeight)
	}
	if o.Center {
		win.Center()
	}
	return win
}

// render builds one frame of the shell: the theme, the bar with its
// controls wired to the window, then the app's view in the room left.
func (sh *shell) render(c *ui.Context) {
	opts := myui.Options{}
	if sh.o.Theme != nil {
		opts = sh.o.Theme()
	}
	t := myui.Apply(c, opts)
	// The view's theme alone does not recolor the native window frame
	// (the 1px edge Windows and macOS draw around it): that follows
	// the app's appearance. Mirror the frame to it, once per change.
	if sh.o.Theme != nil && sh.act.setSource != nil {
		src := themeSource(opts.Appearance)
		if src != sh.lastSource {
			sh.lastSource = src
			sh.act.setSource(src)
		}
	}
	// The maximized state feeds the bar's restore glyph; the read runs
	// on the window's own thread, as the view does.
	if sh.act.isMaximized != nil {
		sh.maximized = sh.act.isMaximized()
	}
	// The window paints behind the view and around its rounded
	// corners. Keep it in the view's own ink: a forced dark appearance
	// must not leave the platform's white along the four edges.
	if sh.o.BackgroundColor == "" {
		if bg := hexOf(t.Background); bg != sh.lastBg {
			sh.lastBg = bg
			if sh.act.setBg != nil {
				_ = sh.act.setBg(bg)
			}
		}
	}
	bar := controls.TitleBarOptions{
		Name:       sh.o.Name,
		Icon:       sh.o.Icon,
		Look:       sh.o.Look,
		Maximized:  &sh.maximized,
		OnMinimize: sh.act.minimize,
		OnMaximize: sh.act.maximize,
		OnClose:    sh.act.close,
	}
	if sh.o.Logo != nil {
		logo := sh.o.Logo
		bar.Logo = func() { logo(c) }
	}
	// The bar owns the top strip; the app's view takes the rest. The
	// column paints the frame's own background: the root would draw
	// the system's light one under the transparent bar, which shows
	// as a white edge on a forced dark frame.
	ui.Column(c).Fill().Background(t.Background).Children(func() {
		controls.NewTitleBar(c, bar)
		if sh.o.View != nil {
			ui.Box(c).Grow(1).Children(func() { sh.o.View(c) })
		}
	})
}

// themeSource maps a view's appearance to the source of the native
// window frame's colors.
func themeSource(a myui.Appearance) mygo.ThemeSource {
	switch a {
	case myui.AppearanceLight:
		return mygo.ThemeLight
	case myui.AppearanceDark:
		return mygo.ThemeDark
	default:
		return mygo.ThemeSystem
	}
}

// hexOf writes c in the #rrggbb syntax the window's background takes.
func hexOf(c ui.Color) string {
	return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
}
