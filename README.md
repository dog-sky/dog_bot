# dog_bot
---
[![codecov](https://codecov.io/gh/dog-sky/dog_bot/branch/main/graph/badge.svg?token=SADKGY8ORK)](https://codecov.io/gh/dog-sky/dog_bot)
---

Dog tracking bot

генерация моков
```
make generate-mocks
```

создание миграций
```
goose -dir migrations create %migrations-name% sql
```

# ToDo
- логированием обмазать
- Составить набор методов (пописы? прогулки? сколько и какой еды было куплено?)
- CI/CD на свой сервак

???
регистрация времени прогулок, чтобы установить время пописа.
Потом получение времен всех последний прогулок.
Получение самого долго промежутка между пописами.
Установление даты покупки корма.
Уведомление, что пора покупать корм.
Уведомление, что пора идти делать пописы
?? интеграция с заказом корма в магазине?
