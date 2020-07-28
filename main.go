package main

import "fmt"

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
/*func main() {
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

}*/

func main() {
	//First Method
	var x [5]float64
	x[0] = 10
	x[1] = 50
	x[2] = 20
	x[3] = 30
	x[4] = 40

	var total1 float64 = 0
	for i := 0; i < len(x); i++ {
		total1 = total1 + x[i]
	}
	fmt.Println("Total1 ", total1)

	//Second Method
	y := [5]float64{10, 20, 30, 40, 50}
	var total2 float64 = 0

	//_ (underscore) is used to tell the compiler that we don't need this
	for _, value := range y {
		total2 += value
	}
	fmt.Println("Total2 ", total2)

}
