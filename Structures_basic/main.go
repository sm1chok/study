package main

import (
	"fmt"
	"time"
)

// custom types (additional topic)
type MyString string //создание кастомного типа на основе уже существующего типа string
type MyInt int

// Strucrures
type Person struct { //что-то типа класса из ООП со своим набором семантически связанных значений
	name string
	age  int
}

func main() {
	/*var slovo MyString = "Hello Joinpa" // создаем переменную со своим кастомным типом
	fmt.Printf("%T %#v \n", slovo, slovo)

	num := 6
	customNum := MyInt(6) // при присваивании для использования кастомного типа делать так, иначе автоматом будет тип int
	fmt.Printf("%T %#v \n", num, num)
	fmt.Printf("%T %#v \n", customNum, customNum)*/

	//create default structure object
	var Artyom Person                       // создание объекта структуры Person
	fmt.Printf("%T %#v \n", Artyom, Artyom) //Здесь присвоились дефлотные значения дял полей name и age

	Artyom = Person{} //просто второй вариант создания объекта с дефлотными значениями
	fmt.Printf("%T %#v \n", Artyom, Artyom)

	// fields accessing
	Artyom.name = "Artyom" // через точку обращаемся к полю структуры
	Artyom.age = 22
	fmt.Println(Artyom)

	//instatnt structure init

	Ulyana := Person{ //инициализируем нашу структуру сразу при создании новой переменной
		name: "Ulyana",
		age:  20,
	}
	fmt.Println(Ulyana)

	//init structures without field names
	Sergey := Person{"Sergy", 60} //Значения буду присваиваться в том порядке, который задан при создании шаблона структуры
	fmt.Println(Sergey)

	// field accessing through pointers

	pSergey := &Sergey
	fmt.Println((*pSergey).age)
	fmt.Println(pSergey.age) //выводит также фактическое значение поля age, а не адресс ячейки памяти хоть это и указатель (некий синтаксический сахар)

	//create pointer to struct
	pRafka := &Person{"Rafka", 2} //Записываем структуру не в переменную, а в указатель
	fmt.Println(pRafka)

	//Anonimus structs
	unnamedStruct := struct { //анонимная структура сразу привязанная к переменной не как type Person
		name, surname, birthDate string
	}{
		name:    "NoName",
		surname: "Pidarasovich",
		//birthDate: fmt.Sprintf("%s", time.Now()),   //sprintf помогает привести к другому типу, %s значит string
		birthDate: time.Now().String(), // второй способ вывода времени и приведеняи его к строке прои помощи метода .String
	}
	fmt.Println(unnamedStruct)
}
