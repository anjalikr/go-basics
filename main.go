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
/*func main() {
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

}*/

//Map
/*func main() {
	x := make(map[string]int)
	x["sunday"] = 1
	x["monday"] = 2
	x["tuesday"] = 3
	fmt.Println(x) //  map[monday:2 sunday:1 tuesday:3]
	delete(x, "monday")
	fmt.Println(x)                              // map[sunday:1 tuesday:3]
	fmt.Println("key for monday:", x["monday"]) //key for monday: 0

	//map can return two values instead of just one.
	//The first value is the result of the lookup,
	//the second tells us whether or not the lookup was successful.

	//First we try to get the value from the map,
	//then if it's successful we run the code inside of the block

	if day, ok := x["monday"]; ok {
		fmt.Println(day, ok) // will not display anyting
	}
	if day, ok := x["tuesday"]; ok {
		fmt.Println(day, ok) //3 true
	}

}*/

//Functions
//We can also name the return type
/*func main() {
	x := f1()
	fmt.Println("Returned value:", x)
	y, z := f2()
	fmt.Println("Multiple value:", y, z)
	total := add(1, 2, 3)
	fmt.Println("Addition of numbers:", total)
	// we can also pass a slice
	slice := []int{1, 2, 3}
	totalSlice := add(slice...)
	fmt.Println("Addition of elemrnts of slice", totalSlice)

}
func f1() (r int) {
	r = 1
	return
}

// Go is capable of returning multiple values from a function
func f2() (int, int) {
	return 4, 5
}

//Variadic functons
//By using ... before the type name of the last parameter you can
//indicate that it takes zero or more of those parameters.
//In this case we take zero or more ints.

func add(num ...int) int {
	var total int = 0
	for _, value := range num {
		total += value
	}
	return total
}*/

//Closure
//A function like this together with the non-local variables it references
//is known as a closure. In this case increment and the variable x form the closure.

/*func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println("CLosure", increment())
	fmt.Println("CLosure", increment())
}*/

//Defer Panic & Recover
//defer which schedules a function call to be run after the function completes
//it keeps our Close call near our Open call so it's easier to understand
/*
f, _ := os.Open(filename)
defer f.Close()
*/
//recover stops the panic and returns the value that was passed to the call to panic
//call to recover will never happen in the case of panic(if inside the main func) because the
//call to panic immediately stops execution of the function.
//Instead we have to pair it with defer:
/*func main() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}*/

//Pointer
//we use the & operator to find the address of a variable.
//&x returns a *int (pointer to an int) because x is an int
func main() {
	x := 0
	zero(&x)
	fmt.Println("Returned value from zero()", x)
	ptr2 := new(int)
	one(ptr2)
	fmt.Println("Returned value from one()", *ptr2)

}
func zero(ptr1 *int) {
	*ptr1 = 25
}
func one(ptr2 *int) {
	*ptr2 = 30
}

//new
//Another way to get a pointer is to use the built-in new function
