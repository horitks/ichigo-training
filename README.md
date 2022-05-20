# ichigo-training

Sorry, this repository(document) is in Japanese.

## Description
いちご王国カレンダー訓練用。

5つの数字から四則演算を使って15を算出するゲームの練習アプリです。

ただ出てきた数字を使って15になるように暗算するだけ。

## Usage
ただ遊びたい場合は環境に合わせたバイナリをダウンロードしてconsoleで実行してください

- [Linux](https://github.com/horitks/ichigo-training/blob/main/src/bin/linux/ichigotraning)
- [Mac](https://github.com/horitks/ichigo-training/blob/main/src/bin/mac/ichigotraning)
- [Windows](https://github.com/horitks/ichigo-training/blob/main/src/bin/windows/ichigotraning)

入力画面が出るので四則演算で入力してください

- 足す: +
- 引く: -
- 掛ける: *
- 割る: /

問題例
```
[2][9][12][1][5]
```

入力例
```
12/3+5*2+1
```


## For Developers
### 構築(build)

```
docker-compose up --build -d
```

### 起動(start up)

```
docker-compose up -d
```

### 確認(ps)

```
docker-compose ps
```

or

```
docker ps
```

### 実行(exec)

```
docker-compose exec golang go run main.go
```

### バイナリの生成(Build)

linux
```
docker-compose exec golang env GOOS=linux GOARCH=amd64 go build -o bin/linux/ichigotraning main.go 
```

mac
```
docker-compose exec golang env GOOS=darwin GOARCH=amd64 go build -o bin/mac/ichigotraning main.go
```

windows
```
docker-compose exec golang env GOOS=windows GOARCH=amd64 go build -o bin/windows/ichigotraning main.go
```

まとめて生成するshell
```
sh build_go_with_docker.sh
```

### バイナリでの実行確認
```
cd src/bin/<仕様OSのディレクトリ>
./ichigotraning
```
