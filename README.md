
# 🚀 Parallel Shortest Path Algorithms in Go

This repository showcases the implementation and optimization of shortest path algorithms using **Go (Golang)**. It highlights how parallel processing can enhance the performance of graph algorithms by leveraging Go's powerful concurrency model.

---

## 🔐 Table of Contents

1. [Overview](#overview)  
2. [Implemented Algorithms](#implemented-algorithms)  
3. [Project Structure](#project-structure)  
4. [Installation](#installation)  
5. [Usage](#usage)  
6. [Performance Analysis](#performance-analysis)  
7. [Contributors](#contributors)  
8. [License](#license)  

---

## Overview

This project implements three shortest path algorithms, with both sequential and parallel versions, to explore the performance gains from parallelization:

- **Dijkstra's Algorithm**  
- **Bellman-Ford Algorithm**  
- **Floyd-Warshall Algorithm**  

The objective is to demonstrate the effectiveness of parallel processing for computationally intensive graph operations, particularly on large datasets.

---

## Implemented Algorithms

### **1. Dijkstra's Algorithm**
- **Purpose**: Computes the shortest path from a single source to all other vertices.  
- **Parallelization**: Splits the graph traversal process across multiple goroutines to enhance performance.

### **2. Bellman-Ford Algorithm**
- **Purpose**: Calculates shortest paths from a single source, allowing for graphs with negative weights.  
- **Parallelization**: Concurrently relaxes edges, significantly reducing computation time for dense graphs.

### **3. Floyd-Warshall Algorithm**
- **Purpose**: Computes all-pairs shortest paths using dynamic programming.  
- **Parallelization**: Distributes row and column operations across multiple goroutines.

---

## Project Structure

```plaintext
shortest_path/
├── Bellman_Ford/                # Bellman_Ford implementation
│   └── bellmanford.go           # Sequential and parallel versions
├── Dijkstra/                    # Dijkstra's implementation
│   └── dijkstra.go              # Sequential and parallel versions
├── Floyd_Warshal/               # Floyd_Warshall implementation
│   └── floydwarshal.go          # Sequential and parallel versions
└── README.md                    # Project documentation
```

---

## Installation

### Prerequisites
- **Go**: [Install Go](https://go.dev/doc/install).  
- **Git**: [Install Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git).  

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/yesetoda/parallel_shortest_path.git
   cd parallel_shortest_path
   ```
2. Navigate to the desired algorithm folder (`Bellman Ford`, `Dijkstra`, or `Floyd Warshal`).

3. Run the program:
   ```bash
   go run <filename.go>
   ```

---

## Usage

### Input
The programs prompt for:
1. **Number of nodes** in the graph.  
2. **Starting node** (for Dijkstra and Bellman-Ford algorithms).  

Graphs are generated randomly with:  
- **Edge weights**: Random integers between 1 and 10.  


### Output
- **Execution times** for sequential and parallel implementations.  

### Example Command
```bash
cd "Bellman Ford"
go run bellmanford.go
```

---

## Performance Analysis

The project compares sequential and parallel implementations to highlight performance improvements.

### **Execution Time for Sequential and Parallel Implementations**
We measured the execution times for both the sequential and parallel implementations of Bellman-Ford, Dijkstra, and Floyd-Warshall algorithms for various graph sizes. The results are summarized in the table below:

### Bellman-Ford Algorithm
| Graph Size (Nodes) | Sequential Time | Parallel Time | Speedup |
|--------------------|----------------|---------------|---------|
| 10                 | 0s             | 0s            | N/A     |
| 100                | 20.96ms        | 9.98ms        | 2.1x    |
| 500                | 1.53s          | 1.09s         | 1.4x    |
| 600                | 2.80s          | 1.71s         | 1.64x   |

### Dijkstra Algorithm
| Graph Size (Nodes) | Sequential Time | Parallel Time | Speedup |
|--------------------|----------------|---------------|---------|
| 10                 | 0s             | 1.02ms        | N/A     |
| 100                | 0s             | 1.07ms        | N/A     |
| 500                | 3.78ms         | 10.65ms       | 0.35x   |
| 600                | 5.48ms         | 7.26ms        | 0.75x   |

### Floyd-Warshall Algorithm
| Graph Size (Nodes) | Sequential Time | Parallel Time | Speedup |
|--------------------|----------------|---------------|---------|
| 10                 | 0s             | 885µs         | N/A     |
| 100                | 4.45ms         | 7.30ms        | 0.61x   |
| 500                | 141.56ms       | 395.71ms      | 0.36x   |
| 600                | 260.77ms       | 216.73ms      | 1.2x    |

### Sample Performance Example
For the **Floyd-Warshall Algorithm**, using 1000 nodes:

```cmd
Enter the number of nodes: 1000
Parallel Time Taken: 1.989603s
Non-Parallel Time Taken: 2.5884075s
```

## Contributors

This project was collaboratively developed by:  
- **Yeneineh Seiba**  
- **Sitotaw Desta**  
- **Bahran Solomon**  
- **Yeabsra Tesfaye**  
- **Dunia Nebil**


Thank you for exploring our project! 🌟
