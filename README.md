# GO-песочница

Первые эксперименты с Go после нескольких лет разработки на PHP.

В качестве тестового проекта выбран таск-трекер.

## Что по планам?
- [x] Скелет API - роуты, контроллеры
- [x] Верхнеуровневое описание сущностей (юзер, таска, статус)
- [x] Подключение к БД PostgreSQL ([gorm](https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL))
- [x] Middlewares и обработка ошибок
- [ ] JWT-авторизация
- [ ] Тестирование
- [ ] Swagger
- [ ] Интеграция со сторонним сервисом, например, тг-бот

## Из чего состоит проект?
Структура может меняться по ходу работы над проектом, но пока так:

```
├── config                    # Глобальные конфиги
│   ├── auth.go               # Хранит текущего юзера и логику аутентификации
│   └── db.go                 # Хранит и управляет глобальным подключением *gorm.DB
├── controllers               # Возможно, логика частично будет вынесена в репозитории
│   ├── authController.go
│   └── tasksController.go
├── models
│   ├── task.go
│   └── user.go
├── services
│   └── jwtService.go
├── util
│   └── util.go               # Всякие хелперы пока тут
└── main.go                   # Роутинг пока тут, но возможно будет вынесет в конфиг
```