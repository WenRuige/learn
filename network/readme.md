#### 网络





> TimeWait怎么产生的?有什么危害?怎么防止?

- 谁先关闭谁产生,如果server端来关闭,则会产生在server端,会持续2个MSL的时长(一般一个msl 30s)
- TimeWait太多产生两类报错,timewait溢出,或者是无法建立新链接(端口号耗尽),如果server端产生,则会引起雪崩无法服务,(例如nginx调用下游服务超时，主动关闭连接，如果访问量很大，可能就产生该问题)
- 避免sever端主动断开链接




> TCP/IP 五元组

源ip地址,目的ip地址,协议号,源端口,目的端口



> 为什么需要三次握手?

为了防止已失效的连接请求报文段突然又传回了服务器端,因而产生错误



> 为什么需要四次挥手

四次挥手常见的状态

- `FIN_WAIT_1`:`FIN_WAIT_1`状态实际上是当`SOCKET`在简历`ESTABLISHED`的状态,它想主动关闭连接,向对方发送了`FIN`报文,此时`SOCKET`即进入了`FIN_WAIT_1`当对方回应`ACK`的时候
- `FIN_WAIT_2`:
- `TIME_WAIT`:表示收到了对方的`FIN`报文,并发送出了`ACK`报文,等2`MSL`后即可回到`CLOSED`可用状态

      