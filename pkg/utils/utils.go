package utils

import (
	"fmt"
	"runtime/debug"
)

func version() string {
	bi, _ := debug.ReadBuildInfo()
	version := (bi.Main.Version)
	return version
}

func PrintVersion() {
	v := version()
	fmt.Println("spgtty in " + v)
}
