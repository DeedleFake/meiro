package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
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

func RandomMaze(w, h int) *Maze {
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

func (c Cell) Walls() (u, d, l, r bool) {
	u = c.paths[up] == nil
	d = c.paths[down] == nil
	l = c.paths[left] == nil
	r = c.paths[right] == nil

	return
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	m := RandomMaze(5, 5)

	fmt.Printf("|")
	for x := 0; x < m.w; x++ {
		fmt.Printf("--|")
	}
	fmt.Println()

	for y := 1; y < m.h; y++ {
		fmt.Printf("|")
		for x := 0; x < m.w; x++ {
			fmt.Printf("  ")

			_, _, _, r := m.At(x, y).Walls()
			if r {
				fmt.Printf("|")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()

		fmt.Printf("|")
		for x := 0; x < m.w; x++ {
			_, d, _, _ := m.At(x, y).Walls()
			if d {
				fmt.Printf("--|")
			} else {
				fmt.Printf("  |")
			}
		}
		fmt.Println()
	}
}
