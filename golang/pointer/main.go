package main

import "fmt"

type Car struct {
	Name string
}

func (c Car) running() {
	c.Name = "BMW"
	fmt.Println(c.Name)
}

func (c *Car) runningRef() {
	c.Name = "BMW"
	fmt.Println(c.Name)
}

func main() {
	a := 10
	fmt.Println("memory ", &a, "(", a, ")")

	var pointer *int = &a
	fmt.Println("memory", pointer, "(", *pointer, ")")

	*pointer = 50
	fmt.Println(*pointer, " = ", a)

	b := &a
	*b = 60
	fmt.Println(*pointer, " = ", a, " = ", *b)

	c := *b
	c = 10000
	fmt.Println(*pointer, " = ", a, " = ", *b, " != ", c)

	fmt.Println("----------------------------------")

	variable := 10
	ret := abc(&variable)
	fmt.Println(variable)
	ret = 100
	fmt.Println(ret, " != ", variable)

	fmt.Println("----------------------------------")

	car := Car{
		Name: "Ka",
	}

	car.running()
	fmt.Println(car.Name)

	fmt.Println("----------------------------------")

	car.runningRef()
	fmt.Println(car.Name)

}

func abc(a *int) int {
	*a = 200
	return *a
}
