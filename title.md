考察项目：nosql、memcache、redis、leveldb、hbase、cassandra、hadoop、mongodb、文档型数据库、zookeeper

语言类：虚析构函数、对象的内存模型、函数对象，仿函数、智能指针、设计模式，桥接模式，策略模式，工厂模式

什么是函数签名

int A(int x, int y)，从接口设计层面保证x和y不会被传反         

算法类：

  内部排序，外部排序，分别列一下

  快速排序，一趟排序后的结果

  归并排序，一堆有序的文件，最终合并成一个有序的文件，加上内存限制

  堆排序，取最小值，最小堆，实现一个定时器

  trie排序，字典树，ip地址判断

  红黑树，val树的区别

  位图算法，加法，乘法，二元方程等

  布隆过滤器

  LRU算法

  倒排索引

  ACID

  CAP理论

  BASE

  一致性协议，paxos，raft

  一致性hash

STL类：

   size和empty的区别

   map的底层数据结构

   内存池

   deque的底层数据结构

 网络类：

  fd的本质

  进程线程模型

  网络io

  磁盘io查看

  程序热点查找

  内存泄漏检测

  write和send的区别

  建立监听套接字

  三次握手，四次握手

  为什么要四次握手？

  全双工，半关闭

  主动关闭，被动关闭

  http协议

  libevent

  io多路复用，select， epoll的触发模式

  长链接，短链接

  什么是rpc，能不能并发

  thrift

  grpc

  brpc

系统类：

  信号的本质

  进程的内存布局

  机械硬盘，ssd

  io优化，随机io，顺序io

  页缓存page cache

  write，read调用

  多线程，线程间通讯

  多进程，无锁编程

  什么是断点

  gdb断点调试，多线程调试，多进程调试

  动态库的版本管理，主版本，次版本

  soname的作用

  软连接，硬链接

  文件锁

  读写锁

  条件变量，生产者，消费者

  xen,kvm,docker

  shell实现单词个数统计并排序

  awk,sed,cut,uniq,sort,tr,tee,

压测的步骤：

  1、步骤 1~6步

  2、什么是吞吐量和并发?

  3、响应时间、错误率

设计一个定时器：

   红黑树

   epoll

   select

设计一个id分配器：

   自增

   分布式

设计一个快速ip查找类：

   字典树

   基数数

版本控制：

  svn分支作用，合并的本质，tag

  git的用法

  gitflow 工作流

公钥私钥：

  公钥私钥的原理

  数字签名

  OAuth2.0

