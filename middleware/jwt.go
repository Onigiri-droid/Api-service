package middleware

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/golang-jwt/jwt/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        authHeader := c.Request().Header.Get("Authorization")
        if authHeader == "" {
            return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Отсутствует токен аутентификации"})
        }

        tokenString := authHeader[len("Bearer "):]
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte("supersecretkey"), nil
        })

        if err != nil || !token.Valid {
            return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Невалидный токен"})
        }

        return next(c)
    }
}
