package domain

import (
	"context"
)

// Service - порт (интерфейс) сервиса
type Service interface {
	GetGreeting(context.Context, string, string) (string, error)
}
