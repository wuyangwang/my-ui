package main

import (
	"fmt"
	"strings"

	"github.com/egoist/mygo/ui"
	displaypkg "github.com/wuyangwang/my-ui/display"
)

// qrCodePage shows the QR code's colors, icon and status overlays.
func (s *showcase) qrCodePage(c *ui.Context, t *ui.Theme) {
	s.sectionNote(c, t, "基本", "库内编码，字节模式，纠错级别 L", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(4)).AlignItems(ui.Center).Children(func() {
				displaypkg.NewQRCode(c, "https://lobehub.com", displaypkg.QRCodeOptions{
					Size: 128,
				})
				displaypkg.NewQRCode(c, "https://mygo.lobehub.com", displaypkg.QRCodeOptions{
					Size: 128, NoBorder: true,
				})
			})
		})
	})
	s.sectionNote(c, t, "状态", "加载与过期覆盖层", func() {
		s.card(c, t, func() {
			ui.Row(c).Gap(t.Space(4)).AlignItems(ui.Center).Children(func() {
				displaypkg.NewQRCode(c, "https://lobehub.com", displaypkg.QRCodeOptions{
					Size: 120, Status: displaypkg.QRLoading,
				})
				displaypkg.NewQRCode(c, "https://lobehub.com", displaypkg.QRCodeOptions{
					Size: 120, Status: displaypkg.QRExpired,
					OnRefresh: func() { s.lastAction = "QR refreshed" },
				})
			})
		})
	})
	s.sectionNote(c, t, "过长的值", "超出版本 1–5 容量时显示错误，不再空白", func() {
		s.card(c, t, func() {
			long := "https://example.com/" + strings.Repeat("a", 160)
			// NewQRCodeCard 多返回一个 Valid()，app 可据此决定是否换一个短值。
			q := displaypkg.NewQRCodeCard(c, long, displaypkg.QRCodeOptions{Size: 120})
			ui.Text(c, "可编码："+fmt.Sprint(q.Valid())).FontSize(c.Theme().FontSize).
				TextColor(t.TextMuted)
		})
	})
}
