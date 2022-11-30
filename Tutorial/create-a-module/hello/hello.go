package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main()  {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	
	message, err := greetings.Hello("Death Vader") 
	if (err != nil) { 
		log.Fatal(err)
	}
	fmt.Println(message)
	
	names := []string{"Dah", "heh", "huh"}
	messages, err := greetings.Hellos(names)
	if (err != nil) {
		log.Fatal(err)
	}
	//fmt.Println(messages)
	printlnMessages(messages)
}

func printlnMessages(messages map[string]string) {
	for _, msg := range messages {
		fmt.Println(msg)
	}
}