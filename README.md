# ШМЯ Еда

## Как запустить программу
1. [Установить версию Go](https://go.dev/doc/install) >= 1.18
2. `go run *.go`

## Пример вывода программы
```
2022/06/23 19:13:20 january holidays tips hypothesis
2022/06/23 19:13:20 ordersWithoutTips = 2754, ordersWithMoreCutlery = 3713, ordersWithFewerCutlery = 2369
2022/06/23 19:13:20 tipsWithMoreCutlery = 294870, tipsWithFewerCutlery = 121300
2022/06/23 19:13:20 average tips for orders:
        with more than two cutlery = 45
        with less than two cutlery = 51
2022/06/23 19:13:20 users with fewer cutlery (<=2) leave more tips on holidays

2022/06/23 19:13:20 another days tips hypothesis
2022/06/23 19:13:20 ordersWithoutTips = 0, ordersWithMoreCutlery = 10653, ordersWithFewerCutlery = 6568
2022/06/23 19:13:20 tipsWithMoreCutlery = 838290, tipsWithFewerCutlery = 339450
2022/06/23 19:13:20 average tips for orders:
        with more than two cutlery = 78
        with less than two cutlery = 51
2022/06/23 19:13:20 users with more cutlery (>2) leave more tips on another days

2022/06/23 19:13:20 users segment saved to shmya_users_segment.csv
2022/06/23 19:13:20 users segment count = 2607
```