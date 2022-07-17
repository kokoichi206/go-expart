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



