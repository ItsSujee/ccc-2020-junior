package main

import (
	"fmt"
	"sort"
)

func main() {
	var x, y int
	var X []int
	var Y []int
	var N int
	fmt.Scanln(&N)

	for i:=0; i<N; i++{
		fmt.Scanf("%d,%d\n", &x, &y)
		X = append(X, x)
		Y = append(Y, y)
	}

	sort.Ints(X)
	sort.Ints(Y)

	fmt.Printf("%d,%d\n", X[0] - 1, Y[0] - 1)
	fmt.Printf("%d,%d\n", X[len(X)-1] + 1, Y[len(X)-1] + 1)
}