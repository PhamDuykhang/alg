package main

import (
	"fmt"
	"sort"
)

type ByWeight []WeighInfo

func (bw ByWeight) Len() int {
	return len(bw)
}
func (bw ByWeight) Less(i, j int) bool {
	return bw[i].W < bw[j].W
}

func (bw ByWeight) Swap(i, j int) {
	bw[i], bw[j] = bw[j], bw[i]
}

type (
	Graph struct {
		NumV      int
		A         [][]int
		WeighList []WeighInfo
	}

	WeighInfo struct {
		S int
		D int
		W int
	}
)

func NewWeightedGraph(numV int) Graph {
	a := make([][]int, numV)

	for i := 0; i < numV; i++ {
		a[i] = make([]int, numV)
	}
	return Graph{
		NumV: numV,
		A:    a,
	}

}

func (g *Graph) AddEdge(src, dest, w int) {
	g.A[src][dest] = w
	g.WeighList = append(g.WeighList, WeighInfo{
		S: src,
		D: dest,
		W: w,
	})
}

func getMin(wl []WeighInfo) WeighInfo {
	fmt.Println(wl)
	m := wl[0]
	for _, w := range wl {
		if w.W < m.W {
			m = w
		}
	}
	return m
}

func (g *Graph) GetConnect(v int) []WeighInfo {
	var w []WeighInfo
	for i, el := range g.A[v] {
		if el > 0 {
			w = append(w, WeighInfo{
				S: v,
				D: i,
				W: el,
			})
		}
	}
	return w
}

func (g *Graph) PrimsMST(src int) []WeighInfo {
	visited := make([]bool, g.NumV)
	visited[src] = true
	ne := g.GetConnect(src)
	var result []WeighInfo
	for i := 0; i < g.NumV; i++ {
		m := getMin(ne)
		for _, el := range ne {
			if !visited[el.D] {
				visited[el.D] = true
				result = append(result, el)
			}
		}
		ne = g.GetConnect(m.D)
		if len(ne)==0{
			break
		}
	}
	return result

}

func (g *Graph) KruskalMST() []WeighInfo {
	sort.Sort(ByWeight(g.WeighList))
	var result []WeighInfo
	dj := InitDisJoin(len(g.WeighList))
	for _, wl := range g.WeighList {
		if !dj.Find(wl.S, wl.D) {
			result = append(result, wl)
			dj.Merge(wl.S, wl.D)
		}
	}
	return result

}

func (g *Graph) PrintWeightList() {
	for _, l := range g.WeighList {
		fmt.Printf("%d (%d,%d)\n", l.W, l.S, l.D)
	}
}

func main() {
	g := NewWeightedGraph(9)
	g.AddEdge(0, 1, 4)
	g.AddEdge(0, 7, 8)
	g.AddEdge(1, 7, 11)
	g.AddEdge(1, 2, 8)
	g.AddEdge(7, 8, 7)
	g.AddEdge(7, 6, 1)
	g.AddEdge(2, 8, 2)
	g.AddEdge(8, 6, 6)
	g.AddEdge(2, 3, 7)
	g.AddEdge(2, 5, 4)
	g.AddEdge(6, 5, 2)
	g.AddEdge(3, 5, 14)
	g.AddEdge(3, 4, 9)
	g.AddEdge(5, 4, 10)

	g.PrintWeightList()
	fmt.Println("-----------Prims----------")
	res := g.PrimsMST(0)
	s := 0
	for _, rl := range res {
		fmt.Printf("Edge %d---%d Cost %d\n", rl.S, rl.D, rl.W)
		s = s + rl.W
	}
	fmt.Println("---------------------")
	fmt.Println("Total cost ", s)

	fmt.Println("----------Kruskal-----------")
	res = g.KruskalMST()
	s=0
	for _, rl := range res {
		fmt.Printf("Edge %d---%d Cost %d\n", rl.S, rl.D, rl.W)
		s = s + rl.W
	}
	fmt.Println("---------------------")
	fmt.Println("Total cost ", s)

}
