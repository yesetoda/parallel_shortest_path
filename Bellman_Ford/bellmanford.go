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

func BellmanFordSequential(graph [][]int, start int) []int {
	n := len(graph)
	dist := make([]int, n)
	for i := range dist {
		dist[i] = INF
	}
	dist[start] = 0

	for i := 0; i < n-1; i++ {
		for u := 0; u < n; u++ {
			for v := 0; v < n; v++ {
				if graph[u][v] != INF && dist[u] != INF && dist[u]+graph[u][v] < dist[v] {
					dist[v] = dist[u] + graph[u][v]
				}
			}
		}
	}
	return dist
}

func BellmanFordParallel(graph [][]int, start int) []int {
	n := len(graph)
	dist := make([]int, n)
	for i := range dist {
		dist[i] = INF
	}
	dist[start] = 0

	numWorkers := runtime.GOMAXPROCS(0) 
	var wg sync.WaitGroup
	chunkSize := (n + numWorkers - 1) / numWorkers

	for i := 0; i < n-1; i++ {
		wg.Add(numWorkers)
		for j := 0; j < numWorkers; j++ {
			go func(workerID int) {
				defer wg.Done()
				startIdx := workerID * chunkSize
				endIdx := startIdx + chunkSize
				if endIdx > n {
					endIdx = n
				}

				for u := startIdx; u < endIdx; u++ {
					for v := 0; v < n; v++ {
						if graph[u][v] != INF && dist[u] != INF && dist[u]+graph[u][v] < dist[v] {
							var mu sync.Mutex
							mu.Lock()
							if dist[u]+graph[u][v] < dist[v] {
								dist[v] = dist[u] + graph[u][v]
							}
							mu.Unlock()
						}
					}
				}
			}(j)
		}
		wg.Wait() 
	}

	return dist
}

func benchmark(graph [][]int, start int) {
	startTime := time.Now()
	_ = BellmanFordSequential(graph, start)
	fmt.Println("Sequential Time Taken:", time.Since(startTime))

	startTime = time.Now()
	_ = BellmanFordParallel(graph, start)
	fmt.Println("Parallel Time Taken:", time.Since(startTime))
}

func main() {
	var nodes, startNode int
	fmt.Print("Enter the number of nodes: ")
	fmt.Scan(&nodes)
	fmt.Print("Enter the starting node: ")
	fmt.Scan(&startNode)

	graph := generateGraph(nodes)

	benchmark(graph, startNode)
}
