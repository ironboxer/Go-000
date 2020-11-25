import random


def Subset(backends, client_id, subset_size):
    subset_count = len(backends) // subset_size
    round = client_id // subset_count
    random.seed(round)
    random.shuffle(backends)
    subset_id = client_id % subset_count
    start = subset_id * subset_size
    return backends[start:start + subset_size]


if __name__ == '__main__':
    # client_id 范围 1 ~ 100000
    client_id = random.randint(1, 100000)
    # 100台backends 编号范围在0~10000之间 不连续 未排序
    backends = random.sample(range(10000), 100)
    # client 只和10台后端连接
    subset_size = 10
    for _ in range(10):
        # 传入为backends的拷贝
        subset = Subset(backends[:], client_id, subset_size)
        print(subset)
