package main

import (
	"fmt"
	"math/rand"
)

const min, max = 1, 5

func main() {
	value := rand.Intn(max-min) + min //генерит случайное число от 0 до 4

	//base switch
	switch value {
	case 1:
		fmt.Println("The value is one")
	case 2, 3:
		fmt.Println("The value is two or three")
	case GetFour():
		fmt.Println("value FOUR")
	default:
		fmt.Println("Do not know which value is it")
	}

	//В других ЯП выполняются все кейсы если не прописать в конце кейса break. В go после выполненичя одного кейса, мы выходим из конструкции switch
	// Но есть команда fallthrough, которая указывается в кейсе и после выполнения этого кейса, выполнится следующий кейс
	//switch with local variable defenition
	switch num := rand.Intn(max-min) + min; num { //Переменная для кейсов определяется сразу в теле свитча перед его началом
	case 1:
		fmt.Println("The num is one")
	case 2, 3:
		fmt.Println("The num is two or three")
	case GetFour():
		fmt.Println("FOUR")
		fallthrough // После выполнения этого кейса, также выполнится следующий кейс
	case 10:
		fmt.Println("WTF")
	default:
		fmt.Println("Do not know which num is it")
	}

	//switch without condition
	switch {
	case value > 2:
		fmt.Println("Value is more than two")
	case value < 2:
		fmt.Println("Value is less than two")
	default:
		fmt.Println("Value is two")

	}
}

// Просто функция, возвращающая 4
func GetFour() int {
	fmt.Println("4itire epta") // Эта строка выведется в лбос случае, если выполнятся кейсы после кейса с этой функцией, т.к. свитч сначала вычисляет все значения кейсов а потом сравнивает значения
	return 4
}
