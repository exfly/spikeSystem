# bench
https://cloud.tencent.com/developer/article/1652536

## redis-benchmark

<!-- https://eli.thegreenplace.net/2019/unix-domain-sockets-in-go/ -->

unix socket : loopback = 2:1

> https://redis.io/topics/benchmarks
redis-benchmark
SET 100200.40 GET 102040.81

redis-benchmark -P 2
SET 194564.20 GET 195320.30

redis-benchmark -P 16
SET 944452.81 GET 1065191.50

## http bench

go run cmd/base/main.go

ab -n 10000 -c 1000 http://localhost:3005/simple

16473.79 - 19010.14

go run main.go

ab -n 10000 -c 1000 http://localhost:3005/buy/ticket

有库存时候，4719（需要跟 redis 交互，减库存）
没有库存的时候 18777（内存操作）

ab -n 10000 -c 10000 http://localhost:3005/buy/ticket
没库存 13000


ab -n 20000 -c 20000 http://localhost:3005/buy/ticket
没库存 1070.34

## postgres

pgbench

600 clients 情况下，大约 7000左右

大致性能为 redis 的十分之一或者更低
