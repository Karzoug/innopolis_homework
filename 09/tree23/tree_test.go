package tree23

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveTree(t *testing.T) {
	tree := New[int]()

	tree.Insert(10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150, 5, 15, 25, 8)
	require.True(t, tree.Search(5))
	require.True(t, tree.Search(30))

	tree.Remove(5)
	tree.Remove(8)
	tree.Remove(30)
	require.False(t, tree.Search(5))
	require.False(t, tree.Search(30))
}

func TestInsertTree(t *testing.T) {
	tree := New[int]()

	tree.Insert(10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150, 5, 15, 25, 8)
	require.True(t, tree.Search(5))
	require.True(t, tree.Search(30))
	require.False(t, tree.Search(7))

	tree.Insert(11)
	require.True(t, tree.Search(11))
}

func TestSearchTree(t *testing.T) {
	tree := New[int]()

	tree.Insert(10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150, 5, 15, 25, 8)
	require.True(t, tree.Search(5))
	require.True(t, tree.Search(30))
	require.False(t, tree.Search(7))
}
