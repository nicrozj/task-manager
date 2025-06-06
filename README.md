# [Task Manager App](https://task.zhorkinigor.ru)

> Приложение для управления задачами с авторизацией, фильтрацией и модальными окнами



![Task Manager](./screenshots/screenshot_1.png)
![Task Manager](./screenshots/screenshot_2.png)

Это полноценное приложение с фронтендом и бэкендом, где пользователь может:

- Создавать, редактировать и удалять задачи
- Фильтровать задачи по статусу и дате
- Пользоваться JWT-аутентификацией

---

## 🧰 Технологии

- **Frontend**: Vue 3 + TypeScript + TailwindCSS
- **Backend**: Go + SQLX + PostgreSQL
- **Аутентификация**: JWT + refresh token (хранится в cookie)
- **Запуск**: Docker Compose

---

## 🚀 Установка и запуск

### Предварительно:

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Запуск:

```bash
docker-compose -f docker-compose.dev.yml up
```

## 🛠 Переменные окружения:

/backend/.env

```env
DB_URL=postgresql://admin:password@postgres:5432/db?sslmode=disable

JWT_SECRET_KEY=secret_key

HTTP_ACCESS_TOKEN_EXPIRE=15
HTTP_REFRESH_TOKEN_EXPIRE=10080

WEB_URL=http://localhost
```

/frontend/.env

```env
VITE_API_BASE_URL=http://localhost/api
```
