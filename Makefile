start_database:
	docker compose up --build

start_development_server:
	HERMES_ENV=development ${GOPATH}/bin/gin run main.go

migrate:
	${GOPATH}/bin/migrate -path ./database/migrations -database "mysql://hermes:secret@tcp(localhost:33060)/hermes_messaging?charset=utf8mb4&parseTime=True&loc=Local" up

test:
	HERMES_ENV=test ${GOPATH}/bin/gotestsum --format testname ./tests/feature
