## 構文解析器

- 扱う文は2つ
  - let, return
  - 左から右に解釈され、処理される
- 残りは全部、式で構成される

## Pratt 構文解析

- 単一のトークンタイプに関連付ける
- 前置演算子（prefix operator）
  - `--5`
- 後置演算子（postfix operator）
  - `foobar++`
- 中間演算子（infix operator）
  - `5 * 46`
- 優先順位（operator precedence, order of operations）
  - 演算子のくっつきやすさ

### Links

- https://journal.stuffwithstuff.com/2011/03/19/pratt-parsers-expression-parsing-made-easy/
