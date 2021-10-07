FROM golang:alpine
ENV GO111MODULE=on \
        CGO_ENABLED=0 \
        GOOS=linux \
        GOARCH=amd64 \
        GOPROXY=https://goproxy.cn,direct \
        PATH="/app/go:${PATH}"

RUN mkdir -p /app/go
COPY . /app/go
WORKDIR /app/go

RUN go mod tidy && go build -o app main.go

EXPOSE 8080

ENTRYPOINT ["app","start"]