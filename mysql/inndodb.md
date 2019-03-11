#### Mysql InnoDB




支持的索引：
* hash(自适应)
* b+tree
* fulltext



##### `InnoDB` 的索引


`InnoDB`的主键索引和行记录是存储在一起的,所以叫做聚集索引

* 没有单独区域存储行记录
* 主键索引的叶子节点,存储主键并且对应行记录(mysiam对应的是指针)




因为这个特性,mysql必须要有主键
* 如果主键定义了(pk),则使用主键作为聚集索引
* 如果没定义主键,会取第一个非空uniq列作为聚集索引
* 否则,`InnoDB`会创建一个隐藏的`row-id`作为聚集索引




`redo`日志为了保证,已提交事务`acid`的特性
`undo`日志为了保证,未提交事务`acid`的特性











[文章](https://mp.weixin.qq.com/s?__biz=MjM5ODYxMDA5OQ==&mid=2651961494&idx=1&sn=34f1874c1e36c2bc8ab9f74af6546ec5&chksm=bd2d0d4a8a5a845c566006efce0831e610604a43279aab03e0a6dde9422b63944e908fcc6c05&scene=21#wechat_redirect)