Here's the complete `README.md` file tailored to your project structure and code:

```markdown
# Parallel Shortest Path Algorithms in Go

This repository implements and optimizes shortest path algorithms using Go (Golang). The project highlights the performance benefits of parallel processing in graph algorithms through Go's concurrency model.

---

## Table of Contents

- [Overview](#overview)
- [Implemented Algorithms](#implemented-algorithms)
- [Project Structure](#project-structure)
- [Installation](#installation)
- [Usage](#usage)
- [Performance Analysis](#performance-analysis)
- [Example Inputs](#example-inputs)
- [Contributors](#contributors)
- [License](#license)

---

## Overview

The project implements three shortest path algorithms both sequentially and in parallel:

1. **Dijkstra's Algorithm**
2. **Bellman-Ford Algorithm**
3. **Floyd-Warshall Algorithm**

The goal is to demonstrate how parallelization can improve performance for computationally intensive graph operations, especially for large graphs.

---

## Implemented Algorithms

### **Dijkstra's Algorithm**
- Finds the shortest path from a single source to all other vertices.
- Parallelization divides the graph traversal process across multiple goroutines.

### **Bellman-Ford Algorithm**
- Computes shortest paths from a single source, accommodating graphs with negative weights.
- Parallelization enables concurrent edge relaxation operations.

### **Floyd-Warshall Algorithm**
- Calculates all-pairs shortest paths.
- Parallelization splits the work across multiple iterations of the graph traversal.

---

## Project Structure

```
shortest_path/
├── Bellman Ford/                # Bellman-Ford implementation
│   └── bellmanford.go           # Sequential and parallel versions
├── Dijkstra/                    # Dijkstra's implementation
│   └── dijkstra.go              # Sequential and parallel versions
├── Floyd Warshal/               # Floyd-Warshall implementation
│   └── floydwarshal.go          # Sequential and parallel versions
└── README.md                    # Project documentation
```

---

## Installation

### Prerequisites
- **Go**: Ensure you have Go installed. Follow the guide at [https://go.dev/doc/install](https://go.dev/doc/install).
- **Git**: Required to clone the repository.

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/yesetoda/parallel_shortest_path.git
   cd parallel_shortest_path
   ```
2. Navigate to the algorithm folder of choice (`Bellman Ford`, `Dijkstra`, or `Floyd Warshal`).

3. Run the code:
   ```bash
   go run <filename.go>
   ```

---

## Usage

### Inputs
Each program prompts the user to provide:
1. Number of nodes in the graph.
2. (For Dijkstra and Bellman-Ford) Starting node for shortest path computation.

Graphs are generated randomly with the following properties:
- **Edge Weights**: Random integers between 1 and 10.
- **Edge Existence**: Determined randomly (50% probability of an edge between any two nodes).

### Outputs
1. Sequential execution time.
2. Parallel execution time.
3. Resulting shortest path distances or distance matrices.

### Example: Running Bellman-Ford
```bash
cd "Bellman Ford"
go run bellmanford.go
```

---

## Performance Analysis

Performance is measured by comparing the execution time of sequential and parallel implementations:
- **Bellman-Ford**: Parallel edge relaxation improves performance for dense graphs.
- **Dijkstra**: Parallelization focuses on concurrent updates to the distance array.
- **Floyd-Warshall**: Significant gains for larger matrices due to concurrent row/column operations.

Example output:
```
Enter the number of nodes: 100
Enter the starting node: 0
Sequential Time Taken: 5.23s
Parallel Time Taken: 2.14s
```

---

## Example Inputs

Graph representation in memory is generated as an adjacency matrix. Here's an example for a 4-node graph:

```
[
    [0, 5, INF, 10],
    [INF, 0, 3, INF],
    [INF, INF, 0, 1],
    [INF, INF, INF, 0]
]
```

---

## Contributors

This project was developed by:
- **Yeneineh Seiba**
- **Sitotaw Desta**
- **Bahran Solomon**
- **Yeabsra Tesfaye**
- **Dunia Nebil**

---
