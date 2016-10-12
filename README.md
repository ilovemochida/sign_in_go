
# README

## 環境構築

```
$ git clone https://github.com/ilovemochida/sign_up_go.git
$ go get github.com/mattn/goemon/cmd/goemon
$ go get github.com/gorilla/mux
$ goemon go run main.go
```

### goemon

goemonはファイルを変更した際に自動でリロードしてくれるやつ

どのようなときにどのファイルをリロードをするかの設定ファイルがgoemon.ymlファイルに記述されているので
新たなフォルダでも作った際はこのymlファイルを書き換えればおk

### gorilla mux

このライブラリはルーティングを楽にしてくれるやつですね