# Как запустить проект:

1. Скачать проект с GitHub
2. Для раскатки зависимостей использую go-migrate ссылка на документацию: https://github.com/golang-migrate/migrate?tab=readme-ov-file


### Установка go-migrate
```
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.darwin-amd64.tar.gz -o migrate.tar.gz
tar -xvf migrate.tar.gz
sudo mv migrate.darwin-amd64 /usr/local/bin/migrate
```

Используя пакетный менеджер
```
brew install golang-migrate
```


1. Перейти в папку проекта
2. Выполнить команду ```make dep```
3. Выполнить команду ```make migrate-up```

Приложение поднято на 10000 порту, можно делать запросы к нему.
Psql поднято на 5432 в докере и 12000 проброшен наружу.