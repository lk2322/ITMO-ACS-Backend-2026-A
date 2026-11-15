# Справочник API

Полный контракт описан в формате OpenAPI 3.0: файл
[openapi.yaml](openapi.yaml) — его можно импортировать в Postman,
Insomnia или сгенерировать по нему клиент.

Ниже — интерактивный просмотр спецификации.

<link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css">
<div id="swagger-ui"></div>
<script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js" crossorigin></script>
<script>
  window.addEventListener('load', function () {
    window.ui = SwaggerUIBundle({
      url: 'openapi.yaml',
      dom_id: '#swagger-ui',
      deepLinking: true,
      docExpansion: 'none',
      defaultModelsExpandDepth: 0,
      presets: [SwaggerUIBundle.presets.apis]
    });
  });
</script>

!!! note "Если блок выше пуст"
    Swagger UI подгружается с CDN — нужен доступ в интернет у браузера.
    Без него спецификацию можно открыть напрямую:
    [openapi.yaml](openapi.yaml).

## Общие правила

**Базовый адрес.** Все методы, кроме `/health`, живут под префиксом
`/api/v1`. Для кластера это `http://192.168.1.245/api/v1`.

**Авторизация.** Защищённые методы требуют заголовок
`Authorization: Bearer <access_token>`. Токен выдают `/auth/register` и
`/auth/login`, срок жизни — 24 часа.

**Формат ошибок** одинаков во всех сервисах:

```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "описание ошибки"
  }
}
```

| Код ответа | Когда возвращается |
|---|---|
| `400` | тело запроса не прошло валидацию |
| `401` | токен отсутствует, истёк или подписан другим ключом |
| `403` | роль пользователя не допускает операцию |
| `404` | объект не найден |
| `409` | конфликт: занятый email или пересечение дат бронирования |

!!! warning "Завершающий слэш у POST /properties"
    Из-за вложенной группы маршрутов в API Gateway создание объекта
    отвечает по адресу `/api/v1/properties/` — запрос на
    `/api/v1/properties` получит `307 Temporary Redirect` на этот адрес.
    Большинство клиентов (Postman, браузер, `curl -L`) следуют редиректу
    автоматически и сохраняют метод и тело, но при написании своего
    клиента это стоит учитывать.
