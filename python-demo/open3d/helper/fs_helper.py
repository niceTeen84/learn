import os
import shutil


def check_file(src: str) -> bool:
    return os.path.exists(src)


def mv(src: str, dst: str) -> None:
    """
    文件重命名
    :param src: 原始文件 path
    :param dst: 重命名文件 path
    :return: None
    """
    if not check_file(src):
        raise Exception(f'file {src} not found')
    if check_file(dst):
        rm(dst)
    os.rename(src, dst)


def rm(src: str) -> None:
    """
    文件删除
    :param src:
    :return:
    """
    if not check_file(src):
        raise Exception(f'file {src} not found')
    os.remove(src)


def du(src: str) -> tuple:
    """
    查看文件大小
    :param src: 文件路径
    :return:
    """
    if not check_file(src):
        raise Exception(f'file {src} not found')
    return shutil.disk_usage(src)


def rename_to_asc(src: str):
    return src.replace('xyzrgb', 'asc')


def pwd() -> str:
    """
    获取当前路径位置
    :return:
    """
    return os.getcwd()
