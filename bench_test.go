package gridsearch

import "testing"

func BenchmarkGridSearch(b *testing.B) {
	g := InitGrid()
	var tmp = FromToBy(-5.12, 5.12, .1)
	g.Append(tmp, tmp)
	g.SetNumGoRoutines(1)
	g.SetZoom(2)
	g.SetNumReturn(2)

	for n := 0; n < b.N; n++ {
		g.Search(Rastrigin)
	}
}
