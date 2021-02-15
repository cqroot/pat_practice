package main

import "fmt"

func main() {
	var num int
	fmt.Scanln(&num)

	for i := 0; i < num/100; i++ {
		fmt.Print("B")
	}
	for i := 0; i < num/10%10; i++ {
		fmt.Print("S")
	}
	for i := 1; i <= num%10; i++ {
		fmt.Print(i)
	}
	fmt.Println()
}
