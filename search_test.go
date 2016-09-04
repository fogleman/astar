package astar

import (
	"fmt"
	"log"
	"math"
	"net/http"
	_ "net/http/pprof"
	"testing"
)

func TestSearch(t *testing.T) {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	g := NewGrid(11, 11)
	for y := 0; y < 9; y++ {
		g.SetWall(3, y, true)
		g.SetWall(7, y, true)
	}
	fmt.Println(g.Search(0, 10))
	for i := 0; i < 10000; i++ {
		g.Search(0, 10)
	}
	fmt.Println(g.Search(0, 10))
}

type Grid struct {
	W, H  int
	Walls []bool
}

func NewGrid(w, h int) *Grid {
	walls := make([]bool, w*h)
	return &Grid{w, h, walls}
}

func (g *Grid) SetWall(x, y int, wall bool) {
	g.Walls[y*g.W+x] = wall
}

func (g *Grid) Edges(node int) []Edge {
	if g.Walls[node] {
		return nil
	}
	edges := make([]Edge, 0, 8)
	x := node % g.W
	y := node / g.W
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx := x + dx
			ny := y + dy
			if nx < 0 || ny < 0 || nx >= g.W || ny >= g.H {
				continue
			}
			index := ny*g.W + nx
			if g.Walls[index] {
				continue
			}
			distance := 1.0
			if dx != 0 && dy != 0 {
				distance = math.Sqrt2
			}
			edge := Edge{index, distance}
			edges = append(edges, edge)
		}
	}
	return edges
}

func (g *Grid) Estimate(src, dst int) float64 {
	x1 := src % g.W
	y1 := src / g.W
	x2 := dst % g.W
	y2 := dst / g.W
	dx := x2 - x1
	dy := y2 - y1
	return math.Sqrt(float64(dx*dx + dy*dy))
}

func (g *Grid) Search(src, dst int) Result {
	return Search(g, src, dst)
}
