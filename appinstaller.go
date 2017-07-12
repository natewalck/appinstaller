// Go app to install a list of OS X Apps

package main

import (
	"flag"
	"log"

	"github.com/natewalck/macosutils"
)

func main() {
	var (
		dmgPath = flag.String("dmg", "", "path to dmg")
	)
	flag.Parse()

	myDmg := macosutils.NewDMG(*dmgPath)
	myDmg.Mount()
	myDmg.GetInstallables()
	if len(myDmg.Pkgs) > 0 {
		for _, p := range myDmg.Pkgs {
			log.Printf("Installing: %v\n", p)
			macosutils.InstallPkg(myDmg.MountPoint, p)
		}
	} else if len(myDmg.Apps) > 0 {
		for _, a := range myDmg.Apps {
			log.Printf("Installing: %v\n", a)
			macosutils.InstallApp(myDmg.MountPoint, a)
		}
	} else {
		log.Print("Nothing found to install.")
	}

	myDmg.Unmount(myDmg.MountPoint)
	log.Printf("Unmounted dmg...")
}
