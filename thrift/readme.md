

#### `thrift`架构








#### 协议

* `TBinaryProtocol`:二进制格式
* `TCompactProtocol`:压缩格式
* `TJSONProtocol`:json格式
* `TSimpleJSONProtocol`:提供json只写协议,生成文件容易通过脚本解析




#### 传输
`thrift`最底层传输可以使用socket,file和zip来实现
* TSocket:使用阻塞的TCP Socket进行数据传输,也是最常见的模式
* THttpTransPort:使用http传输协议进行数据传输
* TFileTransPort:文件(日志)传输,允许client将文件交给server,允许将server收到的数据写到文件中
* TZilbTransPort:与其他的TTransport配合使用，压缩后对数据进行传输，或者将收到的数据解压   




#### 通用

* 底层i/o模块,负责实际的数据传输,传输方式包括(Socket,文件,压缩数据)
* transport层负责以字节   流方式发送和接受消息,是底层i/o模块在thrift框架的实现,
* TProtocol主要负责结构化数据组装成message,或者从message里面读出结构化数据
* TServer负责接受client请求,并将请求转发给processor进行处理,TServer主要任务是高效的接受Client的请求,特别是高并发请求
* Processor 负责对client的请求进行响应,包括RPC请求转发,调用参数解析和用户逻辑调用,返回值写回等操作 



#### Server 模型

* TSimpleServer:使用阻塞i/o单线程服务器,主要用于调试
* TThreadServer:使用阻塞i/o多线程服务器,每一个请求都在线程内处理,并发情况下会有很多线程
* TThreadServer:使用阻塞i/o多线程服务器,使用线程池管理线程
* TNonBlockServer:使用非阻塞i/o多线程服务器,使用少量线程既可以完成大并发量请求响应,必须使用TFrameTransport


#### 参考文献

[挺好的文章](http://www.cnblogs.com/shangxiaofei/p/8504932.html)

https://www.ibm.com/developerworks/cn/java/j-lo-apachethrift/





https://www.cnblogs.com/chenny7/p/4224720.html
http://thrift.apache.org/