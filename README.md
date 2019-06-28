# gridsearch
implementing the grid search with zoom algorithm by using the Go language

## About the package

### Exported functions

```go
//InitGrid initializes the grid
func InitGrid(points ...[]float64) *Grid
```

```go
//FromToBy makes a sequance of floats by some step size
func FromToBy(from, to, by float64) []float64
```

```go
//FromToLen makes a sequance of floats by length
func FromToLen(from, to float64, length int) []float64
```

```go
//Rastrigin computes the Rastrigin fucntion
func Rastrigin(x []float64) float64
```

### Exported struct and its methods

```go
type Grid struct {
//private stuff
}
```

```go
//Dim gets the number of parameters or arguments
func (g Grid) Dim() int
```

```go
//Append appends subgrids to the grid g
func (g *Grid) Append(points ...[]float64)
```

```go
//SetNumGoRoutines sets the number of go routines
func (g *Grid) SetNumGoRoutines(num int) error
```

```go
//SetZoom sets the zoom
//number of (additional) rounds or layers of the zoom-in
func (g *Grid) SetZoom(zoom int) error
```

```go
//SetDecay sets the decay
//representing the decay rate of the grid sizes of the zoom
func (g *Grid) SetDecay(decay float64) error
```

```go
//SetNumReturn sets the number of points to return
//i.e. the smallest points, 1 by default the minimum.
func (g *Grid) SetNumReturn(num int) error
```

The last and most important one
```go
//Search implements the grid search algorithm
func (g Grid) Search(target func([]float64) float64) ([][]float64, []float64)
```


## Tests
After you have put the source code somewhere, go to the folder, and then you can run the test
```
go test
```

The results are as follows:

```
=======================
Test results from TestInit
-----------------------
//Go code:
g := InitGrid([]float64{1, 2, 3}, []float64{1, 3, 2})
fmt.Println(g)
-----------------------
//results:
returned values:                1
go routines:                    1
layers of the zoom-in:          0
decay rate of the grid sizes:   0.5
The grid base is as follows:
1 2 3 
1 2 3 
-----------------------
Go code:
g.Append(FromToBy(1, 0, -.2))
g.Append(FromToLen(1, 0, 5))
g.SetNumGoRoutines(2)
g.SetZoom(1)
g.SetDecay(.8)
g.SetNumReturn(3)
fmt.Println(g)
g.Append(FromToLen(1, 0, 5))
-----------------------
results:
returned values:                3
go routines:                    2
layers of the zoom-in:          1
decay rate of the grid sizes:   0.8
The grid base is as follows:
1 2 3 
1 2 3 
-0.8 -0.6000000000000001 -0.4 -0.2 0 1 
0 0.25 0.5 0.75 1 
4
=======================
Test restuls from TestSearch
-----------------------
//Go code:
g := InitGrid()
var tmp = FromToBy(-5.12, 5.12, .1)
g.Append(tmp, tmp)
g.SetNumGoRoutines(2)
g.SetZoom(2)
g.SetNumReturn(2)
var ret, val = g.Search(Rastrigin)
fmt.Println(ret, val)
fmt.Println(Rastrigin(ret[0]))
-----------------------
results:
[[-1.4970663597679845e-15 0.00015686274509654373] [-1.4970663597679845e-15 0.00015686274509654373]] [4.8816196116518995e-06 4.8816196116518995e-06]
4.8816196116518995e-06
PASS
ok      gridsearch      0.012s
```

For the benchmark, run
```
go test -run=XXX -bench=.
```

and the settings are
```go
g := InitGrid()
var tmp = FromToBy(-5.12, 5.12, .1)
g.Append(tmp, tmp)
g.SetNumGoRoutines(1)
g.SetZoom(2)
g.SetNumReturn(2)
```

I get
```
goos: darwin
goarch: amd64
pkg: gridsearch
BenchmarkGridSearch-4                300           5224740 ns/op
PASS
ok      gridsearch      2.100s
```
for 300 times sequentially within only 2.1 seconds. This can be compared with the last example in

[https://github.com/yukai-yang/zoomgrid]

which runs only once by using 4 cores on the same computer.

## How to use the package

```go
package main

import "github.com/yukai-yang/gridsearch"

func main () {
    var g = gridsearch.InitGrid()
    // do something ...
}
````
