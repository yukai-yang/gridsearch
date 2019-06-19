package gridsearch

import (
	"fmt"
	"math"
	"sort"
	"sync"
)

//utilities including unexported functions and errors

//GridError gives the error message in the package
type GridError struct {
	msg string
}

func (e *GridError) Error() string {
	return e.msg
}

func (g Grid) String() string {
	var grid = ""
	for _, p := range g.base {
		grid = grid + fmt.Sprintf("\n")
		for _, v := range p {
			grid = grid + fmt.Sprintf("%v ", v)
		}
	}

	return "returned values: \t\t" + fmt.Sprintf("%v", g.numReturn) +
		"\ngo routines: \t\t\t" + fmt.Sprintf("%v", g.numGoRoutines) +
		"\nlayers of the zoom-in: \t\t" + fmt.Sprintf("%v", g.zoom) +
		"\ndecay rate of the grid sizes: \t" + fmt.Sprintf("%v", g.decay) +
		"\nThe grid base is as follows:" + grid

}

func buildGrid() {}

//sequentially evaluate the function values
//and store them in the input argument "values"
func evalFunc(target func([]float64) float64, points [][]float64, values []float64) {
	for i, p := range points {
		values[i] = target(p)
	}
}

//evaluate the function values by using go routines
//will call evalFunc inside for each go routine
func goEvalFunc(target func([]float64) float64, points [][]float64, numGo int) []float64 {
	var length = len(points)
	var values = make([]float64, length)
	var size = int(math.Ceil(float64(len(points)) / float64(numGo)))
	var wg sync.WaitGroup

	wg.Add(numGo)
	for i := 0; i < numGo; i++ {
		go func(ii int) {
			from := ii * size
			to := (ii + 1) * size
			if to >= length {
				to = length - 1
			}
			evalFunc(target, points[from:to], values[from:to])
			wg.Done()
		}(i)
	}

	wg.Wait()
	return values
}

//find the first num smallest values in "values"
//"values" will be changed
func findMins(values []float64, num int) []int {
	var tmp = make(map[float64]int)
	var ret = make([]int, num)

	for i := 0; i < len(values); i++ {
		tmp[values[i]] = i
	}

	//change values here...
	sort.Float64s(values)

	for i := 0; i < num; i++ {
		ret[i] = tmp[values[i]]
	}

	return ret
}

func expandGrid(base [][]float64) [][]float64 {
	var grid [][]float64
	return grid
}

//recursively grid search
func recursiveSearch(target func([]float64) float64, base [][]float64, numGo int, zoom int, decay float64, num int) [][]float64 {
	var ret = make([][]float64, num)

	var values = goEvalFunc(target, base, numGo)
	// values has been changed!
	var tmp = findMins(values, num)

	for i := 0; i < num; i++ {
		ret[i] = base[tmp[i]]
	}

	return ret
}
