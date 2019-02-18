```go

type hchan struct {
	qcount   uint           // total data in the queue(当前队列中元素的数量)
	dataqsiz uint           // size of the circular queue (队列可以容纳的元素数量,如果为0表示无缓冲区)
	buf      unsafe.Pointer // points to an array of dataqsiz elements (队列缓冲区)
	elemsize uint16			//(元素的大小)
	closed   uint32			//(是否已被关闭)
	elemtype *_type // element type (元素的类型)
	sendx    uint   // send index (发送元素的序号)
	recvx    uint   // receive index (接受元素的序号)
	recvq    waitq  // list of recv waiters (因为读取chan而被阻塞的G)
	sendq    waitq  // list of send waiters (因为写chan而被阻塞的G)






	// lock protects all fields in hchan, as well as several
	// fields in sudogs blocked on this channel.
	//
	// Do not change another G's status while holding this lock
	// (in particular, do not ready a G), as this can deadlock
	// with stack shrinking.
	lock mutex
}
```


`channel`的底层实际上是一个队列,`sendq&recvq`实际上是waitq的链表


假设我们初始化了一个带缓冲的`buffer`,`ch:=make(chan int,3)`,则初始的值为
```
qcount = 0
dataqsiz = 3
buf = [3]int{0,0,0}
sendx = 0
recvx = 0
```




##### `sudog`
```go
type sudog struct {
	// The following fields are protected by the hchan.lock of the
	// channel this sudog is blocking on. shrinkstack depends on
	// this for sudogs involved in channel ops.

	g *g

	// isSelect indicates g is participating in a select, so
	// g.selectDone must be CAS'd to win the wake-up race.
	isSelect bool
	next     *sudog
	prev     *sudog
	elem     unsafe.Pointer // data element (may point to stack)

	// The following fields are never accessed concurrently.
	// For channels, waitlink is only accessed by g.
	// For semaphores, all fields (including the ones above)
	// are only accessed when holding a semaRoot lock.

	acquiretime int64
	releasetime int64
	ticket      uint32
	parent      *sudog // semaRoot binary tree
	waitlink    *sudog // g.waiting list or semaRoot
	waittail    *sudog // semaRoot
	c           *hchan // channel
}
```