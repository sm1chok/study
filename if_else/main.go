package main

import "fmt"

func main() {
	age := 20

	//Basic if
	if age < 18 {
		fmt.Println("Joinpa tebe a ne club")
	}

	//Short syntax if
	if isChild := isChildren(age); isChild == true { //Создали переменную с значением функции и сразу сравнили с true.
		fmt.Println("Ai maloletka oi maloletka") //Переменная isChild доступна только в блоке if
	}

	//if_else_elseif
	if age < 18 {
		fmt.Println("Zaebal maloy")
	} else if age == 18 {
		fmt.Println("Proneslo proneslo")
	} else {
		fmt.Println("leeeeesgo")
	}

	// && логическое умножение (и)
	if age >= 6 && age <= 18 {
		fmt.Println("Zdarova, Shkila")
	}

	// || логическое сложение (или)
	if age == 14 || age == 20 || age == 40 {
		fmt.Println("Change your passport bitch")
	}

	if !isChildren(age) {
		fmt.Println("Old daun")
	}

	if age == 99 {
		fmt.Println("huita")
	}
}

func isChildren(age int) bool {
	return age < 18
}
