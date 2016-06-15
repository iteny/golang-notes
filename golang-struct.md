golang结构struct
====
* Go中的struct与C中的struct非常相似，并且Go没有class
* 使用type <Name> struct{}定义结构，名称遵循可见性规则
* 支持指向自身的指针类型成员
* 支持匿名结构，可用作成员或定义成员变量
* 匿名结构也可以用于map的值
* 可以使用字面值对结构进行初始化
* 允许直接用过指针来读写结构成员
* 相同类型的成员可进行直接拷贝赋值
* 支持==与！=比较运算符，但不支持>或<
* 支持匿名字段，本质上是定义了以某个类型名为名称的字段
* 嵌入结构作为匿名字段看起来像继承，但不是继承
* 可以使用匿名字段指针

代码演示：
```
package main
import (
	"fmt"
)
type ren struct {
	Name string
	Age  int
}

func main() {
	a := ren{}
	a.Name = "jack"
	a.Age = 28
	fmt.Println(a)
}
//输出结果：
{jack 28}

package main
import (
	"fmt"
)
type ren struct {
	Name string
	Age  int
}
func main() {
	a := ren{
		Name: "jack",
		Age:  19,
	}
	fmt.Println(a)
}
//输出结果：和上面的例子是一样的，只不过这里演示简便的初始化

package main
import (
	"fmt"
)
type ren struct {
	Name string
	Age  int
}
func main() {
	a := ren{
		Name: "jack",
		Age:  19,
	}
	fmt.Println(a)
	A(a)
	fmt.Println(a)
}
func A(r ren) {
	r.Age = 15
	fmt.Println("A=", r)
}
//输出结果：这里很明显是对age值的拷贝
{jack 19}
A= {jack 15}
{jack 19}

package main
import (
	"fmt"
)
type ren struct {
	Name string
	Age  int
}
func main() {
	a := ren{
		Name: "jack",
		Age:  19,
	}
	fmt.Println(a)
	A(&a)
	fmt.Println(a)
}
func A(r *ren) {
	r.Age = 15
	fmt.Println("A=", r)
}
//输出结果：这里就看出，只是指针的一个拷贝
{jack 19}
A= &{jack 15}
{jack 15}

package main
import (
	"fmt"
)
func main() {
	a := struct {
		Name string
		Age  int
	}{
		Name: "jack",
		Age:  19,
	}
	fmt.Println(a)
}
//输出结果：这是一个匿名结构的演示
{jack 19}

package main
import (
	"fmt"
)
type ren struct {
	Name    string
	Age     int
	Contact struct {
		Phone, City string
	}
}
func main() {
	a := ren{Name: "jack", Age: 19}
	a.Contact.Phone = "13020111111"
	a.Contact.City = "shanghai"
	fmt.Println(a)
}
//输出结果：这里演示了结构里面的匿名结构Contact如何初始化
{jack 19 {13020111111 shanghai}}

package main
import (
	"fmt"
)
type ren struct {
	Name string
	Age  int
}
func main() {
	a := ren{"jack", 19}
	fmt.Println(a)
}
//输出结果：这里演示结构里面的匿名字段如何初始化，需要严格按照顺序来初始化，否则报错
{jack 19}

package main
import (
	"fmt"
)
type ren struct {
	Name string
	Age  int
}
func main() {
	a := ren{Name: "jack", Age: 28}
	b := ren{Name: "jack", Age: 27}
	fmt.Println(b == a)
}
//输出结果：这里是相同结构中间的比较，如果把b里面的Age改成28，那样返回的将是true,不同结构名之间的比较，系统会报错，即使他们所有的内部属性都是一样的
false

package main
import (
	"fmt"
)
type human struct {
	Sex int
}
type teacher struct {
	human
	Name string
	Age  int
}
type student struct {
	human
	Name string
	Age  int
}
func main() {
	a := teacher{Name: "jack", Age: 28, human: human{Sex: 1}}
	b := student{Name: "jack", Age: 27, human: human{Sex: 0}}
	fmt.Println(a, b)
}
//输出结果：这里演示一个嵌入结构的初始化案例，由于Go没有继承(所以你就把嵌入结构当作是继承吧)
{{1} jack 28} {{0} jack 27}

package main
import (
	"fmt"
)
type human struct {
	Sex int
}
type teacher struct {
	human
	Name string
	Age  int
}
type student struct {
	human
	Name string
	Age  int
}
func main() {
	a := teacher{}
	b := student{}
	a.Name = "jacktwo"
	a.Age = 15
	a.Sex = 100
	fmt.Println(a, b)
}
//输出结果：另外一种嵌入结构初始化的方法
{{100} jacktwo 15} {{0}  0}

package main
import (
	"fmt"
)
type A struct {
	B
	Name string
}
type B struct {
	Name string
}
func main() {
	a := A{Name: "jack", B: B{Name: "rose"}}
	fmt.Println(a.Name, a.B.Name)
}
//输出结果：这个例子是讲，输出一个结构里面的同名字段，因为嵌入的结构的优先级比原本的低，所以需要这样输出
jack rose
```