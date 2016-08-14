// Go app to install a list of OS X Apps, stored in json format

package main

import (
	"flag"
	"fmt"
)

func main() {
	var applist string
	flag.StringVar(&applist, "a", "", "specify applist to use. ")
	flag.Parse()

	fmt.Printf("applist = %s", applist)
}
