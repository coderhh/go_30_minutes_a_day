/*
sync.Mutex

我们已经看到信道非常适合在各个Go程间进行通信。
但是如果我们并不需要通信呢？比如说，若我们只想保证每次只有一个Go程能够访问一个共享的变量，
从而避免冲突？

这里涉及的概念叫做互斥，我们通常使用互斥锁这一数据结构来提供这种机制。

go标准库中提供了sync.Mutex 互斥锁及其两个方法：

Lock
Unlock
我们可以通过在代码前调用Lock方法，在代码后调用Unlock方法来保证一段代码的互斥执行。
我们也可以用defer语句来保证互斥锁一定会被解锁。
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter 的并发使用是安全的。
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc 增加给定 key 的计数器的值。
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	c.v[key]++
	c.mux.Unlock()
}

// Value 返回给定 key 的计数器的当前值。
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
