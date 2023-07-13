package main

import (
	"os"

	"github.com/lukeorth/maze"
)

func main() {
    m := maze.NewMaze(1000, 1000)

    f, _ := os.Create("test.png")
    defer f.Close()

    m.Png(f, 5)
}
