package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//withOutWait()
	//withWait()
	//writeWithOutConcurrent()
	//writeWithOutMutex()
	//writeWithMutex()
	//readWithMutex()
	readWithRWMutex()
}

func withOutWait() {
	for i := 0; i < 10; i++ {
		go fmt.Println(i + 1)
	}
	fmt.Println("exit")
}
func withWait() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(i + 1)
		}()
	}
	wg.Wait()
	fmt.Println("exit")
}
func writeWithOutConcurrent() {
	start := time.Now()
	var counter int64
	for i := 0; i < 1000; i++ {
		time.Sleep(time.Millisecond)
		counter++
	}
	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}
func writeWithOutMutex() {
	start := time.Now()
	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)
			counter++
		}()
	}
	wg.Wait()
	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}
func writeWithMutex() {
	start := time.Now()
	var counter int64
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)

			mutex.Lock()
			counter++
			mutex.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}
func readWithMutex() {
	start := time.Now()
	var (
		counter int64
		wg      sync.WaitGroup
		mutex   sync.Mutex
	)
	wg.Add(100)

	for i := 0; i < 50; i++ {

		go func() {
			defer wg.Done()

			mutex.Lock()
			time.Sleep(time.Nanosecond)
			_ = counter
			mutex.Unlock()
		}()
		go func() {
			defer wg.Done()

			mutex.Lock()
			time.Sleep(time.Nanosecond)
			counter++
			mutex.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}
func readWithRWMutex() {
	start := time.Now()
	var (
		counter int64
		wg      sync.WaitGroup
		mutex   sync.RWMutex
	)
	wg.Add(100)
	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()

			mutex.RLock()
			time.Sleep(time.Nanosecond)
			_ = counter
			mutex.RUnlock()
		}()
		go func() {
			defer wg.Done()
			mutex.Lock()
			time.Sleep(time.Nanosecond)
			counter++
			mutex.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())

}
