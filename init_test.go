package gridsearch

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	var g = InitGrid([]float64{1, 2, 3}, []float64{1, 3, 2})
	fmt.Println(g)
}
