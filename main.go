package main

import (
    "os"
)

func main() {
    //f, _ := os.Create("test.png")
    //f, _ := os.Create("test.gif")
    //defer f.Close()

    //maze := NewMaze(50, 50, 5)
    //maze.Png(f)   
    MazeAnimation(os.Stdout, 20, 20, 5)
}
