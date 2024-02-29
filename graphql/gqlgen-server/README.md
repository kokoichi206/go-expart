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

## Links

- こちらの本で勉強させてもらってます、ありがとうございます
  - https://zenn.dev/hsaki/books/golang-graphql/viewer/resolveradvanced
