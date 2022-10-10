# NSQ 学习总结

## 安装

```yaml
version: '3'
services:
  nsqlookupd:
    image: nsqio/nsq
    command: /nsqlookupd
    ports:
      - "4160:4160"
      - "4161:4161"
  nsqd:
    image: nsqio/nsq
    command: /nsqd --lookupd-tcp-address=nsqlookupd:4160 --data-path=/data --mem-queue-size=1
    depends_on:
      - nsqlookupd
    ports:
      - "4150:4150"
      - "4151:4151"
    volumes:
      - "d:\\docker-container/nsq-container/data:/data"
  nsqadmin:
    image: nsqio/nsq
    command: /nsqadmin --lookupd-http-address=nsqlookupd:4161
    depends_on:
      - nsqlookupd
    ports:
      - "4171:4171"
```
- `docker-compose up -d` 启动 NSQ
- `docker-compose ps` 查看 NSQ 运行状况
- `docker-compose logs -f` 查看 NSQ 运行日志

## 使用

- 生产一条消息

```bash
curl -d "{\"name\": \"tom\",\"age\": 18}" http://127.0.0.1:4151/pub?topic=renbw&channel=exp
```

- NSQ 的 `channel` 概念类似于传统消息队列（kafka...）的 **消费组**。

- NSQ admin 如果想开启图表模式需要三步设置：
  1. 搭建一个单独的进程用于收集信息, 采用 UDP 协议接受推送数据。eg. [statsdaemon](https://github.com/bitly/statsdaemon)
  2. **nsqd** 启动参数增加 `--statsd-address=<host>:<port>`
  3. **nsqadmin** 启动参数增加 `--graphite-url=http://<host>:<port>`

- NSQ 通过设置 `--mem-queue-size=0` 将所有发布的消息保存在磁盘上。