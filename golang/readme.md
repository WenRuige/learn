
#### `Golang`

`webServer`几种主流的并发模型:
 - 多进程并发模型
 - 多线程: 每个线程一次处理一个请求,在当前请求处理完成之前不会接收其他的请求;高并发环境下;开销比较大
 - 多路i/o复用(select/poll/epoll)
 - 协程:不需要抢占式调度,可以有效的提高任务的并发性
 
 
 
 
 
 
 
 ##### `Golang`的`CSP`并发模型
 - 多线程共享内存
 - `CSP(communicating sequential process)` 并发模型
 
 > 不要以共享内存的方式来通信,而是通过通信来共享内存 (do not communicate by sharing memory instead ,share memory by communicating)
 
 
##### `Golang`调度:


  - `M`:内核线程
  - `G`:`Goroutine`,并发逻辑最小单元,由程序员创建 :)
  - `P`:处理器,执行`G`的上下文环境,每个`P`会维护一个本地的`Goroutine`队列
  
  
 每个`P`除了拥有一个本地的`Goroutine`以外,还有一个全局的`Goroutine`
 
 
 
 
 
 - `P`在初始化的时候由`GOMAXPROCS`决定
 - 获取`G`的顺序,本地队列 -> 全局队列  -> 其他P的队列
 
 