package meiro

import (
	"math/rand"
	"sort"
	"strconv"
)

func Shuffle(data sort.Interface) {
	length := data.Len()

	for i := 0; i < length; i++ {
		i2 := rand.Intn(i + 1)
		data.Swap(i, i2)
	}
}

type Maze struct {
	c    []Cell
	w, h int
}

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

func (m *Maze) randomize() {
	indexes := make(sort.IntSlice, len(m.c))
	for i := range indexes {
		indexes[i] = i
	}
	Shuffle(indexes)

	paths := sort.IntSlice{up, down, left, right}

	for _, i := range indexes {
		c := &m.c[i]

		Shuffle(paths)

		for _, p := range paths {
			n, b := m.neighbor(i, p)
			if (n == nil) || (c.set == n.set) {
				continue
			}

			c.paths[p] = n
			n.paths[b] = c

			old := n.set
			for i := range m.c {
				if m.c[i].set == old {
					m.c[i].set = c.set
				}
			}
		}
	}
}

func (m Maze) At(x, y int) *Cell {
	return &m.c[y*m.w+x]
}

func (m Maze) Width() int {
	return m.w
}

func (m Maze) Height() int {
	return m.h
}

const (
	up = iota
	down
	left
	right
)

type Cell struct {
	paths [4]*Cell

	set int
}

func (c Cell) Up() *Cell {
	return c.paths[up]
}

func (c Cell) Down() *Cell {
	return c.paths[down]
}

func (c Cell) Left() *Cell {
	return c.paths[left]
}

func (c Cell) Right() *Cell {
	return c.paths[right]
}
