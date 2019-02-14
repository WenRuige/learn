##### get


```
{"get",getCommand,2,"rF",0,NULL,1,1,1,0,0},
```

参数:
- `r`:`readCommand (never modify the key space)`
- `F`:`FastCommand`





```
int getGenericCommand(client *c) {
    robj *o;

    if ((o = lookupKeyReadOrReply(c,c->argv[1],shared.nullbulk)) == NULL)
        return C_OK;

    if (o->type != OBJ_STRING) {
        addReply(c,shared.wrongtypeerr);
        return C_ERR;
    } else {
        addReplyBulk(c,o);
        return C_OK;
    }
}
```

`shared.nullbulk`是在`redis`服务器启动时候创建的一个共享变量,


```
robj *lookupKey(redisDb *db, robj *key, int flags) {
    // 在字典中根据key查找字典对象
    dictEntry *de = dictFind(db->dict,key->ptr);
    if (de) {
        // 获取字典对象的值想
        robj *val = dictGetVal(de);

        /* Update the access time for the ageing algorithm.
         * Don't do it if we have a saving child, as this will trigger
         * a copy on write madness. */
        if (server.rdb_child_pid == -1 &&
            server.aof_child_pid == -1 &&
            !(flags & LOOKUP_NOTOUCH))
        {
            if (server.maxmemory_policy & MAXMEMORY_FLAG_LFU) {
                unsigned long ldt = val->lru >> 8;
                unsigned long counter = LFULogIncr(val->lru & 255);
                val->lru = (ldt << 8) | counter;
            } else {
                val->lru = LRU_CLOCK();
            }
        }
        return val;
    } else {
        return NULL;
    }
}
```

执行流程



- 传入参数
- 查询Key
- 判断key是否过期
- 找到value