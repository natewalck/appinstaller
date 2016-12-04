// Go app to install a list of OS X Apps

package main

import (
    "flag"
    "github.com/natewalck/macosutils"
    "log"
)

func main() {
    var (
        dmgPath = flag.String("dmg", "", "path to dmg")
    )
    flag.Parse()

    my_dmg := macosutils.DMG{}
    my_dmg.Mount(*dmgPath)
    my_dmg.GetInstallables()
    if len(my_dmg.Pkgs) > 0 {
        for _, p := range my_dmg.Pkgs {
            log.Printf("Installing: %v\n", p)
            macosutils.InstallPkg(my_dmg.MountPoint, p)
        }
    } else if len(my_dmg.Apps) > 0 {
        for _, a := range my_dmg.Apps {
            log.Printf("Installing: %v\n", a)
            macosutils.InstallApp(my_dmg.MountPoint, a)
        }
    } else {
        log.Print("Nothing found to install.")
    }

    my_dmg.Unmount(my_dmg.MountPoint)
    log.Printf("Unmounted dmg...")
}
