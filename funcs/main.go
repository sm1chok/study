package main

import "fmt"

func main() {
	x := 20
	y := 5
	Sum := func(x, y int) int { return x + y } //Присваиваем переменной значение функции прямо внтури main и передаем в каяестве аргумента другой функции
	// Sub := func(x, y int) int { return x - y }

	fmt.Println(calculate(x, y, Sum))
}

// func Greet() {
// 	fmt.Println("Hello suki")
// }

// func PersonGreet(name string) {
// 	fmt.Printf("Hello, %s\n", name)
// }

// func FioGreet(name, surname string) {
// 	fmt.Printf("Salam, %s %s\n", surname, name)
// }

func calculate(x, y int, action func(x, y int) int) int { //Функция как аргумент
	return action(x, y)
}
