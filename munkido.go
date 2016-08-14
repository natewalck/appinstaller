// Go app to install a list of OS X Apps, stored in json format

package main

import (
	"fmt"
	"github.com/natewalck/appinstaller"
	"os"
)

func main() {
	// Program Name is always the first (implicit) argument
	cmd := os.Args[0]

	fmt.Printf("Program Name: %s\n", cmd)
}
