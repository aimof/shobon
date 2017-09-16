package main

import (
	"fmt"
	"github.com/aimof/shobon/kaomoji"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	rawSize, err := cmd.Output()
	if err != nil {
		log.Fatalln(err)
	}
	size := strings.Split(string(rawSize), " ")
	height, err := strconv.Atoi(size[0])
	if err != nil {
		log.Fatalln(err)
	}
	width, err := strconv.Atoi(strings.Split(size[1], "\n")[0])
	if err != nil {
		log.Fatalln(err)
	}

	var x int
	if width > 74 {
		x = (width - 74) / 2
	}
	if height < 8 {
		printShobon(x, 0, 8)
	} else {
		for i := 0; i < height-kaomoji.SHOBON_HEIGHT; i++ {
			printShobon(x, i, height)
			time.Sleep(150 * time.Millisecond)
		}
	}
}

func printShobon(x, y, height int) {
	var indent string = strings.Repeat(" ", x)
	fmt.Print("\x1b[1J")
	fmt.Print("\x1b[1;1H")
	fmt.Print(strings.Repeat("\n", y))

	fmt.Print(indent + kaomoji.SHOBON0 + "\n")
	fmt.Print(indent + kaomoji.SHOBON1 + "\n")
	fmt.Print(indent + kaomoji.SHOBON2 + "\n")
	fmt.Print(indent + kaomoji.SHOBON3 + "\n")
	fmt.Print(indent + kaomoji.SHOBON4 + "\n")
	fmt.Print(indent + kaomoji.SHOBON5 + "\n")
	fmt.Print(indent + kaomoji.SHOBON6 + "\n")
	fmt.Print(indent + kaomoji.SHOBON7 + "\n")
	fmt.Print(strings.Repeat("\n", height-y-kaomoji.SHOBON_HEIGHT-1))
}
