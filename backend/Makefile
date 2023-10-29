include .env
SOURCE := ./migrations

db-force:
	migrate -source file:$(SOURCE) -database mysql://$(DATABASE_CONNECTION_STRING) force $(V)

db-version:
	migrate -source file:$(SOURCE) -database mysql://$(DATABASE_CONNECTION_STRING) version

db-up:
	migrate -source file:$(SOURCE) -database mysql://$(DATABASE_CONNECTION_STRING) up $(V)

db-down:
	migrate -source file:$(SOURCE) -database mysql://$(DATABASE_CONNECTION_STRING) down $(V)

db-create:
	migrate -source file:$(SOURCE) -database mysql://$(DATABASE_CONNECTION_STRING) create -ext sql -dir $(SOURCE) -seq $(name)

echo-source:
	@echo "$(SOURCE)"

echo-database:
	@echo "$(DATABASE_CONNECTION_STRING)"


SHELL = /bin/sh
RAND = $(shell echo $$RANDOM)
