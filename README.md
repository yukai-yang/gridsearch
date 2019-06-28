# gridsearch
implementing the grid search with zoom algorithm by using the Go language

## Test, build and install
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
You can build
```
go build
```

And you can install
```
go install
```

## How to use the package

```go
package main

import "github.com/yukai-yang/gridsearch"

func main () {
    var g = gridsearch.InitGrid()
    // do something ...
}
````
