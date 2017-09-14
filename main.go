package main

import (
	"fmt"
	"github.com/aimof/shobon/kaomoji"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
	var out io.Writer = os.Stdout
	fmt.Fprint(out, "\x1b[2J")
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
	if height < 8 {
		fmt.Fprintln(out, "\x1b[2J")
		fmt.Print(kaomoji.SHOBON0 + "\n")
		fmt.Print(kaomoji.SHOBON1 + "\n")
		fmt.Print(kaomoji.SHOBON2 + "\n")
		fmt.Print(kaomoji.SHOBON3 + "\n")
		fmt.Print(kaomoji.SHOBON4 + "\n")
		fmt.Print(kaomoji.SHOBON5 + "\n")
		fmt.Print(kaomoji.SHOBON6 + "\n")
		fmt.Print(kaomoji.SHOBON7 + "\n")
	} else {
		for i := 0; i < height-8; i++ {
			fmt.Fprintln(out, "\x1b[2J")
			fmt.Print(strings.Repeat("\n", i))
			whiteSpace := ""
			if width > 74 {
				whiteSpace = strings.Repeat(" ", (width-74)/2)
			}
			fmt.Print(whiteSpace + kaomoji.SHOBON0 + "\n")
			fmt.Print(whiteSpace + kaomoji.SHOBON1 + "\n")
			fmt.Print(whiteSpace + kaomoji.SHOBON2 + "\n")
			fmt.Print(whiteSpace + kaomoji.SHOBON3 + "\n")
			fmt.Print(whiteSpace + kaomoji.SHOBON4 + "\n")
			fmt.Print(whiteSpace + kaomoji.SHOBON5 + "\n")
			fmt.Print(whiteSpace + kaomoji.SHOBON6 + "\n")
			fmt.Print(whiteSpace + kaomoji.SHOBON7 + "\n")
			fmt.Print(strings.Repeat("\n", height-i-kaomoji.SHOBON_HEIGHT))
			time.Sleep(150 * time.Millisecond)
		}
	}
}
