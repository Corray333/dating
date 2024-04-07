.SILENT:
build:
	cd cmd && go build main.go
lint:
	cd backend && golangci-lint run
run: build
	cd cmd && ./main ../.env

goose-up:
	cd migrations && goose postgres "user=postgres password=root host=db port=5432 dbname=dating sslmode=disable" up
goose-down:
	cd migrations && goose postgres "user=postgres password=root host=db port=5432 dbname=dating sslmode=disable" down
swagger:
	swag init -g cmd/main.go