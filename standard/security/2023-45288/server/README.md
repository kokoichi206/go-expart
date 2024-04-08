## x/net

https://pkg.go.dev/golang.org/x/net@v0.24.0

https://pkg.go.dev/golang.org/x/net@v0.24.0?tab=versions

``` sh
# 脆弱性対応版。
go get -u golang.org/x/net@v0.24.0
# go get -u golang.org/x/net@latest

# 脆弱性対応前。
go get -u golang.org/x/net@v0.22.0
```

### 対応後

サーバー

``` sh
❯ make mem
go run *.go -m
Listening on https://localhost:8080
Alloc = 0 MiB   TotalAlloc = 0 MiB      Sys = 7 MiB     NumGC = 0
2024/04/09 01:05:00 http2: server connection error from [::1]:53348: connection error: PROTOCOL_ERROR
```

クライアント

``` sh
2024/04/09 01:05:00 http2: Framer 0x140000f00e0: wrote CONTINUATION stream=1 len=4352
2024/04/09 01:05:00 http2: Framer 0x140000f00e0: wrote CONTINUATION stream=1 len=4352
2024/04/09 01:05:01 write headers error: write tcp [::1]:53348->[::1]:8080: write: broken pipe
exit status 1
make: *** [run] Error 1
```

### 対応前

サーバー

``` sh
make mem
go run *.go -m
Listening on https://localhost:8080
Alloc = 0 MiB   TotalAlloc = 0 MiB      Sys = 7 MiB     NumGC = 0
Alloc = 0 MiB   TotalAlloc = 0 MiB      Sys = 7 MiB     NumGC = 0
Alloc = 1 MiB   TotalAlloc = 4 MiB      Sys = 12 MiB    NumGC = 1
Alloc = 3 MiB   TotalAlloc = 16 MiB     Sys = 12 MiB    NumGC = 4
Alloc = 2 MiB   TotalAlloc = 35 MiB     Sys = 12 MiB    NumGC = 10
Alloc = 3 MiB   TotalAlloc = 36 MiB     Sys = 12 MiB    NumGC = 10
Alloc = 1 MiB   TotalAlloc = 38 MiB     Sys = 12 MiB    NumGC = 11
Alloc = 0 MiB   TotalAlloc = 63 MiB     Sys = 12 MiB    NumGC = 19
Alloc = 3 MiB   TotalAlloc = 110 MiB    Sys = 12 MiB    NumGC = 33
Alloc = 3 MiB   TotalAlloc = 169 MiB    Sys = 12 MiB    NumGC = 50
Alloc = 1 MiB   TotalAlloc = 229 MiB    Sys = 12 MiB    NumGC = 68
Alloc = 1 MiB   TotalAlloc = 276 MiB    Sys = 12 MiB    NumGC = 83
Alloc = 0 MiB   TotalAlloc = 296 MiB    Sys = 12 MiB    NumGC = 89
Alloc = 2 MiB   TotalAlloc = 324 MiB    Sys = 12 MiB    NumGC = 97
Alloc = 2 MiB   TotalAlloc = 384 MiB    Sys = 12 MiB    NumGC = 115
Alloc = 2 MiB   TotalAlloc = 443 MiB    Sys = 12 MiB    NumGC = 133
Alloc = 2 MiB   TotalAlloc = 476 MiB    Sys = 12 MiB    NumGC = 143
Alloc = 0 MiB   TotalAlloc = 510 MiB    Sys = 12 MiB    NumGC = 155
Alloc = 2 MiB   TotalAlloc = 572 MiB    Sys = 12 MiB    NumGC = 173
Alloc = 1 MiB   TotalAlloc = 611 MiB    Sys = 12 MiB    NumGC = 185
Alloc = 2 MiB   TotalAlloc = 648 MiB    Sys = 12 MiB    NumGC = 196
Alloc = 1 MiB   TotalAlloc = 705 MiB    Sys = 12 MiB    NumGC = 213
Alloc = 2 MiB   TotalAlloc = 756 MiB    Sys = 12 MiB    NumGC = 228
Alloc = 2 MiB   TotalAlloc = 810 MiB    Sys = 12 MiB    NumGC = 244
Alloc = 2 MiB   TotalAlloc = 839 MiB    Sys = 12 MiB    NumGC = 253
Alloc = 3 MiB   TotalAlloc = 900 MiB    Sys = 12 MiB    NumGC = 271
Alloc = 3 MiB   TotalAlloc = 934 MiB    Sys = 12 MiB    NumGC = 282
Alloc = 3 MiB   TotalAlloc = 988 MiB    Sys = 12 MiB    NumGC = 297
Alloc = 1 MiB   TotalAlloc = 1009 MiB   Sys = 12 MiB    NumGC = 304
Alloc = 1 MiB   TotalAlloc = 1067 MiB   Sys = 12 MiB    NumGC = 322
Alloc = 0 MiB   TotalAlloc = 1144 MiB   Sys = 12 MiB    NumGC = 345
Alloc = 0 MiB   TotalAlloc = 1181 MiB   Sys = 12 MiB    NumGC = 356
Alloc = 1 MiB   TotalAlloc = 1233 MiB   Sys = 12 MiB    NumGC = 371
Alloc = 2 MiB   TotalAlloc = 1284 MiB   Sys = 12 MiB    NumGC = 387
Alloc = 2 MiB   TotalAlloc = 1324 MiB   Sys = 12 MiB    NumGC = 399
Alloc = 2 MiB   TotalAlloc = 1375 MiB   Sys = 12 MiB    NumGC = 414
Alloc = 3 MiB   TotalAlloc = 1411 MiB   Sys = 12 MiB    NumGC = 425
Alloc = 2 MiB   TotalAlloc = 1413 MiB   Sys = 12 MiB    NumGC = 426
Alloc = 2 MiB   TotalAlloc = 1425 MiB   Sys = 12 MiB    NumGC = 430
Alloc = 0 MiB   TotalAlloc = 1499 MiB   Sys = 12 MiB    NumGC = 453
Alloc = 3 MiB   TotalAlloc = 1542 MiB   Sys = 12 MiB    NumGC = 465
Alloc = 1 MiB   TotalAlloc = 1577 MiB   Sys = 12 MiB    NumGC = 476
Alloc = 3 MiB   TotalAlloc = 1622 MiB   Sys = 12 MiB    NumGC = 489
Alloc = 3 MiB   TotalAlloc = 1669 MiB   Sys = 12 MiB    NumGC = 503
Alloc = 3 MiB   TotalAlloc = 1716 MiB   Sys = 12 MiB    NumGC = 517
Alloc = 2 MiB   TotalAlloc = 1760 MiB   Sys = 12 MiB    NumGC = 531
Alloc = 1 MiB   TotalAlloc = 1810 MiB   Sys = 12 MiB    NumGC = 546
```

クライアント

``` sh
2024/04/09 01:05:00 http2: Framer 0x140000f00e0: wrote CONTINUATION stream=1 len=2626
2024/04/09 01:05:00 http2: Framer 0x140000f00e0: wrote CONTINUATION stream=1 len=2626
2024/04/09 01:05:00 http2: Framer 0x140000f00e0: wrote CONTINUATION stream=1 len=2626
2024/04/09 01:05:00 http2: Framer 0x140000f00e0: wrote CONTINUATION stream=1 len=2626
```
