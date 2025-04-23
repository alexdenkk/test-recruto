# test-recruto
 Тестовое задание

<br>

## Запуск приложения
```bash
git clone https://github.com/alexdenkk/test-recruto.git
cd recruto/deployments/docker-compose
docker-compose up -d
```

<br>

## Конечные точки
### GET `/?name=xxx&message=xxx`
Пример запроса
```bash
chromium http://127.0.0.1:8080/?name=recruto&message=@alexdenkk
```
Пример ответа
```json
200

{
	"message": "Hello recruto! @alexdenkk"
}
```
