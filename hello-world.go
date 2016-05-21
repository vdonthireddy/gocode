package main

import "fmt"

func main() {
	var i1 int
	i1 = 2

	var i2 int
	i2 = 3

	var i3 int

	i3 = i1 + i2

	fmt.Println("i1 is: ", i1);
	fmt.Println("i2 is: ", i2);
	fmt.Println("i3 is: ", i3);

	fmt.Println("i1 is at: ", &i1)
	fmt.Println("i2 is at: ", &i2)
	fmt.Println("i3 is at: ", &i3)

	fmt.Println("Hello World")

	var s1 string;
	s1 = "vijay";
	if (s1 == "vijay") {
		fmt.Println("Hello Vijay!");
	} else {
		fmt.Println("Hello!!!")
	}

	var s2 string;
	s2 = "don"
	if (s2 == "vijay") {

		fmt.Println("Hello Vijay!");
	} else {
		fmt.Println("Hello, you are NOT don")
	}

	a, b := 8,5;
	if (a<b) {
		fmt.Println("a is less than b")
	} else if (a>b) {
		fmt.Println("a is greater than b")
	} else {
		fmt.Println("a is equal to b")
	}

	for i:=0; i<5; i++ {
		fmt.Println("i value is: ", i)
	}

}