PG_CONTAINER_NAME = "test-postgres"
PG_USER = "testemp"
LOCAL_BIN:=$(CURDIR)/bin

run:
		cd cmd/app && go run main.go

up:
		docker-compose up

test: 
	@make setup-db 
	@make test-integration 
	@make db-teardown

setup-db:
		docker-compose -f docker-compose.test.yml down
		docker volume rm avitopvz_v3_postgres_test_volume || true
		docker-compose -f docker-compose.test.yml up -d
		@echo "Ожидание готовности PostgreSQL..."
		@until docker exec $(PG_CONTAINER_NAME) pg_isready -U $(PG_USER) > /dev/null 2>&1; do \
			echo "Postgres не готов, ждём..."; \
			sleep 1; \
 		done


test-integration:
		go  test -tags=integration ./tests/integration/...

db-teardown:
		docker-compose -f docker-compose.test.yml down

unit-test:
		go test -coverpkg=./... -coverprofile=cover ./... && cat cover | grep -v "mock" | grep -v "/internal/storage/" | grep -v "easyjson" | grep -v "/gen/" | grep -v "pvz.pb.go" | grep -v "pvz_grpc.pb.go" > cover.out && go tool cover -func=cover.out
		
install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
		
generate-pvz-api:
		mkdir -p pkg/pvz_v1
		protoc --proto_path proto \
		--go_out=pkg/pvz_v1 --go_opt=paths=source_relative \
		--plugin=protoc-gen-go=bin/protoc-gen-go \
		--go-grpc_out=pkg/pvz_v1 --go-grpc_opt=paths=source_relative \
		--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
		proto/pvz.proto