代码说明:

通过shared workers技术完成浏览器同源的多标签页websocket连接共享

执行以下命令

```
python websockets_server.py
```

并访问

[page1: http://127.0.0.1:8888/index1.html](http://127.0.0.1:8888/index2.html)

[page2: http://127.0.0.1:8888/index2.html](http://127.0.0.1:8888/index2.html)

查看控制台即可看到同步的消息

另外通过浏览器(地址:edge://inspect/#workers或者chrome://inspect/#workers，视浏览器而定)可以查看共享server的打印信息
