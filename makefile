PSQL_URL = postgres://postgres:postgres@localhost:13000/management-system?sslmode=disable


dep:
	@docker-compose up -d
	@sleep 3
	@$(MAKE) migrate-up


down:
	@docker-compose down

migrate-up:
	migrate -path ./migrations -database $(PSQL_URL) up


migrate-down:
	migrate -path ./migrations -database $(PSQL_URL) down



.PHONY: migrate-up migrate-down dep down

