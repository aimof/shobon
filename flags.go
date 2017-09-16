package main

import (
	"flag"
	"fmt"
)

var (
	manyOption    = flag.Bool("many", false, "many shobon")
	reverseOption = flag.Bool("reverse", false, "reverse")
)

func parseFlags() {
	flag.Parse()
	fmt.Println(manyOption)
	fmt.Println(reverseOption)
}
