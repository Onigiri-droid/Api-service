package db

import (
	"api-service/model"
	"context"
	"log"

	"github.com/uptrace/bun"
)

func RunMigrations(db *bun.DB) {
    ctx := context.Background()

    // Создание таблицы только если она еще не существует
    if _, err := db.NewCreateTable().Model((*model.User)(nil)).IfNotExists().Exec(ctx); err != nil {
        log.Fatalf("Ошибка создания таблицы пользователей: %v", err)
    }

    log.Println("Миграция завершена: Таблица пользователей создана (или уже существует)")
}
