# Go Web Project

Пример веб-приложения на Go с правильно организованной структурой проекта.

## Структура проекта

```
project/
├── cmd/
│   └── app/
│       └── main.go          # Точка входа в приложение
├── internal/
│   ├── handlers/            # HTTP обработчики
│   ├── services/            # Бизнес-логика
│   ├── models/              # Структуры данных
│   ├── repositories/        # Работа с базой данных
│   ├── middleware/          # HTTP middleware
│   └── utils/               # Вспомогательные функции
├── pkg/                     # Публичные библиотеки
├── configs/                 # Конфигурационные файлы
├── migrations/              # Миграции базы данных
├── docs/                    # Документация
├── tests/                   # Тесты
├── scripts/                 # Скрипты для сборки и деплоя
├── go.mod
├── go.sum
└── README.md
```

## Запуск приложения

```bash
go run cmd/app/main.go
```

Приложение будет доступно по адресу http://localhost:8080