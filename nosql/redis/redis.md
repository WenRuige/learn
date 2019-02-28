### `Redis`



####`Redis`有哪些数据结构?
 - 字符串(string)
 - 字典(hash)
 - 列表(list)
 - 集合(set)
 - 有序集合(SortedSet)
 - Geo(空间索引)
 - 位图(bitmap)
 
 
 
 


#### redis 分布式锁

实现: redis + lua + setnx + expiretime + uuid (确保是原子操作)

> 为什么要加上uuid？

因为解锁操作只是简单的做了del key,如果客户端在获得锁之后执行业务时间唱过了锁的时间,则这个锁会误解掉其他客户端,所以为了解决这个问题,在每次加锁的时候,
都创建一个uuid,加锁的时候存入此值,并在解锁的时候取出值进行比较,如果匹配才进行del


#### redis是如何做持久化的

`bgsave`做镜像全量持久化,`aof`做增量持久化

`bgsave`会耗费较长时间,不够实时,在停机的时候会导致丢失大量数据,所以需要配合`aof`来使用,`redis`重启时,优先使用`aof`来恢复内存的状态,如果没有`aof`日志,
则会使用`rdb`来进行恢复

`aof`是连续的增量备份,快照是内存数据的二进制序列化形式,在存储上非常紧凑,而`aof`日志记录是文本,这个会变得异常大,所以需要定期进行`aof`重写,给`aof`日志瘦身


`bgsave`的原理是什么:`fork`&`cow`,`fork`是指`redis`通过创建子进程来进行`bgsave`操作,`cow`指的是`copy on write`

`copy on write`:写时拷贝是一种可推迟或者免除拷贝数据的技术,内核此时并不复制整个地址空间,而是让父进程和子进程共享一个拷贝,只有在数据写入的时候,数据才会复制,从而使各个进程拥有自己的拷贝


`aof`重写:`redis`提供`bgwriteaof`指令用于`aof`日志进行瘦身,其原理是开辟一个子进程对内存进行遍历转换成`redis`的操作指令,序列化完毕后再将操作期间产生的增量`aof`日志,追加到新的`aof`日志文件中


#### Redis的同步机制

`Redis`可以使用主从同步,从从同步,第一次同步的时候,主节点进行一次`bgsave`,

#### PipeLine 

`Redis PipeLine`可以一次发送多个命令,并按照顺序返回结果,节省`RTT(ROUND TRIP TIME)`



 
#### string & setnx 和 setex 区别

`setnx`:`set if not exists` 只有当`key`不存在的时候,才可以设置(如果当前中无值,将其设置并返回1,否则返回0)

`setex`:(原子操作)设置值和生存时间,这两个动作会在同一时刻完成,如果key存在会覆盖原有的key.



底层实现:`simple dynamic string (SDS)`

```C
    struct sdshdr{
        int len;        //buf 已经占用的长度
        int free;       //buf 中剩余空间的长度
        char buf[];     //数据空间

    }

```

string内存分配:

当字符串长度小于1M时,扩容是加倍现有的空间,如果超过1M,扩容时只会多扩1M的空间(字符串最大长度为512M)

`SDS`

- 获取字符串长度的复杂度为O(1)
- API是安全的,不会造成缓冲区溢出
- 修改字符串长度N次对多执行N次内存重分配
- 可以保存二进制和文本数据
- 可部分使用<string.h>库中函数



#### List

`QuickList = ZipList + Linkedlist` 

#### Hash(数组 + 链表)



相比于将整个对象序列化后作为string的存储方式,hash能够有效的减少网络传输的消耗


底层实现:hash成员比较少的时候,Redis为了节省内存,会采用类似一维数组的方式来紧凑存储,而不会采用hashMap结构,对应的value redisObject的Encoding为ZipMap
当成员增大的时候会转成HashMap

#### Set

`Redis Set`是无序的,不可重复的string集合


#### SortedSet

`Redis Sorted Set`是有序的不可重复的string集合,`Sorted Set`每个元素都要指定一个score,`Sorted Set`会根据score对元素进行升序排序

实现底层:

`Redis Sorted Set`内部使用的是`HashMap`和跳跃表`SkipList`,来保证数据的存储个有序,`HashMap`里放的成员是到`score`的映射,而跳表里面放的是所有的成员,
排序是依据`HashMap`里面的`score`,使用跳跃表的结构可以获得比较高的查询效率,并且在实现上比较简单



> skipList为什么适用于查找?


`skipList` 就是额外保存了二分查找的中间信息


- skipList 结合了链表和二分查找的思想
- 将原始链表和一些通过'跳跃'生成的链表组成层
- 第0层是原始链表,越上层'跳跃'步距越大,链表元素越少
- 上层链表是下层的子序列
- 查找时从顶层向下,不断减小搜索范围




#### Redis Client

> `Redis`是单线程应用,他是如何与多个客户端建立网络连接并处理命令的?

`Redis`是基于`I/O`多路复用,能够处理多个客户端的请求,`Redis`为每一个连接到`Redis`服务器的客户端创建了一个`RedisClient`结构
这个数据结构包含了每个客户端各自执行的状态和命令.`Redis`服务器使用一个链表来维护多个`redisClient`数据结构






#### Redis 的执行过程


##### 发送命令请求

用户  ->(键入命令)  -> 客户端  -> (将命令请求转换成协议模式) -> 服务器






##### 分布式Redis缓存与数据库一致性

先删除缓存,在更新数据库

**Notice:延时删除**


#### codis

`codis`使用`golang`开发,是一个代理中间件,它和`redis`一样使用`redis`协议提供服务,当客户端向`codis`发送命令的时候
`codis`负责将指令转发到后面`redis`实例来执行,并将返回结果在返回给客户端



`codis`分片原理



#### 过期策略(立即删除 + 定期删除 + 惰性删除)

立即删除:

立即删除能保证内存中数据的最大新鲜度,但是立即删除是对CPU不友好的,因为删除会占用cpu时间,如果刚好碰上CPU很忙,会给CPU造成压力
而且目前redis事件对时间事件处理方式是无序链表

过期`key`集合:`redis`会将每个设置了过期时间的`key`,放到一个独立的字典里面,以后会定时遍历这个字典来删除到期的`key`,除定时遍历删除外,它还会使用惰性策略来删除过期的`key`


**过期时间**:针对过期时间,需要加上一个随机的时间 
- 确保大量请求不会过期打到db上
- 因为大量key过期,redis内部需要时间删除

从库的影响:从库不会进行过期扫描,从库对过期的处理是被动的,主库在`key`到期的时候,会在`aof`文件里面增加一条`del`指令,同步到所有的从库

如果大量的Key还堆积在内存里,导致`redis`内存块耗尽了,此时会走**内存淘汰机制**

#### redis单线程为啥速度还快

- 纯内存操作
- 核心是基于非阻塞的i/o多路复用机制
- 单线程避免了多线程频繁上下文切换






#### Redis对象的Object



- `dictEntry`:redis给每个key-value键值对都分配了一个对应的dictEntry,里面有key和value的指针,next指向下一个dictEntry形成链表,这个指针可以将相同的哈希地址存储在一起
- `sds`:键`key`是以SDS存储
- `redisObject`:值是存储在redisObject中,redisObject中的type字段证明了value的对象模型,ptr指向对象所在的地址



redisObject对象非常重要,redis对象的类型,内部编码,内存回收,共享等功能,都需要redisObject支持,这样的好处是针对不同的使用场景,对5种常用类型设置不同的数据结构实现,从而优化对象在不同场景的使用效率







#### Redis 增量同步 & 快照同步

增量同步:

主节点会将对自己状态产生修改性影响的指令记录在本地内存buffer中,然后异步将buffer中的指令同步到从节点,从节点一边执行同步指令流来达到和主节点一样的状态
一边向主节点反馈自己同步到了哪里

**Notice**
此时有一种场景:因为内存的Buffer是有限的,所以redis主库不能把所有的指令操作都记录在内存Buffer中,redis的复制内存buffer是一个定长的数组,如果数组满了的话,就会从头覆盖前面的内容
如果因为网络不好,从节点无法及时的和主节点同步,redis主节点达到buffer长度可能会将之前的指令覆盖掉,此时从节点无法通过指令流来进行同步,这个时候需要--快照同步



快照同步:

主库进行一次bgsave,从库负责加载生成的rdb文件,加载完成之后再进行增量同步,此时也有一种文体,可能因为某种原因从库加载数据过久
导致主库写入buffer又被再次覆盖,那么主库有需要做bgsave,所以要注意buffer设置的大小





#### `Redis Pipeline`
为什么要使用`pipeline`,对于单一请求的话可以直接进行请求,但是对于多次请求的话,并且如果没多进程的话,
那么必须等待上一条命令应答后在执行,这个中间不仅仅多了`RTT(Round Time Trip)`,而且还频繁的调用系统`i/o`


##### 实现原理
客户端通过一个`TCP`链接发送多条命令


> 从哪方面提升
* `RTT`:`Round Time Trip`节省往返时间
* `I/O`系统调用:因为要进行`Read`操作,节省了从用户态切换至内核态的时间




[mysql-redis双写一致性](https://github.com/doocs/advanced-java/blob/master/docs/high-concurrency/redis-consistence.md)


[redis 执行过程](http://redisbook.com/preview/server/execute_command.html)


[redis 参考](https://redisbook.readthedocs.io/en/latest/internal-datastruct/dict.html)

