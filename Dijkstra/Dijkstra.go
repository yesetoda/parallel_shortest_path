package main1

import (
	"fmt"
	"math"
	"math/rand"
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

func DijkstraSequential(graph [][]int, start int) []int {
	n := len(graph)
	dist := make([]int, n)
	visited := make([]bool, n)

	for i := range dist {
		dist[i] = INF
	}
	dist[start] = 0

	for i := 0; i < n; i++ {
		minDist := INF
		u := -1
		for v := 0; v < n; v++ {
			if !visited[v] && dist[v] < minDist {
				minDist = dist[v]
				u = v
			}
		}

		visited[u] = true

		for v := 0; v < n; v++ {
			if !visited[v] && graph[u][v] != INF && dist[u]+graph[u][v] < dist[v] {
				dist[v] = dist[u] + graph[u][v]
			}
		}
	}
	return dist
}


func DijkstraParallel(graph [][]int, start int) []int {
	n := len(graph)
	dist := make([]int, n)
	visited := make([]bool, n)
	var mu sync.Mutex

	for i := range dist {
		dist[i] = INF
	}
	dist[start] = 0

	for i := 0; i < n; i++ {
		minDist := INF
		u := -1
		for v := 0; v < n; v++ {
			if !visited[v] && dist[v] < minDist {
				minDist = dist[v]
				u = v
			}
		}

		visited[u] = true

		var wg sync.WaitGroup
		for v := 0; v < n; v++ {
			if !visited[v] && graph[u][v] != INF && dist[u]+graph[u][v] < dist[v] {
				wg.Add(1)
				go func(v int) {
					defer wg.Done()
					mu.Lock()
					if dist[u]+graph[u][v] < dist[v] {
						dist[v] = dist[u] + graph[u][v]
					}
					mu.Unlock()
				}(v)
			}
		}
		wg.Wait()
	}
	return dist
}

func benchmark(graph [][]int, start int) {
	startTime := time.Now()
	_ = DijkstraSequential(graph, start)
	fmt.Println("Sequential Time Taken:", time.Since(startTime))

	startTime = time.Now()
	_ = DijkstraParallel(graph, start)
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
