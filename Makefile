build:
	@go build -o bin/api
run:
	@./bin/api
seed:
	@go run seed/seed.go