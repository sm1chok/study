// Указатели - хранят адрес ячейки памяти значения, либо другого указателя (по дефолту значение адреса ячейки = nil)
// Тип указателя обозначается как *type (int, string, bool)
package main

import "fmt"

func main() {
	// default value
	var intPointer *int
	fmt.Printf("%T %#v \n", intPointer, intPointer) // %T - тип данных, %#V - значение

	//not-nil pointer
	//just variable
	var a int64 = 7
	fmt.Printf("%T %#v \n", a, a)

	//get variable pointer
	var pointerA *int64 = &a                                   // можно объявить сразу как pointerA := &a
	fmt.Printf("%T %#v %#v \n", pointerA, pointerA, *pointerA) //Просто выводя значение указателя мы получаем адрес ячейки памяти.
	//Чтобы получить значение на которое ссылается указатель нужно его разименовать с помощью * перед переменной указателя

	//get pointer via "new" keyword
	var newPointer = new(float32) //Создает указзатель, но ячейка памяти не nil. А значение записаное в этой яч. памяти = 0
	fmt.Printf("%T %#v %#v \n", newPointer, newPointer, *newPointer)
	*newPointer = 3 //с помощью разименования можно также изменить значение указателя на которое он ссылается
	fmt.Printf("%T %#v %#v \n", newPointer, newPointer, *newPointer)
}
