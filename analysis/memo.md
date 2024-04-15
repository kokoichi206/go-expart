## [golangtokyo: codelab](https://golangtokyo.github.io/codelab/find-gophers/?index=codelab#0)

- 解析方法
  - 静的解析
    - gofmt/goimports
    - go ver/golint
    - guru
    - gocode
    - errcheck
  - 動的解析
    - `go test -race`
- go と静的解析
  - 文法がシンプル
  - 静的型付け
  - 暗黙の型変換をしない
- go packages
  - go/ast
  - go/build
  - go/constant
  - ...
- Flows
  - ソースコード → トークン
    - 字句解析
    - go/scanner, go/token
  - トークン → 抽象構文木（AST）
    - 構文解析
    - go/parser, go/ast
  - 抽象構文木 → 型情報
    - 型チェック
    - go/types, go/constant
- 字句解析
  - func なのか識別子なのかとか
  - go/parser パッケージが内部で字句解析する
    - 直接字句解析することはあんまないかも
    - 内部で go/scanner を用いている
- **true, false は識別子**
  - リテラルではない！

``` sh
# exp, err := parser.ParseExpr("true") の出力。
     0  *ast.Ident {
     1  .  NamePos: 1
     2  .  Name: "true"
     3  }

$ go doc ast.Ident
package ast // import "go/ast"

type Ident struct {
        NamePos token.Pos // identifier position
        Name    string    // identifier name
        Obj     *Object   // denoted object, or nil. Deprecated: see Object.
}
    An Ident node represents an identifier.

func NewIdent(name string) *Ident
func (x *Ident) End() token.Pos
func (id *Ident) IsExported() bool
func (x *Ident) Pos() token.Pos
func (id *Ident) String() string
```

``` sh
# parser.ParseExpr("v + 1") の出力。
     0  *ast.BinaryExpr {
     1  .  X: *ast.Ident {
     2  .  .  NamePos: 1
     3  .  .  Name: "v"
     4  .  }
     5  .  OpPos: 3
     6  .  Op: +
     7  .  Y: *ast.BasicLit {
     8  .  .  ValuePos: 5
     9  .  .  Kind: INT
    10  .  .  Value: "1"
    11  .  }
    12  }

# 1 などは BasicLit (リテラル)。
$ go doc ast.BasicLit
package ast // import "go/ast"

type BasicLit struct {
        ValuePos token.Pos   // literal position
        Kind     token.Token // token.INT, token.FLOAT, token.IMAG, token.CHAR, or token.STRING
        Value    string      // literal string; e.g. 42, 0x7f, 3.14, 1e-9, 2.4i, 'a', '\x7f', "foo" or `\m\n\o`
}
    A BasicLit node represents a literal of basic type.

func (x *BasicLit) End() token.Pos
func (x *BasicLit) Pos() token.Pos
```

## Links

- [go/ast](https://pkg.go.dev/go/ast)
