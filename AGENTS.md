# AGENTS.md

my-ui is a third-party component library for MyGo native UI (package
`github.com/egoist/mygo/ui`). It is a separate module that depends on mygo's
public API only — nothing here is a fork or a wrapper of mygo internals.

## Hard rules

- **Depend on the public `ui` API only.** Never import
  `github.com/egoist/mygo/internal/...` (and its `internal/platform` etc.):
  internal packages are invisible to this module (one public leak exists
  upstream, `Element.Access(func(*platform.AccessNode))` — do not route
  around it, see docs/research.md "Upstream issues"). The one exception
  is the `window/` subpackage, the frameless-window shell: it imports
  mygo's root package (`NewWindow`/`WindowOptions`/`TitleBarHidden`)
  because the window actions (`Minimize`/`ToggleMaximize`/`Close`) live
  on `mygo.Window`, not in `ui`. No other package — the root `myui`
  included — imports mygo's root.
- **Pin mygo.** The mygo version my-ui targets is fixed per release —
  currently `v0.10.0` (mygo's `go` directive is 1.27.3), and the
  `window/` subpackage imports mygo's root on top of `ui`. Upgrades to
  mygo go through a deliberate edit of `go.mod`, `README.md` and this
  file together.
- **No cgo, no Bun, no npm.** Pure Go. Bun/npm tooling of the mygo repo is
  dev tooling only and does not exist for this project.
- **Colors and metrics come from `c.Theme()` at build time.** Never cache
  colors in package level state; components must follow `SetTheme` and
  system appearance/accent/high-contrast/text-scale changes on the next
  frame. The tokens layer derives everything from a `*ui.Theme`. A theme
  field the framework leaves empty (`Inverse`, `InverseText`) reads as
  transparent: take the framework's own fallback (`myui.InverseColors`),
  not the bare field, or the thing shows as an empty box.
- **Follow ui's binding convention** ("the input is the returned value"):
  components take state as pointers, return their output, and end where
  the next statement can read the edit. No channels as state, no manual
  onChange wiring. Give list-like state a stable `Context.Key`.
- **Overlay content lives in the overlay.** `ui.TooltipBase`,
  `ui.PopoverBase` and `ui.DialogBase` build their panel inside the
  overlay layer, where the context's parent is the layer and not the
  panel: build the content with `panel.Children(func(){ ... })`, or it
  lands in the window's corner. A trigger that opens an overlay is built
  before the overlay, so its click lands in the same frame; the flag it
  writes is the app's, the library only reads the trigger's `Clicked()`.
- **Main thread.** Components are constructed from views that run on the
  main thread; no goroutines may touch UI state. Post work like mygo does.
- **Components live in grouped subpackages, the root is the core.**
  `myui` holds only the shared base (theme, tokens, variant/size/shape,
  motion, i18n, rules). Components are grouped by function: `controls/`,
  `forms/`, `overlays/`, `display/`, with `bind/` for the per-frame handle
  and `window/` for the frameless shell. Component packages dot-import
  the core for the shared vocabulary and qualify each other with named
  imports; the root never imports a component package (no cycles).
  A widget mygo already provides is re-exported as a var only when
  my-ui neither has nor plans a counterpart (`NewToggleGroup`,
  `Segmented`, `display.NewText`); components on the roadmap (Table,
  Spin, DatePicker, …) are deliberately not re-exported — they will
  get my-ui versions.
- **Optional weight stays out of the root package.** Fonts, icon sets and
  AI components live in their own subpackages; importing root must not
  pull them (no `init` chains into subpackages, no `go:embed` in root).
- **Motion respects the desktop.** Use `Animate*`/transitions from ui;
  components carry no timers of their own, so reduced motion and
  background-unwanted frames come free.
- **Every component moves a little, and says how.** A state change
  fades its colors through `Element.Transition`, and the durations come
  from the motion tokens (`MotionFast`/`MotionNormal`/`MotionSlow`).
  Motion belongs to the component that has it: a button's hover, an
  input's focus, a badge that changes. No timers, no goroutines, no
  motion that ignores the desktop's Reduce Motion. The pose animation
  (`AnimatePose`) is not used: it leaves the element it poses invisible
  (see docs/components/motion.md).
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

- **Research before code.** A new component starts with its spec page,
  `docs/components/<name>.md`: the web implementations it can follow,
  what MyGo's public API can and cannot do, how far to go, where the
  component belongs in the packages, and the API it offers. The spec is
  settled before the first line of code, and it is what the component
  is held to.
- **Every component is documented twice.** Go doc comments on every
  exported symbol, and the spec page's API table: one row per prop
  (method, option or field) with what it does. A prop without a row is
  a prop without a doc.
- **A component ships ready to use.** Components are added one by one,
  and each one is done when it can go live: the states, the events, the
  accessibility and the composition a user expects are in it. What the
  framework cannot do is written down as a limit, not left as a
  "later"; what is later is a new component, not a hole in this one.
- One-line constructors + fluent preset methods: `NewXxx(c, ...).Primary().Sm()`
  applies shared `Variant`/`Size` presets and returns the widget; fluent
  options fall through to the widget's own methods after the preset is applied.
  `Variant`/`Size` 是根包共享词表，`Variant` 零值为 `Native`（不涂预设），
  `Secondary` 为其 shadcn 别名。
- `bind.Bind(c)` is the sugar over the constructors: a handle for one
  window, so a view may bind it once and leave the context out of the
  calls. The constructors in the component packages stay the primitive.
- Input-like controls (Input, TextArea, PasswordInput, Select, ComboBox,
  MultiSelect) fill their parent by default; `Width(w)` pins a fixed
  size. Keep that behavior when adding a field-like component.
- Icons use the root's `IconSizeXs/Sm/Md/Lg` scale, never ad-hoc DIPs.
  A clear/close affordance is an icon with a button role
  (`controls.NewIconAction`), not a ButtonBase wrapper: hover and press
  change only the glyph's ink, through the shared StateTransition. In a
  clearable Select the clear icon replaces the chevron.
- Package docs in code comments are written in English (Go ecosystem
  style); the files under `docs/` and `README.md` are written in Chinese.
- New components first: the ordered plan in docs/roadmap.md decides what
  exists. A component that only renames a `ui` widget with no added
  styling or behavior does not belong here.
- my-ui releases are tagged `vX.Y.Z` independently of mygo. The mygo
  version a release is built against is pinned in `go.mod` and recorded
  in `README.md`.

## The icons package

- `icons` holds the Lucide set and `icons/brand` the Simple Icons set:
  one function per icon, each with its own SVG embedded in a
  package-level variable. The linker drops the functions and the SVGs of
  the icons a program never calls, so an app carries only the icons it
  uses.
- The files in `icons_gen.go` and `aliases_gen.go` are generated: change
  the SVGs or `icons/gen/aliases.go` and run `go generate ./icons`
  (and `go generate ./icons/brand`).
- `icons/all` holds the names of the whole set and the icon of every
  name, for an app that shows a picker of icons. It keeps every icon of
  the set in the binary, so it stands apart from the set itself.
- `icons` depends on `github.com/egoist/mygo/ui` only; it does not depend
  on the root package, and the root package does not depend on it.
