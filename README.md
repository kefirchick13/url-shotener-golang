# Разворачивание приложения через docker compose 
docker compose up --build

# Поднятие миграций 
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5438/postgres?sslmode=disable' up
