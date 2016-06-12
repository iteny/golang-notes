golang特有类型切片slice
=====
* 其本身并不是数组，它指向底层的数组<br>
* 作为变长数组的替代方案，可以关联底层数组的局部或全部<br>
* 为引用类型<br>
* 可以直接创建或从底层数组获取生成<br>
* 使用len()获取元素个数,cap()获取容量<br>
* 一般使用make()创建<br>
* 如果多个slice指向相同底层数组，其中一个的值改变会影响全部<br>
* make([]T,len,cap)<br>
* 其中cap可以省略，则和len的值相同<br>
* len表示存数的元素个数，cap表示容量<br>

代码演示：
======
```yanshi
package main
import(
	"fmt"
)
func main(){
	a:=[10]int{1,2,3,4,5,6,7,8,9}
	fmt.Println(a)
	t1:=a[5:len(a)]
	fmt.Println(t1)
}
//输出结果：
[1 2 3 4 5 6 7 8 9 0]
[6 7 8 9 0]

package main
import(
	"fmt"
)
func main(){
	a:=[10]int{1,2,3,4,5,6,7,8,9}
	fmt.Println(a)
	t1:=a[:5]
	fmt.Println(t1)
}
//输出结果：
[1 2 3 4 5 6 7 8 9 0]
[1 2 3 4 5]

package main
import(
	"fmt"
)
func main(){
	ss:=make([]int,22,100)
	fmt.Println(len(ss),cap(ss))
}
//输出结果：分配多少内存容量是固定的，如果不设置容量，编译器会自动把个数当容量
22 100

package main
import(
	"fmt"
)
func main(){
	a:=[]byte{'a','b','c','d','e','f','g','h','i','j','k'}
	ss:=a[2:5]
	fmt.Println(string(ss))
}
//输出结果：
cde

package main
import(
	"fmt"
)
func main(){
	a:=[]byte{'a','b','c','d','e','f','g','h','i','j','k'}
	ss:=a[3:5]
	fmt.Println(string(ss))
}
//输出结果：
de
```
reslice
====
* reslice时索引以被slice的切片为准<br>
* 索引不可以超过被slice的切片的容量cap()的值<br>
* 索引越界不会导致底层数组的重新分配而是引发错误<br>

代码演示：
======
```yanshi
package main
import(
	"fmt"
)
func main(){
	a:=[]byte{'a','b','c','d','e','f','g','h','i','j','k'}
	sa:=a[2:5]
	sb:=sa[3:5]
	fmt.Println(string(sb))
}
//输出结果：
fg

package main
import(
	"fmt"
)
func main(){
	a:=[]byte{'a','b','c','d','e','f','g','h','i','j','k'}
	sa:=a[2:5]
	sb:=sa[1:3]
	fmt.Println(string(sb))
}
//输出结果：
de

package main
import(
	"fmt"
)
func main(){
	a:=[]byte{'a','b','c','d','e','f','g','h','i','j','k'}
	sa:=a[2:5]
	fmt.Println(len(sa),cap(sa))
	sb:=sa[3:5]
	fmt.Println(string(sb))
}
//输出结果：
3 9
fg

package main
import(
	"fmt"
)
func main(){
	a:=[]byte{'a','b','c','d','e','f','g','h','i','j','k'}
	sa:=a[2:5]
	fmt.Println(len(sa),cap(sa))
	sb:=sa[9:11]
	fmt.Println(string(sb))
}
//输出结果：提示我们内存越界
3 9
panic: runtime error: slice bounds out of range

goroutine 1 [running]:
panic(0x4db320, 0xc82000a0c0)
        /root/go/src/runtime/panic.go:481 +0x3e6
main.main()
        /root/golib/src/test/my.go:9 +0x32a
exit status 2
```
append
====
* 可以在slice尾部追加元素<br>
* 可以将一个slice追加在另一个slice尾部<br>
* 如果最终长度未超过追加到slice的容量则返回原始slice<br>
* 如果超过追加到的slice容量则将重新分配数组并拷贝原始数据<br>

代码演示：
```yanshi
package main
import(
	"fmt"
)
func main(){
	s1:=make([]int,3,6)
	fmt.Printf("%p\n",s1)
	s1=append(s1,1,2,3)
	fmt.Printf("%v %p\n",s1,s1)
	s1=append(s1,1,2,3)
	fmt.Printf("%v %p\n",s1,s1)
}
//输出结果：我们可以看出当第2次没有超出容量时，没有重新分配内存！第三次超过了他就会自动重新分配内存
0xc8200160c0
[0 0 0 1 2 3] 0xc8200160c0
[0 0 0 1 2 3 1 2 3] 0xc820048060

package main
import(
	"fmt"
)
func main(){
	a:=[]int{1,2,3,4,5}
	s1:=a[2:5]
	s2:=a[1:3]
	fmt.Println(s1,s2)
	s1[0]=9
	fmt.Println(s1,s2)
}
//输出结果：从上面代码可以看出如果改变某一个slice的值，另外一个也会改变
[3 4 5] [2 3]
[9 4 5] [2 9]

package main
import(
	"fmt"
)
func main(){
	a:=[]int{1,2,3,4,5}
	s1:=a[2:5]
	s2:=a[1:3]
	fmt.Println(s1,s2)
	s2=append(s2,1,1,1,1,1,1,1)
	s1[0]=9
	fmt.Println(s1,s2)
}
//输出结果：当s2用append追加数组，超出了a的容量，那么S2重新分配内存，不再指向数组a了，所以并不会影响s1的结果,s1依旧指向a
[3 4 5] [2 3]
[9 4 5] [2 3 1 1 1 1 1 1 1]
```

copy函数
=====

代码演示：
====
```yanshi
package main
import(
	"fmt"
)
func main(){
	s1:=[]int{1,2,3,4,5,6}
	s2:=[]int{7,8,9}
	copy(s1,s2)
	fmt.Println(s1,s2)
}
//输出结果：我们可以看出s2拷贝到s1，所以只会改变第一个参数
[7 8 9 4 5 6] [7 8 9]

package main
import(
	"fmt"
)
func main(){
	s1:=[]int{1,2,3,4,5,6}
	s2:=[]int{7,8,9}
	copy(s2,s1)
	fmt.Println(s1,s2)
}
//输出结果：上面的例子，可以看出当S2只有3个元素，所以只拷贝s1的前3个元素
[1 2 3 4 5 6] [1 2 3]

package main
import(
	"fmt"
)
func main(){
	s1:=[]int{1,2,3,4,5,6}
	s2:=[]int{7,8,9,1,1,1,1,1}
	copy(s2[2:4],s1[1:3])
	fmt.Println(s2)
}
//输出结果：这个例子很明显，就不多做解释
[7 8 2 3 1 1 1 1]
```