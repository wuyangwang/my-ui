package myui

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestTextAreaSizesGiveRoomierBounds(t *testing.T) {
	height := func(tt *ui.Tester, name string) float32 {
		info, ok := tt.Access(name)
		if !ok {
			t.Fatalf("%s is not on screen", name)
		}
		return info.Bounds.H
	}
	var a, b string
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		ui.Column(c).Gap(c.Theme().Space(2)).Children(func() {
			NewTextArea(c, &a).Sm().Outline().Label("small")
			NewTextArea(c, &b).Md().Outline().Label("medium")
		})
	}
	tt := ui.NewTester(view, 320, 400)
	tt.Advance(time.Second)
	// The text area is three lines high at least, so the lines have to
	// fill it before the size of the text shows.
	for _, name := range []string{"small", "medium"} {
		if err := tt.Click(name); err != nil {
			t.Fatal(err)
		}
		for i := 0; i < 6; i++ {
			tt.Type("line")
			tt.KeyDown(ui.KeyEnter)
			tt.KeyUp(ui.KeyEnter)
		}
		tt.Advance(time.Second)
	}
	if h := height(tt, "small"); h >= height(tt, "medium") {
		t.Errorf("the small text area is not shorter than the middle's: %v against %v", h, height(tt, "medium"))
	}
}

func TestTextAreaGrowsWithItsText(t *testing.T) {
	var value string
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewTextArea(c, &value).Outline().Label("draft").Lines(2, 6)
	}
	tt := ui.NewTester(view, 320, 320)
	tt.Advance(time.Second)
	before, ok := tt.Access("draft")
	if !ok {
		t.Fatal("the text area is not on screen")
	}
	if err := tt.Click("draft"); err != nil {
		t.Fatal(err)
	}
	tt.Type("one")
	for i := 0; i < 5; i++ {
		tt.KeyDown(ui.KeyEnter)
		tt.KeyUp(ui.KeyEnter)
		tt.Type("line")
	}
	tt.Advance(time.Second)
	after, ok := tt.Access("draft")
	if !ok {
		t.Fatal("the text area is not on screen")
	}
	if after.Bounds.H <= before.Bounds.H {
		t.Errorf("the text area did not grow: %v against %v", after.Bounds.H, before.Bounds.H)
	}
	if value == "" {
		t.Error("the text area holds no text")
	}
}

func TestTextAreaOnChangedReportsTheNewValue(t *testing.T) {
	var value string
	var seen string
	view := func(c *ui.Context) {
		Apply(c, Options{})
		NewTextArea(c, &value).Outline().Label("draft").OnChanged(func(v string) {
			seen = v
		})
	}
	tt := ui.NewTester(view, 320, 320)
	tt.Advance(time.Second)
	if err := tt.Click("draft"); err != nil {
		t.Fatal(err)
	}
	tt.Type("hello")
	tt.Advance(time.Second)
	if seen != "hello" {
		t.Errorf("the change reported %q, want %q", seen, "hello")
	}
}

func TestTextAreaDisabledAndReadOnly(t *testing.T) {
	var locked, ro string
	view := func(c *ui.Context) {
		Apply(c, Options{})
		ui.Column(c).Gap(c.Theme().Space(2)).Children(func() {
			NewTextArea(c, &locked).Outline().Label("Locked").Disabled(true)
			NewTextArea(c, &ro).Outline().Label("ReadOnly").ReadOnly(true)
		})
	}
	tt := ui.NewTester(view, 320, 320)
	tt.Advance(time.Second)
	info, ok := tt.Access("Locked")
	if !ok {
		t.Fatal("disabled text area missing")
	}
	if !info.Disabled {
		t.Error("disabled text area not reported as disabled")
	}
	info2, ok := tt.Access("ReadOnly")
	if !ok {
		t.Fatal("readonly text area missing")
	}
	if !info2.ReadOnly {
		t.Error("readonly text area not reported as readonly")
	}
}

func TestTextAreaErrorTextIsReadWithIt(t *testing.T) {
	var value string
	view := func(c *ui.Context) {
		Apply(c, Options{})
		NewTextArea(c, &value).Outline().Label("draft").Invalid(true).ErrorText("太长了")
	}
	tt := ui.NewTester(view, 320, 320)
	tt.Advance(time.Second)
	info, ok := tt.Access("draft")
	if !ok {
		t.Fatal("the text area is not on screen")
	}
	if info.Description != "太长了" {
		t.Errorf("the text area reads %q, want the error message", info.Description)
	}
}

func TestTextAreaClearableEmptiesTheValue(t *testing.T) {
	var value string
	view := func(c *ui.Context) {
		Apply(c, Options{Appearance: AppearanceLight})
		NewTextArea(c, &value).Outline().Label("draft").Clearable()
	}
	tt := ui.NewTester(view, 320, 320)
	tt.Advance(time.Second)
	if err := tt.Click("draft"); err != nil {
		t.Fatal(err)
	}
	tt.Type("hello")
	tt.Advance(time.Second)
	if _, ok := tt.Access("Clear"); !ok {
		t.Fatal("the clear button is not on screen")
	}
	if err := tt.Click("Clear"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if value != "" {
		t.Errorf("the text area holds %q after the clear, want it empty", value)
	}
}
