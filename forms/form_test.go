package forms

import (
	. "github.com/wuyangwang/my-ui"
	testutil "github.com/wuyangwang/my-ui/internal/testutil"
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestFieldShowsItsLabelDescriptionAndError(t *testing.T) {
	var value string
	view := func(c *ui.Context) {
		Apply(c, Options{})
		NewField(c, "邮箱", func() ui.Element {
			return NewInput(c, &value).Outline().Control()
		}).Description("我们只用它发通知").Required(true).Error("邮箱格式不对")
	}
	tt := ui.NewTester(view, 320, 300)
	tt.Advance(time.Second)

	if _, ok := tt.Find("邮箱"); !ok {
		t.Error("the label is not on screen")
	}
	if _, ok := tt.Find("我们只用它发通知"); !ok {
		t.Error("the description is not on screen")
	}
	info, ok := tt.Access("邮箱格式不对")
	if !ok {
		t.Fatal("the error message is not on screen")
	}
	// The error is announced as it appears.
	if info.Role != ui.RoleAlert {
		t.Errorf("the error message is %v, want an alert", info.Role)
	}
}

func TestFieldLabelFocusesTheControl(t *testing.T) {
	var value string
	view := func(c *ui.Context) {
		Apply(c, Options{})
		NewField(c, "邮箱", func() ui.Element {
			return NewInput(c, &value).Outline().Control()
		})
	}
	tt := ui.NewTester(view, 320, 300)
	tt.Advance(time.Second)
	if err := tt.Click("邮箱"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if !tt.Focused("邮箱") {
		t.Error("the label did not focus the control")
	}
}

func TestFieldValidatesAfterItIsLeft(t *testing.T) {
	var value string
	view := func(c *ui.Context) {
		Apply(c, Options{})
		NewField(c, "邮箱", func() ui.Element {
			return NewInput(c, &value).Outline().Control()
		}).Validate(&value, Required("请填写邮箱"))
		ui.Button(c, "别处").Label("别处")
	}
	tt := ui.NewTester(view, 320, 300)
	tt.Advance(time.Second)

	// The field has not been left: it shows nothing, as a form does.
	if _, ok := tt.Access("请填写邮箱"); ok {
		t.Error("the error shows before the field is left")
	}

	if err := tt.Click("邮箱"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if err := tt.Click("别处"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if _, ok := tt.Access("请填写邮箱"); !ok {
		t.Error("the error does not show after the field is left")
	}
}

func TestFieldValidatesTheValueItHolds(t *testing.T) {
	var value string
	view := func(c *ui.Context) {
		Apply(c, Options{})
		NewField(c, "邮箱", func() ui.Element {
			return NewInput(c, &value).Outline().Control()
		}).Validate(&value, Required("请填写邮箱"), Email("邮箱格式不对"))
		ui.Button(c, "别处").Label("别处")
	}
	tt := ui.NewTester(view, 320, 300)
	tt.Advance(time.Second)
	if err := tt.Click("邮箱"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if err := tt.Click("别处"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if _, ok := tt.Access("请填写邮箱"); !ok {
		t.Error("the error of the empty value does not show")
	}

	// A value that holds leaves no error.
	if err := tt.Click("邮箱"); err != nil {
		t.Fatal(err)
	}
	tt.Type("hi@example.com")
	tt.Advance(time.Second)
	if err := tt.Click("别处"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if _, ok := tt.Access("请填写邮箱"); ok {
		t.Error("the error shows although the value holds")
	}
}

func TestFieldKnowsItWasUsed(t *testing.T) {
	var value string
	var field *Field
	view := func(c *ui.Context) {
		Apply(c, Options{})
		field = NewField(c, "邮箱", func() ui.Element {
			return NewInput(c, &value).Outline().Control()
		}).Validate(&value, Required("请填写邮箱"))
		ui.Button(c, "别处").Label("别处")
	}
	tt := ui.NewTester(view, 320, 300)
	tt.Advance(time.Second)
	if field.Touched() {
		t.Error("the field says it was used before it was left")
	}
	if field.Dirty() {
		t.Error("the field says it is dirty before anything was typed")
	}

	if err := tt.Click("邮箱"); err != nil {
		t.Fatal(err)
	}
	tt.Type("hi")
	tt.Advance(time.Second)
	if err := tt.Click("别处"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if !field.Touched() {
		t.Error("the field does not say it was used")
	}
	if !field.Dirty() {
		t.Error("the field does not say it is dirty")
	}
}

func TestPasswordInputHidesAndShowsTheText(t *testing.T) {
	var value string
	view := func(c *ui.Context) {
		Apply(c, Options{})
		NewPasswordInput(c, &value).Outline().Label("密码")
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if err := tt.Click("密码"); err != nil {
		t.Fatal(err)
	}
	tt.Type("secret")
	tt.Advance(time.Second)

	info, ok := tt.Access("密码")
	if !ok {
		t.Fatal("the password field is not on screen")
	}
	if info.Value != "" {
		t.Errorf("the hidden password shows %q", info.Value)
	}
	if _, ok := tt.Access("Show password"); !ok {
		t.Fatal("the button that shows the password is not on screen")
	}
	if err := tt.Click("Show password"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	info2, ok := tt.Access("密码")
	if !ok {
		t.Fatal("the password field is not on screen")
	}
	if info2.Value != "secret" {
		t.Errorf("the shown password holds %q, want %q", info2.Value, "secret")
	}
}

func TestPasswordInputKeepsTheEditorWhenShown(t *testing.T) {
	var value string
	var input *PasswordInput
	view := func(c *ui.Context) {
		Apply(c, Options{})
		input = NewPasswordInput(c, &value).Outline().Label("密码")
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	if err := tt.Click("密码"); err != nil {
		t.Fatal(err)
	}
	tt.Type("secret")
	tt.Advance(time.Second)
	if value != "secret" {
		t.Fatalf("the value is %q, want secret", value)
	}
	// The eye button keeps the focus on the editor, so typing can go on
	// after a show/hide: the editor is not rebuilt.
	if err := tt.Click("Show password"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if !input.Visible() {
		t.Error("the text is still hidden after the eye was pressed")
	}
	if !tt.Focused("密码") {
		t.Error("the editor lost the focus on show")
	}
	if value != "secret" {
		t.Errorf("the value changed on show: %q", value)
	}
	if err := tt.Click("Hide password"); err != nil {
		t.Fatal(err)
	}
	tt.Advance(time.Second)
	if input.Visible() {
		t.Error("the text is still shown after the eye was pressed again")
	}
	if value != "secret" {
		t.Errorf("the value changed on hide: %q", value)
	}
}

func TestPasswordInputTakesItsOwnIconAndColor(t *testing.T) {
	var value string
	red := ui.Hex("#ff0000")
	view := func(c *ui.Context) {
		Apply(c, Options{})
		NewPasswordInput(c, &value).Outline().Label("密码").
			EyeIcons(testIcon, testIcon).EyeColor(red)
	}
	tt := ui.NewTester(view, 320, 240)
	tt.Advance(time.Second)
	info, ok := tt.Access("Show password")
	if !ok {
		t.Fatal("the button that shows the password is not on screen")
	}
	// The icon of the button is drawn in the color the app chose.
	b := info.Bounds
	r, g, bl := testutil.Pixel(tt, b.X+b.W/2, b.Y+b.H/2)
	if r != red.R || g != red.G || bl != red.B {
		t.Errorf("the icon drew (%d, %d, %d), want the color the app chose", r, g, bl)
	}
}
