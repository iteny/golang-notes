判断语句if
====
* 条件表达式没有括号<br>
* 支持一个初始化表达式(可以是并行方式)<br>
* 左大括号必须和条件语句或else在同一行<br>
* 支持单行模式<br>
* 初始化语句中的变量为block级别，同时隐藏外部同名变量<br>
* 1.0.3版本中的编译器BUG<br>

if语句演示
====
```yanshi
func main(){
	a := true
	if a,b,c := 1,2,3; a+b+c > 6{
		fmt.Println("大于6")
	}else{
		fmt.Println("小于等于6")
		fmt.Println(a)
	}
	fmt.Println(a)
}
//输出结果：
小于等于6
1
true
```
循环语句for
=====
* Go只有for一个循环语句关键字，但支持3种形式<br>
* 初始化和步进表达式可以是多个值<br>
* 条件语句每次循环都会被重新检查，因此不建议在条件表达式内使用函数，尽量提前计算好好条件并以变量或常量代替<br>
* 左大括号必须和条件语句在同一行<br>

for语句第一种形式，无限循环
====
```wuxian
package main
import(
	"fmt"
)
func main() {
	a := 1
	for{
		a++
		if a>3{
			break
		}
		fmt.Println(a)
	}
	fmt.Println("到这里程序结束")
}
```
for语句第二种形式，for自带条件表达式
====
```tiaojian
package main
import(
	"fmt"
)
func main() {
	a := 1
	for a<=4{
		a++
		fmt.Println(a)
	}
	fmt.Println("到这里程序结束")
}
```
for语句第三种形式，计数器的形式
====
```jishuqi
package main
import(
	"fmt"
)
func main() {
	a := 1
	for i:=1;i<3;i++{
		a++
		fmt.Println(a)
	}
	fmt.Println("到这里程序结束")
}
```
选择语句switch
====
* 可以使用任何类型或表达式作为条件语句<br>
* 不需要写break,一旦条件符合自动终止<br>
* 如希望继续执行下一个case,需使用fallthrough语句<br>
* 支持一个初始化表达式(可以是并行方式),左侧需跟分号<br>
* 左大括号必须和条件语句在同一行<br>

switch语句的演示
```yanshi
func main(){
	a := 1
	switch a{
	case 0:
		fmt.Println("a=0")
	case 1:
		fmt.Println("a=1")
	}
	fmt.Println(a)
}
func main(){
	a := 1
	switch {
	case a >= 0:
		fmt.Println("a=0")
	case a >= 1:
		fmt.Println("a=1")
	}
	fmt.Println(a)
}
func main(){
	switch a := 1;{
	case a >= 0:
		fmt.Println("a=0")
	case a >= 1:
		fmt.Println("a=1")
	}
	fmt.Println(a)
}
func main(){
	a := 2
	switch a{
	case 0:
		fmt.Println("a=0")
	case 1:
		fmt.Println("a=1")
	default:
		fmt.Println("都不符合条件")
	}
}
```
跳转语句goto,break,continue
=====
* 三个语法都可以配合标签使用<br>
* 标签名区分大小写,若不使用会造成编译错误<br>
* break与continue配合标签可用于多层循环的跳出<br>
* goto是调整执行位置,与其它2个语句配合标签的结果并不相同<br>

跳转语句的演示：
====
```yanshi
func main(){
	LABEL1:
	for{
		for i:=0;i<10;i++{
			if i>2{
				break LABEL1
			}else{
				fmt.Println(i)
			}
		}
	}
}
//输出结果：如果不加标签，此处将无限循环
0
1
2

func main(){
	for{
		for i:=0;i<10;i++{
			if i>2{
				goto LABEL1
			}else{
				fmt.Println(i)
			}
		}
	}
	LABEL1:
	fmt.Println("OK")
}
//输出结果：注意标签必须在for下面，否则死循环
0
1
2
OK

func main(){
	LABEL1:
	for i:=0;i<10;i++{
		for{
			fmt.Println(i)
			continue LABEL1
		}
	}
	fmt.Println("OK")
}
//输出结果：程序走到continue会自动停止无限循环
0
1
2
3
4
5
6
7
8
9
OK
```

