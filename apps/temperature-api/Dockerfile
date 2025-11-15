# Этап сборки
FROM golang:1.23-alpine as builder

WORKDIR /app

# Копируем файлы зависимостей для кэширования
COPY go.mod go.sum ./ 

# Загружаем зависимости
RUN go mod download

# Копируем исходный код
COPY . /app

# Собираем приложение
RUN go build -o temperature-api main.go

# Финальный этап
FROM alpine:latest

WORKDIR /app

# Копируем собранное приложение из этапа сборки
COPY --from=builder /app/temperature-api .
EXPOSE 8081
ENTRYPOINT ["./temperature-api"]
# Устанавливаем необходимые пакеты для runtime
RUN apk --no-cache add ca-certificates

# Определяем команду для запуска приложения
CMD ["./temperature-api"]