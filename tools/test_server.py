import aiohttp
import asyncio


async def main(idx):
    async with aiohttp.ClientSession() as session:
        url = 'http://192.168.31.200:80/api/v1/analysis?key=ковид&city=Асбест'
        async with session.get(url) as resp:
            print("Status = {}, idx = {}".format(resp.status, idx))
            print(await resp.text())


async def run_main_times():
    tasks = []
    for i in range(1000):
        tasks.append(asyncio.ensure_future(main(i)))
    await asyncio.gather(*tasks)

asyncio.run(run_main_times())
