// Package all holds the whole brand set: the names of its icons and the
// icon of every name.
//
// Import it only when the app needs the whole set, as a picker of
// brands does: it keeps every icon of the set in the binary, where the
// set itself keeps only the icons the app calls.
//
//	all.Names            // the names of the icons
//	all.ByName("GitHub") // the icon of that name
package all
