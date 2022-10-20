package main

import (
    "os"
)

func main() {
    f, _ := os.Create("test.png")
    defer f.Close()

    maze := NewMaze(1000, 1000, 5)
    maze.Png(f)   
}
