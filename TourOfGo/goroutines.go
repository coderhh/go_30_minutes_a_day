/*
Go 程

Go程是由Go运行时管理的轻量级线程
go f(x,y,x)
会启动一个新的Go程并执行
f(x,y,z)

f,x,y 和 z 的求值发生在当前的Go程中，而f的执行发生在新的Go程中。
Go程在相同的地址空间中运行，因此在访问共享的内存时必须进行同步。sync包提供了这种能力，不过在Go
中并不经常用到，因此还有其它的办法。
*/

package main

import (
	"fmt"
	"time"
)

func say(s string)  {
	for i := 0; i < 5; i++{
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main()  {
	go say("world")
	say("hello")
}