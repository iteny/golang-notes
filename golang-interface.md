接口interface
====
* 接口是一个或多个方法签名的集合
* 只要某个类型拥有该接口的所有方法签名，即算实现该接口，无需显示声明显示了哪个接口，这称为structural typing
* 接口只有方法声明，没有实现，没有数据字段
* 接口可以匿名嵌入其他接口或嵌入到结构中
* 将对象赋值给接口时，会发生拷贝，而接口内部存储的是指向这个复制品的指针，即无法修改复制品的状态，也无法获取指针
* 只有当接口存储的类型和对象都为nil时，接口才等于nil
* 接口调用不会做receiver的自动转换
* 接口同样支持匿名字段方法
* 接口也可实现类似OOP中的多态
* 空接口可以作为任何类型数据的容器

代码演示：
```
package main
import (
	"fmt"
)
type USB interface {
	Name() string
	Connect()
}
type PhoneConnecter struct {
	name string
}
func (pc PhoneConnecter) Name() string {
	return pc.name
}
func (pc PhoneConnecter) Connect() {
	fmt.Println("Connected=", pc.name)
}
func main() {
	a := PhoneConnecter{"PhoneConnecter"}
	a.Connect()
	Disconnect(a)
}
func Disconnect(usb USB) {
	fmt.Println("Disconnected.")
}
//输出结果：可以看我们调用Disconnect传入一个a USB类型，成功输出，说明这个接口被成功实现了
Connected= PhoneConnecter
Disconnected.

package main
import (
	"fmt"
)
type USB interface {
	Name() string
	Connecter
}
type Connecter interface {
	Connect()
}
type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}
func (pc PhoneConnecter) Connect() {
	fmt.Println("Connected=", pc.name)
}
func main() {
	a := PhoneConnecter{"PhoneConnecter"}
	a.Connect()
	Disconnect(a)
}
func Disconnect(usb USB) {
	fmt.Println("Disconnected.")
}
//输出结果：这里演示嵌入一个接口
Connected= PhoneConnecter
Disconnected.

package main
import (
	"fmt"
)
type USB interface {
	Name() string
	Connecter
}
type Connecter interface {
	Connect()
}
type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}
func (pc PhoneConnecter) Connect() {
	fmt.Println("Connected", pc.name)
}
func main() {
	a := PhoneConnecter{"PhoneConnecter"}
	a.Connect()
	Disconnect(a)
}
func Disconnect(usb USB) {
	if pc, ok := usb.(PhoneConnecter); ok {
		fmt.Println("Disconnected:", pc.name)
		return
	}
	fmt.Println("Unknown decive.")
}
//输出结果：
Connected PhoneConnecter
Disconnected: PhoneConnecter

package main
import (
	"fmt"
)
type USB interface {
	Name() string
	Connecter
}
type Connecter interface {
	Connect()
}
type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}
func (pc PhoneConnecter) Connect() {
	fmt.Println("Connected", pc.name)
}
func main() {
	a := PhoneConnecter{"PhoneConnecter"}
	a.Connect()
	Disconnect(a)
}
func Disconnect(usb interface{}) {
	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("Disconnected:", v.name)
	default:
		fmt.Println("Unknown decive.")
	}
}
//输出结果：当你使用空接口的时候就可以用switch
Connected PhoneConnecter
Disconnected: PhoneConnecter

package main
import (
	"fmt"
)
type USB interface {
	Name() string
	Connecter
}
type Connecter interface {
	Connect()
}
type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}
func (pc PhoneConnecter) Connect() {
	fmt.Println("Connected", pc.name)
}
func main() {
	pc := PhoneConnecter{"PhoneConnecter"}
	var a Connecter
	a = Connecter(pc)
	a.Connect()
	pc.name = "pc"
	a.Connect()
}
func Disconnect(usb interface{}) {
	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("Disconnected:", v.name)
	default:
		fmt.Println("Unknown decive.")
	}
}
//输出结果：这里例子知道，a拿到的只是值的拷贝
Connected PhoneConnecter
Connected PhoneConnecter

package main
import (
	"fmt"
)
type USB interface {
	Name() string
	Connecter
}
type Connecter interface {
	Connect()
}
type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}
func (pc PhoneConnecter) Connect() {
	fmt.Println("Connected", pc.name)
}
func main() {
	var a interface{}
	fmt.Println(a == nil)
	var p *int = nil
	a = p
	fmt.Println(a == nil)
}
func Disconnect(usb interface{}) {
	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("Disconnected:", v.name)
	default:
		fmt.Println("Unknown decive.")
	}
}
//输出结果：
true
false
```
