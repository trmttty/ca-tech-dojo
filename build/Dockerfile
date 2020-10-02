# # Start from golang base image
# FROM golang:1.15.2-alpine3.12 as builder

# WORKDIR /app

# COPY .. .
# RUN go mod download 
# RUN go build -o start_app /app/cmd/game-api/server.go

# # Start a new stage from scratch
# FROM alpine:3.12
# COPY --from=builder /app/start_app /bin/start_app
# EXPOSE 800
# CMD [ "/bin/start_app" ]



FROM golang:1.15-alpine

WORKDIR /ca-tech-dojo

ENV GO111MODULE=on

EXPOSE 8080
