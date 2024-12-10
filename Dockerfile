FROM golang:alpine

# Убедимся, что у нас есть необходимые инструменты
RUN apk update && apk add --no-cache \
    bash \
    curl \
    postgresql-client \
    build-base \
    git

# Установить версию Go (дополнительно, если нужно)
RUN go version

# Установить GOPATH (опционально)
ENV GOPATH=/

# Копируем исходный код приложения
COPY ./ ./

# Делаем скрипт для ожидания Postgres исполняемым
RUN chmod +x wait-for-postgres.sh

# Загружаем зависимости Go
RUN go mod download

# Сборка Go-приложения
RUN go build -o todo-app ./cmd/main.go

# Указываем команду запуска приложения
CMD ["./todo-app"]
