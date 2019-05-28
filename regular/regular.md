#### 正则的学习




> 什么是`AWK`

`AWK`是一门解释行编程语言,在文本处理领域是非常强大的.
它的名字来源于三位作者的姓氏  `Alfred Aho` `Peter Weinberger` 和 `Brian Kernighan`





> `AWK`程序结构

* `BEGIN`语句块

    ```bash

      BEGIN {awk-commands}
    ```
    `begin`语句块在程序开始的时候执行,在这里可以初始化变量,


* `BODY`语句块
    ```bash
    /pattern/ {awk-commands}
    ```

    `body`语句块没有关键字


* `END` 语句块
    ```
       END {awk-commands}
    ```
    `END`语句块在最后执行