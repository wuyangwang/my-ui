# AGENTS.md

my-ui is a third-party component library for MyGo native UI (package
`github.com/egoist/mygo/ui`). It is a separate module that depends on mygo's
public API only — nothing here is a fork or a wrapper of mygo internals.

## Hard rules

- **Depend on the public `ui` API only.** Never import
  `github.com/egoist/mygo/internal/...` (and its `internal/platform` etc.):
  internal packages are invisible to this module (one public leak exists
  upstream, `Element.Access(func(*platform.AccessNode))` — do not route
  around it, see docs/research.md "Upstream issues").
- **Pin mygo.** The mygo version my-ui targets is fixed per release —
  currently `v0.8.3` (mygo's `go` directive is 1.27.3). While no code
  imports mygo yet, `go mod tidy` keeps `go.mod` free of the require
  line; the first component code adds it and, from then on, upgrades to
  mygo go through a deliberate edit of `go.mod`, `README.md` and this
  file together.
- **No cgo, no Bun, no npm.** Pure Go. Bun/npm tooling of the mygo repo is
  dev tooling only and does not exist for this project.
- **Colors and metrics come from `c.Theme()` at build time.** Never cache
  colors in package level state; components must follow `SetTheme` and
  system appearance/accent/high-contrast/text-scale changes on the next
  frame. The tokens layer derives everything from a `*ui.Theme`.
- **Follow ui's binding convention** ("the input is the returned value"):
  components take state as pointers, return their output, and end where
  the next statement can read the edit. No channels as state, no manual
  onChange wiring. Give list-like state a stable `Context.Key`.
- **Main thread.** Components are constructed from views that run on the
  main thread; no goroutines may touch UI state. Post work like mygo does.
- **Optional weight stays out of the root package.** Heavy or rare things
  (fonts, icon sets, AI components) live in subpackages; importing root
  must not pull them (no `init` chains into subpackages, no
  `go:embed` in root). Every subpackage stands alone beside `myui`.
- **Motion respects the desktop.** Use `Animate*`/transitions from ui;
  components carry no timers of their own, so reduced motion and
  background-unwanted frames come free.
- **Keep embedded assets licensed.** Lucide SVGs (ISC), simple-icons
  (CC0), fonts (OFL) only; record additions and their licenses in the
  subpackage README.

## Commands

From the module root:

```sh
go build ./... && go vet ./...                       # every change
GOOS=linux go build ./... && GOOS=linux go vet ./... # cross-compile
GOOS=windows go vet ./...                            # (CI matrix mirrors mygo)
GOOS=js GOARCH=wasm go vet ./...                     # native UI in browsers
go test ./...                                        # ui.Tester based, no real windows
```

`ui.Tester` runs without real windows or a GPU (fake backend), so tests are
CI-safe on all three desktop OSes plus js/wasm.

## Testing rules (inherited from mygo)

- Tests never pass or fail on wall-clock time: wait for the condition or
  event; a `ui.Tester`'s clock is advanced, never slept against.
- Measure work by counts, not durations.
- Cover states per widget: default, hover, pressed, disabled, read-only,
  themed (light and dark), and RTL where layout mirrors.

## Conventions

- One-line constructors with `Variant`/`Size`/`Inset` enums; fluent options
  fall through to the widget's own methods after the preset is applied.
- Package docs in code comments are written in English (Go ecosystem
  style); the files under `docs/` and `README.md` are written in Chinese.
- New components first: the gap table in docs/roadmap.md decides what
  exists. A component that only renames a `ui` widget with no added
  styling or behavior does not belong here.
- Version releases of my-ui correspond one-to-one with the mygo version
  they are built against.
