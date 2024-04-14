## memo

- 対データ競合
  - チャネル
  - sync
    - こっちを使えるなら使う
    - チャネルは人類には早い
- sync
  - Mutex
    - そのパスに入っていい goroutine を1つに絞る
- sync.Once は1回読んでみると良い
  - 1.12 -> 1.13 で大きな変更が入った
- close を2回呼ぶと panic
  - sync.Once とか使う
- 何かを疑う
  - メソッドが変数に代入
  - if 文じゃなくて switch 文を使っていたら
- 静的解析
  - https://github.com/skeletonlabs/skeleton
  - inspect.Preorder

``` sh
go doc ast.CallExpr

$ go doc ast.SelectorExpr
package ast // import "go/ast"

type SelectorExpr struct {
        X   Expr   // expression
        Sel *Ident // field selector
}
    A SelectorExpr node represents an expression followed by a selector.

func (x *SelectorExpr) End() token.Pos
func (x *SelectorExpr) Pos() token.Pos
```

``` go
fmt.Printf("%[1]T %[1]v\n", n.Fun)
```

- https://cs.opensource.google/go/go/+/master:src/cmd/compile/internal/loopvar/loopvar.go;l=39

## Links

- [Looking inside a Race Detector](https://www.infoq.com/presentations/go-race-detector/)
- https://github.com/golang/tools
