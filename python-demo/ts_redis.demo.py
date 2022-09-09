from redis.commands import timeseries as ts
import redis


METRICS_KEY = 'sensor_1'


def create_ts_key(cli):
    ts_cli = ts.TimeSeries(cli)
    meta_data = {
        'retention_msecs': 86400_000,  # one day TTL
        'chunk_size': 1048576,  # 最大 1MB
        'labels': {
            'desc': 'just for fun',
            'date': '2022-09-09'
        }

    }
    ts_cli.create(METRICS_KEY, **meta_data)
    print('done')


# pip install redis
def main():
    pool = redis.ConnectionPool(port=6378)
    cli = redis.Redis(connection_pool=pool)
    if not cli.exists(METRICS_KEY):
        create_ts_key(cli)
    # add sample data
    ts_cli = ts.TimeSeries(cli)
    meta = {
        'duplicate_policy': 'max'
    }
    for i in range(1000):
        ts_cli.add(METRICS_KEY, '*', 36.5, **meta)


if __name__ == '__main__':
    main()
