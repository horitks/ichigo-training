#!/bin/sh

set -euo pipefail

trap catch ERR

function catch {
    printf '\033[31m%s\033[m\n' 'catch: error! Have you started up a docker?'
}

echo 'lunux build ...'
docker-compose exec golang env GOOS=linux GOARCH=amd64 go build -o bin/linux/ichigotraning main.go 

echo 'darwin build ...'
docker-compose exec golang env GOOS=darwin GOARCH=amd64 go build -o bin/mac/ichigotraning main.go

echo 'windows build ...'
docker-compose exec golang env GOOS=windows GOARCH=amd64 go build -o bin/windows/ichigotraning main.go

printf '\033[32m%s\033[m\n' 'build succeeded'
