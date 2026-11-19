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
		"dropHint":       "Drag a file here, or use the button",
		"open":           "Open",
		"qrTooLong":      "Value too long to encode",
		"user":           "User",
		"assistant":      "Assistant",
		"effort":         "Reasoning effort",
		"effortLow":      "Low",
		"effortMedium":   "Medium",
		"effortHigh":     "High",
		"temperature":    "Temperature",
		"topP":           "Top P",
		"presetPrecise":  "Precise",
		"presetBalanced": "Balanced",
		"presetCreative": "Creative",
		"searchModels":   "Search models…",
		"noModels":       "No matching models",
		"noResults":      "No results found",
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
		"dropHint":       "拖入文件，或点击按钮选择",
		"open":           "打开",
		"qrTooLong":      "内容过长，无法生成二维码",
		"user":           "你",
		"assistant":      "助手",
		"effort":         "思考强度",
		"effortLow":      "低",
		"effortMedium":   "中",
		"effortHigh":     "高",
		"temperature":    "温度",
		"topP":           "Top P",
		"presetPrecise":  "精确",
		"presetBalanced": "均衡",
		"presetCreative": "创意",
		"searchModels":   "搜索模型…",
		"noModels":       "没有匹配的模型",
		"noResults":      "没有匹配的结果",
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
		"dropHint":       "拖入檔案，或點擊按鈕選擇",
		"open":           "開啟",
		"qrTooLong":      "內容過長，無法產生二維碼",
		"user":           "你",
		"assistant":      "助手",
		"effort":         "思考強度",
		"effortLow":      "低",
		"effortMedium":   "中",
		"effortHigh":     "高",
		"temperature":    "溫度",
		"topP":           "Top P",
		"presetPrecise":  "精確",
		"presetBalanced": "均衡",
		"presetCreative": "創意",
		"searchModels":   "搜尋模型…",
		"noModels":       "沒有符合的模型",
		"noResults":      "沒有符合的結果",
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
