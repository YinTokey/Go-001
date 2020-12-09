启动时开启 http server 监听和 linux 信号监听，然后使用 control + C 触发中断。整个过程控制台打印的结果如下。

```
启动http监听 
启动linux信号监听 
^C
 收到 linux 信号: interrupt
http 关闭 
http server 发生错误 http: Server closed 
```

