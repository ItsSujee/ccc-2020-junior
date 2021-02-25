package main

import "fmt"

func main() {
	// [S, M, L]
	var input [3] int
	fmt.Scanln(&input[0])
	fmt.Scanln(&input[1])
	fmt.Scanln(&input[2])

	weights := [3] int{1, 2, 3}

	var sum int
	for i:=0; i<3; i++ {
		sum += + input[i] * weights[i]
	}

	var output string
	if sum >= 10 {
		output = "happy"
	} else {
		output = "sad"
	}
	fmt.Println(output)
}