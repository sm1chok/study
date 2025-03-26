package main

import (
	"fmt"
	"time"
)

func main() {
	//increment / decrement
	i := 0
	i += 2
	i++ //увеличение на еденицу/ Так же работает с минусом

	// fmt.Println(VzrivJopi(3))

	var j int
	for j = 0; j < 3; j++ {
		//fmt.Println(j)
	}

	sum := 0
	for sum < 20 {
		sum += 4
		//fmt.Println(sum)
	}

	// бесконечный цикл
	// for {
	// 	fmt.Println("AHAHAHAHAHAHAHA")
	// }

	for z := 0; z <= 10; z++ {
		if z%2 == 1 {
			continue // скипает все действия после и выбивает в начало цикла
		}
		fmt.Println(z)
	}

Label1: //Метка прикрепленная к внешнему циклу
	for i := 0; i <= 10; i++ {
		for j := 0; j <= 5; j++ {
			fmt.Println("I:", i, "J:", j)
			if i >= 5 {
				continue Label1 //Теперь continue выбрасывает нас в цикл, к которому прикреплена метка
			}
			fmt.Println("Puk")
		}
	}

	for i := 0; i <= 200; i++ {
		fmt.Println(i)
		if i >= 10 {
			break // в отличии от continue заканчивает цикл полностью (также действуют метки)
		}
	}

}

func VzrivJopi(k int) string {
	for j := k; j > 0; j-- {
		fmt.Println("До взрыва жопы:", j)
		time.Sleep(1 * time.Second)
	}
	return "Boom!"
}
