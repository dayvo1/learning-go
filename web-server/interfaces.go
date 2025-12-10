package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof! my name is " + d.Name
}

func (c Cat) Speak() string {
	return "Meow! my name is " + c.Name
}

func MakeAnimalSpeak(a Animal) {
	fmt.Println(a.Speak())
}

// func main() {
// 	dog := Dog{Name: "Buddy"}
// 	cat := Cat{Name: "Whiskers"}

// 	MakeAnimalSpeak(dog)
// 	MakeAnimalSpeak(cat)
// }
