数组Array
===
* 定义数组的格式：var <varName> [n]<type>,n>=0<br>
* 数组长度也是类型的一部分，因此具有不同长度的数组为不同的类型<br>
* 注意区分指向数组的指针和指针数组<br>
* 数组之间可以使用==或!=进行比较，但不可以用<或><br>
* 可以使用new来创建数组，此方法返回一个指向数组的指针<br>
* Go支持多维数组

下面演示指针数组：
=====
```yanshi
package main
import(
	"fmt"
)
func main(){
	x,y:=1,2
	a:=[2]*int{&x,&y}
	fmt.Println(a)
}
//输出结果：
[0xc82000a2a8 0xc82000a2d0]
```
下面演示指向数组的指针
====
```yanshi
package main
import(
	"fmt"
)
func main(){
	a:=[...]int{19: 1}
	var p *[20]int = &a
	fmt.Println(p)
}
//输出结果：
&[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1]

package main
import(
	"fmt"
)
func main(){
	a := new([20]int)
	fmt.Println(a)
}
//输出结果：
&[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]

package main
import(
	"fmt"
)
func main(){
	a := [20]int{}
	a[2] = 2
	fmt.Println(a)
	p := new([20]int)
	p[2] = 2
	fmt.Println(p)
}
//输出结果：
[0 0 2 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
&[0 0 2 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
```
数组之间比较：需要相同的类型和长度
====
```yanshi
package main
import(
	"fmt"
)
func main(){
	a := [2]int{1,2}
	b := [2]int{1,2}
	fmt.Println(a==b)
}
//输出结果：
true
```
多维数组
====
~~~yanshi
package main
import(
	"fmt"
)
func main(){
	a:=[2][3]int{
		{1,2},
		{1,2,3}}
	fmt.Println(a)
}
//输出结果：
[[1 2 0] [1 2 3]]
```
Go版的冒泡排序演示
====
```yanshi
package main
import(
	"fmt"
)
func main(){
	a:=[10]int{5,12,25,15,7,28,10,6,14,19}
	fmt.Println(a)
	num:=len(a)
	for i := 0; i < num; i++ {
		for j := i+1; j < num; j++ {
			if a[i]>a[j]{
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	fmt.Println(a)
}
//输出结果：
[5 12 25 15 7 28 10 6 14 19]
[5 6 7 10 12 14 15 19 25 28]
```