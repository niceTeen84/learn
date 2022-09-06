import asyncio
import time
from concurrent.futures import ProcessPoolExecutor


def calc(num: int):
    print(f'start cpu heavy task {num}')
    time.sleep(10)
    print(f'task {num} done')
    return num


async def main(ev_loop):
    """
    the main process function
    :param ev_loop:
    :return:
    """
    print('start main process')
    th_pool = ProcessPoolExecutor(max_workers=16)
    data = await asyncio.gather(*(ev_loop.run_in_executor(th_pool, calc, num) for num in range(32)))
    print(data)


if __name__ == '__main__':
    loop = asyncio.get_event_loop()
    loop.run_until_complete(main(loop))
