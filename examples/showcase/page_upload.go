package main

import (
	"github.com/egoist/mygo/ui"
	"github.com/wuyangwang/my-ui/forms"
)

// uploadPage demonstrates a drop zone, with an optional app-supplied file
// picker. The library never calls mygo.OpenFile: that API lives in mygo's
// root package, which a component library must not import. The app passes
// the chosen path back through OnPick.
func (s *showcase) uploadPage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "拖放上传", "可拖入文件，也可注入 OnPick 在 app 代码里打开文件选择器", func() {
		s.card(c, t, func() {
			forms.NewUpload(c, forms.UploadOptions{
				Accept: func(path string) bool { return true },
				OnFiles: func(paths []string) {
					if len(paths) > 0 {
						s.lastAction = paths[0]
					}
				},
				OnPick: func() (string, bool) {
					// 真正的 app 在这里调用 mygo.OpenFile：
					//   path, ok := mygo.OpenFile(mygo.OpenFileOptions{...})
					//   return path, ok
					// 选择器属于 app 代码，不属于库；示例直接给一个路径演示回调。
					return "/demo/report.pdf", true
				},
			})
			ui.Text(c, "最近选择："+s.lastAction).TextColor(t.TextMuted)
		})
	})
}
