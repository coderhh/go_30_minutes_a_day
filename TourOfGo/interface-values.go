/*
接口值

接口也是值。它们可以像其它值一样传递。

接口值可以用作函数的参数或返回值。

在内部，接口值可以看做包含值和具体类型的元组：
（value, type）

接口值保存了一个具体底层类型的具体值。
接口值调用方法时会执行其底层类型的同名方法。
*/

package main
import (
	"fmt"
	"math"
)

// I : interface
type I interface {
	M()
}

// T : struct 
type T struct {
	S string
}

// M : func
func (t *T) M()  {
	fmt.Println(t.S)
}

// F : F
type F float64

// M : M
func (f F) M()  {
	fmt.Println(f)
}

func main()  {
	var i I
	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

}

func describe(i I)  {
	fmt.Println("(%v, %T)\n", i, i)
}