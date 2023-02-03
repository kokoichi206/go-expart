- go.mod の module 名は教えてあげた方が良さそう
- 0 件のときの考慮ができておらず、json ではなく nil が返ってきてしまった
  - `http://localhost:8080/group` -> nil
  - 初期値としてなんか入れといても良かったかもしれない
    - `INSERT INTO groups VALUES (1, 'group_01');`


Go で API サーバーを作りたいのですが、以下の条件でコードを生成していただけますでしょうか？
・クリーンアーキテクチャを意識し、適切にパッケージ分割を行う
　・DB 層においては interface に依存した作りにしてください
・単体テストの記載もお願いします
・プロジェクトのパッケージ名は "chat-gpt" でお願いします
・適切な godoc コメントの記載もお願いします
・web フレームワークは gin を使う
・ローカルには postgresql が 5432 のポート番号で起動している
　・ユーザー名は root, パスワードは rootpassword, データベース名は postgresql とする
　・groups という初期テーブルが存在しており、id が int で name が VARCHAR(20) で定義されている
・作って欲しいエンドポイントは次の3つです
　・/health にアクセスした時は {"status": "ok"} という json を response body とする
　・/hello にアクセスした時は name というクエリパラメーターを受け取り、{"greeting": "hello <name>"} という json の response body とする。ただし <name> には受け取ったクエリパラメーターを入れてください
　・/group にアクセスした時は、postgresql に定義された groups というテーブルのグループ情報を全権取得し、json として返す。ただし 0 件の時は空配列として返してください。
　　・例: [{"id": 1, "name": "a"}, {"id": 2, "name": "abb"}]
　　・例: []
よろしくお願いします。
