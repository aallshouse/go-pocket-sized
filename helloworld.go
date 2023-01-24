package main

import "fmt"

func main() {
	greeting := greet("en")
	fmt.Println(greeting)
}

type language string

func greet(l language) string {
	switch l {
	case "en":
		return "Hello, World!"
	case "es":
		return "Hola, Mundo!"
	case "fr":
		return "Bonjour, le monde!"
	default:
		return ""
	}
}
