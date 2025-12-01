package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("what is today's lucky number?")

	c := make(chan int)
	go getLuckyNum(c)

	num := <-c

	fmt.Printf("Today's your lucky number is %d!\n", num)

	// 使い終わったチャネルはcloseする
	close(c)
}

func getLuckyNum(c chan<- int) {
	fmt.Println("...")

	// ランダム占い時間
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	time.Sleep(time.Duration(r.Intn(3000)) * time.Millisecond)

	num := r.Intn(10)
	c <- num
}
