<h5>Установка утилиты migrate для миграций</h5>

```bash
https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md
```

<h5>Запуск БД в Докере</h5>

```bash
sudo docker run --name todo-app -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres
```

<h5>Создание Миграций</h5>

```bash
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up
```

<h5>Зайти в базу данных</h5>

```bash
sudo docker exec -it a2bc6a20d679 bin/bash
psql -U postgres
```