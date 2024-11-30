# Music library API for Effective Mobile GO interview

## Описание
Этот репозиторий содержит тестовое задание на собеседование в компанию EM: апи для музыкально библиотеки.
## Установка

1. Клонировать репозиторий:
    ```shell
        git clone https://github.com/karrless/em-interview.git
    ```
2. Перейти в созданный репозиторий:
    ```shell
        cd em-interview
    ```
3. Установить всё зависимости:
    ```shell
        go mod tidy
    ```
4. Создать файл .env в корне проекта на примере файла example.env:
    ```.env
    POSTGRES_HOST=localhost # PostgreSQL host
    POSTGRES_PORT=5432 # PostgreSQL port
    POSTGRES_USER=postgres # PostgreSQL user
    POSTGRES_PASSWORD=postgres # PostgreSQL password
    POSTGRES_DB=postgres # PostgreSQL database
    POSTGRES_SSLMODE=disable # PostgreSQL SSL mode

    EXTERNAL_API_URL=http://localhost:9090 # External API URL for song info
    SERVER_HOST=localhost # Server host
    SERVER_PORT=8080 # Server port
    DEBUG=false # Debug mode
    ```
5. Запустить приложение:
   1. Используя `Makefile`:
        ```shell
            make run
        ```
   2. Используя `go`:
        ```shell
            go run ./cmd/main/main.go
        ```

## Swagger
Документация будет доступна по `/swagger/index.html`


## Дополнение
- Пагинация песней выполнена на основе `limit` и `offset`: где `limit` - максимальное количество, `offset` - смещение от начала
- Пагинация по куплетам выполнена не была, так как не было достаточной ифнормации о том, как именно внешнее API предтавляет текст песни (есть ли некоторые идентификаторы различных куплетов или нет), а именно - просто текстом с переносом строк (судя по swagger документации)
