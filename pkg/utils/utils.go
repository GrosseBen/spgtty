package utils

import "runtime/debug"

func Version() string {
	bi, _ := debug.ReadBuildInfo()
	version := (bi.Main.Version)
	return version
}
