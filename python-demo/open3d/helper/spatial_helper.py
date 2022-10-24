"""
空间计算 helper
"""
import open3d as o3d
import numpy as np
import colorsys
from scipy.spatial.distance import pdist

import pandas as pd

from dataclasses import dataclass


@dataclass
class EdgeInfo:
    x_len: float
    y_len: float
    z_len: float
    coord_list: np.ndarray


def compute_two_points_dist(a: tuple, b: tuple):
    """
    两点间距离
    :param a: 点 a
    :param b: 点 b
    :return: float
    """
    return pdist(np.asarray([a, b]))


def compute_segment_dist(df: pd.DataFrame, formula: tuple):
    """
    计算所有点到平面的距离
    :param df: dataframe
    :param formula: 平面方程 (a, b, c, d) ax + by + cz + d = 0
    :return: 直接改变 df 增加一列 dist
    """
    a, b, c, d = formula
    df['dist'] = df['x'] * a + df['y'] * b + df['z'] * c + d


def compute_pcd_edge(pcd: o3d.geometry.PointCloud) -> EdgeInfo:
    """
    计算点云边界
    :param pcd:
    :return:
    """
    box: o3d.geometry.AxisAlignedBoundingBox = pcd.get_axis_aligned_bounding_box()
    x, y, z = box.get_extent()
    pts = np.asarray(box.get_box_points())
    return EdgeInfo(x, y, z, pts)


def compute_axis_normal():

    pass


def pick_color(pcd: o3d.geometry.PointCloud, hue: float, lightness=1., saturation=1.):
    """
    选取颜色
    :param pcd: 点云数据
    :param hue: 色值
    :param lightness: 亮度
    :param saturation: 饱和度
    :return:
    """
    sample_pcd: o3d.geometry.PointCloud = pcd.random_down_sample(0.2)
    # sample_pcd: o3d.geometry.PointCloud = pcd.uniform_down_sample(5)
    # pcd.cluster_dbscan()
    df = pd.DataFrame(data=np.hstack((np.asarray(sample_pcd.points), np.asarray(sample_pcd.colors))),
                      columns=['x', 'y', 'z', 'r', 'g', 'b'])
    df[['h', 'l', 's']] = df.apply(lambda row: colorsys.rgb_to_hls(row['r'], row['g'], row['b']),
                                   axis=1,
                                   result_type='expand')
    res: pd.DataFrame = df.query(f'h > {hue - 8 * 0.01} & h < {hue + 4 * 0.01} & s > 0.2 and s < 0.9')

    arr, case_pcd = res.to_numpy(dtype='float32')[:, :6], o3d.geometry.PointCloud()
    del df
    del sample_pcd
    case_pcd.points = o3d.utility.Vector3dVector(arr[:, :3])
    case_pcd.colors = o3d.utility.Vector3dVector(arr[:, 3:])

    case_pcd, _ = case_pcd.remove_statistical_outlier(100, 1)
    labels = np.asarray(case_pcd.cluster_dbscan(0.1, 100))

    mx = np.max(labels)
    for i in range(mx + 1):
        ind = np.where(labels == i)[0]
        part_pcd: o3d.geometry.PointCloud = case_pcd.select_by_index(ind)
        part_pcd, _ = part_pcd.remove_statistical_outlier(200, 1)
        c: np.ndarray = part_pcd.get_center()
        print(c)

        # o3d.io.write_point_cloud(f'scan_data/faro/inner-wall-part{i}.pcd', part_pcd, compressed=True)
        o3d.visualization.draw_geometries([part_pcd])
        # calc_with_and_height(part_pcd)
        print(f'part {i} done')

    o3d.visualization.draw_geometries([case_pcd])
    pass



if __name__ == '__main__':

    pass

    # pcd = o3d.io.read_point_cloud('../scan_data/faro/inner-wall.pcd')
    # # # pick_color(pcd, 0.16)
    # compute(pcd)
    # 色相 hue 60 rad 60 / 360
    # 亮度 lightness 0.5
    # 饱和度 saturation 25 - 90



    # df = pd.read_csv('../scan_data/faro/out-wall.xyz', sep=' ', names=['x', 'y', 'z', 'r', 'g', 'b'])
    # # df['hue'] = colorsys.rgb_to_hls(df['r'], df['g'], df['b'])
    # df['hue'] = df.apply(lambda row: colorsys.rgb_to_hls(row['r'], row['g'], row['b'])[0], axis=1)
    # res = df.query('hue > 0.14 & hue < 0.17')
    # print()
    # hls = colorsys.rgb_to_hls(226/ 255, 217/255, 150/255)
    # print(hls)
    # # 219;213;151
    #
    # hls = colorsys.rgb_to_hls(219 / 255, 213 / 255, 151 / 255)
    # print(hls)
    #
    # # 115;92;83
    # hls = colorsys.rgb_to_hls(115 / 255, 92 / 255, 83 / 255)
    # print(hls)
    #
    # # 199;95;100
    # hls = colorsys.rgb_to_hls(199 / 255, 95 / 255, 100 / 255)
    # print(hls)
    # for i in range(101):
    #     rgb = colorsys.hls_to_rgb(0.16666666, 0.25, 1.0)
    #     print(rgb)