package main

import (
	"math/rand"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type cell struct {
	drawable uint32

	alive     bool
	aliveNext bool
}

func makeCells() [][]*cell {
	rand.Seed(time.Now().UnixNano())
	intcolumns := int(columns)
	introws := int(rows)

	cells := make([][]*cell, introws)
	for x := 0; x < introws; x++ {
		cells[x] = make([]*cell, intcolumns)
		for y := 0; y < intcolumns; y++ {
			c := newCell(x, y)

			c.alive = rand.Float64() < threshold
			c.aliveNext = c.alive

			cells[x][y] = c
		}
	}

	return cells
}

func newCell(x, y int) *cell {
	points := make([]float32, len(square))
	copy(points, square)

	sizex := 2.0 / rows
	sizey := 2.0 / columns
	positionx := (float32(x) * sizex) - 1
	positiony := (float32(y) * sizey) - 1

	for ix := 0; ix < len(points); ix += 3 {
		points[ix] = points[ix]*sizex + positionx
		points[ix+1] = points[ix+1]*sizey + positiony
	}

	return &cell{
		drawable: makeVao(points),
	}
}

func (c *cell) draw() {
	if !c.alive {
		return
	}

	gl.BindVertexArray(c.drawable)
	gl.DrawArrays(gl.TRIANGLE_FAN, 0, 4)
}

// checkState determines the state of the cell for the next tick of the game.
func (c *cell) checkState(cells [][]*cell, x int, y int) {
	c.alive = c.aliveNext
	c.aliveNext = c.alive

	liveCount := liveNeighbors(cells, x, y)
	if c.alive {

		if liveCount < 2 {
			// 1. Any live cell with fewer than two live neighbours dies, as if caused by underpopulation.
			c.aliveNext = false

		} else if liveCount > 3 {
			// 2. Any live cell with more than three live neighbours dies, as if by overpopulation.
			c.aliveNext = false

		} else {
			// 3. Any live cell with two or three live neighbours lives on to the next generation.
			c.aliveNext = true
		}

	} else {
		// 4. Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
		if liveCount == 3 {
			c.aliveNext = true
		}
	}
}

// liveNeighbors returns the number of live neighbors for a cell.
func liveNeighbors(cells [][]*cell, cx int, cy int) int {
	var liveCount int
	add := func(x, y int) {
		// If we're at an edge, check the other side of the board.
		if x == len(cells) {
			x = 0
		} else if x == -1 {
			x = len(cells) - 1
		}
		if y == len(cells[x]) {
			y = 0
		} else if y == -1 {
			y = len(cells[x]) - 1
		}

		if cells[x][y].alive {
			liveCount++
		}
	}

	add(cx-1, cy)   // To the left
	add(cx+1, cy)   // To the right
	add(cx, cy+1)   // up
	add(cx, cy-1)   // down
	add(cx-1, cy+1) // top-left
	add(cx+1, cy+1) // top-right
	add(cx-1, cy-1) // bottom-left
	add(cx+1, cy-1) // bottom-right

	return liveCount
}
