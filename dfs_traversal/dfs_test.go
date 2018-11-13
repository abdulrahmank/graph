package dfs_traversal_test

import (
	"github.com/abdulrahmank/graph/dfs_traversal"
	"github.com/abdulrahmank/graph/dfs_traversal/dynamic_programming"
	"github.com/abdulrahmank/graph/dfs_traversal/iteration"
	"github.com/abdulrahmank/graph/dfs_traversal/recursion"
	"testing"
)

func TestDepthFirstSearch(t *testing.T) {

	nodes := make(map[string]dfs_traversal.NodeList)
	nodes["a"] = dfs_traversal.NodeList{"b", "c", "d"}
	nodes["b"] = dfs_traversal.NodeList{"e", "f"}
	nodes["c"] = dfs_traversal.NodeList{"g"}
	nodes["e"] = dfs_traversal.NodeList{"h", "i"}
	expectedResult := []string{"a","b","e", "h", "i", "f", "c", "g", "d"}

	t.Run("Should traverse given tree in depth first search using recursion", func(t *testing.T) {
		result := recursion.Dfs("a", nodes)

		for index, node := range expectedResult {
			if result[index] != node {
				t.Errorf("Expected %s but was %s", node, result[index])
			}
		}
	})

	t.Run("Should traverse given tree in depth first search using iteration", func(t *testing.T) {
		result := iteration.Dfs("a", nodes)

		for index, node := range expectedResult {
			if result[index] != node {
				t.Errorf("Expected %s but was %s", node, result[index])
			}
		}
	})

	t.Run("Should traverse given tree in depth first search using Dynamic programming", func(t *testing.T) {
		result := dp.Dfs("a", nodes)

		for index, node := range expectedResult {
			if result[index] != node {
				t.Errorf("Expected %s but was %s", node, result[index])
			}
		}
	})
}

func BenchmarkDFS(b *testing.B) {

	nodes := make(map[string]dfs_traversal.NodeList)
	nodes["a"] = dfs_traversal.NodeList{"b", "c", "d"}
	nodes["b"] = dfs_traversal.NodeList{"e", "f"}
	nodes["c"] = dfs_traversal.NodeList{"g"}
	nodes["e"] = dfs_traversal.NodeList{"h", "i"}

	b.Run("Recursion", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			recursion.Dfs("a", nodes)
		}
	})

	b.Run("Iteration", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			iteration.Dfs("a", nodes)
		}
	})

	b.Run("Dynamic programming", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			dp.Dfs("a", nodes)
		}
	})
}

