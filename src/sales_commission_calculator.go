package main

import "fmt"

func main() {
	var sum float32
	var item float32
	var value float32

	for number := 1; number < 10; number++ {
		fmt.Scanln(&item)
		sum += item
		value = 0.9 * sum
		fmt.Printf("%-15s%n%d%-15s%n%f ", number, "Value \n", value)
	}

}
