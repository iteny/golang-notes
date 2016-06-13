Go指针
=====
	Go虽然保留了指针，但与其它编程语言不同的是，在Go当中不支持指针运算以及"->"运算符，而直接采用"."选择符来操作指针目标对象的成员
	* 操作符"&"取变量地址，使用"*"通过指针间接访问目标对象
	* 默认值为nil而非NULL

递增递减语句
======
	在Go当中，++与--是作为语句而并不是作为表达式(所以不能放在变量右边，只能是单独的一行)

指针演示
=====
```yanshi
package main
import(
	"fmt"
)
func main() {
	a := "string"
	var p *string = &a
	fmt.Println("a的地址为：",p)
}
//代码输出结果：
a的地址为： 0xc82000a2d0
```
