package main

import (
	"fmt"
	"math"
	"encoding/binary"
)

func main() {
	fmt.Println(S(1732913, []int{11,7,1}))
	fmt.Println(s2(1732913, []int{11,7,1}))
}

// S computes the smallest number of coins needed to create change for the given sum.
// It returns a map from coin denominations to coin counts.
// sum must be positive.
// denoms must be in decreasing order, with 1 as the smallest element.

// Essentially, S minimizes    x1 + x2 + ... + xn
//              subject to     d1*x1 + d2*x2 + ... + dn*xn = sum
func S(sum int, denoms []int) map[int]int {
	//fmt.Printf("S(%v, %v)\n", sum, denoms)
	if sum == 0 {
		return makeZeroMap(denoms)
	}
	denoms = filterGreaterThan(denoms, sum)
	d1 := pop(&denoms)
	x1 := int(math.Floor(float64(sum)/float64(d1)))
	rem := sum % d1

	sol := S(rem, denoms)
	sol[d1] = x1
	//fmt.Printf("d1=%v, sol=%v\n", d1, sol)

	for {
		if len(denoms) == 0 {
			break
		}
		d2 := denoms[0]
		rem += d1
		x1--
		best := x1 + int(math.Ceil(float64(rem)/float64(d2)))
		if best >= sumValues(sol) {
			break
		}
		solNext := S(rem, denoms)
		solNext[d1] = x1
		if (sumValues(solNext) < sumValues(sol)) {
			sol = solNext
			//fmt.Printf("d1=%v, sol=%v\n", d1, sol)
		}
	}
	return sol
}

// makeZeroMap creates a map containing the given keys, with values set to zero.
func makeZeroMap(keys []int) map[int]int {
	m := make(map[int]int)
	for _, v := range keys {
		m[v] = 0
	}
	return m
}

// sumValues returns the sum of the values in the given map.
func sumValues(m map[int]int) int {
	sum := 0
	for _, value := range m {
		sum += value
	}
	return sum
}

// filterGreaterThan creates a slice containing all elements in descendingValues that are <= maxValue.
// descendingValues must be in decreasing order.
func filterGreaterThan(descendingValues []int, maxValue int) []int {
	i := 0
	for {
		if i >= len(descendingValues) {
			break
		}
		if descendingValues[i] <= maxValue {
			break
		}
		i++
	}
	return descendingValues[i:]
}

// pop removes and returns the first element of a slice.
func pop(arr *[]int) int {
	first := (*arr)[0]
	*arr = (*arr)[1:]
	return first
}

// s2 is a standard implementation of S, for comparison.
// It uses top-down dynamic programming.
// It's based on https://www.cs.usfca.edu/~galles/visualization/DPChange.html
func s2(sum int, denoms []int) map[int]int {
	args := serializeS2Args(sum, denoms)
	memo := s2Table[args]
	if memo != nil {
		return clone(memo)
	}
	if sum == 0 {
		ret := makeZeroMap(denoms)
		s2Table[args] = ret
		//fmt.Printf("s2(%v,%v)=%v\n", sum, denoms, ret)
		return clone(ret)
	}
	var best map[int]int
	for _, d := range denoms {
		var nextTry map[int]int
		if d <= sum {
			nextTry = s2(sum - d, denoms)
			nextTry[d] += 1
		}
		if best == nil || sumValues(best) > sumValues(nextTry) {
			best = nextTry
		}
	}
	s2Table[args] = best
	//fmt.Printf("s2(%v,%v)=%v\n", sum, denoms, best)
	return clone(best)
}

// Memoization table for function s2.
var s2Table map[string]map[int]int = make(map[string]map[int]int)

func serializeS2Args(sum int, denoms []int) string {
	b := make([]byte, 8 + 8 * len(denoms))

	binary.LittleEndian.PutUint64(b[0:], uint64(sum))
	
	for i := 0; i < len(denoms); i++ {
		binary.LittleEndian.PutUint64(b[8*(i+1):], uint64(denoms[i]))
	}
	return string(b)
}

func clone(m map[int]int) map[int]int {
	ret := make(map[int]int)
	for k, v := range m {
		ret[k] = v
	}
	return ret
}

