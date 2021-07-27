package main

import "fmt"

func my_math(num int, num1 int) string {
	sol := num * num1
	// returning as a string
	return fmt.Sprint(sol)
}

// Args can be passed into functions like this if the type of both args are the same
// Instead of specifying each type of the arg, you can pass both args and then their type as last
func my_math1(num, num1 int) int {
	sol := num * num1
	// returning as an integer
	return sol
}

func myLoop() string {
	var my_var string = "champion"
	// counting the length of strings
	for i := 0; i < len(my_var); i++ {
		// Counts the length of the string
		fmt.Println(i)
		/* Prints out the letters in the string
		using the string method to convert it to letters so it
		doesn't print the unicode of the letters
		*/
		fmt.Println(string(my_var[i]))
	}
	return "Complete"
}

func newLoop() string {
	myString := "Hello World"
	for i, j := range myString {
		// Prints out the index alongside the unicode characters
		// fmt.Println(i, j)
		// Prints out the index alongside the letters in the string
		fmt.Println(i, string(j))
	}
	return "Loop done and out"
}

// looping through an array
func myArrLoop() string {
	/* Using the elipses(...) saves one the stress of having to update the length
	of the array/slice if the array/slice changes in length*/
	newArr := [...]string{"cat", "dog", "rabbit", "Squirrel"}
	for i := 0; i < len(newArr); i++ {
		fmt.Println(newArr[i])
	}
	return "Done"
}

// Another method of creating loops
// Here the variable is defined outside the scope of the loop
func anotherLoop() int {
	i := 0
	for ; i < 5; i++ {
		fmt.Println(i)
	}
	return i
}

// Using pointers to point to a memory allocation of an object
func testPointer() string {
	var myVariable int = 10
	var myPointer *int = &myVariable
	fmt.Println(*myPointer, myVariable)
	return "Run out"
}

type Person interface {
	name() string
}

type user struct {
	firstName string
	lastName  string
}

func test(g Person) {
	fmt.Println(g)
	fmt.Println(g.name())
}

func (myUser *user) name() string {
	return myUser.firstName + " " + myUser.lastName
}

func main() {
	my_math(5, 6)
	fmt.Println(my_math1(5, 8))
	myLoop()
	myArrLoop()
	anotherLoop()
	newLoop()
	testPointer()

	newData := user{firstName: "Foo", lastName: "Bar"}
	firstData := &newData
	test(firstData)
}
