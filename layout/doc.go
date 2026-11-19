// Package layout holds my-ui's page furniture: grids, scroll containers,
// headers, rails and floating panels. It depends on mygo's public ui API
// only, and on the root package for the shared theme vocabulary.
//
// The four names NewGrid, NewScroll, NewSidebar and NewSplit replace the
// unchanged mygo re-exports that used to live in display/reexports.go. They
// keep the mygo signatures, so existing calls compile, but paint themselves
// from c.Theme() on every frame with shadcn-style tokens (Background,
// Surface, Border, Text, Accent, TextSize and Radius). The richer
// components in this package (ScrollArea, SideNav, DraggablePanel, …) add
// the behavior my-ui promises.
package layout
