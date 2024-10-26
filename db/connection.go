package db

import (
    "database/sql"
    "log"

    "github.com/uptrace/bun"
    "github.com/uptrace/bun/dialect/pgdialect"        // Поддержка PostgreSQL
    "github.com/uptrace/bun/driver/pgdriver"          // Драйвер PostgreSQL
)

func InitDB() *bun.DB {
    // Строка подключения к базе данных
    dsn := "postgres://api_user:api_password@postgres-db:5432/api_database?sslmode=disable"
    
    // Инициализация драйвера для PostgreSQL
    sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
    
    // Создание объекта базы данных через bun
    db := bun.NewDB(sqldb, pgdialect.New())

    // Проверяем подключение к базе
    if err := db.Ping(); err != nil {
        log.Fatalf("Не удалось подключиться к базе данных: %v", err)
    }

    log.Println("Подключение к базе данных успешно установлено")
    
    return db
}
