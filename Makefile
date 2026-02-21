build: # сборка и запуск утилиты
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

lint: # проверка кода линтером golangci-lint
	golangci-lint run

lint-fix: # автоисправления и форматирование
	golangci-lint run --fix
	
test: # запуск тестов
	go test -v ./...
