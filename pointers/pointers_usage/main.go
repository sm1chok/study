package main

import "fmt"

// Where to use pointers?
func main() {
	//a) side effect
	num := 3
	square(num)
	fmt.Println(num) //значение num не изменится, т.к. в квадрат возводится локальная пременнная num функции square

	squarePointer(&num) //Кладем в аргумент функции адрес ячейки нашего значения num
	fmt.Println(num)

	//b) empty value flag (chek)
	var wallet1 *int                // создаем nil указатель как пример несуществующего кошелька
	fmt.Println(hasWallet(wallet1)) //в аргумент передается адрес т.к. wallet1 - указатель

	wallet2 := 100
	fmt.Println(hasWallet(&wallet2)) // Передаем адрес при помощи &

	wallet3 := 0
	fmt.Println(hasWallet(&wallet3))

}

func square(num int) {
	num *= num
}

func squarePointer(num *int) { // принимает в аргумент указатель интовой переменной
	*num *= *num //возводим в квадрат с использованием разименования, чтобы возвести именно значение, лежащее по данному в аргументе адресу
}

func hasWallet(money *int) bool { //Функция проверяет, лежит ли по полученному в аргумент адресу, какое либо значение или эта ячейка памяти пустая
	return money != nil // суть именно в проверке не баланса условного кошелька, а в проверке его существования вообще
}
