/*
带缓冲的信道

信道可以是带缓冲的。将缓冲长度作为第二个参数提供给make
来初始化一个带缓冲的信道：
ch := make(chan int, 100)

仅当信道的缓冲区填满后，向其发送数据时才会阻塞。
当缓冲区为空时，接收方会阻塞。

*/

package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
