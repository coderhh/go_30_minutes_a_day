/*
图像

image 包定义了Image接口：

package image

type Image Interface {
	ColorModel() color.Model
	Bounds() Rectangle
	At(x, y int) color.Color
}
注意： Bounds 方法的返回值Rectangle实际上是一个image.Rectangle, 它在image 包中声明。
color.Color 和 color.Model 类型也是接口，但是通常因为直接使用预定义的实现 image.RGBA 
和 image.RGBAModel 而被忽视了。这些接口和类型由 image/color 包定义。
*/

package main

import (
	"fmt"
	"image"
)

func main()  {
	m := image.NewRGBA(image.Rect(0,0,100,100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0,0).RGBA())
}