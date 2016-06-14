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
	m := make(map[int]string)
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
```