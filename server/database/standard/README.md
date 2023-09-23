## コネクション数確認

``` sql
SELECT pid, usename, application_name, state, query_start FROM pg_stat_activity;
```

## 登場人物

- [sql.DB](https://github.com/golang/go/blob/152ffca82fa53008bd2872f7163c7a1885da880e/src/database/sql/sql.go#L402-L438)
- [sql.driverConn](https://github.com/golang/go/blob/152ffca82fa53008bd2872f7163c7a1885da880e/src/database/sql/sql.go#L456-L472)
  - [ci](https://github.com/golang/go/blob/152ffca82fa53008bd2872f7163c7a1885da880e/src/database/sql/sql.go#L461)
- [sql/driver.Conn](https://github.com/golang/go/blob/152ffca82fa53008bd2872f7163c7a1885da880e/src/database/sql/driver/driver.go#L223-L248)
  - 各ドライバーが実装
  - sql.DB が、コネクションプールの管理をスレッドセーフにしてくれているため、このコネクションはその辺の考慮不要

pq で実装 driver.Conn を実装している構造体。

``` go
type conn struct {
	c         net.Conn
	...
}
```

`net.Conn` の抽象化もエグい。。。
