import http.server
import socketserver
import os

# 指定静态文件的目录（可根据实际情况修改）
directory = 'static'

# 设置服务器的地址和端口
server_address = ('127.0.0.1', 8888)

# 创建静态文件处理器
handler = http.server.SimpleHTTPRequestHandler

# 切换到指定目录
os.chdir(directory)

# 开启文件服务器
def runserver():
    with socketserver.TCPServer(server_address, handler) as httpd:
        print(f"Server running on http://{server_address[0]}:{server_address[1]}/")
        httpd.serve_forever()