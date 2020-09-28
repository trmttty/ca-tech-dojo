# Start from golang base image
# FROM golang:1.15.2-alpine3.12 as builder

# WORKDIR /app

# COPY . .
# RUN go mod download 
# RUN go build -o start_app

# # Start a new stage from scratch
# FROM alpine:3.12
# COPY --from=builder /app/start_app /bin/start_app
# EXPOSE 80
# CMD [ "/bin/start_app" ]



FROM golang:1.14-alpine

# 作業ディレクトリの作成・設定
WORKDIR /ca-tech-dojo

# Go Modules を有効化
ENV GO111MODULE=on

# COPY go.mod .
# COPY go.sum .

# go.mod 内のパッケージをダウンロード
# RUN go mod download

EXPOSE 8080
