#### Tell

如何讲述项目?


比如说面试官问你,最有难度的一个项目,如何给面试官讲述



拿ZONE来说


* 首先讲述ZONE背景(这是一个基于LBS的商圈服务)
* 应用场景(计算在线时长,获取商圈信息,商家吸附策略,营销发券圈etc)
* 商圈如何产生(策略切地块产生)
* 算法实现(射线法,点到线段距离,geohash等)



射线法实现,一条射线穿过多边形交点为偶数个则在多边形外,否则则在多边形内

以下是需要考虑几种临界情况

* 点在圈上
* 点在多边形顶点上
* 射线刚好经过凸多边形两条相邻边的交点


* 判断点是否在线上,计算点与两个多边形顶点的连线斜率是否相等
* 点是否与多边形顶点重合
* 点穿越实现(线段的两个顶点分别在射线的两侧)




点到商圈(多边形)距离

* 点到多边形距离计算
* 点到直线距离

Q:TRIP-NODE 是一个什么服务?

A:TRIP-NODE 是一个数据服务,负责指导骑手在当前节点做什么事情,
用于: 引擎,骑手,查询历史行程


在其中担当什么角色? 负责行程服务代码编写和搭建.表的设计?



* 缓存一致性问题
* 大数据量查询问题
* 并发改造优化
* msgp 使用: 为什么用msgp,msgp为何比json小
* redis pipeline



数据库分表,如何分

数据库索引:

* 数据库索引是B+ TREE

什么是b+tree
* 多叉树












[文章](https://juejin.im/post/5c63f52a6fb9a049d519ffa1)



https://www.nowcoder.com/discuss/129660?type=0&order=0&pos=20&page=1
https://www.nowcoder.com/discuss/145338?type=2
https://leetcode-cn.com/explore/interview/card/top-interview-quesitons-in-2018/263/simulation/1197/

https://halfrost.com/http/
https://zhuanlan.zhihu.com/p/32540678
https://halfrost.com/tcp_ip/