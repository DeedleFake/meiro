// meiro is a package for the creation and manipulation of mazes.
package meiro

import (
	"github.com/DeedleFake/meiro/shuffle"
	"sort"
	"strconv"
)

// A Maze is a rectangular maze.
type Maze struct {
	c    []Cell
	w, h int
}

// Random returns a randomized maze of the given width and height. All
// of the outer walls of the maze are enabled; in other words, there
// are no entrances or exits.
//
// TODO: Make it possible to modify the maze more directly from other
// packages so that clients can, among other things, add entrances and
// exits.
func Random(w, h int) *Maze {
	m := &Maze{
		c: make([]Cell, w*h),

		w: w,
		h: h,
	}

	for i := range m.c {
		m.c[i].set = i
	}

	m.randomize()

	return m
}

func (m Maze) neighbor(i, p int) (n *Cell, b int) {
	switch p {
	case up:
		if i < m.w {
			return nil, -1
		}
		return &m.c[i-m.w], down

	case down:
		if i >= len(m.c)-m.w {
			return nil, -1
		}
		return &m.c[i+m.w], up

	case left:
		if i%m.w == 0 {
			return nil, -1
		}
		return &m.c[i-1], right

	case right:
		if (i+1)%m.w == 0 {
			return nil, -1
		}
		return &m.c[i+1], left
	}

	panic("Unexpected path value: " + strconv.FormatInt(int64(p), 10))
}

// randomize randomizes the maze using Kruskal's method. It is only
// designed to work on a maze with all of the walls enabled.
func (m *Maze) randomize() {
	indexes := make(sort.IntSlice, len(m.c))
	for i := range indexes {
		indexes[i] = i
	}
	shuffle.Shuffle(indexes)

	paths := sort.IntSlice{up, down, left, right}

	for _, i := range indexes {
		c := &m.c[i]

		shuffle.Shuffle(paths)

		for _, p := range paths {
			n, b := m.neighbor(i, p)
			if (n == nil) || (c.set == n.set) {
				continue
			}

			c.paths[p] = n
			n.paths[b] = c

			// TODO: Replace with graph traversal.
			old := n.set
			for i := range m.c {
				if m.c[i].set == old {
					m.c[i].set = c.set
				}
			}
		}
	}
}

// At returns the cell at (x, y).
func (m Maze) At(x, y int) *Cell {
	return &m.c[y*m.w+x]
}

// Width returns the width of the maze.
func (m Maze) Width() int {
	return m.w
}

// Height returns the height of the maze.
func (m Maze) Height() int {
	return m.h
}

const (
	up = iota
	down
	left
	right
)

// A Cell represents a single room of the maze.
type Cell struct {
	paths [4]*Cell
	set   int
}

// Up returns the cell just north of c.
func (c Cell) Up() *Cell {
	return c.paths[up]
}

// Down returns the cell just south of c.
func (c Cell) Down() *Cell {
	return c.paths[down]
}

// Left returns the cell just west of c.
func (c Cell) Left() *Cell {
	return c.paths[left]
}

// Right returns the cell just east of c.
func (c Cell) Right() *Cell {
	return c.paths[right]
}
