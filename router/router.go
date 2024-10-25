package router

import (
    "api-service/handler"
    "api-service/middleware"
    "github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, userHandler *handler.UserHandler) {
    // Публичные маршруты
    e.POST("/register", userHandler.Register)
    e.POST("/login", userHandler.Login)

    // Защищенные маршруты
    e.GET("/users", userHandler.GetUsers, middleware.JWTMiddleware)// Получение списка всех пользователей (с JWT)
    e.POST("/users", userHandler.CreateUser, middleware.JWTMiddleware)// Создание нового пользователя (без JWT)
    e.PUT("/users/:id", userHandler.UpdateUser, middleware.JWTMiddleware) // Обновление пользователя (с JWT)
    e.DELETE("/users/:id", userHandler.DeleteUser, middleware.JWTMiddleware) // Удаление пользователя (с JWT)
}

    
