# ffmpeg 推流

## 安装软件
1. 安装 Nginx ffmpeg libnginx-mod-rtmp

```shell
sudo apt install nginx ffmpeg libnginx-mod-rtmp -y
```
2. 配置 Nginx

```shell
# 注释http内部的include
#include /etc/nginx/conf.d/*.conf;
# 外部包含 rtmp.conf
include /etc/nginx/conf.d/*.conf;
```
3. 推流

```shell
ffmpeg -re -i demo.mp4 -vcodec copy -f flv rtmp://localhost:1935/rtmplive/home
```

4. 测试
安装 VLC 打开拉流地址 `rtmp://192.168.0.150:1935/rtmplive/home`

## 摄像头推流
1. 获取设备名称

```shell
./ffmpeg.exe -list_devices true -f dshow -i dummy
```

2. 摄像头推流

```shell
./ffmpeg.exe -f dshow -i video="Integrated Camera" -vcodec libx264 -acodec copy -preset:v ultrafast -tune:v zerolatency -f flv rtmp://192.168.0.150:1935/rtmplive/home
```

3. 测试

```shell
 real-time buffer [Integrated Camera] [video input] too full or near too full (136% of size: 3041280 [rtbufsize parameter])! frame dropped!
```
VLC 拉流时会出现上述错误

[stackoverflow ref](https://stackoverflow.com/questions/45643572/ffmpeg-problems-with-real-time-buffer)
```shell
# -rtbufsize 1G
./ffmpeg.exe -f dshow -rtbufsize 1G -i video="Integrated Camera" -vcodec libx264 -acodec copy -preset:v ultrafast -tune:v zerolatency -f flv rtmp://192.168.0.150:1935/rtmplive/home
```