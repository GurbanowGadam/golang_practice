package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	/*go*/
	func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()

	//go
	func() {
		for i := 10; i < 25; i++ {
			fmt.Println(i)
		}
	}()

	elapsedTime := time.Since(start)

	fmt.Println("Total Time For Execution: " + elapsedTime.String())

	time.Sleep(time.Second)
}
