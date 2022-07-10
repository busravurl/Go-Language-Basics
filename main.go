package main

import "fmt"

func main() {
	fmt.Println("Hello")
	fmt.Print("Helloo")
	fmt.Print("!")

	//////////////////////

	var intro string = "helloo"
	fmt.Println(intro)

	var num int = 15
	fmt.Println(num)
	fmt.Println(100 + (100 * 10 / 100))
	fmt.Println()

	var num2 float32 = 0.1
	fmt.Println(num2)
	fmt.Println(100 + 100*num2)

	num3 := 25.2
	fmt.Println(num3)
	fmt.Printf("data type : %T\n", num3)

	//////////////////////

	var state bool

	var text1 string = "hello"
	var text2 string = "world"

	state = text1 != text2
	fmt.Println(state)
	fmt.Println(!state)
}
