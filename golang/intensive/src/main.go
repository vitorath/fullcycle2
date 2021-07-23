package main

import (
	"fmt"

	"github.com/vitorath/golang-intensive/src/utility"
	myhelper "github.com/vitorath/golang-intensive/src/utility/helper"
)

func main() {
	fmt.Println("Hello, World")
	fmt.Println(utility.GetMyName())
	fmt.Println(utility.GetName())
	fmt.Println(myhelper.GetName())
}
