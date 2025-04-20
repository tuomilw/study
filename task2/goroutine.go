package task2

import (
	"fmt"
	"sync"
	"time"
)

/*
*
编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数
*/
func RunGoroutineTask1() {
	var wg = sync.WaitGroup{}
	wg.Add(2) //加入两个协程计数器
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			if i%2 != 0 {
				fmt.Println("1到10之间的奇数:", i)
			}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Println("1到10之间的偶数:", i)
			}
		}
	}()
	wg.Wait()
}

/*
*
设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
*/
func RunGoroutineTask2() {
	tasks := []func(){
		func() {
			time.Sleep(1 * time.Second) // 任务1
		},
		func() {
			time.Sleep(2 * time.Second) // 任务2
		},
		func() {
			time.Sleep(3 * time.Second) // 任务3
		},
	}
	durationArray := schedule(tasks...)
	for i, duration := range durationArray {
		fmt.Printf("任务 %d 耗时：%v \n", i+1, duration)
	}
}

func schedule(tasks ...func()) []time.Duration {
	type result struct {
		index    int           // 记录任务原始顺序
		duration time.Duration // 记录任务执行耗时
	}
	resultInChan := make(chan result, len(tasks)) //带缓冲的通道
	for index, vFunc := range tasks {
		// 使用立即执行函数避免闭包陷阱
		// 通过参数传递当前索引和任务副本（关键细节）
		go func(index int, vFuncParam func()) {
			start := time.Now()           // 记录任务开始时间
			vFuncParam()                  // 这里执行的是传入的vFuncParam参数，最终都是在执行我们在RunGoroutineTask2函数中定义的匿名函数
			duration := time.Since(start) // 计算耗时
			resultInChan <- result{
				index:    index,    // 保留原始索引
				duration: duration, // 计算精确耗时
			}
		}(index, vFunc) // 显式传递循环变量（避免共享变量问题）
	}
	// 结果收集切片
	resultDurationArray := make([]time.Duration, len(tasks))

	timeout := time.After(5 * time.Second)
	// 循环接收所有任务结果（循环次数=任务数量）
	for range tasks {
		// 使用select语句接收结果（可扩展超时控制）
		select {
		case result := <-resultInChan: //接收chan的数据
			// 按照原始索引位置存储结果（关键设计）
			// 保证无论协程完成顺序如何，结果都能正确对应
			resultDurationArray[result.index] = result.duration
		case <-timeout:
			fmt.Println("超时处理")
			return nil
		}
	}
	return resultDurationArray // 返回有序结果
}
