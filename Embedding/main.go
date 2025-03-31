package main

import "fmt"

func main() {
	Colliding()
}

/*
func Shadowing() {
	person1 := Businessman{Person{"Bobby", 23}, 20000, "Jake"}
	fmt.Printf("%T %#v \n", person1, person1)

	fmt.Println(person1.Person.Age) //Длинная запись
	fmt.Println(person1.Age)        // Укороченная. У типа businessman нету поля Age, но у встроенного Person - есть
	fmt.Println(person1.Money)

	person1.PrintName() //так метод вызовется для типа businessman и выведет Jake (происходит shadowing переменной внутри типа Person)
	//но если бы у businessman не было поля name, то вызывая метод PrintName, мы бы все равно выводили name но уже из встроенного типа Person
	person1.Person.PrintName() // так метод зайдет уже непостредственно в часть структуры встроенной из Person и выведет Bobby
}
*/

func Colliding() {
	person1 := Businessman{Person{"Bobby", 23}, 20000, "Jake", WorkExp{"Manager", 7}}
	fmt.Printf("%T %#v \n", person1, person1)
	//fmt.Println(person1.Age) // го не может определить какой эйдж вытщаить т.к. оба поля эйдж находятся на одинаковой глубине
	//первый Age внутри Person второй Age внутри WorkExp. Это называется colliding
	fmt.Println(person1.Person.Age)
	fmt.Println(person1.WorkExp.Age) // Единственный вариант избежать colliding указать путь до Age явно

}

type Person struct {
	Name string
	Age  int
}
type Businessman struct {
	Person //Можно встраивать один тип в другой и тогда у него появляются все его поля значений. Называется композицией
	Money  int
	Name   string
	WorkExp
}

func (b Businessman) PrintName() {
	fmt.Println(b.Name)
}

func (p Person) PrintName() {
	fmt.Println(p.Name)
}

type WorkExp struct {
	Name string
	Age  int
}
