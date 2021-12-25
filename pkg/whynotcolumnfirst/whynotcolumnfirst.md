### 循环遍历先行后列真的好吗？

废话不说，先贴结果！ 
```go
goos: darwin
goarch: arm64
pkg: neekity/go-run-OR-you-run/pkg/whynotcolumnfirst
BenchmarkRowFirst4096-8               58          20118547 ns/op
BenchmarkColumnFirst4096-8             6         167679465 ns/op
BenchmarkRowFirst2048-8              282           4248085 ns/op
BenchmarkColumnFirst2048-8           282           4247841 ns/op
PASS
ok      neekity/go-run-OR-you-run/pkg/whynotcolumnfirst 6.847s
```

有意思的是**2048\*2048数组**先行还是先列区别不大，**4096\*4096数组**先行的速度有显著提升。

在计算机中还有比内存访问更快的部件：CPU缓存，而CPU缓存都是抓取某个内存块。
如果存储的是数组这种数据结构，这个块上的数据几乎都是同行不同列，所以先遍历行能有效命中缓存减少耗时。
当然如果你的CPU缓存够大，一次性能把整个数组全部放在CPU缓存中，那我得说一句，**老板真有钱！**