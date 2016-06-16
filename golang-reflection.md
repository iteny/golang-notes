反射reflection
=====
* 反射可大大提高程序的灵活性，使得interface{}有更大的发挥空间
* 反射使用TypeOf和ValueOf函数从接口中获取目标对象信息
* 反射会将匿名字段作为独立字段(匿名字段本质)
* 想要利用反射修改对象状态，前提是interface.data是settable,即pointer-interface
* 通过反射可以"动态调用方法"

代码演示：
```
package main
import (
	"fmt"
	"reflect"
)
type User struct {
	Id   int
	Name string
	Age  int
}
func (u User) Hello() {
	fmt.Println("Hello world.")
}
func main() {
	u := User{1, "OK", 12}
	Info(u)
}
func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())
	v := reflect.ValueOf(o)
	fmt.Println("Fields:")
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}
}
//输出结果：这个案例是对某一个结构反射出一系列的属性
Type: User
Fields:
    Id: int = 1
  Name: string = OK
   Age: int = 12
 Hello: func(main.User)

package main
import (
	"fmt"
	"reflect"
)
type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello() {
	fmt.Println("Hello world.")
}
func main() {
	u := User{1, "OK", 12}
	Info(&u)
}
func Info(o interface{}) {
	t := reflect.TypeOf(o)

	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("反射类型错误")
		return
	}

	fmt.Println("Type:", t.Name())
	v := reflect.ValueOf(o)
	fmt.Println("Fields:")
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}
}
//输出结果：这个案例说明当你传入的类型不对(这里是个指针),就不会反射成功
反射类型错误

package main
import (
	"fmt"
	"reflect"
)
type User struct {
	Id   int
	Name string
	Age  int
}
type Manager struct {
	User
	title string
}

func main() {
	m := Manager{User: User{1, "OK", 12}, title: "123"}
	t := reflect.TypeOf(m)
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1}))
}
//输出结果：本案例-通过FieldByIndex方法找到结构Manager的匿名字段User的里面的信息
reflect.StructField{Name:"Name", PkgPath:"", Type:(*reflect.rtype)(0x4cc580), Tag:"", Offset:0x8, Index:[]int{1}, Anonymous:false}


package main
import (
	"fmt"
	"reflect"
)
type User struct {
	Id   int
	Name string
	Age  int
}
type Manager struct {
	User
	title string
}
func main() {
	m := Manager{User: User{1, "OK", 12}, title: "123"}
	t := reflect.TypeOf(m)
	fmt.Printf("%#v\n", t.Field(0))
}
//输出结果：反射出Manager结构的信息
reflect.StructField{Name:"User", PkgPath:"", Type:(*reflect.rtype)(0x4f4e00), Tag:"", Offset:0x0, Index:[]int{0}, Anonymous:true}

package main
import (
	"fmt"
	"reflect"
)
func main() {
	x := 123
	x = 555
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(999)
	fmt.Println(x)
}
//输出结果：基本类型的操作演示
999

package main
import (
	"fmt"
	"reflect"
)
type User struct {
	Id   int
	Name string
	Age  int
}
func main() {
	u := User{1, "OK", 12}
	Set(&u)
	fmt.Println(u)
}
func Set(o interface{}) {
	v := reflect.ValueOf(o)
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("XXX")
		return
	} else {
		v = v.Elem()
	}
	if f := v.FieldByName("Name"); f.Kind() == reflect.String {
		f.SetString("BYEBYE")
	}
}
//输出结果：
{1 BYEBYE 12}

package main
import (
	"fmt"
	"reflect"
)
type User struct {
	Id   int
	Name string
	Age  int
}
func main() {
	u := User{1, "OK", 12}
	Set(&u)
	fmt.Println(u)
}
func Set(o interface{}) {
	v := reflect.ValueOf(o)
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("XXX")
		return
	} else {
		v = v.Elem()
	}
	f := v.FieldByName("Name1")
	if !f.IsValid() {
		fmt.Println("BAD")
		return
	}
	if f.Kind() == reflect.String {
		f.SetString("BYEBYE")
	}
}
//输出结果：当没有找到Name1时，就输出BAD，并不修改值，本案例结构当中的对象，通过反射对他的值进行修改
BAD
{1 OK 12}

package main
import (
	"fmt"
)
type User struct {
	Id   int
	Name string
	Age  int
}
func (u User) Hello(name string) {
	fmt.Println("Hello", name, ",my name is", u.Name)
}
func main() {
	u := User{1, "OK", 12}
	u.Hello("jack")
}
//输出结果：这个案例是正常调用函数，下面通过反射来调用函数
Hello jack ,my name is OK

package main
import (
	"fmt"
	"reflect"
)
type User struct {
	Id   int
	Name string
	Age  int
}
func (u User) Hello(name string) {
	fmt.Println("Hello", name, ",my name is", u.Name)
}
func main() {
	u := User{1, "OK", 12}
	v := reflect.ValueOf(u)
	mv := v.MethodByName("Hello")
	args := []reflect.Value{reflect.ValueOf("jack")}
	mv.Call(args)
}
//输出结果：这里我们看出上一个案例输出同样的结果，但是这里是通过反射调用
Hello jack ,my name is OK
```