package main

import (
	"fmt"
	"github.com/aimof/shobon/kaomoji"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
	parseFlags()
	height, width, err := measureSize()
	if err != nil {
		log.Println(err)
	}
	switch {
	case *manyOption:
		printManyShobon(height, width)
	default:
		printDefaultShobon(height, width)
	}
}

func printManyShobon(height, width int) {
	rand.Seed(time.Now().UnixNano())
	var indents = make([]string, 0, height)
	for i := 0; i < height; i++ {
		if width > len(kaomoji.SHOBON_ORIGINAL)+2 {
			x := rand.Intn(width - len(kaomoji.SHOBON_ORIGINAL) - 2)
			indents = append(indents, strings.Repeat(" ", x))
		}
		if *reverseOption {
			fmt.Print(strings.Repeat("\n", height-i-1))
			for j := range indents {
				fmt.Print(indents[j] + kaomoji.SHOBON_ORIGINAL + "\n")
			}
		} else {
			for j := range indents {
				fmt.Print(indents[i-j] + kaomoji.SHOBON_ORIGINAL + "\n")
			}
			fmt.Print(strings.Repeat("\n", height-i-1))
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
		for i := 0; i < height-kaomoji.SHOBON_HEIGHT; i++ {
			if *reverseOption {
				printShobon(x, height-kaomoji.SHOBON_HEIGHT-i-1, height)
			} else {
				printShobon(x, i, height)
			}
			time.Sleep(150 * time.Millisecond)
		}
	}
}

func measureSize() (height, width int, err error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	rawSize, err := cmd.Output()
	if err != nil {
		return 0, 0, err
	}
	size := strings.Split(string(rawSize), " ")
	height, err = strconv.Atoi(size[0])
	if err != nil {
		return 0, 0, err
	}
	width, err = strconv.Atoi(strings.Split(size[1], "\n")[0])
	if err != nil {
		return 0, 0, err
	}
	return height, width, nil
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
