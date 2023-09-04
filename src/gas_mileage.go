package main

import "fmt"

func main() {
	var gallonCounter float32
	var milesCounter float32
	for i := 1; i <= 10; i++ {
		fmt.Println("Enter the gallons used (-1 to end): ")
		var gallons float32
		fmt.Scanln(&gallons)
		gallonCounter++
		fmt.Println("Enter the miles driven: ")
		var miles float32
		fmt.Scanln(&miles)
		milesCounter++
		milePerGallon := miles / gallons
		fmt.Println("The miles/gallons for this tank was ", milePerGallon)
		if gallons == -1 {
			break
		}
	}
	average := gallonCounter / milesCounter
	fmt.Println("The overall average miles/ gallon was ", average)

}
