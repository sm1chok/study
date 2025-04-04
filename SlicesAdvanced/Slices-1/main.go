package main

import "fmt"

func main() {
	//variadicFunctions()
	//sliceToArray()
	//passToFunction()
	//sliceWithNew()
	//getSlice()
	//copySlice()
	deleteElement()
}

// функции с неограниченным числом параметров
func variadicFunctions() {
	showAllElements(1, 2, 3)
	showAllElements(1, 2, 3, 4, 5, 6, 7)

	firstSlice := []int{5, 6, 7, 9}
	secondSlice := []int{2, 2, 8}

	showAllElements(firstSlice...) // чтобы передать в функцию с неограниченным числом параметров слайс, нкжно написать в аргументе sliceName... (раскладывает слайс на элементы)

	newSlice := append(firstSlice, secondSlice...) //добавляет все элементы одного слайса в другой
	showAllElements(newSlice...)
}

/* слайс - это по факту структура, которая выглядит так:
type _slice struct {
	elements unsafe.Pointer	// Указатель на массив с определенным типом данных
	len int					// Количество элементов
	cap int					// Вместимость слайса
}
*/

// можно конвертировать слайсы в указатель на массив
func sliceToArray() {
	slaysik := []int{1, 2}
	fmt.Printf("Type: %T; Value: %#v;\n", slaysik, slaysik)
	fmt.Printf("Length: %d; Capacity: %d;\n", len(slaysik), cap(slaysik))

	arraychik := (*[2]int)(slaysik) // конвертируем наш слайс в указатель на массив, который содержит элементы слайса
	fmt.Printf("Type: %T; Value: %#v;\n", arraychik, arraychik)
	fmt.Printf("Length: %d; Capacity: %d;\n", len(arraychik), cap(arraychik))
}

func passToFunction() {
	initialSLice := []int{2, 3}
	fmt.Printf("Type: %T; Value: %#v;\n", initialSLice, initialSLice)
	fmt.Printf("Length: %d; Capacity: %d;\n", len(initialSLice), cap(initialSLice))

	chageValue(initialSLice) // значение слайса поменяется и снаружки функции, а не только внутри, т.к. в слайсе хранится указатель на массив данных. Мы как бы сразу передаем в аргумент функции указатель на массив
	fmt.Printf("Type: %T; Value: %#v;\n", initialSLice, initialSLice)
	fmt.Printf("Length: %d; Capacity: %d;\n", len(initialSLice), cap(initialSLice))

	newSlice := append(initialSLice, 1337) // при превышении вместимости слайса, его вместимость capacity увеличивается в два раза и создает новый указатель, ссылающийся на новый массив в памяти устройства
	fmt.Printf("Type: %T; Value: %#v;\n", newSlice, newSlice)
	fmt.Printf("Length: %d; Capacity: %d;\n", len(newSlice), cap(newSlice))

	newSlice2 := appendValues(newSlice)

	//Знвчение newSlice не изменится т.к. внутри функции appendValues, создалась ссылка на новый слайс. А newSlice2 соответственно изменится т.к. туда мы передали ссылку на наш новый слайс из функции
	//То есть при добавлении значений в слайс через функцию, всегда создается новый слайс (новая ссылка на новый массив)
	// чтобы наш изначчальный слай изсенился, можно было бы сделать так: newSlice := appendValues(newSlice)
	fmt.Printf("Type: %T; Value: %#v;\n", newSlice, newSlice)
	fmt.Printf("Length: %d; Capacity: %d;\n", len(newSlice), cap(newSlice))

	fmt.Printf("Type: %T; Value: %#v;\n", newSlice2, newSlice2)
	fmt.Printf("Length: %d; Capacity: %d;\n", len(newSlice2), cap(newSlice2))
}

func sliceWithNew() {
	slicePointer := new([]int) //создает пустой УКАЗАТЕЛЬ на слайс (nil). Это практически никогда не используется, лучше использовать make
	fmt.Printf("Type: %T; Value: %#v;\n", slicePointer, *slicePointer)
	fmt.Printf("Length: %d; Capacity: %d;\n", len(*slicePointer), cap(*slicePointer))

	newSlice := append(*slicePointer, 1, 2, 3)
	fmt.Printf("Type: %T; Value: %#v;\n", newSlice, newSlice)
	fmt.Printf("Length: %d; Capacity: %d;\n", len(newSlice), cap(newSlice))
}

func getSlice() {
	intArr := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("Type: %T; Value: %#v;\n\n", intArr, intArr)

	intSlice := intArr[1:3] // получаем слайс из массива. Не включаетс в себя последний элемент. То есть мы получим элементы с индексом 1 и 2
	fmt.Printf("Type: %T; Value: %#v;\n", intSlice, intSlice)
	fmt.Printf("Length: %d; Capacity: %d;\n\n", len(intSlice), cap(intSlice)) //Capacity высчитываетя без начальных элементов. Тоесть при intArr[1:3] мы игнорируем элемент с нулевым индексом, а все остальные берем и формируем вместимость

	fullSlice := intArr[:] // Позволяет передать в слайс все элеменыт массива. Также можно брать [:x] или [x:] (от начала до границы или от границы до конца)
	fmt.Printf("Type: %T; Value: %#v;\n", fullSlice, fullSlice)
	fmt.Printf("Length: %d; Capacity: %d;\n\n", len(fullSlice), cap(fullSlice))

	sliceFromSlice := fullSlice[:3]
	fmt.Printf("Type: %T; Value: %#v;\n", sliceFromSlice, sliceFromSlice)
	fmt.Printf("Length: %d; Capacity: %d;\n\n", len(sliceFromSlice), cap(sliceFromSlice))

	//т.к. все созданные слайсы ссылаются на массив intArr, при изменении его элементов, будут менться и все слайсы
	intArr[2] = 500
	fmt.Printf("Type: %T; Value: %#v;\n", intArr, intArr)
	fmt.Printf("Type: %T; Value: %#v;\n", intSlice, intSlice)
	fmt.Printf("Type: %T; Value: %#v;\n", fullSlice, fullSlice)
	fmt.Printf("Type: %T; Value: %#v;\n", sliceFromSlice, sliceFromSlice)

}

func copySlice() {
	destination := make([]string, 0, 2)
	source := []string{"Jopa", "Popa", "Pidor"}

	//функция copy - копирует элементы из одного массива в другой и возвращает количество скопированных элементов в виде int
	fmt.Printf("Copied: %d\n", copy(destination, source)) // копируется число элементо равное len(destination). Т.к. слайс в который мы копируем - пустой (len(destination) = 0), в destination будет скопировано ноль элементов
	fmt.Printf("Type: %T; Value: %#v;\n", destination, destination)
	fmt.Printf("Length: %d; Capacity: %d;\n\n", len(destination), cap(destination))

	destination = make([]string, 2, 3) //создаем слайс с двумя элементами, чтобы можно было скопировать в него 2 значения
	fmt.Printf("Copied: %d\n", copy(destination, source))
	fmt.Printf("Type: %T; Value: %#v;\n", destination, destination)
	fmt.Printf("Length: %d; Capacity: %d;\n\n", len(destination), cap(destination))

	destination = make([]string, len(source)) //если не указать вместимость, она по дефолту равняется длине
	fmt.Printf("Copied: %d\n", copy(destination, source))
	fmt.Printf("Type: %T; Value: %#v;\n", destination, destination)
	fmt.Printf("Length: %d; Capacity: %d;\n\n", len(destination), cap(destination))

	// чтобы сделать нормальную копию можно воспользоваться следующей связкой append + make
	rightCopy := append(make([]string, 0, len(source)), source...)
	fmt.Printf("Type: %T; Value: %#v;\n", rightCopy, rightCopy)
	fmt.Printf("Length: %d; Capacity: %d;\n\n", len(rightCopy), cap(rightCopy))

}

func deleteElement() {
	slice := []int{1, 2, 3, 4, 5}
	i := 2 // чтобы удалить элемент нужно заранее знать его индекс в нашем случае это 2
	fmt.Printf("Type: %T; Value: %#v;\n", slice, slice)
	fmt.Printf("Length: %d; Capacity: %d;\n\n", len(slice), cap(slice))

	// 1 способ удаления with append
	slice = append(slice[:i], slice[i+1:]...) // берем все значения до i, пропускаем i-ый элемент и добавляем все что после него. Если перезаписать в новый слайс, то исходынй слайс сломается
	fmt.Printf("Type: %T; Value: %#v;\n", slice, slice)
	fmt.Printf("Length: %d; Capacity: %d;\n\n", len(slice), cap(slice))

	// 2 способ удаления with copy
	slice = []int{1, 2, 3, 4, 5}
	slice = slice[:i+copy(slice[i:], slice[i+1:])]
	fmt.Printf("Type: %T; Value: %#v;\n", slice, slice)
	fmt.Printf("Length: %d; Capacity: %d;\n\n", len(slice), cap(slice))

}

// ...int - значит что функция может принимать сколько кгодно элементов типа int
func showAllElements(values ...int) {
	for _, val := range values {
		fmt.Printf("Value: %d\n", val)
	}
	fmt.Println("---------------------")
}

func chageValue(slice []int) {
	slice[1] = 28
}

func appendValues(slice []int) []int {
	slice = append(slice, 4, 5)
	fmt.Printf("Type: %T; Value: %#v;\n", slice, slice)
	fmt.Printf("Length: %d; Capacity: %d;\n", len(slice), cap(slice))
	return slice
}
