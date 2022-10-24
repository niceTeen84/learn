import pandas as pd
import numpy as np
import open3d as o3d
import colorsys
from const.scan_state import COLUM_LIST


def conv_dataframe(pcd: o3d.geometry.PointCloud, names=COLUM_LIST, data_type='float32') -> pd.DataFrame:
    """
    点云转化为 pandas data frame
    :param pcd:
    :param names:
    :param data_type:
    :return:
    """
    return pd.DataFrame(data=np.hstack((np.asarray(pcd.points), np.asarray(pcd.colors))),
                        columns=names, dtype=data_type)


def conv_pcd(data) -> o3d.geometry.PointCloud:
    """
    data frame 转化为点云
    :param data: n * 6 二维矩阵
    :return: pcd
    """
    pcd = o3d.geometry.PointCloud()
    if len(data) == 0: return pcd

    arr = data if not isinstance(data, pd.DataFrame) else data.to_numpy()[:, :6]

    pcd.points = o3d.utility.Vector3dVector(arr[:, :3])
    pcd.colors = o3d.utility.Vector3dVector(arr[:, 3:])
    return pcd


def conv_rgb_to_hsl(df: pd.DataFrame) -> None:
    """
    rgb 转化 hls
    在原有 dataframe 上增加  3 列
    :param df:
    :return:
    """
    df[['h', 'l', 's']] = df.apply(lambda row: colorsys.rgb_to_hls(row['r'], row['g'], row['b']),
                                   axis=1,
                                   result_type='expand')


def combine_pcd(*pcds: o3d.geometry.PointCloud):
    """
    合并多个点云
    :param pcds: 点云集合
    :return: 合并后的点云数据
    """
    merged_pcd = o3d.geometry.PointCloud()
    if len(pcds) == 0: return merged_pcd

    point_list = [pcd.points for pcd in pcds]
    color_list = [pcd.color for pcd in pcds]

    merged_pcd.points = o3d.utility.Vector3dVector(np.concatenate(point_list))
    merged_pcd.colors = o3d.utility.Vector3dVector(np.concatenate(color_list))

    return merged_pcd


if __name__ == '__main__':
    pass
