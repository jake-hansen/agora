package main

import (
	"flag"
	"fmt"
	"github.com/jake-hansen/agora/config"
	"os"
)

func main() {
	var environment *string = flag.String("e", "dev", "environment to run in")
	flag.Usage = func() {
		fmt.Println("Usage: serve -e {environment}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
}
