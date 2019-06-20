package gridsearch

//functions to be exported

import (
	"math"
	"sort"
)

//InitGrid initializes the grid
func InitGrid(points ...[]float64) Grid {
	var g = Grid{}
	g.base = make([][]float64, len(points))
	for i, p := range points {
		var tmp = make([]float64, len(p))
		copy(tmp, p)
		sort.Float64s(tmp)
		g.base[i] = tmp
	}
	g.numGoRoutines = 1
	g.zoom = 0
	g.decay = .5
	g.numReturn = 1
	return g
}

//FromToBy makes a sequance of floats by some step size
func FromToBy(from, to, by float64) []float64 {
	if (to-from)*by < 0 {
		panic("The sign of 'by' is not appropriate.")
	}

	var length = int(math.Ceil((to-from)/by)) + 1
	var points = make([]float64, length)
	var tmp = from
	for i := 0; i < length; i++ {
		points[i] = tmp
		tmp += by
	}
	return points
}

//FromToLen makes a sequance of floats by length
func FromToLen(from, to float64, length int) []float64 {
	if length <= 0 {
		panic("The 'length' is not appropriate.")
	}
	var points = make([]float64, length)
	var by = (to - from) / float64(length-1)
	var tmp = from
	for i := 0; i < length; i++ {
		points[i] = tmp
		tmp += by
	}
	return points
}
