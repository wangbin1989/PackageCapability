// 程序需要CAP: CAP_SETFCAP CAP_DAC_OVERRIDE

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		usage()
		return
	}

	packageName := os.Args[1]
	relPath := os.Args[2]
	capabilities := os.Args[3]

	log.Printf("package_name: %s, relpath: %s, capabilities: %s.", packageName, relPath, capabilities)

	UpdateCapabilities(packageName, relPath, capabilities)
}

func usage() {
	fmt.Println("Usage: pkgcap <package_name> <relpath> <capability1>,<capability2>,...")
}
