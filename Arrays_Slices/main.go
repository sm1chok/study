package main

import "fmt"

func main() {
	//arrays()
	slices()
}

func arrays() {
	var intArr [3]int // так задается массив. внутри квадратных скобок - кол-во элементов массива
	fmt.Printf("Type: %T; Value: %#v \n", intArr, intArr)

	intArr[0] = 5 //изменение элемента массива с индексом 0
	intArr[1] = 9
	fmt.Printf("Type: %T; Value: %#v \n", intArr, intArr)

	intArrShort := [3]int{1, 2, 23} //сокращенный вариант создания массива и сразу инициализация его элементов
	fmt.Printf("Type: %T; Value: %#v \n", intArrShort, intArrShort)

	people := [2]Person{ // создание массива из объектов
		{
			Name: "Joinpa",
			Age:  21,
		},
		{
			Name: "Sisya",
			Age:  26,
		},
	}
	fmt.Printf("Type: %T; Value: %#v \n", people, people)

	// можно создаь массив не задавая явное кол-во жлементов в нем. go сам поймет, сколько эл-тов мы добавили

	strArr := [...]string{"Jija", "Kaka", "Tripoponi", "Gutalini"}
	fmt.Printf("Type: %T; Value: %#v \n", strArr, strArr)

	//Для массива, его длина и вместимость всегда одинаковы и определяются при его объявлении
	fmt.Printf("Length:	%d; Capacity: %d;\n", len(strArr), cap(strArr)) // len(x) и cap(x) - выводят длину и вместимость массива соответственно. Длина - сколько занчений в массиве сейчас. Фактически всегда равны с capacity

	for i := 0; i < len(strArr); i++ {
		fmt.Printf("Index: %d Value: %s\n", i, strArr[i])
	}

	// сокращенная запись for для перебора элементов массива

	for idx, value := range intArr { // В 1ю переменную (idx) приходит индекс, во 2ю приходит значение элемента (value)
		fmt.Printf("Index: %d Value: %d\n", idx, value)
	}
}

func slices() {
	// в массиве внутри квадратных скобок все присутствует какое-то число, а в слайсах они всегда пустые
	var defaultSlice []int // слайсы (динамические массивы) могут содержать любое количество элементов 
	fmt.Printf("Type: %T; Value: %#v \n", defaultSlice, defaultSlice)
	fmt.Printf("Length: %d; Capacity: %d;\n", len(defaultSlice), cap(defaultSlice))
	fmt.Println("__________________________________")

	sliceLiteral := []string {"Chichiski", "Govno", "Gergerger"}
	fmt.Printf("Type: %T; Value: %#v \n", sliceLiteral, sliceLiteral)
	fmt.Printf("Length: %d; Capacity: %d;\n", len(sliceLiteral), cap(sliceLiteral))
	fmt.Println("__________________________________")

	// с помощью make можно создать слайс у кторого будут разные значения length и capacity (length <= cap)
	sliceByMake := make([]int, 0, 5) // 0 - длина, 5 - вместимость
	fmt.Printf("Type: %T; Value: %#v \n", sliceByMake, sliceByMake)
	fmt.Printf("Length: %d; Capacity: %d;\n", len(sliceByMake), cap(sliceByMake))
	fmt.Println("__________________________________")

	//Добавляем новые значения в слайс
	sliceByMake = append(sliceByMake, 1, 2, 3, 4, 5)
	fmt.Printf("Type: %T; Value: %#v \n", sliceByMake, sliceByMake)
	fmt.Printf("Length: %d; Capacity: %d;\n", len(sliceByMake), cap(sliceByMake))
	fmt.Println("__________________________________")

	sliceByMake = append(sliceByMake, 6) //здесь добавляется один жлемент за пределами ваместимости, поэтому аместимость автоматически увеличится в 2 раза
	fmt.Printf("Type: %T; Value: %#v \n", sliceByMake, sliceByMake)
	fmt.Printf("Length: %d; Capacity: %d;\n", len(sliceByMake), cap(sliceByMake))
	fmt.Println("__________________________________")

	for idx, value := range sliceByMake{
		fmt.Printf("Index: %d; Value: %d;\n", idx, value)
	}
}

type Person struct{
	Name string
	Age  int
}
