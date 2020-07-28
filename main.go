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

//Arrays

/*func main() {
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

}*/

//Slice
func main() {
	x := make([]float64, 5)
	fmt.Println(x) //[ 0 0 0 0 0 ]

	y := make([]float64, 5, 10)
	//10 represents the capacity of the underlying array
	fmt.Println(y) //[ 0 0 0 0 0 ]

	arr := [5]float64{1, 2, 3, 4, 5}
	z1 := arr[0:5]
	z2 := arr[1:5]
	z3 := arr[:]
	fmt.Println(z1) //[1 2 3 4 5]
	fmt.Println(z2) //[2 3 4 5]
	fmt.Println(z3) //[1 2 3 4 5]

	//Slice Functions Append and Copy
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5, 6)
	fmt.Println(slice1, slice2) //[1 2 3] [1 2 3 4 5 6]

	slice3 := []int{10, 20, 30}
	slice4 := make([]int, 3)
	//copy(slice3, slice4)
	copy(slice4, slice3) //[10 20 30] [10 20 30]
	fmt.Println(slice3, slice4)

}
