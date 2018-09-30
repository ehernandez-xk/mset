package mset

import "fmt"

var (
	// VersionMajor is for an API incompatible changes
	VersionMajor = 0
	// VersionMinor is for functionality in a backwards-compatible manner
	VersionMinor = 0
	// VersionPatch is for backwards-compatible bug fixes
	VersionPatch = 1
)

// Version is the specification version that the package types support.
var Version = fmt.Sprintf("%d.%d.%d",
	VersionMajor, VersionMinor, VersionPatch)
