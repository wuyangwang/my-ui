// Package all holds the whole icon set: the names of its icons and the
// icon of every name.
//
// Import it only when the app needs the whole set, as a picker of
// icons does: it keeps every icon of the set in the binary, where the
// set itself keeps only the icons the app calls.
//
//	all.Names            // the names of the icons
//	all.ByName("save")   // the icon of that name
package all
