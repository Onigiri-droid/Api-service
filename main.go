package main

import (
    "api-service/db"
    "api-service/handler"
    "api-service/router"
    "github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()

    // Подключение к базе данных
    dbConn := db.InitDB()
    defer dbConn.Close()

    // Миграции
    db.RunMigrations(dbConn)

    // Создаём обработчик пользователей
    userHandler := &handler.UserHandler{DB: dbConn}

    // Настройка маршрутов
    router.SetupRoutes(e, userHandler)

    // Запуск сервера
    e.Logger.Fatal(e.Start(":8080"))
}
