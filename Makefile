.PHONY: db-up db-down

db-up:
	@echo "Running database migrations..."
	goose postgres "host=localhost port=5432 user=postgres password=postgres dbname=learning sslmode=disable" up

db-down:
	@echo "Rolling back database migrations..."
	goose postgres "host=localhost port=5432 user=postgres password=postgres dbname=learning sslmode=disable" down
