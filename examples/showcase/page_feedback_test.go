package main

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

// TestTheAlertPageShowsItsAlerts takes the alert page for a drive: the
// alerts show, with their kinds.
func TestTheAlertPageShowsItsAlerts(t *testing.T) {
	s := &showcase{page: "alert"}
	tt := ui.NewTester(s.view, 960, 620)
	tt.Advance(time.Second)
	for _, what := range []string{"这是一条提示", "保存前看一眼", "操作成功", "出错了"} {
		if _, ok := tt.Find(what); !ok {
			t.Errorf("the alert %q is not on screen", what)
		}
	}
}

// TestTheToastPageShowsAToast: a click shows a toast at the corner of
// the window.
func TestTheToastPageShowsAToast(t *testing.T) {
	s := &showcase{page: "toast"}
	tt := ui.NewTester(s.view, 960, 620)
	tt.Advance(time.Second)
	if err := tt.Click("成功"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if _, ok := tt.Find("操作成功"); !ok {
		t.Fatal("the toast is not on screen")
	}
}

// TestTheDrawerPageOpensItsPanel: a click opens the panel at the edge of
// the window.
func TestTheDrawerPageOpensItsPanel(t *testing.T) {
	s := &showcase{page: "drawer"}
	tt := ui.NewTester(s.view, 960, 620)
	tt.Advance(time.Second)
	if err := tt.Click("打开抽屉"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if _, ok := tt.Find("面板里的内容"); !ok {
		t.Fatal("the panel of the drawer is not on screen")
	}
	if !s.drawerOpen {
		t.Error("the drawer's flag is not true")
	}
}
