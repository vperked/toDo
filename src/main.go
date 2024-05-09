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
		first  string
		second string
		third  string
	)
	fmt.Scan(&first)
	fmt.Println("1." + first)
	fmt.Scan(&second)
	fmt.Println("2. " + second)
	fmt.Scan(&third)
	fmt.Println("3. " + third)
	fmt.Println(third)
	fileTomfoolery(first, second, third)
	timeLimit := time.Hour * 24
	fmt.Println("Come back in", timeLimit)
	duration := time.Second
	time.Sleep(duration)
	os.Exit(2)
}

func fileTomfoolery(first string, second string, third string) {
	create, err := os.Create("list.txt")
	if err != nil {
		log.Err(err)
	}
	defer create.Close()
	_, err = create.WriteString(fmt.Sprintf("%s\n%s\n%s\n", first, second, third))
	if err != nil {
		log.Err(err)
	}
}
