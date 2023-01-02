start_database:
	docker compose up --build -d

start_development_server:
	CATAPI_ENV=development ${GOPATH}/bin/gin run main.go

migrate:
	${GOPATH}/bin/migrate -path ./database/migrations -database "mysql://catapi:secret@tcp(localhost:33060)/catapi?charset=utf8mb4&parseTime=True&loc=Local" up

test:
	CATAPI_ENV=test ${GOPATH}/bin/gotestsum --format testname ./tests/feature
