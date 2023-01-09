package main

import (
	"fmt"
	"sync"
	"time"
)

func Sleepsort(numbers []int) (res []int) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	for _, number := range numbers {
		wg.Add(1)
		go func(num int) {
			time.Sleep(time.Duration(num) * time.Second)
			mu.Lock()
			res = append(res, num)
			mu.Unlock()
			wg.Done()
		}(number)
	}
	wg.Wait()
	return res
}

func main() {
	fmt.Println(Sleepsort([]int{1}))
	fmt.Println(Sleepsort([]int{2, 1}))
	fmt.Println(Sleepsort([]int{5, 4, 2, 3, 1}))
	fmt.Println(Sleepsort([]int{2, 1, 1, 2}))
}
