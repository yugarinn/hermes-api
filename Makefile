start_database:
	docker compose up --build -d

start_development_server:
	PIGEON_ENV=development ${GOPATH}/bin/gin run main.go

migrate:
	${GOPATH}/bin/migrate -path ./database/migrations -database "mysql://pigeon:secret@tcp(localhost:33060)/pigeon_messaging?charset=utf8mb4&parseTime=True&loc=Local" up

test:
	PIGEON_ENV=test ${GOPATH}/bin/gotestsum --format testname ./tests/feature
