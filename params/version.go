package params

import (
	"fmt"
)

const (
	VersionMajor = 0 // Major version component of the current release
	VersionMinor = 0 // Minor version component of the current release
	VersionPatch = 1 // Patch version component of the current release
	VersionMeta  = "beta"
)

// Version holds the textual version string.
var Version = func() string {
	v := fmt.Sprintf("%d.%d.%d", VersionMajor, VersionMinor, VersionPatch)
	if VersionMeta != "" {
		v += "-" + VersionMeta
	}
	return v
}()

func VersionWithCommit(gitCommit string) string {
	vsn := Version
	if len(gitCommit) >= 8 {
		vsn += "-" + gitCommit[:8]
	}
	return vsn
}
