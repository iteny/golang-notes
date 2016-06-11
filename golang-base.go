//当前程序的包名
package main
//导入其他的包
import(
	"fmt"
)
//常量的定义
const PI = 3.14
//全局变量的声明与赋值
var name = "gopher"
//一般类型声明
type newType int
//结构声明
type gopher struct{}
//接口声明
type golang interface{}
//由main函数作为程序入口点启动
func main() {
	fmt.Println("Hello World!你好，世界！")
}
//Go语言可见性规则，使用大小写来决定该常量，变量，类型，接口，结构或函数是否可以被外部包所调用
//根据约定，函数名首字幕小写即为private(私有的),函数名首字母大写即为public(公共的)

