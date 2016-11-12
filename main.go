package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	nameFlag := flag.String("name", "[name]", "Put your name here")
	booleanFlag := flag.Bool("really", false, "Add flag if it is true")
	flag.Parse()

	var reader = bufio.NewReader(os.Stdin)
	fmt.Print("Hello, " + *nameFlag + ", please enter some text: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error:", err)
	}
	fmt.Println(text)

	fmt.Println("Really:", *booleanFlag)
}
