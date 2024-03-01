``` sh
go install github.com/99designs/gqlgen@latest

gqlgen version
v0.17.44
```

``` sh
go get -u github.com/99designs/gqlgen
gqlgen init
```

``` graphql
query {
  todos {
    id
    text
    done
  }
}
```

``` graphql
mutation {
  createTodo(input:{
    text:"sample-todo"
    userId:"user-id-test"
  }){
    id
    text
    done
    user {
      id
      name
    }
  }
}
```

`gqlgen generate` では `gqlgen init` とは違い server.go は生成されない。

``` sh
gqlgen generate

sqlboiler sqlite3
```

## query

- GraphQL で取得するフィールドはすべてスカラ型になってる必要がある
- 独自型
  - DateTime
- リゾルバの分割
  - **分割されたリゾルバの実行順**が重要になてくる
  - ネストが浅い順に呼ばれ、それらが実行時引数に渡されてそう
    - その情報はクエリ時に使える！
  - オーバーフェッチを防ぐ
  - 発行されるSQLクエリを簡潔に保つ
- N + 1 問題
  - Dataloader
    - N 個のクエリを IN 句で1個にまとめる！！
    - https://github.com/graph-gophers/dataloader
  - 仕組み
    - 検索条件の後すぐにクエリを投げるのではなく、一旦待機する
    - 複数個の条件がたまってから, IN 句で投げる！

```
  ): User @isAuthenticated
```

``` sh
go get -u github.com/graph-gophers/dataloader/v7
```
