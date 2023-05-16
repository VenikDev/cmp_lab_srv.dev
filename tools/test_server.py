import aiohttp
import asyncio


async def main():
    async with aiohttp.ClientSession() as session:
        url = 'http://localhost:80/api/v1/analysis?key=Кровь&city=Асбест'
        async with session.get(url) as resp:
            print(resp.status)
            print(await resp.text())


async def run_main_times():
    tasks = []
    for i in range(10):
        tasks.append(asyncio.ensure_future(main()))
    await asyncio.gather(*tasks)

asyncio.run(run_main_times())
