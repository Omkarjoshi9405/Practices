package main

import (
	"fmt"
	"sync"
)

func odd(odd1 chan bool, even1 chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 49; i += 2 {
		<-odd1
		fmt.Println(i)
		even1 <- true
	}
}

func even(odd1 chan bool, even1 chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 50; i += 2 {
		<-even1
		fmt.Println(i)
		odd1 <- true
	}
}

func main() {
	var wg sync.WaitGroup
	odd1 := make(chan bool)
	even1 := make(chan bool)

	fmt.Println("channel created")

	wg.Add(2)

	go odd(odd1, even1, &wg)
	go even(odd1, even1, &wg)

	odd1 <- true

	wg.Wait()

	fmt.Println("Promgram ended")

}
