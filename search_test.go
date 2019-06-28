package gridsearch

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	fmt.Println("=======================")
	fmt.Println("Test restuls from TestSearch")
	fmt.Println("-----------------------")
	fmt.Println("//Go code:")
	fmt.Println("g := InitGrid()")
	fmt.Println("var tmp = FromToBy(-5.12, 5.12, .1)")
	fmt.Println("g.Append(tmp, tmp)")
	fmt.Println("g.SetNumGoRoutines(2)")
	fmt.Println("g.SetZoom(2)")
	fmt.Println("g.SetNumReturn(2)")
	fmt.Println("var ret, val = g.Search(Rastrigin)")
	fmt.Println("fmt.Println(ret, val)")
	fmt.Println("fmt.Println(Rastrigin(ret[0]))")
	fmt.Println("-----------------------")
	fmt.Println("results:")

	g := InitGrid()
	var tmp = FromToBy(-5.12, 5.12, .1)
	g.Append(tmp, tmp)
	g.SetNumGoRoutines(2)
	g.SetZoom(2)
	g.SetNumReturn(2)

	var ret, val = g.Search(Rastrigin)
	fmt.Println(ret, val)
	fmt.Println(Rastrigin(ret[0]))
	//fmt.Println(Rastrigin(ret[1]))
}
