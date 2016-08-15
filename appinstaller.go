// Go app to install a list of OS X Apps

package main

import (
	"flag"
	"github.com/natewalck/macosutils"
)

func main() {
	var (
		dmgPath = flag.String("dmg", "", "path to dmg")
	)
	flag.Parse()

	mountpoint := macosutils.MountDmg(*dmgPath)
	macosutils.InstallApp(mountpoint)
	macosutils.UnmountDmg(mountpoint)
}
