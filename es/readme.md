#### es



`es`中存储数据的基本单位是索引

```
index -> type -> mapping -> document -> field
```




- `index`定义:相当于`mysql`关系型数据库中的`database`,是一个独立的`Lucene`实例
- `type`定义: 相当于关系型数据库的`table`,一个索引有多个`type`
- `document`定义:相当于关系型数据库的`row`,文档由字段构成,是一个`json`对象
- `field`定义:相当于关系型数据库的`column`,每个字段都有它对应的类型
- `Mapping`定义:相当于关系型数据库的`schema`,一个`type`对应一个`mapping`






`es`写数据过程

- 客户端选择一个`node`发送请求过去,这个`node`就是`coordinating node`(协调节点)
- `coordinating node`对`document`进行路由,将请求转发给对应的`node`
- 实际`node`上的`primary shard` 处理请求,然后将数据同步到`replica node`
- `coordinating node` 如果发现`primary node`和`replica node`都搞定之后,就返回响应给客户端