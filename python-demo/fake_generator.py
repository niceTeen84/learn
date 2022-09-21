from faker import Faker
import random


def main():
    # "zh_CN"
    fk = Faker()
    for _ in range(10000):
        print(f"('{fk.name()}', {random.randint(0, 90)}, {random.randint(30, 100)}),")


if __name__ == '__main__':
    main()
