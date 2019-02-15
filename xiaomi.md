#### 工程

* Mysql 最左原则
* Redis 为什么快
* Redis i/o模型 epoll select poll (fd)
* Redis 存储1k数据为什么比3k数据快
* Redis 持久化
* localCache实现原理,如何更新localCache
* 如何设计一个类似json压缩比较高的

* golang map 底层实现
* golang slice array区别
* golang 值传递,引用传递
* golang 遍历一个特别大的map如何去做优化

* thrift数据传输如何做的压缩,protobuf呢



#### 算法

* 5张牌顺子
* 删除K位数字









### Answer

> Redis 为什么快?

* 纯内存操作
* 使用多路复用I/O模型(网络采用epoll,多路复用效率高)
* 采用单线程,不必考虑上下文切换和竞争条件,不用考虑锁等条件


> golang slice和array


`golang array`

* `golang`类型是值类型(值传递)
* `golang`的长度也是`type`的一部分



`golang slice`
* `slice`是一个引用类型,是一个动态指向数组的指针(引用传递)
* `slice`是一个不定长的,总是指向数组`array`的数据结构
* 如果切片容量小于1024,那么扩容的时候`slice`的`cap`就会翻番,一旦超过1024,增长因子变成1.25













