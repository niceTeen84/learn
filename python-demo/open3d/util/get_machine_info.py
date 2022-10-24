#!/usr/bin/env python
"""
常用的系统工具函数
"""

import socket
import subprocess
import uuid
import os
import sys
from subprocess import check_output, check_call
import re


def get_mac_address():
    """
    获取本机 MAC 地址
    :return:
    """
    mac = uuid.UUID(int=uuid.getnode()).hex[-12:]
    return str(':'.join([mac[e:e + 2] for e in range(0, 11, 2)])).upper()


def get_process_pid():
    """
    获取 python 进程 pid
    :return:
    """
    return os.getpid()


def get_local_ip():
    """
    获取本地内网 ip
    :return:
    """
    hostname = socket.gethostname()
    return socket.gethostbyname(hostname)


def exec_cmd(cmd: list):
    # stdout 禁止打印
    ret = check_call(cmd, stdout=subprocess.DEVNULL)
    if ret == 0:
        return str(check_output(cmd), 'utf-8').strip()


def get_java_version():
    """
    检查 java version
    :return:
    """
    res = exec_cmd(['java', '--version'])
    m = re.match('^java\s([.\d]+)\s[-\w]+', res)
    if m:
        return m.group(1)


if __name__ == '__main__':
    print(get_java_version(), file=sys.stderr)
