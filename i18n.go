package myui

import (
	"strings"
	"sync"

	"github.com/egoist/mygo/ui"
)

// The strings the controls show, by key, in the languages the library
// ships.
var catalog = map[string]map[string]string{
	"en": {
		"clear":          "Clear",
		"close":          "Close",
		"showPassword":   "Show password",
		"hidePassword":   "Hide password",
		"minimize":       "Minimize",
		"maximize":       "Maximize",
		"restore":        "Restore",
		"previous":       "Previous page",
		"next":           "Next page",
		"searchCommands": "Search commands…",
		"noCommands":     "No matching commands",
		"dropFiles":      "Drop files here",
		"dropHint":       "Drag and drop only; no file picker",
	},
	"zh-Hans": {
		"clear":          "清除",
		"close":          "关闭",
		"showPassword":   "显示密码",
		"hidePassword":   "隐藏密码",
		"minimize":       "最小化",
		"maximize":       "最大化",
		"restore":        "还原",
		"previous":       "上一页",
		"next":           "下一页",
		"searchCommands": "搜索命令…",
		"noCommands":     "没有匹配的命令",
		"dropFiles":      "把文件拖到这里",
		"dropHint":       "拖放区只接收文件，不打开文件选择器",
	},
	"zh-Hant": {
		"clear":          "清除",
		"close":          "關閉",
		"showPassword":   "顯示密碼",
		"hidePassword":   "隱藏密碼",
		"minimize":       "最小化",
		"maximize":       "最大化",
		"restore":        "還原",
		"previous":       "上一頁",
		"next":           "下一頁",
		"searchCommands": "搜尋命令…",
		"noCommands":     "沒有符合的命令",
		"dropFiles":      "把檔案拖到這裡",
		"dropHint":       "只接受拖放，不開啟檔案選擇器",
	},
}

// mu guards the catalog, which an app may change.
var mu sync.RWMutex

// Message returns the string of the key in the language of the frame.
//
// The language is the locale of the frame, as the app sets it with
// App.SetLocale: a tag with a region ("zh-Hans-CN") falls back to its
// language ("zh-Hans"), and a language the library has no strings for
// falls back to English.
//
// An app changes the strings of the library with SetMessages.
func Message(c *ui.Context, key string) string {
	return message(c.Locale(), key)
}

// SetMessages adds the strings of the app for a language, which take the
// place of the library's own for that language. Call it before the
// first frame.
func SetMessages(locale string, msgs map[string]string) {
	mu.Lock()
	defer mu.Unlock()
	if catalog[locale] == nil {
		catalog[locale] = map[string]string{}
	}
	for k, v := range msgs {
		catalog[locale][k] = v
	}
}

// message is the string of the key in the language of the tag, as the
// chain falls back: the tag, its language, English.
func message(tag, key string) string {
	mu.RLock()
	defer mu.RUnlock()
	for _, t := range fallback(tag) {
		if s, ok := catalog[t][key]; ok {
			return s
		}
	}
	return key
}

// fallback is the chain of the languages a tag falls back to: the tag
// itself, then its language without its region, then English.
func fallback(tag string) []string {
	out := []string{tag}
	for {
		i := strings.LastIndexAny(tag, "-_")
		if i <= 0 {
			break
		}
		tag = tag[:i]
		out = append(out, tag)
	}
	return append(out, "en")
}
