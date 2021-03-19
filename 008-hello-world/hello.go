package main

import "fmt"

func main() {
	fmt.Println("Hello")
	foo()
	fmt.Println("after foo")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("i = ", i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("inside foo")
}

func bar() {
	fmt.Println("this is code exit")
}

/*
int func foo1() {
	int a = 1
	return a
}
*/
