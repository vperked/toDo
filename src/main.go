package main

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog/log"
)

func main() {
	fmt.Println("Welcome to your toDo list!")
	fmt.Println("Please write what you need to do: ")
	program()
}

func program() {
	var (
		first    string
		second   string
		third    string
		finished string
	)
	fmt.Scan(&first)
	fmt.Scan(&second)
	fmt.Scan(&third)

	fmt.Println("1. ", first)
	fmt.Println("2. ", second)
	fmt.Println("3. ", third)
	fmt.Println("Please come back when you are done.")
	create, err := os.Create("list.txt")
	if err != nil {
		log.Err(err)
	}
	defer create.Close()
	_, err = create.WriteString(fmt.Sprintf("%s\n%s\n%s\n", first, second, third))
	if err != nil {
		log.Err(err)
	}
	fmt.Scan(finished)
	math := time.Now().Day() + 1
	if finished == "done" {
		fmt.Println("Come back in ", math)
	} else {
		fmt.Println("Invalid")
	}
	err = os.Remove("list.txt")
	if err != nil {
		log.Err(err)
	}
}
