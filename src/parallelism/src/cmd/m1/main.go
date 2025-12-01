package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("what is today's lucky number?")
	
	go getLuckyNum()

	time.Sleep(time.Second * 5)
}

func getLuckyNum() {
	fmt.Println("...")

	// 占いにかかる時間はランダム
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	time.Sleep(time.Duration(r.Intn(3000)) * time.Millisecond)

	num := r.Intn(10)
	fmt.Printf("Today's your lucky number is %d!\n", num)
}