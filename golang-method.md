方法method
=====
* Go中虽没有class,但依旧有method
* 通过显示说明receiver来实现与某个类型的组合
* 只能为同一个包中的类型定义方法
* receiver可以是类型的值或者指针
* 不存在方法重载
* 可以使用值或指针来调用方法，编译器会自动完成转换
* 从某种意义上来说，方法是函数的语法糖，因为receiver其实就是方法所接收的第一个参数(Method Value vs.Method Expression)
* 如果外部结构和嵌入结构存在同名方法，则优先调用外部结构的方法
* 类型别名不会拥有底层类型所附带的方法
* 方法可以调用结构中的非公开字段

代码演示：
```
package main
import (
	"fmt"
)
type A struct {
	Name string
}
type B struct {
	Name string
}
func main() {
	a := A{}
	a.Print()
	b := B{}
	b.Print()
}
func (a A) Print() {
	fmt.Println("A")
}
func (b B) Print() {
	fmt.Println("B")
}
//输出结果：这个例子可以看出，类型和方法的绑定
A
B

package main
import (
	"fmt"
)
type TZ int
func main() {
	var a TZ
	a.Print()
}
func (a *TZ) Print() {
	fmt.Println("TZ")
}
//输出结果：这段代码告诉我们Go语言的灵活之处,可以类型绑定方法
TZ

package main
import (
	"fmt"
)
type A struct {
	name string
}
func main() {
	a := A{}
	a.Print()
	fmt.Println(a.name)
}
func (a *A) Print() {
	a.name = "123"
	fmt.Println(a.name)
}
//输出结果：和class其实一样，方法是可以访问私有的字段,但是要注意字段名小写，只可以在当前包下
123