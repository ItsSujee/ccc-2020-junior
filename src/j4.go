package main

import (
	"fmt"
	"math"
)

func main() {
	var text string
	var str string
	output := "no"

	fmt.Scanf("%s\n",&text)
	fmt.Scanf("%s\n",&str)

	textLen := len(text)
	strLen := len(str)

	//List of every possible cycle of str
	var p []string
	for i:=0; i<strLen; i++{
		letters := i + strLen
		remainder := int(math.Mod(float64(letters), float64(strLen)))
		buffer := str[i:] + str[0:remainder]
		p = append(p, buffer)
	}

	//List every str length substring in text
	steps := textLen - strLen
	var n []string
	for i:=0; i < steps; i++ {
		n = append(n, text[i:i+strLen])
	}

	//Output yes if any of the cycles of str exist in the list of all possible
	for _, v1 := range n {
		for _,v2 := range p {
			if v1 == v2 {
				output = "yes"
				break
			} else {
				continue
			}
		}
	}

	fmt.Println(output)
}
