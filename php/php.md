####









#### php-fpm


fpm实现是创建一个master进程,在master进程中创建并且监听socket,然后fork出多个子进程
这些子进程accept请求,子进程的处理非常简单,它在启动后阻塞在accept




[php](https://github.com/pangudashu/php7-internal/blob/master/1/fpm.md)





