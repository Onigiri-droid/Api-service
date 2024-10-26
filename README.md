# Api-service_1.0

Стек:
Go, PostgreSQL, Echo, bun, golang-jwt

Описание проекта:
Api-сервис предоставляет возможность зарегистрировать и авторизовать пользователя с использованием jwt-токена, токен дает возможность запрашивать список пользователей, а так же создавать новых, обновлять и удалять. Данные хранятся в PostgreSQL.

Для подключения к бд перейдите по пути db\connection.go и настройте переменную dsn := "postgres://username:password@postgres-db:5432/dbname?sslmode=disable"

Значения:
username: имя пользователя базы данных,
password: пароль пользователя базы данных,
postgres-db: имя контейнера с базой данных (указано в docker-compose.yml),
dbname: имя базы данных, к которой вы хотите подключиться.

Руководство к API с использованием Postman:

Регистрация нового пользователя.
http://localhost:8080/register
{
    "name": "user",
    "email": "user@gmail.com",
    "password": "password"
}
Ответ:
{
    "id": 1,
    "name": "user",
    "email": "user@gmail.com",
    "password": "Хэш-код"
}

Авторизация пользователя.
http://localhost:8080/login
{
    "email": "user@gmail.com",
    "password": "password"
}
Ответ:
{
    "token": "Хэш-код для закрытых заросов"
}

Для следующих 4-х запросов понадобится токен (/login).

Получение списка пользователей.
http://localhost:8080/users
Ответ:
{
        "id": 7,
        "name": "User",
        "email": "user@gmail.com",
        "password": "Хэш-код"
    },
    {
        "id": 17,
        "name": "Test",
        "email": "user123@gmail.com",
        "password": "Хэш-код"
    },

Создание нового пользователя.
http://localhost:8080/users
{
    "name": "user",
    "email": "user@yandex.ru",
    "password": "password"
}
Ответ:
{
    "id": 4,
    "name": "user",
    "email": "user@yandex.ru",
    "password": "password"
}

Обновление пользователя.
http://localhost:8080/users/7
{
    "name": "Иван"
}
Ответ:
{
    "id": 7,
    "name": "Иван",
    "email": "user@gmail.com",
    "password": "Хэш-код"
}

Удаление пользователя с конретным id.
http://localhost:8080/users/7
