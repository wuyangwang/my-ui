// Package icons holds the Lucide icon set for MyGo apps.
//
// Every icon is a function that returns the parsed SVG:
//
//	ui.Icon(c, icons.Save())
//
// The SVG of an icon is parsed once, on the first call, and kept for
// the next one.
//
// The icons come from Lucide, which is under the ISC license; see
// LICENSE in this directory for the license and its notice.
package icons

import (
	"sync"

	"github.com/egoist/mygo/ui"
)

//go:generate go run ./gen -svg svg -out icons_gen.go -aliases-out aliases_gen.go

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
