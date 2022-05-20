package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	// VERY Custom logic for generating an english greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

//func printGreeting(eb englishBot) {
//	fmt.Println(eb.getGreeting())
//}

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}