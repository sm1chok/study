package main

import "fmt"

func main() {
	definition()
}

type Square struct {
	Side int
}

func (s Square) Perimetr() { //метод создается как обычная функция но после func в скобках указывается ресивер (структура или тип для которого мы создаем метод)
	fmt.Printf("%T %#v \n", s, s) // с помощью перменной s можно обращаться к данным нашкей структуры s.Side = ...
	fmt.Printf("Пириметр фигуры: %d \n", s.Side*4)
}

func (s *Square) Scale(multiplyer int) { //Здесь ресивер не типа value, а типа pointer
	fmt.Printf("%T %#v \n", s, s)
	s.Side *= multiplyer
	fmt.Printf("%T %#v \n", s, s)
}

func (s Square) WrongScale(multiplyer int) { //Этот метод не будет менять параметры нашей переменной, вне этого метода
	fmt.Printf("%T %#v \n", s, s) //так как value значение в отличии от поинтера просто копируется в этот метод, изменяется, а с значением снаружи метода ничего не происходит
	s.Side *= multiplyer
	fmt.Printf("%T %#v \n", s, s)
}

func definition() {
	square := Square{4} //создаем объект структуры Square с заданной в аргументе стороной
	pSquare := &square  // поинтер на наш объект
	square2 := Square{2}

	// square.Perimetr() //метод вызывается так: struct.method
	square2.Perimetr()
	// pSquare.Scale(10)

	// также в go ничего не мешает у поинтера вызывать value reciver, а у обычного значения метод с pointer reciver (синтаксический сахар)
	pSquare.Perimetr() //идентично (*pSquare).Perimetr()
	square.Scale(5)    //идентично (&square).Scale(5). это значение поменялось не только внутри метода т.к. использовали поинтер ресивер

	square.WrongScale(5) // значение увеличится только внутри метода
	fmt.Println(square)  //поэтому это значение все также равно 20
	fmt.Println(pSquare) //это значнеие также осталось не тронутым

	//Вывод: value ресивер меняет значения только внутри метода, а с поинтер ресивером можно изменять значения и за пределами метода
	//Type не должен являться указателем, чтобы можно было добавить метод

	//Метод можео добавить для T(type) или *T при условиях:
	// 1) Т находится в данном пакете
	// 2) Т не является указателем
	// 3) Т - конкретнйы не интерфейсный тип (type puki interface)
	// 4) T - не является встроенным типом
	// 5) T является объявленным типом
}
