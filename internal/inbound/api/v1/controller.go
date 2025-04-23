package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"recruto/internal/domain"
)

// Controller - структура контроллера
type Controller struct {
	Service domain.Service
}

// New - создание экземпляра структуры контроллера
func New(s domain.Service) *Controller {
	return &Controller{
		Service: s,
	}
}

// RegisterEndpoints - регистрация эндпоинтов для контроллера
func (ct *Controller) RegisterEndpoints(engine *gin.Engine) {
	engine.GET("/", ct.GetGreeting)
}

// GetGreeting - Endpoint для фукнции domain.Service.GetGreeting 
func (ct *Controller) GetGreeting(c *gin.Context) {
	// Получение переменных из URL query
	name, name_ok := c.GetQuery("name")
	message, message_ok := c.GetQuery("message")

	if !name_ok || !message_ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "нет необходимых url параметров",
		})
		return
	}

	// Вызов сервиса
	greeting, err := ct.Service.GetGreeting(
		c.Request.Context(),
		name,
		message,
	)

	// В случае ошибки при выполнении запроса, отправка ее клиенту
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Ответ в формате json
	c.IndentedJSON(http.StatusOK, gin.H{
	 	"message": greeting,
	})
}
