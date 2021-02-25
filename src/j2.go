package main

import "fmt"

func main() {
	var population int
	var sick int
	var rate int
	fmt.Scanln(&population)
	fmt.Scanln(&sick)
	fmt.Scanln(&rate)

	infectious := sick
	day := 0

	for sick <= population {
		sick = sick + infectious * rate
		infectious = infectious * rate
		day++
	}

	fmt.Println(day)
}
