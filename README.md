# ШМЯ Еда

## Как запустить программу
1. [Установить версию Go](https://go.dev/doc/install) >= 1.18
2. `make build`
3. `./bin/shmya-eda`

## Пример вывода программы
```
2022/06/12 20:16:53 january holidays tips hypothesis
2022/06/12 20:16:53 ordersWithoutTips = 2754, ordersWithMoreCutlery = 4163, ordersWithFewerCutlery = 2664
2022/06/12 20:16:53 tipsWithMoreCutlery = 329880, tipsWithFewerCutlery = 136990
2022/06/12 20:16:53 average tips for orders:
        with more than two cutlery = 47
        with less than two cutlery = 51
2022/06/12 20:16:53 users with fewer cutlery (<=2) leave more tips on holidays

2022/06/12 20:16:53 another days tips hypothesis
2022/06/12 20:16:53 ordersWithoutTips = 0, ordersWithMoreCutlery = 10203, ordersWithFewerCutlery = 6273
2022/06/12 20:16:53 tipsWithMoreCutlery = 803280, tipsWithFewerCutlery = 323760
2022/06/12 20:16:53 average tips for orders:
        with more than two cutlery = 78
        with less than two cutlery = 51
2022/06/12 20:16:53 users with more cutlery (>2) leave more tips on another days

2022/06/12 20:16:53 users segment saved to shmya_users_segment.csv
2022/06/12 20:16:53 users segment count = 2541
```