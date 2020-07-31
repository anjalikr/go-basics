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
//new
//Another way to get a pointer is to use the built-in new function
/*func main() {
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
}*/

//Struct

/*type Circle struct {
	x, y, r float64
}

func main() {
	var circle1 Circle //Initilization
	circle2 := new(Circle)
	circle3 := Circle{x: 0, y: 0, r: 4}
	circle4 := Circle{0, 0, 6}
	fmt.Println(circle1.x, circle1.y, circle1.r)
	fmt.Println(circle2.x, circle2.y, circle2.r)
	fmt.Println(circle3.x, circle3.y, circle3.r)
	fmt.Println(circle4.x, circle4.y, circle4.r)
	fmt.Println("Area By Value:", circleAreaByValue(circle3))
	fmt.Println("Area By Reference:", circleAreaByReference(&circle4))
	fmt.Println("Area By Method:", circle3.circleAreaMethod())

	//The is-a relationship works this way intuitively:
	//People can talk, an android is a person, therefore an android can talk.
	android1 := new(Android)
	fmt.Println("Android:", android1)
	android1.Person.Talk()

}

func circleAreaByValue(c Circle) float64 { // Call by value
	return math.Pi * c.r * c.r
}

func circleAreaByReference(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

//Methods
//In between the keyword func and the name of the function we've added a “receiver”.
//The receiver is like a parameter – it has a name and a type – but by creating the
//function in this way it allows us to call the function using the . operator
//This is much easier to read, we no longer need the & operator
//(Go automatically knows to pass a pointer to the circle for this method)

func (c *Circle) circleAreaMethod() float64 {
	return math.Pi * c.r * c.r
}

//Embedded Types
//A struct's fields usually represent the has-a relationship.
// For example a Circle has a radius

type Person struct {
	name string
}

func (p *Person) Talk() {
	fmt.Println("My name is ", p.name)
}

//We use the type (Person) and don't give it a name.
// When defined this way the Person struct can be accessed using the type name
//Create new Android struct
type Android struct {
	Person
	modell string
}

//Interface
// we can use interface types as arguments to functions:
//Interfaces can also be used as fields*/

//Goroutine
//A goroutine is a function that is capable of running concurrently with other functions.
//Normally when we invoke a function our program will execute all the statements in a
//function and then return to the next line following the invocation. With a goroutine
//we return immediately to the next line and don't wait for the function to complete.

/*func main() {
	go f1()
	for i := 0; i < 10; i++ {
		go f1()
	}
	//This is why the call to the Scanln function has been included; without it the program
	//would exit before being given the opportunity to print all the numbers.
	var input string
	fmt.Scanln(&input)

}

//prints out the numbers from 0 to 10, waiting between 0 and 250 ms after each one.
//The goroutines should now run simultaneously.
func f1() {
	for i := 0; i < 10; i++ {
		fmt.Println(" ", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}*/

//Channels
//Channels provide a way for two goroutines to communicate with one another and
//synchronize their execution.
//The <- (left arrow) operator is used to send and receive messages on the channel.
// c <- "ping" means send "ping". msg := <- c means receive a message and store it in msg
//When pinger attempts to send a message on the channel it will wait until printer
// is ready to receive the message. (this is known as blocking)

/*func main() {

	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)
	var input string
	fmt.Scanln(&input)

}

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}
func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}*/

//Go has a special statement called select which works like a switch but for channels
/*func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			}
		}
	}()
	var input string
	fmt.Scanln(&input)
}*/

/*func main() {
	arr := []float64{1, 2, 3}
	avg := math.Average(arr)
	fmt.Println("Average", avg)

}*/

//We also only use the short name math when we reference functions from our library.
//If we wanted to use both libraries in the same program Go allows us to use an alias

//Type composition embedde type and overriding methods of embedded type

type Person struct {
	name     string
	age      int64
	location string
}

func (p *Person) PrintName() {
	fmt.Println("Name is", p.name)
}

func (p *Person) PrintDetails() {
	fmt.Println("Name:", p.name)
	fmt.Println("Age::", p.age)
	fmt.Println("Location:", p.location)
}

type Admin struct {
	Person //type embedding for composition
	Roles  []string
}

//overridng methods of embedded type
func (a *Admin) PrintDetails() {
	a.Person.PrintDetails()
	fmt.Println("Roles are")
	for _, value := range a.Roles {
		fmt.Println(value)
	}
}
func main() {
	ram := Admin{
		Person{"Ram", 32, "Kochi"},
		[]string{"Manage team", "Manage tasks", "Assign tasks"},
	}
	ram.PrintName()
	ram.PrintDetails()
}
