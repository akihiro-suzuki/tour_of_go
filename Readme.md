### Goのお勉強プロジェクト
https://go.dev/doc/tutorial/getting-started で勉強します。

### 前提
- GOPATH
  - goがinstallされているパス
- gopls
  -  Go が公式サポートしている Language Server

### learn1
- 64bit cpuではintもint64も同じ
- 定数の場合は、const宣言が必要。:= で書くことはできない
- forでwhileも表現できる
- go のswitchはcaseごとになる(各caseにbreakが入ってるみたいになる)
- if文代わりに switch {} と書くこともできる
### 備忘
-  git config --global --add safe.directory /go/src/app