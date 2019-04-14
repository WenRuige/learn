#### `Mac docker install`

略



#### `docker install centos`

1.下载`centOS 7`镜像
```
docker pull centos:7
```
2.进入到`centOS`镜像内
```
docker run --privileged=true -it centos:7
```
3.安装一些软件
```
like zsh/git/nginx/ etc
```
4.保存作案现场
```
1. docker ps
2. docker commit (container id) centos:7
```


##### 一些参数

目录挂载`-v`


`docker run --privileged=true -it  -v /Users/gewenrui/Desktop/docker:/home centos:7`




##### 遇到的问题

1.登录问题.

`docker login` 输入的用户名非邮箱,而是`username`
