package task2

import (
	"fmt"
	"sync"
)

/*
*
编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
*/
func MetexTask() {
	var mutex sync.Mutex
	var waitGroup sync.WaitGroup
	count := 0
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go func() {
			mutex.Lock()
			for j := 0; j < 1000; j++ {
				count++
			}
			mutex.Unlock()
			waitGroup.Done()
		}()

	}
	waitGroup.Wait()
	fmt.Println(count)
}

/*
*
使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
*/
func AtomicTask() {
	var waitGroup sync.WaitGroup
	count := 0
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				count++
			}
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
	fmt.Println(count)
}
