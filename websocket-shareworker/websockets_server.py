import asyncio
import datetime
import time
import concurrent

import websockets
from file_server import runserver

executor = concurrent.futures.ThreadPoolExecutor()
# 提交任务到线程池执行
future = executor.submit(runserver)


# 服务器端的处理逻辑
async def handle_message(websocket, path):
    await websocket.send("success!")
    try:
        while True:
            print(str(datetime.datetime.now()))
            time.sleep(1)
            await websocket.send(str(datetime.datetime.now()))
    except websockets.ConnectionClosedError:
        await websocket.close()


# 创建WebSocket服务器并启动
start_server = websockets.serve(handle_message, "localhost", 8765)

asyncio.get_event_loop().run_until_complete(start_server)
print("runserver on port:", "http://127.0.0.1:8765")
asyncio.get_event_loop().run_forever()


