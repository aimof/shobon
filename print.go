package main

import (
	"fmt"
	"github.com/aimof/shobon/kaomoji"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func printBooon(height, width int) {
	if width < 2*kaomoji.BOOON_WIDTH {
		fmt.Println("Need more space to fly!")
		return
	}
	line := []rune(" " + kaomoji.BOOON + strings.Repeat(" ", width-kaomoji.BOOON_WIDTH-1))

	for loop := 0; loop < 3; loop++ {
		for i := 0; i < len(line); i++ {
			fmt.Print("\x1b[1J")
			fmt.Print("\x1b[" + strconv.Itoa(height/2+1) + ";1H")
			if i == 0 {
				fmt.Print(string(line) + "\n")
			} else {
				fmt.Print(string(line[len(line)-i:]) + string(line[0:len(line)-i]))
			}
			fmt.Print("\x1b[" + strconv.Itoa(height-1) + ";1H")
			time.Sleep(25 * time.Millisecond)
		}
	}
}

func printJumpingShobon(height, width int) {
	if height < 2*kaomoji.SHOBON_HEIGHT {
		fmt.Println("sorry, your terminal is too small to jump.")
		return
	}
	var x int = 0
	if width > kaomoji.SHOBON_WIDTH {
		x = (width - kaomoji.SHOBON_WIDTH) / 2
	}
	if height < 8 {
		printShobon(x, 0, 8)
	} else {
		for jump := 0; jump < 10; jump++ {
			highest := float64(height - kaomoji.SHOBON_HEIGHT)
			for i := height - kaomoji.SHOBON_HEIGHT - 1; i >= 0; i-- {
				var shobonHeight float64 = float64(height - kaomoji.SHOBON_HEIGHT - i)
				printShobon(x, i, height)
				time.Sleep(time.Duration(math.Pow((shobonHeight-highest/2)/(highest/2), 2) * float64(40*time.Millisecond)))
			}
			for i := 0; i < height-kaomoji.SHOBON_HEIGHT-1; i++ {
				var shobonHeight float64 = float64(height - kaomoji.SHOBON_HEIGHT - i)
				printShobon(x, i, height)
				time.Sleep(time.Duration(math.Pow((shobonHeight-highest/2)/(highest/2), 2) * float64(40*time.Millisecond)))
			}
		}
	}
}

func printManyShobon(height, width int) {
	rand.Seed(time.Now().UnixNano())
	var indents = make([]string, 0, height)
	for i := 0; i < height-1; i++ {
		fmt.Print("\x1b[1J")
		fmt.Print("\x1b[1;1H")
		if width > kaomoji.SHOBON_ORIGINAL_WIDTH+2 {
			x := rand.Intn(width - kaomoji.SHOBON_ORIGINAL_WIDTH - 1)
			indents = append(indents, strings.Repeat(" ", x+1))
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

func printTooManyShobon(height, width int) {
	if width < kaomoji.SHOBON_ORIGINAL_WIDTH*5 {
		fmt.Println("This terminal is too narrow to gether.")
		fmt.Println(strings.Repeat(kaomoji.SHOBON_ORIGINAL, width))
		return
	}
	rand.Seed(time.Now().UnixNano())
	var lines []string = make([]string, 0, height)
	for i := 0; i < height-2; i++ {
		fmt.Print("\x1b[1J")
		fmt.Print("\x1b[1;1H")
		if i%2 == 0 {
			space0 := rand.Intn(width/3 - kaomoji.SHOBON_ORIGINAL_WIDTH + 1)
			space1 := rand.Intn(width/3-kaomoji.SHOBON_ORIGINAL_WIDTH+1) + width/3 - space0 - kaomoji.SHOBON_ORIGINAL_WIDTH
			space2 := rand.Intn(width/3-kaomoji.SHOBON_ORIGINAL_WIDTH+1) + width*2/3 - space0 - space1 - kaomoji.SHOBON_ORIGINAL_WIDTH*2
			indents := []string{
				strings.Repeat(" ", space0),
				strings.Repeat(" ", space1),
				strings.Repeat(" ", space2),
			}
			var line string = strings.Repeat(" ", (width%3)/2)
			for _, indent := range indents {
				line = line + indent + kaomoji.SHOBON_ORIGINAL
			}
			line = line + "\n"
			lines = append(lines, line)
		} else {
			space0 := rand.Intn(width/4 - kaomoji.SHOBON_ORIGINAL_WIDTH + 1)
			space1 := rand.Intn(width/4-kaomoji.SHOBON_ORIGINAL_WIDTH+1) + width/4 - space0 - kaomoji.SHOBON_ORIGINAL_WIDTH
			space2 := rand.Intn(width/4-kaomoji.SHOBON_ORIGINAL_WIDTH+1) + width/2 - space0 - space1 - kaomoji.SHOBON_ORIGINAL_WIDTH*2
			space3 := rand.Intn(width/4-kaomoji.SHOBON_ORIGINAL_WIDTH+1) + width*3/4 - space0 - space1 - space2 - kaomoji.SHOBON_ORIGINAL_WIDTH*3
			indents := []string{
				strings.Repeat(" ", space0),
				strings.Repeat(" ", space1),
				strings.Repeat(" ", space2),
				strings.Repeat(" ", space3),
			}
			var line string = strings.Repeat(" ", (width%4)/2)
			for _, indent := range indents {
				line = line + indent + kaomoji.SHOBON_ORIGINAL
			}
			line = line + "\n"
			lines = append(lines, line)
		}

		if *reverseOption {
			fmt.Print(strings.Repeat("\n", height-i-3))
			for j := range lines {
				fmt.Print(lines[j])
			}
		} else {
			for j := range lines {
				fmt.Print(lines[i-j])
			}
			fmt.Print(strings.Repeat("\n", height-i-3))
		}
		time.Sleep(100 * time.Millisecond)
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
