package display

import (
	"testing"

	. "github.com/wuyangwang/my-ui"

	"github.com/egoist/mygo/ui"
)

func TestResultShowsTitleDescriptionAndStatus(t *testing.T) {
	for _, status := range []ResultStatus{ResultSuccess, ResultWarning, ResultInfo, ResultError} {
		view := func(c *ui.Context) {
			Apply(c, Options{Appearance: AppearanceLight})
			NewResult(c, ResultOptions{Status: status, Title: "完成", Description: "细节在这里"})
		}
		tt := ui.NewTester(view, 320, 240)
		if _, ok := tt.Find("完成"); !ok {
			t.Fatalf("status %d: the title is missing", status)
		}
		if _, ok := tt.Find("细节在这里"); !ok {
			t.Fatalf("status %d: the description is missing", status)
		}
	}
}
