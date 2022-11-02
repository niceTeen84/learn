from os import path

def main():
    file_path = 'D:\\qingdao\\SCENE_LT_2022.1.0.9661_Setup.exe'
    # 获取文件名
    print(path.basename(file_path))
    # 得到一个元组 ('D:\\qingdao', 'SCENE_LT_2022.1.0.9661_Setup.exe')
    print(path.split(file_path))
    # 获取文件夹的名称 D:\qingdao
    print(path.dirname(file_path))
    # 获取一个元组 ('D:\\qingdao\\SCENE_LT_2022.1.0.9661_Setup', '.exe')
    # 用来得到文件扩展名成
    print(path.splitext(file_path))
    # print(path.splitunc(file_path))
    print(path.splitdrive(file_path))


if __name__ == '__main__':
    main()
