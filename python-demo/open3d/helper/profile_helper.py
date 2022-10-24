import time
import logging
import numpy as np


__all__ = ['start', 'stop']


def start(desc=''):
    if desc:
        content = f'{desc} start'
        print(content)
        logging.info(content)
    return time.time()


def stop(s, desc=''):
    content = f'{desc} stopped cost {np.round(time.time() - s, 4)} seconds'
    print(content)
    logging.info(content)
