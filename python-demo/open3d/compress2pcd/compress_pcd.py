import open3d as o3d
import pandas as pd
import numpy as np
import os, re


SCAN_DATA_ROOT = 'D:\\qingdao\\青岛\\导出'
COLUMNS = ['x', 'y', 'z', 'r', 'g', 'b']

FLOAT = 'float32'
INT16 = 'int16'
SPACE = ' '


def convert_rgb(f_path: str) -> np.ndarray:
    """
    convert source file rgb color value divide 255
    and color range is 0 - 1
    :param f_path: the scan file abs path
    :return: numpy ndarray
    """
    type_dict = {elm: FLOAT if idx < 3 else INT16 for idx, elm in enumerate(COLUMNS)}
    df: pd.DataFrame = pd.read_csv(f_path,
                                   sep=SPACE,
                                   names=COLUMNS,
                                   dtype=type_dict)
    for k in COLUMNS[3:]:
        df[k] = df[k] / 255
    return df.to_numpy(dtype=FLOAT)


def convert_to_pcd(arr: np.ndarray, out: str) -> None:
    """
    convert numpy ndarray to open3d geometry point cloud object
    :param arr: numpy array dtype is float32
    :param out export pcd compressed file path
    :return: None
    """
    if arr is None or len(arr) == 0: return
    pcd = o3d.geometry.PointCloud()
    pcd.points = o3d.utility.Vector3dVector(arr[:, :3])
    pcd.colors = o3d.utility.Vector3dVector(arr[:, 3:])
    o3d.io.write_point_cloud(out, pcd, compressed=True, print_progress=True)


def main():
    # walk dir
    for root, _, files in os.walk(SCAN_DATA_ROOT, topdown=False):
        for name in files:
            out, abs_path = re.sub('(xyz)$', 'pcd', name), str(os.path.join(root, name))
            arr = convert_rgb(abs_path)
            convert_to_pcd(arr, f'{root}\\{out}')
            del arr


if __name__ == '__main__':
    main()
