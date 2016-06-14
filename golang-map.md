Go语言map类型
=====
* 类似其他语言中的哈希表或者字典，以key-value形式存储数据
* Key必须是支持==或!=比较运算的类型，不可以是函数，map或slice
* Map查找比线性搜索快很多，但比使用索引访问数据的类型慢100倍
* Map使用make()创建，支持:=这种简写方式
* make([keyType]valueType,cap)，cap表示容量，可省略
* 超出容量时会自动扩容，但尽量提供一个合理的初始值
* 使用len()获取元素个数
* 键值对不存在时自动添加，使用delete()删除某键值对
* 使用for range对map和slice进行迭代操作

代码的演示：
```
package main
import (
	"fmt"
)
func main() {
	m := make(map[int]string) //初始化map
	fmt.Println(m)
	m[1] = "OK"
	fmt.Println(m)
	a := m[1]
	fmt.Println(a)
}
//输出结果：
map[]
map[1:OK]
OK

package main
import (
	"fmt"
)
func main() {
	var m map[int]map[int]string
	m = make(map[int]map[int]string)
	a, ok := m[2][1]   //多返回值的应用
	if !ok {
		m[2] = make(map[int]string)
	}
	m[2][1] = "GOOD"
	a, ok = m[2][1]
	fmt.Println(a, ok)
}
//输出结果：
GOOD true

//下面的例子是一个slice下map的迭代操作，相当于其他语言的foreach
package main
import (
	"fmt"
)
func main() {
	sm := make([]map[int]string, 5)
	for _, v := range sm {
		v = make(map[int]string, 1)  //所有v是一个拷贝，并不是影响sm的值
		v[1] = "OK"
		fmt.Println(v)
	}
	fmt.Println(sm)
}
//输出结果：
map[1:OK]
map[1:OK]
map[1:OK]
map[1:OK]
map[1:OK]
[map[] map[] map[] map[] map[]]

package main
import (
	"fmt"
)
func main() {
	sm := make([]map[int]string, 5)
	for i := range sm {
		sm[i] = make(map[int]string, 1)
		sm[i][1] = "OK"
		fmt.Println(sm[i])
	}
	fmt.Println(sm)
}
//输出结果：
map[1:OK]
map[1:OK]
map[1:OK]
map[1:OK]
map[1:OK]
[map[1:OK] map[1:OK] map[1:OK] map[1:OK] map[1:OK]]

package main
import (
	"fmt"
	"sort" //排序库
)
func main() {
	m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	s := make([]int, len(m))
	i := 0
	for k, _ := range m {
		s[i] = k
		i++
	}
	sort.Ints(s)
	fmt.Println(s)
}
//输出结果：如果不用sort.Ints那么结果将是无序的
[1 2 3 4 5]

//下面的例子是一个键名和键值的交换
package main
import (
	"fmt"
)
func main() {
	m1 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e:", 6: "f", 7: "g"}
	m2 := map[string]int{}
	for k, v := range m1 {
		m2[v] = k
	}
	fmt.Println(m1, m2)
}
//输出结果：
map[6:f 7:g 1:a 2:b 3:c 4:d 5:e:] map[d:4 e::5 f:6 g:7 a:1 b:2 c:3]
```