goroutine：教练我想学并发筛选素数

教练：别折腾了，你看单线程全场耗时0.07毫秒，我看你能有7毫秒就谢天谢地了

单线程：😎


##使用协程并发进行素数筛选法真的快得飞起吗？
素数的刷选方法非常适用于channel的设计理念，一想到使用协程并发筛选素数岂不是能瞬秒传统单线程，心里乐开了花。
***
**PrimeA和PrimeB使用并发实现素数的筛选，PrimeS采用单线程进行素数的筛选**
***
PrimeA的核心逻辑如下代码
``` golang
for {
    prime := <- origin
    next := make(chan int)
    go func(prime int, origin chan int, next chan int) {
        for i := range origin {
            if i%prime != 0 {
                next <- i
            }
        }
        close(next)
    }(prime, origin, next)
    origin = next
}
```
`origin`和`next`分别代表需要筛选的和筛选后的数字channel，`prime`是当前的素数。每次刷选的逻辑如下:
1. 从`origin`取出将要刷选的第一个数作为素数
2. 陆续将`origin`中的数字取出与当前素数进行判断是否能整除，不能则放入`next`，最后将通道关闭
3. 下一次循环输入即为`next`
4. 反复进行上述三步

***
PrimeB的核心思路与PrimeA一致，只不过采用了递归的思想。注意对于无缓存的channel要新起一个协程进行阻塞。
***
###三个函数的运行时间基准测试如下 
```text
goos: darwin
goarch: arm64
pkg: neekity/go-run-OR-you-run/pkg/prime
BenchmarkPrimeA-8            698           1639675 ns/op
BenchmarkPrimeB-8            735           1633192 ns/op
BenchmarkPrimeS-8         196308              5274 ns/op
PASS
ok      neekity/go-run-OR-you-run/pkg/prime     4.111s
```
单线程的处理方式简直甩协程并发方式好几条街啊。

**古语有云，并发非并行，多线程上下文切换消耗资源，古人诚不欺我！**