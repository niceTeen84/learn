import os

from sqlalchemy import create_engine
from dotenv import load_dotenv


def main():
    load_dotenv()
    engine = create_engine(f"mysql+pymysql://{os.getenv('USER')}:{os.getenv('PASSWORD')}@{os.getenv('HOST')}:{os.getenv('PORT')}/{os.getenv('DATABASE')}",
                           pool_size=20, max_overflow=10)
    with engine.connect() as conn:
        res = conn.execute('select version()')
        r = res.fetchone()
        print(r)


if __name__ == '__main__':
    main()
