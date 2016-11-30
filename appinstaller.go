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

    my_dmg := macosutils.DMG{}
    my_dmg.Mount(*dmgPath)
    macosutils.InstallApp(my_dmg.MountPoint)
    my_dmg.Unmount(my_dmg.MountPoint)
}
