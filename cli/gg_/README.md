## やりたいこと

名前何にしよかな

- Git でその人の更新行を全取得する
  - ユーザー名を指定
  - Token ありかなしかで private or public を分けられそう？
  - output をありなし
  - 最終的にやりたいのは 1 日分の更新量

## 全リポジトリ取得

### [List public repositories](https://docs.github.com/ja/rest/repos/repos#list-public-repositories)

まずこっちの対応

```sh
# これトークンいるんかいな
curl \
  -H "Accept: application/vnd.github+json" \
  -H "Authorization: token <TOKEN>" \
  https://api.github.com/repositories
```

### [List repositories for the authenticated user](https://docs.github.com/ja/rest/repos/repos#list-repositories-for-the-authenticated-user)

どうやって Token を読み込ませるか考慮（他サービス検索）

```sh
curl \
  -H "Accept: application/vnd.github+json" \
  -H "Authorization: token <TOKEN>" \
  https://api.github.com/user/repos
```
curl  https://api.github.com/repos/kokoichi206/til/readme

curl \
  -H "Accept: application/vnd.github+json" \ 
  -H "Authorization: token <TOKEN>" \
  https://api.github.com/repos/OWNER/REPO/stats/punch_card

curl -H "Accept: application/vnd.github+json" https://api.github.com/repos/kokoichi206/til/stats/punch_card
