# ichigo-training

## Description
いちご王国カレンダー訓練

## Usage
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
