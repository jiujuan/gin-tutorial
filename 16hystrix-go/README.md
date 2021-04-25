# hystrix- 简介

[hysteix-go](https://github.com/afex/hystrix-go) 灵感是来自于 Netflix 开源的项目：https://github.com/Netflix/Hystrix。

## 用 hystrix command 命令执行代码：

```go
hystrix.Go("my_command", func() error {
    // talk to other services
    return nil
}, nil)
```

第一个参数：command 名称，可以把它当成一个独立的服务

第二个参数：处理正常的逻辑

第三个参数：处理调用失败，执行这里的逻辑。就是说上面第二个参数处理逻辑失败就调用这里。保底操作。
          如果服务错误率高导致熔断开启，那么之后的请求也会直接调用此函数。

## 定义 fallback 行为：

```go
hystrix.Go("my_command", func() error {
    // talk to other services
    return nil
}, func(err error) error {
    // do this when services is down
    return nil
})
```
hystrix 有2个方法， Do 和 Go 方法。

Go 方法是异步的方式。Do 方法是同步的方式。

2 个方法使用的方式一样。

## 等待输出操作：

```go
output := make(chan bool, 1)
errors := hystrix.Go("my_command", func() error {
	// talk to other services
	output <- true
	return nil
}, nil)

select {
case out := <-output:
	// success
case err := <-errors:
	// failure
}
```

## 配置操作：

```go
hystrix.ConfigureCommand("my_command", hystrix.CommandConfig{
	Timeout:               1000,  // 执行超时
	MaxConcurrentRequests: 100,   // 最大并发数
	ErrorPercentThreshold: 25,   // 错误百分比
})
```

还有一些其他参数：
- SleepWindow：当熔断器被打开后，SleepWindow 的时间就是控制过多久后去尝试服务是否可用了
- RequestVolumeThreshold： 一个统计窗口 10 秒内请求数量。达到这个请求数量后才去判断是否要开启熔断
- ErrorPercentThreshold： 错误百分比，请求数量大于等于 RequestVolumeThreshold 并且错误率到达这个百分比后就会启动熔断

不设置这些参数，就会启用默认参数。


> 其他相关问题都可以到 https://github.com/afex/hystrix-go 去查看

## simple 代码例子

编写一个简单的程序，
hystrix.go 编写 hystrix。
client/client.go 编写请求 hystrix.go。

写完后，运行 go run hystrix.go。

在运行：go run client.go，
输出：
```shell
read body:  ret error
read body:  ret error
read body:  ret error
read body:  ret success
read body:  ret success
read body:  ret success
read body:  ret success
read body:  ret success
read body:  ret success
read body:  ret success
read body:  ret success
read body:  ret success
read body:  ret success
finished
```

hystrix.go 里设置最大请求数 10， 而 client.go 最大请求数是 13， 返回数据失败了 3 个，说明程序运行成功。
