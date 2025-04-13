package main

import (
	"fmt"
	"time"
)

func main() {
	// nilChannel()
	// unbufferedChannel()
	// bufferedChannel()
	forRange()

}

//channels (каналы) - инструмент коммуникации, который позволяет обмениваться данными между горутинами
// структурно выглядит примерно так:
/* type chan struct {
	mx sync.mutex
	buffer []T //буффер
	readers	[]Goroutines //очередь - читающая горутина
	writers	[]Goroutines // очередь - пишущая горутина
}
*/

func nilChannel() {
	var nilChannel chan int //int после задания типа переенной chan означает, что данные типа int можно отпралять и получать из нашего канала
	// длина канала - кол-во эллементов находящихся в буффере, вместимость - вместимость буффера
	fmt.Printf("Len: %d Cap: %d \n", len(nilChannel), cap(nilChannel))

	// write to nil cannel (blocks forever)
	// nilChannel <- 1 //Запись в nil канал. Т.к. у нас нету буффера значение 1 некуда поместить, программа зависает на этом моменте и мы получаем deadlock

	//read from nil channel blocks forever (deadlock)
	//<-nilChannel

	//closing nil channel will raise a panic
	// close(nilChannel())
}

// небуфферизированный - значит длина буфера равна 0 (буффер не используется)
func unbufferedChannel() {
	unbufferedChannel := make(chan int) // создание небуфферезированного канала (создается с помощью make)
	fmt.Printf("Len: %d Cap: %d \n", len(unbufferedChannel), cap(unbufferedChannel))

	//blocks untill smb reads (чтобы записать какое-то значение, в очереди должна быть читающая горутина (также и с чтением))
	// unbufferedChannel <- 1 //deadlock (cause nobody reads)

	//blocks untill smb writes
	// <-unbufferedChannel // deadlock

	//blocks on reading then writes (after 1 second)
	go func(chanForWriting chan<- int) {
		time.Sleep(time.Second)
		chanForWriting <- 1
	}(unbufferedChannel)

	value := <-unbufferedChannel
	fmt.Println(value)

	//blocks on writing then read (after 1 second)
	//Встает в очередь на чтение, потом встает в очередь на запись и сразу записывает, после читает
	go func(chanForReading <-chan int) {
		time.Sleep(time.Second)
		value := <-chanForReading
		fmt.Println(value)
	}(unbufferedChannel)

	unbufferedChannel <- 2
	//чтобы в пустой буффер можно было записать или прочитать из него, обязательно существование как минимум двух горутин. Одна на запись, другая на чтение

	// Close channel
	/*	go func() {
			time.Sleep(time.Microsecond * 500)
			close(unbufferedChannel)
		}()
		go func() {
			time.Sleep(time.Second)
			fmt.Println(<-unbufferedChannel)
		}()
		//write to closed chan - panics
		unbufferedChannel <- 3 */

	//close of closed chan -> panics
	/*close(unbufferedChannel)
	close(unbufferedChannel)*/

	//можно создать канал только для записи или чтения
	// readChan := make(<-chan int)
	// writeChan := make(chan<- int)
}

func bufferedChannel() {
	bufferedChannel := make(chan int, 2) // 2м аргументом передается размер буффера
	fmt.Printf("Len: %d Cap: %d \n", len(bufferedChannel), cap(bufferedChannel))

	//doesnt block while buffer not full
	bufferedChannel <- 1
	bufferedChannel <- 2
	fmt.Printf("Len: %d Cap: %d \n", len(bufferedChannel), cap(bufferedChannel))

	//doesnt block, while buffer not empty
	fmt.Println(<-bufferedChannel)
	fmt.Println(<-bufferedChannel)
	fmt.Printf("Len: %d Cap: %d \n", len(bufferedChannel), cap(bufferedChannel))
}

func forRange() {
	bufferedChannel := make(chan int, 3)

	numbers := []int{5, 6, 7, 8}

	//show all elements with for

	go func() {
		for _, num := range numbers {
			bufferedChannel <- num
		}
		close(bufferedChannel) // Если закрыть канал, а потом считать значение из него, мы вернем просто дефолтное значение 0
	}()

	//чтобы после закрытие не считывать дефолтные значения можно использовать следующую конструкцию
	for {
		v, ok := <-bufferedChannel // В такой конструкции вторая переменная ok показывает, закрыт какнал или открыт
		fmt.Println(v, ok)
		if !ok { //
			break
		}
	}

	//show with for range buffered
	bufferedChannel = make(chan int, 3)

	go func() {
		for _, num := range numbers {
			bufferedChannel <- num
		}
		close(bufferedChannel)
	}()

	for v := range bufferedChannel { //Здесь конструкция for range сразу считывает значения в переменную v, а также эта конструкция сама проверяет открытость или закртытость канала и считывает все значения до его закрытия
		fmt.Println("buffered", v)
	}

	//Same thing for unbudderedChannel

	unbufferedChannel := make(chan int)

	go func() {
		for _, num := range numbers {
			unbufferedChannel <- num
		}
		close(unbufferedChannel)
	}()

	for v := range unbufferedChannel {
		fmt.Println("unbuffered", v)
	}

}
