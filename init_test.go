package gridsearch

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	g := InitGrid([]float64{1, 2, 3}, []float64{1, 3, 2})
	fmt.Println(g)

	var tmp []float64
	var err error

	tmp, err = FromToBy(0, 1, -.2)
	if err != nil {
		fmt.Println(err)
	}

	tmp, err = FromToBy(1, 0, -.2)
	if err != nil {
		fmt.Println(err)
	}
	g.Append(tmp)

	tmp, err = FromToLen(1, 0, 0)
	if err != nil {
		fmt.Println(err)
	}

	tmp, err = FromToLen(1, 0, 5)
	if err != nil {
		fmt.Println(err)
	}
	g.Append(tmp)

	g.SetNumGoRoutines(2)
	g.SetZoom(1)
	g.SetDecay(.8)
	g.SetNumReturn(3)

	fmt.Println(g)
	fmt.Println(g.Dim())
}
