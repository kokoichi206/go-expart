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


```
  ): User @isAuthenticated
```
