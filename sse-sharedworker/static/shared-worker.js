// 记个数
let count = 0;
// 把每个连接的端口存下来
const ports = [];

let source = undefined;
// 连接函数 每次创建都会调用这个函数
onconnect = (e) => {
    console.log("这里是共享线程展示位置");
    // 获取端口
    const port = e.ports[0];
    // 把丫存起来
    ports.push(port);
    // 监听方法
    port.onmessage = (msg) => {
        // 这边的console.log是看不到的 debugger也是看不到的 需要在线程里面看
        console.log("共享线程接收到信息：", msg.data);
    };
};

if (!source){
    source = new EventSource('http://127.0.0.1:8081/events/');
    source.onmessage = function(event){
        // console.log(e);
        ports.forEach((p) => {
            // 循环向所有端口广播
            console.log(typeof(event.data))
            console.log(event.data)
            p.postMessage(event.data);
        });
    }
}