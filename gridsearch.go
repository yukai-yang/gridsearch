package gridsearch

import (
	"sort"
)

//defining the main structures and interfaces

//Grid will be exported
//but not its member
type Grid struct {
	base          [][]float64
	numGoRoutines int
	zoom          int
	decay         float64
	numReturn     int
}

//GridSearcher is an interface defining the functions
type GridSearcher interface {
	Dim() int
	Append(...[]float64)
	SetNumGoRoutines(int) error
	SetZoom(int) error
	SetDecay(float64) error
	SetNumReturn(int) error
	Search(func([]float64) float64) [][]float64
}

//Dim gets the number of parameters or arguments
func (g Grid) Dim() int {
	return len(g.base)
}

//Append appends subgrids to the grid g
func (g *Grid) Append(points ...[]float64) {
	var tmp = make([][]float64, len(points))
	for i, p := range points {
		tmp[i] = make([]float64, len(p))
		copy(tmp[i], p)
		sort.Float64s(tmp[i])
	}
	g.base = append(g.base, tmp...)
}

//SetNumGoRoutines sets the number of go routines
func (g *Grid) SetNumGoRoutines(num int) error {
	if num <= 0 {
		return &GridError{"The number of go routines is not positive."}
	}
	g.numGoRoutines = num
	return nil
}

//SetZoom sets the zoom
//number of (additional) rounds or layers of the zoom-in
func (g *Grid) SetZoom(zoom int) error {
	if zoom < 0 {
		return &GridError{"The zoom layer is not appropriate."}
	}
	g.zoom = zoom
	return nil
}

//SetDecay sets the decay
//representing the decay rate of the grid sizes of the zoom
func (g *Grid) SetDecay(decay float64) error {
	if decay < 0 || decay > 1 {
		return &GridError{"The decay is not in between 0 and 1."}
	}
	g.decay = decay
	return nil
}

//SetNumReturn sets the number of points to return
//i.e. the smallest points, 1 by default the minimum.
func (g *Grid) SetNumReturn(num int) error {
	if num < 1 {
		return &GridError{"The number of points to return is not appropriate."}
	}
	g.numReturn = num
	return nil
}

//Search implements the grid search algorithm
func (g Grid) Search(target func([]float64) float64) ([][]float64, []float64) {
	return recursiveSearch(target, g.base, g.numGoRoutines, g.zoom, g.decay, g.numReturn)
}
