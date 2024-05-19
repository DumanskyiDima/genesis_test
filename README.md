Цей проект - виконане тестове завдання для відбору на курс Software Engineering School 4.0 від Genesis Academy
Автор: Думанський Дмитро

Основні використані технології: GoLang ([Gin](https://github.com/gin-gonic/gin) HTTP web framework), MongoDB та Gmail smtp.

Використано зовнішній API від [НБУ](https://bank.gov.ua/ua/open-data/api-dev)


#### Запуск за допомогою Docker:

```bash
# Build & Create Docker Containers
docker-compose up -d
```

API працює на порті 8080:
#### ENDPOINTS:

- `GET /rate` Повкртає курс валюти USD/UAH

---

- `POST /subscribe` Створює підписку на щоденну розсилку для обраного email

---


Структура Проекту
-----------------

```
├── database            # Логіка взаємодії з базою даних
├── logs
├── models              # моделі даних
├── services            # логіка взаємодії з зовнішніми сервісами (API, SMTP)
└── views               # основна логіка обробки запитів на 
```

