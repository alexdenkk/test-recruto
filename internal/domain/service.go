package domain

import (
	"context"
	"fmt"
)

// service - структура сервиса
type service struct {}

// New - создание экземпляра структуры service,
// соответствующей порту Service
func New() Service {
	return &service{}
}

// GetGreeting - Получение приветсвия
// Параметры:
// ctx: Контекст
// name: Имя пользователя
// message: Сообщение
func (s *service) GetGreeting(
	ctx context.Context, name, message string,
) (string, error) {
	return fmt.Sprintf("Hello %s! %s", name, message), nil
}
