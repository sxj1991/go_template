FROM golang:1.20-bullseye as builder

RUN mkdir -p /workspace
WORKDIR /workspace/main

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOPROXY=https://goproxy.cn,direct go build -ldflags="-w -s" -o main .


FROM centos:7
WORKDIR /usr/local/demo
COPY --from=builder /workspace/main .
EXPOSE 8088
CMD ["./main"]