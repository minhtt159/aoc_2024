package main

import (
	"sort"
)

// Function to find all maximal cliques using the Bron-Kerbosch algorithm
func bronKerbosch(R, P, X []string, graph map[string][]string) [][]string {
	var cliques [][]string

	// Base case: when P and X are empty, we have found a clique
	if len(P) == 0 && len(X) == 0 {
		// We found a clique, so add it to the result
		clique := make([]string, len(R))
		copy(clique, R)
		cliques = append(cliques, clique)
		return cliques
	}

	// Iterate over all nodes in P
	for i := 0; i < len(P); i++ {
		node := P[i]
		// Create the new sets
		newR := append([]string(nil), R...) // Copy R
		newR = append(newR, node)

		// Create new sets P' and X'
		newP := intersect(P, graph[node])
		newX := intersect(X, graph[node])

		// Recursively call Bron-Kerbosch with the updated sets
		cliques = append(cliques, bronKerbosch(newR, newP, newX, graph)...)

		// Move node from P to X (since we have processed it)
		P = remove(P, node)
		X = append(X, node)
	}

	return cliques
}

// Helper function to compute the intersection of two slices
func intersect(a, b []string) []string {
	set := make(map[string]bool)
	var result []string

	// Create a set from slice a
	for _, val := range a {
		set[val] = true
	}

	// Add common elements from b to the result
	for _, val := range b {
		if set[val] {
			result = append(result, val)
		}
	}
	return result
}

// Helper function to remove a node from a slice
func remove(slice []string, node string) []string {
	for i, val := range slice {
		if val == node {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

// Function to find all maximal cliques in a graph
func findCliques(graph map[string][]string) [][]string {
	// Initialize R, P, and X
	var R []string
	var P []string
	var X []string

	// Populate P with all nodes in the graph
	for node := range graph {
		P = append(P, node)
	}

	// Sort P to ensure we process the nodes in a consistent order (for efficiency)
	sort.Strings(P)

	// Run the Bron-Kerbosch algorithm
	return bronKerbosch(R, P, X, graph)
}
