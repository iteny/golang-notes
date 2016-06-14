golang函数function
======
* Go函数不支持嵌套,重载和参数默认值
* 但支持以下特性
	无需声明原型，不定长度变参，多返回值，命名返回值参数
	匿名函数，闭包
* 定义函数使用关键字func,且左大括号不能另起一行
* 函数也可以作为一种类型使用

代码演示：
=====
```
package main
import (
	"fmt"
)
func main() {
	A(1, 2, 3, 4, 5, 6, 7)

}
func A(a ...int) {
	fmt.Println(a)
}
//输出结果：这是不定长参数的用法(注意，它必须是参数列表中最后一个参数)
[1 2 3 4 5 6 7]

package main
import (
	"fmt"
)
func main() {
	s1 := []int{1, 2, 3, 4}
	A(s1)
	fmt.Println(s1)
}
func A(s []int) {
	s[0] = 5
	s[1] = 6
	s[2] = 7
	s[3] = 8
	fmt.Println(s)
}
//输出结果：如果传递一个slice进去，会影响到main函数本身的s1,并不是传递一个指针，而是内存地址的拷贝
[5 6 7 8]
[5 6 7 8]

package main
import (
	"fmt"
)
func main() {
	a := 1
	A(a)
	fmt.Println(a)
}
func A(a int) {
	a = 2
	fmt.Println(a)
}
//输出结果：可以看到int,string型传递进去，并不会改变
2
1

package main
import (
	"fmt"
)
func main() {
	a := 1
	A(&a)
	fmt.Println(a)
}
func A(a *int) {
	*a = 2
	fmt.Println(*a)
}
//输出结果：可以看到当我们拷贝的地址时，就可以改变a的值
2
2

package main
import (
	"fmt"
)
func main() {
	a := A
	a()
}
func A() {
	fmt.Println("func A")
}
//输出结果：可以看到函数是一种类型，他是a的赋值，所以调用小a等于调用了大A
func A

package main
import (
	"fmt"
)
func main() {
	a := func() {
		fmt.Println("func a")
	}
	a()
}
//输出结果：这个例子就是一个匿名函数
func a

package main
import (
	"fmt"
)
func main() {
	f := closure(10)
	fmt.Println(f(1))
	fmt.Println(f(2))
}
func closure(x int) func(int) int {
	fmt.Printf("%p\n", &x)
	return func(y int) int {
		fmt.Printf("%p\n", &x)
		return x + y
	}
}
//输出结果：这是一个闭包的例子
0xc0820361c8
0xc0820361c8
11
0xc0820361c8
12
```

defer
=====
* 执行方式类似其他语言中的析构函数，在函数体执行结束后按照调用顺序的相反顺序逐个执行
* 即时函数发生严重错误也会执行
* 支持匿名函数的调用
* 常用于资源清理，文件关闭，解锁以及记录时间等操作
* 通过与匿名函数配合可在return之后修改函数计算结果
* 如果函数体内某个变量作为defer时匿名函数的参数，则在定义defer时已经获得了拷贝，否则是引用某个变量的地址
* Go没有异常机制，但有panic/recover模式来处理错误
* Panic可以在任何地方引发，但recover只有在defer调用的函数中有效

代码演示：
====
```
package main
import (
	"fmt"
)
func main() {
	fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
}
//输出结果：
a
c
b

package main
import (
	"fmt"
)
func main() {
	for i := 0; i < 6; i++ {
		defer fmt.Println(i)
	}
}
//输出结果：
5
4
3
2
1
0

package main
import (
	"fmt"
)
func main() {
	for i := 0; i < 6; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}
//输出结果：这个i引用的是地址；这是一个defer,匿名函数，闭包的一个例子
6
6
6
6
6
6

package main
import (
	"fmt"
)
func main() {
	A()
	B()
	C()
}
func A() {
	fmt.Println("func A")
}
func B() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}
func C() {
	fmt.Println("func C")
}
//输出结果：
func A
recover in B
func C

package main
import (
	"fmt"
)
func main() {
	A()
	B()
	C()
}
func A() {
	fmt.Println("func A")
}
func B() {
	panic("panic in B")
}
func C() {
	fmt.Println("func C")
}
//输出结果：
func A
panic: panic in B

goroutine 1 [running]:
panic(0x4cbf40, 0xc082036210)
	e:/go/src/runtime/panic.go:481 +0x3f4
main.B()
	F:/golib/src/test/my.go:16 +0x6c
main.main()
	F:/golib/src/test/my.go:9 +0x20
exit status 2

package main
import (
	"fmt"
)
func main() {
	var fs = [4]func(){}
	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i=", i)
		defer func() { fmt.Println("defer_closure i=", i) }()
		fs[i] = func() { fmt.Println("closure i=", i) }
	}
	for _, f := range fs {
		f()
	}
}
//输出结果：这是一个defer，闭包的例子
closure i= 4
closure i= 4
closure i= 4
closure i= 4
defer_closure i= 4
defer i= 3
defer_closure i= 4
defer i= 2
defer_closure i= 4
defer i= 1
defer_closure i= 4
defer i= 0
```