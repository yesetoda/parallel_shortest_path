package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

const INF = math.MaxInt32

func generateGraph(size int) [][]int {
	rand.Seed(time.Now().UnixNano())

	graph := make([][]int, size)
	for i := range graph {
		graph[i] = make([]int, size)
		for j := range graph[i] {
			if i == j {
				graph[i][j] = 0
			} else if rand.Float32() > 0.5 {
				graph[i][j] = rand.Intn(10) + 1
			} else {
				graph[i][j] = INF
			}
		}
	}
	return graph
}

func parallelFloydWarshall(graph [][]int) [][]int {
	n := len(graph)
	dist := make([][]int, n)
	for i := range graph {
		dist[i] = make([]int, n)
		copy(dist[i], graph[i])
	}

	numWorkers := runtime.GOMAXPROCS(0)

	for k := 0; k < n; k++ {
		var wg sync.WaitGroup

		chunkSize := (n + numWorkers - 1) / numWorkers

		for w := 0; w < numWorkers; w++ {
			start := w * chunkSize
			end := start + chunkSize
			if end > n {
				end = n
			}

			wg.Add(1)
			go func(start, end int) {
				defer wg.Done()
				for i := start; i < end; i++ {
					for j := 0; j < n; j++ {
						if dist[i][k] != INF && dist[k][j] != INF && dist[i][k]+dist[k][j] < dist[i][j] {
							dist[i][j] = dist[i][k] + dist[k][j]
						}
					}
				}
			}(start, end)
		}
		wg.Wait()
	}

	return dist
}

func FloydWarshall(graph [][]int) [][]int {
	n := len(graph)
	dist := make([][]int, n)
	for i := range graph {
		dist[i] = make([]int, n)
		copy(dist[i], graph[i])
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k] != INF && dist[k][j] != INF && dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	return dist
}


func benchmark(graph [][]int) {
	pstart := time.Now()
	_ = parallelFloydWarshall(graph)
	fmt.Println("Parallel Time Taken: ", time.Since(pstart))

	npstart := time.Now()
	_ = FloydWarshall(graph)
	fmt.Println("Non-Parallel Time Taken: ", time.Since(npstart))
}

func main() {
	var nodes int
	fmt.Print("Enter the number of nodes: ")
	fmt.Scan(&nodes)
	graph := generateGraph(nodes)
	benchmark(graph)
}
