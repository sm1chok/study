package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//withoutWait()
	//withWait()
	//wrongAdd()
	//withoutMutex()
	//withMutex()
	readWithMutex()
	readWithRWMutex()
}

// WaitGroup - механизм ожидания завершения группы задач

func withoutWait() {
	for i := 0; i < 10; i++ {
		go fmt.Println(i)
	}

	fmt.Println("Exit")
}

func withWait() {
	var wg sync.WaitGroup
	// числа буду выводитсья в рандомном порядке т.к. планировщик задач сам решает в каком порядке выполнять горутины
	for i := 0; i < 10; i++ {
		wg.Add(1) // Принимает кол-во задач, которые надо принять в нашу группу. Здесь на каждую итерацию добавляем по 1 вейтгруппе
		// wg.Add нужно добавлять до самой горутины (внутри горутины - нельзя, т.К. горутина может не выполниться)

		go func(i int) {
			fmt.Println(i)
			wg.Done() // После выполнения задачи нужно прописать Done, чтобы дать программе понять, что мы выполнили поставленную задачу
		}(i)
	}

	wg.Wait() // блокирует основную горутину  и ждет пока в нашей waitGroup не останется невыполненных задач
	// таким образом WaitGroup позволяет дождать выполнения всех горутин
	fmt.Println("exit")
}

func wrongAdd() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {

		go func(i int) {
			wg.Add(1) //Нельязя использовать Add внутри горутины. Горутины могут не успеть выполнитсья, а когда цикл дойдет до wg.Wait() в этот момент ни одна таска может не успеть добавиться в вейтгруппу и функция завершится
			// нужно добавлять Add заранее до горутины
			defer wg.Done()

			fmt.Println(i)
		}(i)
	}
	wg.Wait()
	fmt.Println("exit")
}

// Data Race (Гонка) - обращение к одним и тем же данным из разных программ/тредов/горутин, где одно из обращений - запись

func withoutMutex() {
	start := time.Now()
	var counter int
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)
			counter++ //Здесь из-за конкурентности может произойти ситуация 555 + 1 = 556 // 555 + 1 = 556
			//То есть сразу две горутины приняли в себя значение 555 и выполнившись выдали 556, из-за этого в конце выводится не 1000, а меньше

		}()
	}

	wg.Wait()
	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

// Mutex и RWMutex - механизм получения исключительной блокировки (для решения проблемы Data Race)
// m.Lock() / m.Unlock()

func withMutex() {
	start := time.Now()
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex // создали переменную для mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)

			// Этот мьютекс доступен для всех горутин, но используется в моменте лишь в одной.
			mu.Lock()
			counter++ //Код между Lock и Unlock блокируется для всех горутин кроме одной, позволяя вносить изменения только одной горутине
			mu.Unlock()
			//Остальные 999 горутин ждут, пока какая-нибудь из ник не получит доступ к mutex. То есть все операции с увеличением значения счетчика буду происходить последовательно
		}()
	}

	wg.Wait()
	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func readWithMutex() {
	start := time.Now()
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	//все 100 действий будут выполняться последовательно т.к. mutex одинаково работает и для записи и для чтения
	wg.Add(100)
	for i := 0; i < 50; i++ {
		//50 операций чтения
		go func() {
			defer wg.Done()

			mu.Lock()

			time.Sleep(time.Nanosecond)
			_ = counter

			mu.Unlock()

		}()

		//50 операций записи
		go func() {
			defer wg.Done()

			mu.Lock()

			time.Sleep(time.Nanosecond)
			counter++

			mu.Unlock()

		}()
	}

	wg.Wait()
	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

// при помощи
func readWithRWMutex() {
	start := time.Now()
	var counter int
	var wg sync.WaitGroup
	var mu sync.RWMutex

	wg.Add(100)
	for i := 0; i < 50; i++ {
		//50 операций чтения
		go func() {
			defer wg.Done()

			mu.RLock() //отдельный mutex для чтения. Это не эксклюзивная блокировка, позволяет читать все значения сразу

			time.Sleep(time.Nanosecond)
			_ = counter

			mu.RUnlock()

			//операции чтения можно выполнять быстрее с помощью RWMutex
		}()

		//50 операций записи
		go func() {
			defer wg.Done()

			mu.Lock() // Здесь все еще эксклюзивная блокировка и при ее вызове все блокировки для чтения также ждут завершения этой блокировки

			time.Sleep(time.Nanosecond)
			counter++

			mu.Unlock()

		}()
	}

	wg.Wait()
	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}
