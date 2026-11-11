// Command gen writes the Go functions of the icon set.
//
// It reads the SVG files of the set and writes one function per icon,
// each with the SVG embedded in it, so that an icon a program does not
// call is left out of its binary. It also writes the functions of the
// aliases in aliases.go.
//
// Usage, from the icons directory:
//
//	go generate ./icons
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"os"
	"sort"
	"strings"
)

func main() {
	svgDir := flag.String("svg", "svg", "the directory of the SVG files")
	out := flag.String("out", "icons_gen.go", "the file of the icon functions")
	aliasOut := flag.String("aliases-out", "aliases_gen.go", "the file of the alias functions")
	allOut := flag.String("all-out", "", "the file of the names of the set and the icons of them")
	pkg := flag.String("pkg", "icons", "the name of the package")
	set := flag.String("set", "Lucide", "the name of the set, for the docs")
	data := flag.String("data", "", "the JSON of the set, with the titles of its icons")
	importPath := flag.String("import", "", "the import path of the set, for the file of the names")
	flag.Parse()

	files, err := svgFiles(*svgDir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "gen:", err)
		os.Exit(1)
	}
	titles, err := titlesOf(*data)
	if err != nil {
		fmt.Fprintln(os.Stderr, "gen:", err)
		os.Exit(1)
	}

	names := map[string]string{}  // the Go name of an icon to its file
	byFile := map[string]string{} // the name of an SVG file to its Go name
	skipped := 0
	for _, f := range files {
		name := goName(f)
		if title, ok := titles[strings.TrimSuffix(f, ".svg")]; ok {
			if n := nameOf(title); n != "" {
				name = n
			}
		}
		if prev, ok := names[name]; ok {
			// The set has an icon under two names, as arrow-down-0-1 and
			// arrow-down-01 are one icon: the first name wins.
			fmt.Fprintf(os.Stderr, "gen: %s is the same icon as %s, so it is left out\n", f, prev)
			skipped++
			continue
		}
		names[name] = f
		byFile[strings.TrimSuffix(f, ".svg")] = name
	}

	if err := writeIcons(*out, names, *pkg, *set); err != nil {
		fmt.Fprintln(os.Stderr, "gen:", err)
		os.Exit(1)
	}
	if *aliasOut != "" {
		if err := writeAliases(*aliasOut, names, byFile, *pkg, *set); err != nil {
			fmt.Fprintln(os.Stderr, "gen:", err)
			os.Exit(1)
		}
	}
	if *allOut != "" {
		if err := writeAll(*allOut, names, *importPath); err != nil {
			fmt.Fprintln(os.Stderr, "gen:", err)
			os.Exit(1)
		}
	}
	fmt.Printf("gen: %d icons", len(names))
	if *aliasOut != "" {
		fmt.Printf(", %d aliases", len(aliases))
	}
	if skipped > 0 {
		fmt.Printf(", %d names left out", skipped)
	}
	fmt.Println()
}

// titlesOf reads the JSON of a set, whose icons have a title and a slug,
// and returns the title of every icon by its slug.
func titlesOf(path string) (map[string]string, error) {
	if path == "" {
		return nil, nil
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var icons []struct {
		Title string `json:"title"`
		Slug  string `json:"slug"`
	}
	if err := json.Unmarshal(data, &icons); err != nil {
		return nil, err
	}
	out := map[string]string{}
	for _, i := range icons {
		out[i.Slug] = i.Title
	}
	return out, nil
}

// svgFiles lists the SVG files of a directory, by name.
func svgFiles(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var files []string
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".svg") {
			continue
		}
		files = append(files, e.Name())
	}
	sort.Strings(files)
	return files, nil
}

// goName turns the name of an SVG file into the name of its function:
// "arrow-up-right.svg" becomes ArrowUpRight, "loader-2.svg" becomes
// Loader2.
func goName(file string) string {
	base := strings.TrimSuffix(file, ".svg")
	var b strings.Builder
	for _, word := range strings.FieldsFunc(base, func(r rune) bool {
		return r == '-' || r == '_' || r == ' ' || r == '.'
	}) {
		if b.Len() > 0 && word[0] >= '0' && word[0] <= '9' {
			// A number rides along with the word before it: Loader2.
			b.WriteString(word)
			continue
		}
		b.WriteString(strings.ToUpper(word[:1]))
		b.WriteString(word[1:])
	}
	name := b.String()
	if name == "" {
		return ""
	}
	if name[0] >= '0' && name[0] <= '9' {
		// A name may not start with a number.
		name = "N" + name
	}
	return name
}

// nameOf turns the title of an icon into the name of its function,
// keeping the spelling of the title: "GitHub" becomes GitHub, "Node.js"
// becomes Nodejs.
func nameOf(title string) string {
	var b strings.Builder
	for _, r := range title {
		switch {
		case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z', r >= '0' && r <= '9':
			b.WriteRune(r)
		}
	}
	name := b.String()
	if name == "" {
		return ""
	}
	// A name starts with a capital, as the names of the set's own icons.
	name = strings.ToUpper(name[:1]) + name[1:]
	if name[0] >= '0' && name[0] <= '9' {
		// A name may not start with a number.
		name = "N" + name
	}
	return name
}

// writeAll writes the names of the set and the icon of every one of
// them, so that an app can show the whole set: the file keeps every
// icon of the set in the binary, so it stands apart from the set.
func writeAll(path string, names map[string]string, importPath string) error {
	var b bytes.Buffer
	b.WriteString("// Code generated by \"go generate ./icons\"; DO NOT EDIT.\n\n")
	b.WriteString("package all\n\n")
	b.WriteString("import (\n")
	b.WriteString("\t\"github.com/egoist/mygo/ui\"\n")
	fmt.Fprintf(&b, "\t%q\n", importPath)
	b.WriteString(")\n\n")

	keys := make([]string, 0, len(names))
	for k := range names {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	b.WriteString("// Names are the names of the icons of the set.\n")
	b.WriteString("var Names = []string{\n")
	for _, name := range keys {
		fmt.Fprintf(&b, "\t%q,\n", name)
	}
	b.WriteString("}\n\n")

	b.WriteString("// byName is the icon of every name of the set.\n")
	b.WriteString("var byName = map[string]func() *ui.SVG{\n")
	for _, name := range keys {
		fmt.Fprintf(&b, "\t%q: %s.%s,\n", name, pkgAlias(importPath), name)
	}
	b.WriteString("}\n\n")

	b.WriteString("// ByName returns the icon of that name, or nil when the set has\n")
	b.WriteString("// none of that name.\n")
	b.WriteString("func ByName(name string) *ui.SVG {\n")
	b.WriteString("\tfn, ok := byName[name]\n")
	b.WriteString("\tif !ok {\n\t\treturn nil\n\t}\n")
	b.WriteString("\treturn fn()\n")
	b.WriteString("}\n")
	return writeGo(path, &b)
}

// pkgAlias is the name the set's package is known by in the file of the
// names.
func pkgAlias(importPath string) string {
	if i := strings.LastIndex(importPath, "/"); i >= 0 {
		return importPath[i+1:]
	}
	return importPath
}

// writeGo writes the Go source, formatted as gofmt would.
func writeGo(path string, b *bytes.Buffer) error {
	src, err := format.Source(b.Bytes())
	if err != nil {
		return err
	}
	return os.WriteFile(path, src, 0o644)
}

// writeIcons writes the function of every icon.
func writeIcons(path string, names map[string]string, pkg, set string) error {
	var b bytes.Buffer
	b.WriteString("// Code generated by \"go generate ./icons\"; DO NOT EDIT.\n\n")
	b.WriteString("package " + pkg + "\n\n")
	b.WriteString("import (\n\t_ \"embed\"\n\n\t\"github.com/egoist/mygo/ui\"\n)\n\n")

	keys := make([]string, 0, len(names))
	for k := range names {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, name := range keys {
		file := names[name]
		doc := strings.TrimSuffix(file, ".svg")
		fmt.Fprintf(&b, "//go:embed %s\n", "svg/"+file)
		fmt.Fprintf(&b, "var svg%s []byte\n\n", name)
		fmt.Fprintf(&b, "// %s returns the %s %q icon.\n", name, set, doc)
		fmt.Fprintf(&b, "func %s() *ui.SVG {\n", name)
		fmt.Fprintf(&b, "\treturn parse(%q, svg%s)\n", doc, name)
		b.WriteString("}\n\n")
	}
	return writeGo(path, &b)
}

// writeAliases writes the function of every alias, pointing at the icon
// it names. An alias that the set has a name for is left out: the
// set's own name wins.
func writeAliases(path string, names map[string]string, byFile map[string]string, pkg, set string) error {
	var b bytes.Buffer
	b.WriteString("// Code generated by \"go generate ./icons\"; DO NOT EDIT.\n\n")
	b.WriteString("package " + pkg + "\n\n")
	b.WriteString("import \"github.com/egoist/mygo/ui\"\n\n")

	keys := make([]string, 0, len(aliases))
	for k := range aliases {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, alias := range keys {
		target := aliases[alias]
		fn, ok := byFile[target]
		if !ok {
			return fmt.Errorf("the alias %s points at %s, which the set does not have", alias, target)
		}
		if _, ok := names[alias]; ok {
			fmt.Fprintf(os.Stderr, "gen: the alias %s is a name of the set, so it is left out\n", alias)
			continue
		}
		fmt.Fprintf(&b, "// %s is an alias of the %s %q icon.\n", alias, set, target)
		fmt.Fprintf(&b, "func %s() *ui.SVG {\n", alias)
		fmt.Fprintf(&b, "\treturn %s()\n", fn)
		b.WriteString("}\n\n")
	}
	return writeGo(path, &b)
}
