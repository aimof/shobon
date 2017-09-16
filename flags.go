package main

import (
	"flag"
	"fmt"
)

var (
	manyOption = flag.Bool("many", false, "many shobon")
)

func parseFlags() {
	flag.Parse()
	fmt.Println(manyOption)
}
