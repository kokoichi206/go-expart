## つらい

- 配列操作
  - js とか
    - 宣言的な map, filter とかの操作
  - Go
    - 手続き型の操作
    - 記述量が多いだけで、コードの糸も分かりづらい
      - HOW を重ねただけ。。
- nil の扱い
  - kotlin とか null が入る場合
    - コンパイラがチェックできる
    - 実行時に適切な処理がされる
- エラー
  - error 型が文字列型と大差がない
    - エラーを扱うのに十分な情報を持たない
  - err != nil はコンパイルエラーにならない。。。
  - e.g.
    - try catch は実質 go to と同じ

## 中身

- Go は手続き型の言語と捉えると良さそう
  - struct を中途半端なオブジェクト指向と捉えると不便
  - 少ないデータ構造に対して、たくさんの関数を生やす方がうまくいきそう？
    - c.f. それぞれのデータ型を作るオブジェクト指向
- ジェネレータ？
  - 再利用性の高いコードを手軽に作るのが難しい性質のせい？

## 改善

- org
  - https://github.com/uber
  - https://github.com/samber
- ROP: Railway Oriented Programming
  - https://fsharpforfunandprofit.com/rop/
