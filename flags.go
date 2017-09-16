package main

import (
	"flag"
	"fmt"
)

var (
	booonOption   = flag.Bool("booon", false, "booon (required japanese monospaced font.  ja:en = 2:1)")
	manyOption    = flag.Bool("many", false, "many shobon")
	reverseOption = flag.Bool("reverse", false, "reverse")
	jumpOption    = flag.Bool("jump", false, "jump! (ignore other options)")
)

func parseFlags() {
	flag.Parse()
	fmt.Println(manyOption)
	fmt.Println(reverseOption)
	fmt.Println(jumpOption)
}
