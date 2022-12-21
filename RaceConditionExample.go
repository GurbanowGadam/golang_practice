package main

import (
	"fmt"
	"sync"
)

type Hit struct {
	count int
	sync.Mutex
}

var wg sync.WaitGroup

func main() {

	hit := &Hit{}

	count := 2000

	for i := 0; i < count; i++ {
		wg.Add(1)

		go func(hit *Hit) {

			//hit.Lock()
			hit.count++
			//hit.Unlock()

			wg.Done()

		}(hit)
	}

	wg.Wait()
	fmt.Println("Olması Gereken", count, "\nŞimdi Elimizdeki ", hit.count, "\nKayıp", count-hit.count)

}
