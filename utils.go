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
			to := from + size
			if to > length {
				to = length
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
func findMins(values []float64, num int) ([]int, []float64) {
	var tmp = make(map[float64]int)
	var index = make([]int, num)
	var val = make([]float64, num)

	for i := 0; i < len(values); i++ {
		tmp[values[i]] = i
	}

	//change values here...
	sort.Float64s(values)

	for i := 0; i < num; i++ {
		val[i] = values[i]
		index[i] = tmp[values[i]]
	}

	return index, val
}

func indexadd(index []int, pos int, lengths []int) {
	if lengths[pos]-index[pos] > 1 {
		index[pos] = index[pos] + 1
	} else {
		if len(index)-pos > 1 {
			index[pos] = 0
			indexadd(index, pos+1, lengths)
		}
	}
}

func expandGrid(base [][]float64) [][]float64 {
	var width = len(base)
	var lengths = make([]int, width)
	var length = 1
	for i, p := range base {
		lengths[i] = len(p)
		length *= lengths[i]
	}
	var grid = make([][]float64, length, length)
	var index = make([]int, width, width)

	for i := 0; i < length; i++ {
		grid[i] = make([]float64, width)
		for j := 0; j < width; j++ {
			grid[i][j] = base[j][index[j]]
		}
		indexadd(index, 0, lengths)
	}

	return grid
}

func buildSubBase(center []float64, base [][]float64, decay float64) [][]float64 {
	var length = int(math.Ceil(float64(len(base)) * decay))
	var subgrid = make([][]float64, length, length)

	return subgrid
}

//recursively grid search
func recursiveSearch(target func([]float64) float64, base [][]float64, numGo int, zoom int, decay float64, num int) ([][]float64, []float64) {
	var ret = make([][]float64, num)
	var width = len(base)

	var grid = expandGrid(base)
	var values = goEvalFunc(target, grid, numGo)
	// values has been changed!
	var mins, val = findMins(values, num)
	for i := 0; i < num; i++ {
		ret[i] = make([]float64, width, width)
		copy(ret[i], grid[mins[i]])
	}
	//ret and val are ready

	if zoom > 0 {
		var tmpr [][]float64
		var tmpv []float64
		var tmp = make([][]float64, num, num)
		for i := 0; i < num; i++ {
			tmp[i] = make([]float64, width, width)
			copy(tmp[i], ret[i])
		}

		for _, c := range tmp {
			tmpr, tmpv = recursiveSearch(target, buildSubBase(c, base, decay), numGo, zoom-1, decay, num)
			ret = append(ret, tmpr...)
			val = append(val, tmpv...)
		}

		mins, val = findMins(val, num)
		tmp = make([][]float64, num, num)
		for i := 0; i < num; i++ {
			tmp[i] = make([]float64, width, width)
			copy(tmp[i], ret[mins[i]])
		}

		ret = tmp

	}

	return ret, val
}
