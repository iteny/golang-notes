并发concurrency
====
* 很多人都是冲着Go大肆宣扬的高并发而忍不住跃跃欲试，但其实从源码的解析来看，goroutine只是由官方实现的超级"线程池"而已。不过话说回来，每个实例4-5KB的栈内存占用和由于实现机制而大幅减少的创建和销毁开销，是制造Go号称的高并发的根本原因。另外，goroutine的简单易用，也在语言层面上给予开发者巨大的便利。
* 大神说过并发不是并行
* 并发主要由切换时间片来实现"同时"运行，而并行则是直接利用多核实现多线程的运行，但Go可以设置使用核数，以发挥多核计算机的能力。
* goroutine奉行通过通信来共享内存；而不是共享内存来通信
Channel
======
* Channel是goroutine沟通的桥梁，大都是阻塞同步的
* 通过make创建,close关闭
* Channel是引用类型
* 可以使用for range来迭代不断操作Channel
* 可以设置单向或双向通道(如果用make创建Channel那么他就是一个双向通道)
* 可以设置缓存大小，在未被填满前不会发生阻塞
Select
====
* 可处理一个或多个Channel的发送与接收
* 同时有多个可用的Channel时按随机顺序处理
* 可用空的Select来阻塞main函数
* 可设置超时

代码演示：
```
package main
import (
	"fmt"
	"time"
)
func main() {
	go Go()
	time.Sleep(2 * time.Second)
}
func Go() {
	fmt.Println("Go Go Go")
}
//输出结果：运行一个并发，简单的演示
Go Go Go

package main
import (
	"fmt"
)
func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("GO GO GO")
		c <- true
	}()
	<-c
}
//输出结果：这里不需要人为的close关闭,因为当程序结束，所有的资源都会自己释放
GO GO GO

package main
import (
	"fmt"
)
func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("GO GO GO")
		c <- true
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
}
//输出结果：这里案例要对这个Channel关闭，否则死锁，程序崩溃退出
GO GO GO

package main
import (
	"fmt"
	"runtime"
)
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go Go(c, i)
	}
	for i := 0; i < 10; i++ {
		<-c
	}
}
func Go(c chan bool, index int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	c <- true
}
//输出结果：利用Channel完成10个任务的并发
9 49999995000001
6 49999995000001
4 49999995000001
7 49999995000001
5 49999995000001
0 49999995000001
8 49999995000001
3 49999995000001
2 49999995000001
1 49999995000001

package main
import (
	"fmt"
	"runtime"
	"sync"
)
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Go(&wg, i)
	}
	wg.Wait()
}
func Go(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	wg.Done()
}
//输出结果：这个例子和上面一样，不过用的是同步包实现的
9 49999995000001
1 49999995000001
0 49999995000001
5 49999995000001
7 49999995000001
8 49999995000001
6 49999995000001
2 49999995000001
4 49999995000001
3 49999995000001

package main
import (
	"fmt"
	"runtime"
)
func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to c
}
func main() {
	a := []int{7, 2, 8, -9, 4, 0}
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan int, 10)

	go sum(a[:len(a)/2], c)
	go sum(a[1:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	x, y, z := <-c, <-c, <-c // receive from c
	fmt.Println(x, y, z, x+y)
}
//输出结果：这里例子有些奇怪，感觉Channel接收是伪随机的
-5 17 10 12

package main
import (
	"fmt"
)
func main() {
	c1, c2 := make(chan int), make(chan string)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c1=", v)
			case v, ok := <-c2:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c2=", v)
			}
		}
	}()
	c1 <- 1
	c2 <- "hello"
	c1 <- 3
	c2 <- "nihao"
	close(c1)
	close(c2)
	<-o
}
//输出结果：演示了select的简单应用
c1= 1
c2= hello
c1= 3
c2= nihao

package main
import (
	"fmt"
)
func main() {
	c := make(chan int)
	go func() {
		for v := range c {
			fmt.Println(v)
		}
	}()
	for {
		select {
		case c <- 0:
		case c <- 1:
		}
	}
}
//输出结果：这里例子向channel中无限插入0,1,然后无限随机读出
0
1
0
0
1

package main
import (
	"fmt"
	"time"
)
func main() {
	c := make(chan bool)
	select {
	case v := <-c:
		fmt.Println(v)
	case <-time.After(3 * time.Second):
		fmt.Println("TimeOut")
	}
}
//输出结果：这是一个select超时的简单例子
TimeOut

package main
import (
	"fmt"
)
var c chan string
func Pingpong() {
	i := 0
	for {
		fmt.Println(<-c)
		c <- fmt.Sprintf("From Pingpong:Hi,#%d", i)
		i++
	}
}
func main() {
	c = make(chan string)
	go Pingpong()
	for i := 0; i < 10; i++ {
		c <- fmt.Sprintf("From main:Hello,#%d", i)
		fmt.Println(<-c)
	}
}
//输出结果：信息相互交互的例子
From main:Hello,#0
From Pingpong:Hi,#0
From main:Hello,#1
From Pingpong:Hi,#1
From main:Hello,#2
From Pingpong:Hi,#2
From main:Hello,#3
From Pingpong:Hi,#3
From main:Hello,#4
From Pingpong:Hi,#4
From main:Hello,#5
From Pingpong:Hi,#5
From main:Hello,#6
From Pingpong:Hi,#6
From main:Hello,#7
From Pingpong:Hi,#7
From main:Hello,#8
From Pingpong:Hi,#8
From main:Hello,#9
From Pingpong:Hi,#9
```