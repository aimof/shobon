package main

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	parseFlags()
	height, width, err := measureSize()
	if err != nil {
		log.Println(err)
	}
	switch {
	case *booonOption:
		printBooon(height, width)
	case *jumpOption:
		printJumpingShobon(height, width)
	case *manyOption:
		printManyShobon(height, width)
	case *tooManyOption:
		printTooManyShobon(height, width)
	default:
		printDefaultShobon(height, width)
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
