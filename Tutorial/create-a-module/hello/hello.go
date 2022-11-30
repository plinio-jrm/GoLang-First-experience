package main

import (
	"fmt"
	"example.com/greetings"
)

func main()  {
	message := greetings.Hello("Shitness")
	fmt.Println(message)
}