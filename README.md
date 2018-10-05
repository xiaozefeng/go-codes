# learning golang
##  golang type conversion is displayed
### 指针
- 接口变量自带指针
- 接口变量同样采用值传递，几乎不需要使用接口的指针
- 指针接受者实现只能以指针方式使用；值接受者都可

### 测试
- 分离的测试数据和测试逻辑
- 明确的出错信息
- 可以部分失败
- go语言的语法使得我们更易实现表格驱动测试

#### 代码覆盖率
> go test -coverprofile=c.out
> go tool cover -html=test.out

#### 性能测试
> go test -bench .

#### 性能调优
> go test -bench . -cpuprofile=cpu.out
> go tool pprof cpu.out
> 输入 web


### goroutine
#### 携程 Coroutine
- 轻量级 "线程"
- 非抢占式多任务处理，由携程主动交出控制权
- 编译器/解释器/虚拟机层面的多任务
- 多个携程可能在一个或多个线程上运行
- 子程序时携程的一个特例

#### 检测数据冲突
> go run -race xx.go

