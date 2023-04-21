# goods-manager-api
JSON-RPC API для управления товарами на складе с использованием PostgreSQL в качестве хранилища

## Инструкция по запуску сервиса
Установите migrate такой командой: `go install github.com/golang-migrate/migrate/v4/cmd/migrate`

Запустите сервис командой `make compose-up`

## Инструкция по запуску тестов
Установите библиотеку gomock `go get github.com/golang/mock/gomock`

Введите команду `make run-test`

## Запрос на отмену резервирования
На вход: 
1. Версия jsonrpc 2.0
2. ID
3. Метод: Router.CancelReservation
4. Параметры: массив идентификаторов товаров и идентификатор склада (на случай, если у нас несколько складов)

curl -X POST
     -H 'Content-Type: application/json'
     -d '{"jsonrpc":"2.0","id":"1","method":"Router.CancelReservation","params":[{"item_ids": [1, 5], "storage_id": 2}]}'
     http://localhost:8080

На выходе:
    HTTP/1.1 200 OK
    Content-Type: application/json; charset=utf-8
    X-Content-Type-Options: nosniff
    Date: Fri, 21 Apr 2023 12:00:38 GMT
    Content-Length: 34
    
    {
        "result": {},
        "error": null,
        "id": 1
    }

## Запрос на резервирование товара (мой случай предусматривает как один, так и несколько товаров)

На вход:
1. Версия jsonrpc 2.0
2. ID
3. Метод: Router.ReserveItems 
4. Параметры: массив идентификаторов товаров и идентификатор склада (на случай, если у нас несколько складов)

curl -X POST
    -H 'Content-Type: application/json'
    -d '{"jsonrpc":"2.0","id":"1","method":"Router.ReserveItems","params":[{"item_ids": [1, 5], "storage_id": 2}]}'
    http://localhost:8080

На выходе:
    HTTP/1.1 200 OK
    Content-Type: application/json; charset=utf-8
    X-Content-Type-Options: nosniff
    Date: Fri, 21 Apr 2023 12:00:38 GMT
    Content-Length: 34

    {
        "result": {},
        "error": null,
        "id": 1
    }

## Запрос на список оставшихся товаров (незарезервированных)

На вход:
1. Версия jsonrpc 2.0
2. ID
3. Метод: Router.ItemsAmount 
4. Параметры: идентификатор склада

Запрос:
    curl -X POST
        -H 'Content-Type: application/json'
        -d '{"jsonrpc":"2.0","id":"1","method":"Router.ItemsAmount","params":[{"storage_id": 2}]}'
        http://localhost:8080

На выходе:
    HTTP/1.1 200 OK
    Content-Type: application/json; charset=utf-8
    X-Content-Type-Options: nosniff
    Date: Fri, 21 Apr 2023 12:42:48 GMT
    Content-Length: 44
    
    {
        "result": {
            "amount": 4
        },
        "error": null,
        "id": 1
    }


### Необходимые улучшения: защитить базу данных от конкурентного чтения/записи


# Комментарии по заданию
Я слегка изменила структуру таблицы Товар, убрав из неё столбец Количество и добавив его в таблицу items_storage.
Это нужно было для реализации логики работы с товарами, которые одновременно могут находиться на нескольких складах,
то есть, в таблице items_storage отображается товар, склад, на котором он находится, его количество на складе 
и количество зарезервированных единиц товара на конкретном складе
