package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("what is today's lucky number?")

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		getLuckyNum()
	}()

	wg.Wait()
}

func getLuckyNum() {
	fmt.Println("...")

	// 占いにかかる時間はランダム
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	time.Sleep(time.Duration(r.Intn(3000)) * time.Millisecond)

	num := r.Intn(10)
	fmt.Printf("Today's your lucky number is %d!\n", num)
}
