# Определяем базовый образ для сборки
ARG GO_VERSION=1.21

# Стадия сборки для разработки
FROM golang:${GO_VERSION}-alpine as dev

WORKDIR /app/

RUN apk --no-cache add ca-certificates gcc g++ make git curl nodejs npm

COPY ./ ./

RUN go mod download
RUN CGO_ENABLED=0 go install -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv@latest
RUN CGO_ENABLED=0 go build -gcflags "all=-N -l" -o main .

EXPOSE 9999 4000

# Установка Air для hot reload
RUN CGO_ENABLED=0 go install github.com/cosmtrek/air@latest

# Обновление PATH для включения директории с Air
ENV PATH $PATH:/go/bin

#CMD ["air"]
#CMD ["./main"]

# Стадия сборки для продакшена
FROM golang:${GO_VERSION}-alpine as builder

WORKDIR /app/

RUN apk --no-cache add ca-certificates

COPY ./ ./

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

# Окончательный образ для продакшена
FROM alpine:latest as prod

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/calculator/main .

EXPOSE 9999

CMD ["./main"]