## Go

Go 公式が fomat を用意していることで「自転車置き場の理論」を避けることができている。

標準 lint ツール。
lint ツールはコードの静的解析を行い、次のような点を指摘してくれる。

- コード上でバグが発生しそうな部分
- スタイルがふぞろいな点
- Go らしくない書き方

``` sh
go vet main.go
go vet .
go vet ./...

# Go らしくないコーディングスタイルを警告してくれる
golint

go vet ./...; golint ./...

golint -set_exit_status ./...
```

go get golang.org/x/lint/golint

go vet, golint は CI に組み込むのが良い！

golint には検査の厳しさを調整するオプションがある。
数値が小さいほど厳しくなり、デフォルトでは 0.8。

``` sh
golint -min_confidence=0.1 ./...
```

ドキュメント閲覧ツール

go doc と godoc。

``` sh
go get golang.org/x/tools/cmd/godoc

go doc -all fmt
godoc
# http://localhost:6060
```

gopls（Go Please）は Go が公式サポートしている Language Server 

godef は定義ジャンプのためのツール

### name

リポジトリ名に go- という prefix をつける。go-myproj がプロジェクト名で、パッケージは myproj とする。

実行可能なドキュメントのようなかたちでタスクランナーを充実させる。


## Best practice

- panic よりエラーハンドリング
- map を避ける
    - 可能な限り struct できちんと type 定義
    - map はスレッドセーフではない
- reflect を避ける
- 継承より委譲
- go のコードを読もう
    - ghq get golang/go
- build
    - ビルド時にバージョン情報や git のリビジョン情報をバイナリに埋め込むのも良いテクニック



## 社内ツール

Ruby, Perl, Python といったスクリプト言語では、ユーザに各プログラミング言語のランタイムインストールを要求しないといけない。

## マルチプラットフォームを意識する上で守るべきルール
- 積極的に path/filepath を使う
    - パスのセパレータ / やリストセパレータ : など
- defer を使いまくる
    - defer は呼び出された順番とは逆順に実行される！
    - エラーチェックの直後に defer
- UTF-8 を使う

## TUI: Text-based User Interface
termbox とか使ってみ

## OS 固有の処理への対応

runtime.GOOS or Build Constraints を用いたビルド時の OS 振り分け。

file_${GOOS}.go
file_${GOARCH}.go
file_${GOOS}_${GOARCH}.go

プロセスのデーモン化を提供するツールまたはシステム。

## シングルバイナリにこだわる

go-bindata, go-assets などのアセットツールを使って、シングルバイナリにこだわりたい。

statik で静的なファイルをアセットとしてバイナリに埋め込める。

``` go
// マジックコマンドであり、一般的には go build の際に必要となる Go の
// ソースファイルを自動的に生成する目的で使用される。
// go generate && go build
// go: generate statik
```

## Windows 用

Windows アプリケーションの起動時にコマンドプロンプトが表示されるコンソールアプリケーションと、表示されない windows アプリケーションがある。
go コマンドでは -H オプションでこのモードを変換できる。

``` sh
go build -ldflags="-H windowsgui"
```

マルチプラットフォームな GUI を作るには

- GoQt
- ui
- go-qml
- go-gtk
- walk
- shiny


## 実用的なアプリケーションの条件とは
- どのような機能を持っているかが容易に調べられる
- パフォーマンスが良い
- 多様な入出力を扱える
- 人間にとって扱いやすい形式で入出力できる
- メンテナンス性が高い
- 想定外の場合に安全に処理を停止できる


## バージョン管理

``` sh
#!/bin/sh

GIT_VER=$(git describe --tags)
go build -ldflags "-X main.version=${GIT_VER}"
```

go-latest の比較

- GitHub のタグ
- HTML のメタタグ
    - Better
- JSON API
    - Better


## I/O

Go ではランタイムでの自動的なバッファリングは行われない。（LL 言語では、出力先によっては行われるものも結構ある！）

go-isatty で、出力先が端末かどうかを判断している。端末であればバッファリングを行わず、端末でなければ bufio で出力を行う、のがよさそう。




