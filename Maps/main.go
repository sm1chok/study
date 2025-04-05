package main

import "fmt"

func main() {
	// Мапа - это типа данных вида - ключ: значение (хэш таблица)
	var defaultValue map[int64]string // в скобках указывапется тип ключа мапы, а после - тип значения мапы

	fmt.Printf("Type: %T; Value: %#v;\n", defaultValue, defaultValue) //дефолтное значение
	fmt.Printf("Length: %d;\n\n", len(defaultValue))                  //длина

	//map by make (чтобы значение было не nil)
	mapByMake := make(map[string]string)
	fmt.Printf("Type: %T; Value: %#v;\n\n", mapByMake, mapByMake) //дефолтное значение

	//map by make with cap
	mapWithCap := make(map[string]string, 3)                        //через заптую указывается типа вместимость. У мап нет свойства capacity.
	fmt.Printf("Type: %T; Value: %#v;\n\n", mapWithCap, mapWithCap) //дефолтное значение

	//map by literal
	mapByLiteral := map[string]int{"Jopa": 21, "Pisya": 33}
	fmt.Printf("Type: %T; Value: %#v;\n", mapByLiteral, mapByLiteral)
	fmt.Printf("Length: %d;\n\n", len(mapByLiteral))

	//map by new
	mapByNew := *new(map[string]string) //new создает указатель на мапу, поэтому используем разименовывание
	fmt.Printf("Type: %T; Value: %#v;\n\n", mapByNew, mapByNew)

	//insert value
	mapByMake["First"] = "Sosalo"
	fmt.Printf("Type: %T; Value: %#v;\n", mapByMake, mapByMake)
	fmt.Printf("Length: %d;\n\n", len(mapByMake))

	//update value
	mapByMake["First"] = "Konchalo"
	fmt.Printf("Type: %T; Value: %#v;\n", mapByMake, mapByMake)
	fmt.Printf("Length: %d;\n\n", len(mapByMake))

	//get map value
	fmt.Println(mapByLiteral["Jopa"])

	//get deafult map value
	fmt.Println(mapByLiteral["notExistingKey"]) //Если использовать несуществующий ключ, выведет дефолтное значение

	//chek value existence
	value, ok := mapByLiteral["notExistingKey"] // при получении значения из мапы в первую переменную записывается ее значение, а во вторую записывается флаг существования true/false
	fmt.Printf("Value: %d; Is exist: %t; \n\n", value, ok)
	value, ok = mapByLiteral["Jopa"]
	fmt.Printf("Value: %d; Is exist: %t; \n\n", value, ok)

	//delete map value
	delete(mapByMake, "First")
	fmt.Printf("Type: %T; Value: %#v;\n", mapByMake, mapByMake)
	fmt.Printf("Length: %d;\n\n", len(mapByMake))

	//map iteration
	mapForIteration := map[string]int{"First": 1, "Second": 2, "Third": 3, "Fourth": 4}

	for key, val := range mapForIteration {
		fmt.Printf("Key: %s; Value: %d\n\n", key, val) // Последовательность вывода всех элементов для мапы всегда рандомна
	}

	// unique values
	users := []User{ //Создаем слайс из нашей структуры User
		{
			Id:   1,
			Name: "Huisos",
		},
		{
			Id:   45,
			Name: "Suka",
		},
		{
			Id:   57,
			Name: "Keke",
		},
		{
			Id:   45,
			Name: "Suka",
		},
	}

	uniqueUsers := make(map[int64]struct{}, len(users)) //мапа в корторую мы будем складывать только уникальные элементы

	for _, user := range users { //пробегаемся по всем элементам слайса без их id
		if _, ok := uniqueUsers[user.Id]; !ok { // проверяем есть ли внутри uniqueUsers ключ Id. Если нет (!ok = True), тогда добавляем в наш uniqueUsers
			uniqueUsers[user.Id] = struct{}{} // кладем пустую структуру в значение, т.к. при создании мапы задали ей значение пустой структуры
		}
	}

	fmt.Printf("Type: %T; Value: %#v;\n\n", uniqueUsers, uniqueUsers)

	// find by value (создаем мапу внутри которйо в качестве ключа лежит айди юзера из слайса, а в качестве значения - все значения структуры User, тоесть айди и имя)
	usersMap := make(map[int64]User, len(users))
	for _, user := range users {
		if _, ok := usersMap[user.Id]; !ok {
			usersMap[user.Id] = user
		}
	}

	fmt.Println(usersMap)
	fmt.Println(findInSlice(57, users))
	fmt.Println(findInMap(45, usersMap))
}

type User struct {
	Id   int64
	Name string
}

// функция ищет в слайсе по айди. Очень долгое время поиска нужного значения т.к. если значение последнее, мы пробегаемся по всем значениям. Сложность - O(n)
func findInSlice(id int64, users []User) *User {
	for _, user := range users {
		if user.Id == id {
			return &user
		}
	}
	return nil
}

// функция ищет в мапе по айди. время поиска - O(1)
func findInMap(id int64, usersMap map[int64]User) *User {
	if user, ok := usersMap[id]; ok {
		return &user
	}
	return nil
}
