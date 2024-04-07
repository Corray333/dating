.SILENT:
build:
	cd backend && make build
lint:
	cd backend && make lint
run: build
	cd backend && make run