package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func (eb englishBot) getGreeting() string {
	//Very custom logic to return greeting in english
	return "Helloe there"
}

func (sb spanishBot) getGreeting() string {
	//very custom logic to rerurn greeting in spanish
	return "Hola"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// we will get an error that two functions cannot be declared with the same name, no overloading allowed
//func printGreeting(sb spanishBot) {
//fmt.Println(sb.getGreeting())
