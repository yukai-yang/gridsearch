package gridsearch

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	fmt.Println("=======================")
	fmt.Println("Restuls from TestSearch")
	g := InitGrid()

	var tmp = FromToBy(-5.12, 5.12, .1)
	g.Append(tmp, tmp)

	g.SetNumGoRoutines(1)
	g.SetNumReturn(3)

	var ret, val = g.Search(Rastrigin)
	fmt.Println(ret, val)
	fmt.Println(Rastrigin(ret[0]))
	//fmt.Println(Rastrigin(ret[1]))

}
