package helper

import (
	"fmt"
	"runtime"
)

const (
	onWindows = runtime.GOOS == "windows"
)

// PlatformIdentifier converts identifier for the current platform.
func PlatformIdentifier(identifier string) string {
	// From https://golang.org/pkg/runtime/#GOARCH
	// GOOS is the running program's operating system target: one of darwin, freebsd, linux, and so on.
	// GOARCH is the running program's architecture target: one of 386, amd64, arm, s390x, and so on.
	return fmt.Sprintf("%s_%s/%s", runtime.GOOS, runtime.GOARCH, identifier)
}

// MandatoryUpdates returns mandatory updates that should be loaded on install
// or reset.
func MandatoryUpdates() (identifiers []string) {
	// Binaries
	if onWindows {
		identifiers = []string{
			PlatformIdentifier("core/portmaster-core.exe"),
			PlatformIdentifier("kext/portmaster-kext.dll"),
			PlatformIdentifier("kext/portmaster-kext.sys"),
			PlatformIdentifier("start/portmaster-start.exe"),
			PlatformIdentifier("notifier/portmaster-notifier.exe"),
			PlatformIdentifier("notifier/portmaster-snoretoast.exe"),
		}
	} else {
		identifiers = []string{
			PlatformIdentifier("core/portmaster-core"),
			PlatformIdentifier("start/portmaster-start"),
			PlatformIdentifier("notifier/portmaster-notifier"),
		}
	}

	// Components, Assets and Data
	identifiers = append(
		identifiers,

		// User interface components
		PlatformIdentifier("app/portmaster-app.zip"),
		"all/ui/modules/portmaster.zip",
		"all/ui/modules/assets.zip",

		// Filter lists data
		"all/intel/lists/base.dsdl",
		"all/intel/lists/intermediate.dsdl",
		"all/intel/lists/urgent.dsdl",

		// Geo IP data
		"all/intel/geoip/geoipv4.mmdb.gz",
		"all/intel/geoip/geoipv6.mmdb.gz",
	)

	return identifiers
}

// AutoUnpackUpdates returns assets that need unpacking.
func AutoUnpackUpdates() []string {
	return []string{
		PlatformIdentifier("app/portmaster-app.zip"),
	}
}
