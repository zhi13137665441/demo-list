代码说明:

通过shared workers技术完成浏览器同源的多标签页websocket连接共享

执行以下命令

```
go run main.go --port==1234
go run main.go --port==1235
go run main.go --port==1236
```
把./nginx/nginx.conf放到nginx对应目录重启nginx，开启websocket的负载均衡
(注意备份本机的nginx配置哦,静态文件路由配置本项目下static)

并访问[http://localhost/static/index.html](http://localhost/static/index.html)

查看控制台即可看到同步的消息
