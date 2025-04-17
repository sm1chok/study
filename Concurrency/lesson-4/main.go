package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//baseSelect()
	gracefulShutdown()
}

//Select - block, unblock, default

func baseSelect() {

	bufferedChan := make(chan string, 3) //если разсер буффера 2, тогда обе операции неблокирующие и выполнится рандомный кейс
	bufferedChan <- "first"

	select { //похоже на case
	case str := <-bufferedChan: //неблокирующее действие (выполнится в первую очередь т.к. неблокирующее)
		fmt.Println("read", str)
	case bufferedChan <- "second": //блокирующее действие
		fmt.Println("write", <-bufferedChan, <-bufferedChan)
	}

	unbuffChan := make(chan int)

	go func() {
		time.Sleep(time.Second)
		unbuffChan <- 1
	}()

	select {
	case bufferedChan <- "third":
		fmt.Println("unblocking writing")
	case val := <-unbuffChan:
		fmt.Println("blocking writing", val)
	case tm := <-time.After(time.Millisecond * 500):
		fmt.Println("time's up")
		fmt.Println(tm)
	default:
		fmt.Println("default case")
	}

	resultChan := make(chan int)
	timer := time.After(time.Microsecond * 1000) //timer создается вне цикла, потому что внутри цикла он будет всегда обновлятся и никогда не закончится

	go func() {
		defer close(resultChan)

		for i := 0; i < 1000; i++ {

			select {
			case <-timer:
				fmt.Println("time's up")
				return //return здесь позволяет завершить всю анонимную функцию
			default:
				time.Sleep(time.Nanosecond)
				resultChan <- i + 1

			}
		}
	}()

	for v := range resultChan {
		fmt.Println(v)
	}

}

// usecase "Graceful Shutdown" - возможность завершить программу, предварительно дав ей время на выполнение каких-либо действий перед завершением
func gracefulShutdown() {
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	timer := time.After(time.Second * 10)

	for {
		select {
		case <-timer:
			fmt.Println("time's up")
			return
		case sig := <-sigChan:
			fmt.Println("Stopped by signal:", sig) // при получении сигнала (ctrl + c), программа завершается
			return
		}
	}
}
