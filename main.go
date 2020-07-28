package main

import (
	"fmt"
)

/* For Loop
func main() {
	fmt.Println("Enter a number")
	var num int
	fmt.Scanf("%d", &num)

	fmt.Println(num)
	for i := 1; i < num; i++ {
		if i%2 == 0 {
			fmt.Println("Even", i)
		} else {
			fmt.Println("Odd ", i)
		}
	}
}*/

//Switch
func main() {
	fmt.Println("Enter a number")
	var num int
	fmt.Scanf("%d", &num)
	switch num % 2 {
	case 0:
		fmt.Println("Even")
	case 1:
		fmt.Println("Odd")
	default:
		fmt.Println("Unknown")
	}

}
