package main

import "fmt"

func main() {

	var i int = 1

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 0; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}
}
