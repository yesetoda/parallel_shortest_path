
# 🚀 Parallel Shortest Path Algorithms in Go

This repository showcases the implementation and optimization of shortest path algorithms using **Go (Golang)**. It highlights how parallel processing can enhance the performance of graph algorithms by leveraging Go's powerful concurrency model.

---

## 📑 Table of Contents

1. [Overview](#overview)  
2. [Implemented Algorithms](#implemented-algorithms)  
3. [Project Structure](#project-structure)  
4. [Installation](#installation)  
5. [Usage](#usage)  
6. [Performance Analysis](#performance-analysis)  
7. [Example ](#example)  
8. [Contributors](#contributors)  


---

## 📝 Overview

This project implements three shortest path algorithms, with both sequential and parallel versions, to explore the performance gains from parallelization:

- **Dijkstra's Algorithm**  
- **Bellman-Ford Algorithm**  
- **Floyd-Warshall Algorithm**  

The objective is to demonstrate the effectiveness of parallel processing for computationally intensive graph operations, particularly on large datasets.

---

## ⚙️ Implemented Algorithms

### **1. Dijkstra's Algorithm**
- Purpose: Computes the shortest path from a single source to all other vertices.  
- Parallelization: Splits the graph traversal process across multiple goroutines to enhance performance.

### **2. Bellman-Ford Algorithm**
- Purpose: Calculates shortest paths from a single source, allowing for graphs with negative weights.  
- Parallelization: Concurrently relaxes edges, significantly reducing computation time for dense graphs.

### **3. Floyd-Warshall Algorithm**
- Purpose: Computes all-pairs shortest paths using dynamic programming.  
- Parallelization: Distributes row and column operations across multiple goroutines.

---

## 📂 Project Structure

```plaintext
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

## 🛠️ Installation

### Prerequisites
- **Go**: [Install Go](https://go.dev/doc/install).  
- **Git**: [Install Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git).  

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/yesetoda/shortest_path.git
   cd shortest_path
   ```
2. Navigate to the desired algorithm folder (`Bellman Ford`, `Dijkstra`, or `Floyd Warshal`).

3. Run the program:
   ```bash
   go run <filename.go>
   ```

---

## 🚀 Usage

### Input
The programs prompt for:
1. **Number of nodes** in the graph.  
2. **Starting node** (for Dijkstra and Bellman-Ford algorithms).  

Graphs are generated randomly with:  
- **Edge weights**: Random integers between 1 and 10.    

### Output
- **Execution times** for sequential and parallel implementations.  
- **Shortest path results** (distances or matrices).  

### Example Command
```bash
cd "Bellman Ford"
go run bellmanford.go
```

---

## 📊 Performance Analysis

The project compares sequential and parallel implementations to highlight performance improvements.  
**Key Findings**:
- Parallel Bellman-Ford: Faster for dense graphs due to concurrent edge relaxation.  
- Parallel Dijkstra: Benefits from shared distance updates across goroutines.  
- Parallel Floyd-Warshall: Notable gains for larger adjacency matrices.  

### example 
```plaintext
Enter the number of nodes: 100
Enter the starting node: 0
Sequential Time Taken: 5.23s
Parallel Time Taken: 2.14s
```


## 👥 Contributors

This project was collaboratively developed by:  
- **Yeneineh Seiba**  
- **Sitotaw Desta**  
- **Bahran Solomon**  
- **Yeabsra Tesfaye**  
- **Dunia Nebil**

---

Thank you for exploring our project! 🌟
