#### goroutine





当我们创建一个`goroutine`的时候,会先放到全局运行队列中,等待go运行调度器进行调度,把他们分配给一个逻辑处理器,
并放到这个逻辑处理器对应的本地队列中,最终等待被逻辑处理器执行.




#### 协程和线程的区别


协程和线程区别是不再被内核调度,而是交给了程序自己
而线程是将自己交给内核调度,





#### 进程/线程与协程之间的区别



* 进程/线程都需要抢占式调度


[goroutine文章](https://www.cnblogs.com/zkweb/p/7815600.html)
[书籍](https://docs.hacknode.org/gopl-zh/ch13/ch13-05.html)
[goroutine文章2](http://www.sizeofvoid.net/goroutine-under-the-hood/)