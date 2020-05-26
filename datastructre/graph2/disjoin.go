package main

type (
	DisJoin struct {
		size       []int
		Arr        []int
		numElement int
	}
)

func InitDisJoin(n int) DisJoin {
	a := make([]int, n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = i
		s[i] = i
	}
	return DisJoin{
		size:       s,
		Arr:        a,
		numElement: n,
	}
}

func root(a []int, i int) int {
	for a[i] != i {
		i = a[i]
	}
	return i
}

func (dj *DisJoin) Merge(a, b int) {
	dj.weighJoint(a, b)
}

func (dj *DisJoin) weighJoint(i, j int) {
	rootI := root(dj.Arr, i)
	rootJ := root(dj.Arr, j)
	if dj.size[rootI] < dj.size[rootJ] {
		dj.Arr[rootI] = rootJ
		dj.size[rootJ] = dj.size[rootJ] + 1
	} else {
		dj.Arr[rootJ] = rootI
		dj.size[rootI] = dj.size[rootI] + 1
	}
}

func (dj *DisJoin) Find(a, b int) bool {
	if root(dj.Arr, a) == root(dj.Arr, b) {
		return true
	}
	return false
}
