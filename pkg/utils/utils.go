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
	bi, _ := debug.ReadBuildInfo()
	fmt.Println(bi.Main.Version)
	fmt.Println("of " + bi.Main.Path + " compiled in " + bi.GoVersion)
	//fmt.Println(bi.Path)
}
