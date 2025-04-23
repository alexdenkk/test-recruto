package app

import (
	"recruto/internal/domain"
	"recruto/internal/inbound/api/v1"
	"recruto/config"
	"github.com/gin-gonic/gin"
	"log"
)

// App - структура приложения
type App struct {
	Engine *gin.Engine
	Config *config.Config
}

// New - создание экземпляра структуры App
// Параметры:
// cfg: Конфигурация приложения
func New(cfg *config.Config) *App {
	service := domain.New()
	controller := api.New(service)

	engine := gin.Default()

	controller.RegisterEndpoints(engine)
	
	return &App{
		Engine: engine,
		Config: cfg,
	}
}

// Run - запуск приложения
func (app *App) Run() {
	if err := app.Engine.Run(app.Config.Host); err != nil {
		log.Fatal(err)
	}
}
