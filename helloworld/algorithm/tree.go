package algorithm

import (
	"fmt"
	"math"
)

type Node[T any] struct {
	Left  *Node[T] `json:"left"`
	Right *Node[T] `json:"right"`
	Value T        `json:"value"`
}

func newM[T any](length int) [][]T {
	cache := make([][]T, length)
	for i := range cache {
		cache[i] = make([]T, length)
	}
	return cache
}

func PrintMat[T any](mat [][]T) {
	for _, arr := range mat {
		fmt.Println(arr)
	}
}

func Solution(a []float64, b []float64) ([][]int, float64, [][]float64) {
	if len(a)+1 != len(b) {
		panic("Illegal Access Parameters")
	}
	n := len(a)
	w := newM[float64](n)
	m := newM[float64](n)
	s := newM[int](n)
	for space := 0; space < n; space++ {
		for i := 0; i < n-space; i++ {
			j := i + space
			if space == 0 {
				w[i][j] = a[i] + b[i] + b[i+1]
			} else {
				w[i][j] = w[i][j-1] + a[j] + b[j+1]
			}
			m[i][j] = math.MaxInt
			for k := i; k <= j; k++ {
				t := float64(0)
				if k < j {
					t += m[k+1][j]
				} else {
					t += b[k+1]
				}
				if i < k {
					t += m[i][k-1]
				} else {
					t += b[i]
				}
				if m[i][j] > t {
					m[i][j] = t
					s[i][j] = k
				}
			}
			m[i][j] += w[i][j]
		}
	}

	return s, m[0][n-1], m
}

func Treeify[T int | float64 | float32](arr []T, m [][]int, begin, end int) *Node[T] {
	if begin > end {
		return nil
	}
	k := m[begin][end]
	node := &Node[T]{
		Left:  Treeify(arr, m, begin, k-1),
		Right: Treeify(arr, m, k+1, end),
		Value: arr[k],
	}

	return node
}
