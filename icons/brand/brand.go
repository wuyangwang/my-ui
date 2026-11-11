// Package brand holds the brand icon set for MyGo apps: the logos of
// the companies and the projects.
//
// Every icon is a function that returns the parsed SVG:
//
//	ui.Icon(c, brand.GitHub())
//
// The SVG of an icon is parsed once, on the first call, and kept for
// the next one.
//
// The icons come from Simple Icons, which is under the CC0 license; see
// LICENSE.md in this directory for the license and its notice.
package brand

import (
	"sync"

	"github.com/egoist/mygo/ui"
)

//go:generate go run ../gen -svg svg -out icons_gen.go -pkg brand -set simple-icons -data data.json -aliases-out= -all-out all/all_gen.go -import github.com/wuyangwang/my-ui/icons/brand

// cache holds the icons parsed so far, by name.
var cache = map[string]*ui.SVG{}

// mu guards the cache.
var mu sync.Mutex

// parse parses an icon's SVG once and keeps it for the next call.
//
// The SVG of an icon is embedded in the variable that holds it, so an
// icon a program never calls is left out of its binary: the linker
// drops the function and the SVG with it.
func parse(name string, data []byte) *ui.SVG {
	mu.Lock()
	defer mu.Unlock()
	if s, ok := cache[name]; ok {
		return s
	}
	s := ui.MustParseSVG(data)
	cache[name] = s
	return s
}
