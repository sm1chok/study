package main

import "fmt"

func main() {
	//defered funcs (отложенные функции)

	// defer fmt.Println(1) // выполняется только после окончания работы основной функции. До этого момента все дефер функции лежат в стеке и выполняются в обратном порядке (т.к. это стек)
	// defer fmt.Println(2)
	// defer fmt.Println(sum(2, 3))

	// deferValues()

	fmt.Println("Exit")

	makePanic()
}

func sum(x, y int) (sum int) {
	defer func() { //эта отложенная функция выполнится после сложения (после return). Может менять именно именованные возвращаемые значения например как здесь (sum int)
		sum *= 2
	}()

	sum = x + y
	return
}

// func() {}() - это анонимная функция. После фигурных скобок ставят круглые скобки, чтобы сразу вызвать эту функцию

func deferValues() {
	for i := 0; i < 10; i++ {
		defer fmt.Println("First", i)
	}

	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Println("Second", i) // выполняется правильно т.к. с версси пo 1.22 на каждой итерации цикла создается новая переменная i, а не просто изменяется ее значение.
			//поэтому defer функция обращается каждый раз к разным переменным i, хранящим разные значения в разных ячейках памяти
		}()

	}

}

func makePanic() {
	defer func() { // для отлова паники создается дефер функция, чтобы она выполнилась в конце программы
		panicValue := recover() //для отлова используется recover() - возвращает значение паники
		fmt.Println(panicValue)
	}()

	panic("Some panic") //Вызывает панику, принимает пустой интерфейс (можно положить любой тип данных)
	fmt.Println("Jopa") //этот код не выведется т.к. находится после паники
}
