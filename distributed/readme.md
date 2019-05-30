#### 数据库的一致性实现
* 原子性(`Atomicity`)
* 一致性(`Consistency`)
* 隔离性(`Isolation`)
* 持久性(`Durability`)


#### CAP理论

* 一致性(`Consistency`):客户端知道一系列的操作都会同时生效
* 可用性(`Availability`):每个操作都必须认可以可预期的响应结束
* 分区容错性(`Partition tolerance`):即使出现单个组件无法可用,操作依然可以完成



#### BASE理论
* `Basic Availability`:(基本可用)
* `Soft State`:(软状态)
* `Eventually Consistency`:(最终一致性)


`base`理论是对`cap`中的一致性和可用性进行一个权衡的结果
理论的核心思想是:我们无法做到强一致,但是每个应用都可以根据自身的业务特点,采用适当的方式使系统达到最终一致性



#### 分布式事务

##### 两阶段提交(2PC)




##### 补偿事务(TCC)



##### 提供回滚接口



##### 本地消息表










[参考文档-cap和base理论实践](https://www.cnblogs.com/savorboard/p/distributed-system-transaction-consistency.html)

[infoQ分布式一致性](https://www.infoq.cn/article/solution-of-distributed-system-transaction-consistency)

[数据缓存一致性](https://weibo.com/ttarticle/p/show?hmsr=toutiao.io&id=2309404022116222639373  )