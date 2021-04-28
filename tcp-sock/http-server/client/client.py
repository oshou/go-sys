#!/usr/bin/python

import asyncio
import time
import requests
import logging


async def run(loop):
    def req():
        requests.get('http://192.168.10.3:8080')
        time.sleep(100)

    async def run_req(i):
        return await loop.run_in_executor(None, req(), i)

    tasks = [run_req(i) for i in range(100)]
    return await asyncio.gather(*tasks)


logging.basicConfig(level=logging.DEBUG)
st = time.time()
loop = asyncio.get_event_loop()
print(loop.run_until_complete(run(loop)))
print(time.time() - st)
