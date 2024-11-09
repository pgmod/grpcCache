# Используем официальный образ Go для сборки
FROM golang:alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./

# Устанавливаем зависимости
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем проект
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o server .
# Создаём финальный минималистичный образ
FROM scratch

COPY --from=builder /app/server /server


# Определяем команду по умолчанию
CMD ["./server"]