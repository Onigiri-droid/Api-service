package model

// Структура пользователя
type User struct {
    ID       int    `json:"id" bun:",pk,autoincrement"`
    Name     string `json:"name,omitempty"`
    Email    string `json:"email,omitempty"`
    Password string `json:"password,omitempty"`
}

// Структура для запроса на логин
type LoginRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
}
