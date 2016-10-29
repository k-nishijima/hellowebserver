# 第4回Golang勉強会 in Okinawa （2016年10月29日）

お題：コマンドラインから起動する自作Webサーバを作ってみよう

# CLIの起動方法

これでいいの？

```
$ go run cmd/hellowebserver/main.go
```

# flagを使ってコマンドライン・オプション追加

```
$ go run cmd/hellowebserver/main.go -help
Usage of /var/folders/58/vsrs9b517qv9gm_vr8nchz380000gn/T/go-build355279879/command-line-arguments/_obj/exe/main:
  -port int
        use (default 3000)
exit status 2
```

# cmdをバイナリとしてインストール

```
$ cd cmd/hellowebserver/
$ go install
$ hellowebserver -help
Usage of hellowebserver:
  -port int
    	use (default 3000)
```
