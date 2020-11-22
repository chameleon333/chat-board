# ベースとなるDockerイメージ指定
FROM golang:latest
RUN go get github.com/derekparker/delve/cmd/dlv \
&& go get github.com/go-sql-driver/mysql

# コンテナログイン時のディレクトリ指定
WORKDIR /var/www/chat-board/backend
