package main

import (
	"fmt"
	"sync"
	"time"
)

func func_1(wg *sync.WaitGroup) {
	time.Sleep(5 * time.Second)
	fmt.Println("Func_1 5 saniye bekledikden sonra tamamlandi")
	wg.Done()
}

func func_2(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	fmt.Println("Func_2 2 saniye bekledikden sonra tamamlandi")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func_1(&wg)
	go func_2(&wg)
	fmt.Println("Func lar ile birlikde saniyede basladi ")
	wg.Wait()

	fmt.Println("Func lar bitti!")
}
