/*
错误

Go程序使用error值来表示错误状态。
与fmt.Stringer 类似，error 类型是一个内建接口：
type error interface {
	Error() string
}
(与fmt.Stringer类似，fmt包打印值时也会满足error。)
通常函数会返回一个error值， 调用的它的代码应当判断这个错误
是否等于nil来进行错误处理。

i, err := strconv.Atoi("42")
if err != nil {
	fmt.Println("couldn't convert number: %v\n", err)
	return
}
fmt.Println("Converted integer:", i)

error 为nil时表示成功；非nil的error表示失败。
*/

package main
import(
	"fmt"
	"time"
)

// MyError  :
type MyError struct {
	When time.Time
	What string
}
func (e *MyError) Error() string  {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error  {
	return &MyError{
		time.Now(),
		"It didn't work",
	}
}

func main()  {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}