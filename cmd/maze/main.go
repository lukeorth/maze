package main

import (
	"flag"
	"log"
	"os"

	"github.com/lukeorth/maze"
)

var (
    width int
    height int
    scale int
    out string
)

func init() {
    flag.IntVar(&width, "w", 10, "maze width")
    flag.IntVar(&height, "h", 10, "maze height")
    flag.IntVar(&scale, "s", 1, "maze scale (number of pixels per cell)")
    flag.StringVar(&out, "o", "maze.png", "output image path")
}

func main() {
    // gracefully handle errors and exit
    var err error
    defer func() {
        if err != nil {
            log.Fatalln(err)
        }
    }()

    flag.Parse()

    m := maze.NewMaze(width, height)

    f, _ := os.Create(out)
    defer f.Close()

    m.Png(f, scale)
}
