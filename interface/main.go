package main

import "fmt"

func main() {
	//interfaceValues()
	PolymorphismAndAssertion()
}

//Значение интерфейсного типа хранит в себе другие значения неинтерфейсного типа а например string, int, bool etc.
// таким образо значение интерфейсного типа хранит в себе информацию о конкретном типе (int) данных и о значении этого типа данных (26)

func interfaceValues() {
	var runner Runner
	fmt.Printf("%T %#v \n", runner, runner) //мы не задали конкретного значения и типа поэтому выводится дефолтные значения (nil)

	if runner == nil {
		fmt.Println("Runner is nil")
	}

	// runner = int64(1) - нельзя т.к. в типе int64 нет метода Run() поэтому он не имплементирует интерфейс Runner
	// fmt.Println(runner.Run()) //будет выдавать ошибку, так как у интерфейсного занчения без конкретного типа и значения нельзя вызвать методы его интерфейса

	var unnamedRunner *Human // для примера создаем указатель структуры Human, чтобы значение было nil, а тип был
	fmt.Printf("%T %#v \n", unnamedRunner, unnamedRunner)

	runner = unnamedRunner //Теперь значение интерфейсного типа приобрело знание о типе данных, но все еще нет конкретного значения
	fmt.Printf("%T %#v \n", runner, runner)
	if runner == nil {
		fmt.Println("Runner is nil") //Эта строка не выведется т.к. у нас уже есть как минимум тип нашего интерфейсного значения
	}

	// fmt.Println(runner.Run()) // также нельзя вызвать методы интерфейса т.е. недоопределено значение

	namedRunner := &Human{"Joinpa"}
	fmt.Printf("%T %#v \n", namedRunner, namedRunner)

	runner = namedRunner
	fmt.Printf("%T %#v \n", namedRunner, namedRunner) //Теперь у интерфейсного значения есть и конкретный тип и конкретное значение

	fmt.Println(runner.Run()) // Наконец можно вызвать метод Run() интерфейса Runner у его интерфейсного значения т.к. определены и тип и значение

	//empty interface (чтобы имплементировать не нужно реализовывать никаких методов)
	var emptyInterface interface{} = namedRunner
	fmt.Printf("%T %#v \n", emptyInterface, emptyInterface) //можно положить оъект с какими-то методами

	emptyInterface = int64(12) // можно положить данные любого типа как например int64 или bool т.к. этот интервейс по дефолту имплементирован
	fmt.Printf("%T %#v \n", emptyInterface, emptyInterface)

	emptyInterface = bool(true)
	fmt.Printf("%T %#v \n", emptyInterface, emptyInterface)
	//то есть например можно положить в аргумент функции значение ЛЮБОГО типа, если аргументу задать как тип пустой интерфейс
}

// Указание в аргументе функции вместо типа - интерфейса, позволяет нам не ограничивать функцию только определенным типом значений
// а дает возможность использовать в качестве аргумента любое значение имплементирующее указанный интерфейс
// это является полиморфизмом

// Так например пустой интерфейс позволяет вводить значение с любым типом, а не пустой только значения имплементирующие данный интерфейс
func PolymorphismAndAssertion() {
	var runner Runner
	fmt.Printf("Type: %T Vlue: %#v \n", runner, runner)

	mike := &Human{"Mike"}
	runner = mike
	Polymorphism(runner)

	donald := &Duck{"Donald", "Duck"}
	runner = donald
	Polymorphism(runner) // функция принимает в аргумент значения с разными типами 1) Human 2) Duck, но так как они оба имплементируют интерфейс - все заебись

	//Чтобы понять какой именно тип значения пришел в нашу функцию (Human или Duck) можно использовать Type Assertion
	//Это позволяет в зависимости от полученного типа, выполнить разные действия для разных типов значений
	//То есть дял утки - одно, а для человека - другое

	runner = mike

	typeAssertion(runner)
}

//////////////////////////////////////////////////////////////////////////////////////////////

type Runner interface { //интерфейс создается похожим с структурой образом
	Run() string // в фигурных скобках передаются различные методы нашего инфтерфейса
}

type Flyer interface {
	Fly() string
}

type Swimmer interface {
	Swim() string
}

type Ducker interface { // интерфейсы можно встраивать в другие интерфейсы
	Runner
	Flyer
	Swimmer
}

type Human struct {
	Name string
}

type Duck struct {
	Name    string
	Surname string
}

func (d Duck) Run() string {
	return fmt.Sprintf("Бежит утка %s", d.Name)
}

func (h Human) Run() string { //Чтобы имплементировать интерфейс, нужно просто реализовать все его методы в значении или структуре
	return fmt.Sprintf("Человек %s бегает", h.Name) // то есть для структуры Human объявили метод Run, значит мы заимплементировали интерфейс Runner этой структурой
}

func Polymorphism(runner Runner) { //создаем функцию, которая вместо конкретного типа имеет интерфейс Runner, поэтому у значения не обязательно должен быть конкретный тип, а оно просто должно имплементировать интерфейс
	fmt.Println(runner.Run()) //то есть чтобы удовлетворить данной функции, аргумент должен имплементировать интерфейс Runner
}

func typeAssertion(runner Runner) {
	fmt.Printf("Type: %T Vlue: %#v \n", runner, runner)
	//type assertion конструкция. Мы проверяем если тип runner (тип указанный в скобках) соответствует типу Human и в зависимости от этого ok = true/false
	// Если ok = true, в переменную human запишется конкретное значение интерфейсного значения runner. если ok = false - блок скипается
	if human, ok := runner.(*Human); ok {
		fmt.Printf("Type: %T Vlue: %#v \n", human, human)
		human.writeCode() //human теперь явялется указателем на структуру Human и мы можем вызывать его методы

	}
	if duck, ok := runner.(*Duck); ok {
		fmt.Printf("Type: %T Vlue: %#v \n", duck, duck)
		fmt.Println(duck.Fly())

	}
	// Конструкцию из if выше можно записать короче и проще с помощью switch

	switch v := runner.(type) { // в переменную v кладется конкретное значение, если type совпадает
	case *Human:
		fmt.Println(v.Run())
	case *Duck:
		fmt.Println(v.Swim())
	default:
		fmt.Printf("Type: %T Vlue: %#v \n", v, v)

	}
}

func (h Human) writeCode() {
	fmt.Println("Человек пишет код")
}

func (d Duck) Fly() string {
	return "Утка летает"
}

func (d Duck) Swim() string {
	return "Утка плавает"
}
