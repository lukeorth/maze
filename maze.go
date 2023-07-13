package maze

import (
	"math/rand"
	"time"
)

type maze struct {
    cells   []*cell
    cols    int
    rows    int
}

type cell struct {
    x       int
    y       int
    walls   uint
    visited bool
    current bool
}

func NewMaze(cols int, rows int) *maze {
    // init maze
    m := &maze{
        cells: make([]*cell, cols * rows),
        cols: cols,
        rows: rows,
    }

    // init cells
    for y := 0; y < rows; y++ {
        for x := 0; x < cols; x++ {
            m.cells[cols * y + x] = &cell{x, y, 15, false, false}
        }
    }

    m.backTrack(0, 0, make([]*cell, 0))

    return m
}

func (m *maze) backTrack(x int, y int, seen []*cell) *cell {
    c := m.cellAt(x, y)

    neighbors := []*cell{ 
        m.cellAt(x, y - 1),
        m.cellAt(x + 1, y), 
        m.cellAt(x, y + 1), 
        m.cellAt(x - 1, y),
    }
    rand.Seed(time.Now().UnixNano())
    random := rand.Intn(4)
    for i := range neighbors {
        n := neighbors[(random + i) % 4]
        if n != nil && !n.visited {
            n.visited = true
            n.current = true
           
            c.removeWall(n)
            seen = append(seen, c)
            m.backTrack(n.x, n.y, seen)
            return n
        }
    }
    if len(seen) > 0 {
        c := seen[len(seen)-1]
        seen = seen[:len(seen)-1]
        m.backTrack(c.x, c.y, seen)
    }
    return nil
}

func (m *maze) cellAt(x int, y int) *cell {
    // return nil if index is invalid
    if x < 0 || y < 0 || x > m.cols - 1 || y > m.rows - 1 {
        return nil
    }
    return m.cells[m.cols * y + x]
}

func (c *cell) direction(n *cell) uint {
    // direction for getting from Cell(c) to Cell(n)
    if c.x > n.x {          // left
        return 1
    } else if c.x < n.x {   // right
        return 4
    } else if c.y > n.y {   // top
        return 8
    } else {                // bottom
        return 2
    }
}

func (c *cell) removeWall(n *cell) {
    // remove wall between Cell(c) and Cell(n)
    dir := c.direction(n)
    c.walls = c.walls & ^dir
    n.walls = n.walls & ^(dir >> 2 | dir << 2)
}
