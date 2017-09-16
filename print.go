package main

import (
	"fmt"
	"github.com/aimof/shobon/kaomoji"
	"math/rand"
	"strings"
	"time"
)

func printManyShobon(height, width int) {
	rand.Seed(time.Now().UnixNano())
	var indents = make([]string, 0, height)
	for i := 0; i < height-1; i++ {
		fmt.Print("\x1b[1J")
		fmt.Print("\x1b[1;1H")
		if width > len(kaomoji.SHOBON_ORIGINAL)+2 {
			x := rand.Intn(width - len(kaomoji.SHOBON_ORIGINAL) - 2)
			indents = append(indents, strings.Repeat(" ", x))
		}
		if *reverseOption {
			fmt.Print(strings.Repeat("\n", height-i-2))
			for j := range indents {
				fmt.Print(indents[j] + kaomoji.SHOBON_ORIGINAL + "\n")
			}
		} else {
			for j := range indents {
				fmt.Print(indents[i-j] + kaomoji.SHOBON_ORIGINAL + "\n")
			}
			fmt.Print(strings.Repeat("\n", height-i-2))
		}

		time.Sleep(time.Millisecond * 150)
	}

}

func printDefaultShobon(height, width int) {
	var x int = 0
	if width > kaomoji.SHOBON_WIDTH {
		x = (width - kaomoji.SHOBON_WIDTH) / 2
	}
	if height < 8 {
		printShobon(x, 0, 8)
	} else {
		for i := 0; i < height-kaomoji.SHOBON_HEIGHT-1; i++ {
			if *reverseOption {
				printShobon(x, height-kaomoji.SHOBON_HEIGHT-i-2, height)
			} else {
				printShobon(x, i, height)
			}
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