import zlib
import pymysql
from pymysql.cursors import DictCursor
from dotenv import load_dotenv
import os

CHARSET = 'utf-8'
TEST_STR = 'hello 🦍'

fn_val = lambda k: os.getenv(k.upper())


def main():
    load_dotenv()
    info = {k: fn_val(k) if k != 'port' else int(fn_val(k)) for k in ['host', 'user', 'password', 'database', 'port']}
    conn = pymysql.connect(**info)
    with conn:
        with conn.cursor(DictCursor) as cur:
            cur.execute(f"select compress(repeat('{TEST_STR}', 100)) res")
            res = cur.fetchone()
            # mysql compress
            """
            Nonempty strings are stored as a 4-byte length of the uncompressed string (low byte first),
             followed by the compressed string. If the string ends with space, an extra . 
             character is added to avoid problems with endspace trimming 
             should the result be stored in a CHAR or VARCHAR column.
              (However, use of nonbinary string data types such as CHAR or VARCHAR to store
               compressed strings is not recommended anyway because character set conversion may occur.
                Use a VARBINARY or BLOB binary string column instead.)
            """
            result = zlib.decompress(res['res'][4:]).decode(CHARSET)
    # 对比
    r = zlib.compress((TEST_STR * 100).encode(CHARSET))
    print('bytes to hex: ', r.hex())
    print('revert hex to bytes: ', bytes(bytearray.fromhex(r.hex())))
    assert result == zlib.decompress(r).decode(CHARSET)


if __name__ == '__main__':
    main()
