package span

import (
	"Graph/graph"
	"fmt"
	"time"
)

func BenchMark() {
	var start = time.Now()
	var edges, size, err = readGraph() //IO就是慢！！！
	if err != nil {
		fmt.Println("Illegal Input")
		return
	}
	var roads = transform(edges, size)
	fmt.Printf("Prepare Graph [%d vertexes & %d edges] in %v\n", size, len(edges), time.Since(start))

	start = time.Now()
	ret1, fail := Kruskal(edges, size)
	var tm1 = time.Since(start)
	if fail {
		fmt.Println("fail")
	}

	start = time.Now()
	ret2, fail := Prim(roads)
	var tm2 = time.Since(start)
	if fail {
		fmt.Println("fail")
	}

	if ret1 != ret2 {
		fmt.Printf("Kruskal[%d] != Prim[%d]\n", ret1, ret2)
	} else {
		fmt.Println("Kruskal:", tm1)
		fmt.Println("Prim:   ", tm2)
	}
}

func readGraph() (edges []graph.Edge, size int, err error) {
	var total int
	_, err = fmt.Scan(&size, &total)
	if err != nil || size < 2 || size > total {
		return []graph.Edge{}, 0, err
	}
	edges = make([]graph.Edge, total)
	for i := 0; i < total; i++ {
		_, err = fmt.Scan(&edges[i].A, &edges[i].B, &edges[i].Dist)
		if err != nil {
			return []graph.Edge{}, 0, err
		}
	}
	return edges, size, nil
}

func transform(edges []graph.Edge, size int) [][]graph.Path {
	var roads = make([][]graph.Path, size)
	for _, path := range edges {
		roads[path.A] = append(roads[path.A], graph.Path{Next: path.B, Dist: path.Dist})
	}
	return roads
}
