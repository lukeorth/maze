package maze

import (
	"image"
    "image/color"
	"image/png"
	"io"
	"math"
)

func (m *maze) Png(w io.Writer, scale int) {
    // dimensions
    width := m.cols * (2 * scale) + scale
    height := m.rows * (2 * scale) + scale

    // setup image
    img := image.NewRGBA(image.Rect(0, 0, width, height))

    // draw cells
    for _, c := range m.cells {
        c.drawPNG(img, scale)
    }
    png.Encode(w, img)
}

func (c *cell) drawPNG(img *image.RGBA, scale int) {
    // weight of cell wall in pixels
    weight := int(math.Ceil(float64(scale) / 4))

    // cell corners (including walls)
    x1 := c.x * (scale + weight)
    x2 := x1 + (2 * weight) + scale - 1
    y1 := c.y * (scale + weight)
    y2 := y1 + (2 * weight) + scale - 1

    black := color.RGBA{0, 0, 0, 255}
    white := color.RGBA{255, 255, 255, 255}

    for x := x1; x <= x2; x++ {
        for y := y1; y <= y2; y++ {
            // initialize all pixels to black
            img.Set(x, y, black)
            // set pixels to white (open) where needed
            if x > x1 && x < x2 {
                // cell body
                if y > y1 && y < y2 {
                    img.Set(x, y, white)
                }
                // top wall
                if y < y1 + weight && c.walls & 8 == 0 {
                    img.Set(x, y, white)
                }
                // bottom wall
                if y > y2 - weight && c.walls & 2 == 0 {
                    img.Set(x, y, white)
                }
            }
            if y > y1 && y < y2 {
                // left wall
                if x < x1 + weight && c.walls & 1 == 0 {
                    img.Set(x, y, white)
                }
                // right wall
                if x > x2 - weight && c.walls & 4 == 0 {
                    img.Set(x, y, white)
                }
            }
        }
    }
}
