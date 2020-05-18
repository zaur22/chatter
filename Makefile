

# tests запускает все виды тестирования
tests: tests-unit tests-integrations test-build

# tests-unit запускает юнит тесты
tests-unit:
	docker-compose -f docker-compose.yml -f docker-compose.test.yml down --volumes
	docker-compose -f docker-compose.yml -f docker-compose.test.yml up --abort-on-container-exit unit

# tests-integrations запускает интеграционные тесты
tests-integrations:
	docker-compose -f docker-compose.yml -f docker-compose.test.yml down --volumes
	docker-compose -f docker-compose.yml -f docker-compose.test.yml up --abort-on-container-exit postgres local-migrations
	docker-compose -f docker-compose.yml -f docker-compose.test.yml up --abort-on-container-exit postgres integration

# test-build запускает сборку проекта auro
# этот тест нужен для проверки возможности вообще компилировать проект
test-build:
	docker-compose -f docker-compose.yml -f docker-compose.test.yml down --volumes
	docker-compose -f docker-compose.yml -f docker-compose.test.yml up --abort-on-container-exit build

# gen-grpc генерирует grpc код из спецификации ./api/grpc/aura.proto
# Требуется установленный protoc и модуль для генерации go кода.
# Подробнее https://grpc.io/docs/quickstart/go/
gen-grpc:
	protoc -I ./api ./api/grpc/chatter.user.proto --go_out=plugins=grpc:pkg/rpc/
	protoc -I ./api ./api/grpc/chatter.chat.proto --go_out=plugins=grpc:pkg/rpc/
