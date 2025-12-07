# Практическая работа #8
## Работа с MongoDB: подключение, создание коллекции, CRUD-операции
## Саттаров Булат Рамилевич ЭФМО-01-25

---
## Требования
- Go **≥ 1.21**
- Docker и Docker Compose
- `curl` **или** любой HTTP-клиент (Postman и т.п.)

## Команды запуска
Запуск контейнера
``` bash

docker compose up -d
```
Копирование файл с примером переменных окружения:

``` bash

cp .env .env
```

Запуск:
``` bash

go run ./cmd/api
```

## Скриншоты

- POST /notes создание заметки

![post.png](docs/screenshots/post.png)

- GET /notes получение списка заметок

![get.png](docs/screenshots/get.png)

- GET /notes/{id} получение по id

![getbyid.png](docs/screenshots/getbyid.png)

- PATCH /notes/{id} обновление заметки

![update.png](docs/screenshots/update.png)

- DELETE /notes/{id} удаление заметки

![delete.png](docs/screenshots/delete.png)

## Доп задания
- Текстовый поиск по search

![img_1.png](docs/screenshots/img_1.png)
![img_2.png](docs/screenshots/img_2.png)

- Поле expiresAt - TTL индекс

![img.png](docs/screenshots/img.png)

## Тестирование

![img_3.png](docs/screenshots/img_3.png)