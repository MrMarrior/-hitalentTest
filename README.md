# Chat API

Тестовое задание для Hitalent

---


## Запуск проекта

1. Зайти в корневую папку проекта
2. Запустите команду: docker-compose up --build -d
3. Запустите миграции: docker-compose exec app goose -dir migrations postgres "host=postgres user=chatuser password=chatpass dbname=chatdb sslmode=disable" up
4. Сервер доступен на http://localhost:8080
