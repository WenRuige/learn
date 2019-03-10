#### golang 内存

`golang`在内存管理的时候,使用了`tcmalloc`



内存管理:
* 内存创建
    * 定长内存分配(`bitmap`),(`链表分配`)
    * 变长内存分配
* 内存回收







定长内存分配:



变长内存分配:

我们把所有的变长记录进行取整,比如分配7字节,则取整到8字节,但是并不是简单的往2^n次幂取整,而是按照
`8,16,32,48,64,80`




分配大对象如何分配呢?

上面讲的是基于`Page`,分配小于`Page`的对象,但是如果分配的对象大于一个`page`,
我们就需要多个`Page`来分配了.


多个连续的配置组成`span`,`span`中记录起始的`page`编号,以及`page`的数量






`Cache`:线程是私有的,每个线程都有一个`cache`,用来存放小对象

```go



//	mheap: the malloc heap, managed at page (8192-byte) granularity(间隔,尺寸).
//	mspan: a run of pages managed by the mheap. (mheap管理mspan)
//	mcentral: collects all spans of a given size class.
//	mcache: a per-P cache of mspans with free space.
//	mstats: allocation statistics.(统计)

```


#### `Tcmalloc`

`Thread Cache Malloc`将内存分解为多层,从而减小内存的粒度`Tcmalloc`将内存管理分为线程内存以及页堆两部分。