package main

import (
	"recruto/app"
	"recruto/config"
)

// main - точка входа
func main() {
	// Подгрузка конфигурации
	cfg := config.Load()
	// Создание приложения
	recrutoApp := app.New(cfg)
	// Запуск приложения
	recrutoApp.Run()
}
