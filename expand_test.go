package gridsearch

import (
	"fmt"
	"testing"
)

func TestExpand(t *testing.T) {
	g := InitGrid([]float64{1, 2, 3}, []float64{6, 4, 5})

	tmp := expandGrid(g.base)

	fmt.Println(tmp)
}
