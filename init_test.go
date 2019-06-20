package gridsearch

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	fmt.Println("=======================")
	fmt.Println("Results from TestInit")
	g := InitGrid([]float64{1, 2, 3}, []float64{1, 3, 2})
	fmt.Println(g)

	g.Append(FromToBy(1, 0, -.2))

	g.Append(FromToLen(1, 0, 5))

	g.SetNumGoRoutines(2)
	g.SetZoom(1)
	g.SetDecay(.8)
	g.SetNumReturn(3)

	fmt.Println(g)
	fmt.Println(g.Dim())
}
