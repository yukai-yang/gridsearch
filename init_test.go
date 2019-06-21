package gridsearch

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	fmt.Println("=======================")
	fmt.Println("Test results from TestInit")
	fmt.Println("-----------------------")
	fmt.Println("//Go code:")
	fmt.Println("g := InitGrid([]float64{1, 2, 3}, []float64{1, 3, 2})")
	fmt.Println("fmt.Println(g)")
	fmt.Println("-----------------------")
	fmt.Println("//results:")
	g := InitGrid([]float64{1, 2, 3}, []float64{1, 3, 2})
	fmt.Println(g)

	fmt.Println("-----------------------")
	fmt.Println("Go code:")
	fmt.Println("g.Append(FromToBy(1, 0, -.2))")
	fmt.Println("g.Append(FromToLen(1, 0, 5))")
	fmt.Println("g.SetNumGoRoutines(2)")
	fmt.Println("g.SetZoom(1)")
	fmt.Println("g.SetDecay(.8)")
	fmt.Println("g.SetNumReturn(3)")
	fmt.Println("fmt.Println(g)")
	fmt.Println("g.Append(FromToLen(1, 0, 5))")
	fmt.Println("-----------------------")
	fmt.Println("results:")
	g.Append(FromToBy(1, 0, -.2))
	g.Append(FromToLen(1, 0, 5))

	g.SetNumGoRoutines(2)
	g.SetZoom(1)
	g.SetDecay(.8)
	g.SetNumReturn(3)

	fmt.Println(g)
	fmt.Println(g.Dim())
}
