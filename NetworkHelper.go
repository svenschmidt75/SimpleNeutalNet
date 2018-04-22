package main

import (
	"math"
	"math/rand"
)

func max(lhs int, rhs int) int {
	if lhs < rhs {
		return rhs
	}
	return lhs
}

func min(lhs int, rhs int) int {
	if lhs < rhs {
		return lhs
	}
	return rhs
}

func GenerateRandomIndices(size int) []int {
	// generate random permutation
	perm := rand.Perm(size)
	return perm
}

func GetError(outputActivations []float64, a []float64) float64 {
	var err float64
	for idx := range a {
		d1 := outputActivations[idx]
		d2 := a[idx]
		err += (d1 - d2) * (d1 - d2)
	}
	return math.Sqrt(err)
}

func GetClass(a []float64) int {
	var index int
	var value float64 = -1
	for idx := range a {
		if a[idx] > value {
			value = a[idx]
			index = idx
		}

	}
	return index
}
