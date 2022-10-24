import open3d as o3d
import pandas as pd

from const.scan_state import ScanSate
from helper import fs_helper as fs
from helper.state_handler import die

colum_fmt = ['x', 'y', 'z', 'r', 'g', 'b']
dtype = 'float32'


def check(file_path: str) -> None:
    """
    检查文件情况
    :param file_path:
    :return:
    """
    if not fs.check_file(file_path):
        die(ScanSate.NOT_FOUND, f'can not found {file_path}')


def read_pcd(file_path: str) -> o3d.geometry.PointCloud:
    """
    加载 pcd 格式点云
    :param file_path: 路径
    :return:
    """
    check(file_path)
    try:
        return o3d.io.read_point_cloud(file_path)
    except Exception as e:
        die(ScanSate.NOT_PARSED, f'can not parse {file_path} {e}')


def read_txt(file_path: str, fmt=''.join(colum_fmt)) -> o3d.geometry.PointCloud:
    """
    加载文本格式的点云文件
    :param file_path:
    :param fmt: 点云格式
    :return:
    """
    check(file_path)
    try:
        return o3d.io.read_point_cloud(file_path, format=fmt)
    except Exception as e:
        die(ScanSate.NOT_PARSED, f'can not parse {file_path} {e}')


def read_csv(file_path: str, colum=frozenset(colum_fmt)) -> pd.DataFrame:
    """
    加载 文本格式为 df
    :param file_path:
    :param colum:
    :return:
    """
    check(file_path)
    try:
        return pd.read_csv(file_path, sep=' ', names=colum, dtype=dtype)
    except Exception as e:
        die(ScanSate.NOT_PARSED, f'can not parse {file_path} {e}')


__all__ = [read_pcd, read_txt, read_csv]

