MIGRATIONS_FOLDER=db/migrations
CA_FILE := ./certs/ca.pem

# Create a new migration -> make migrate-create name="create_entity_table
migrate-create:
	goose -dir $(MIGRATIONS_FOLDER) create $(name) sql

migrate-up: # -> make migrate-up
	goose -dir $(MIGRATIONS_FOLDER) -certfile $(CA_FILE) up

migrate-down: # -> make migrate-down
	goose -dir $(MIGRATIONS_FOLDER) -certfile $(CA_FILE) down

# Rollback all migrations and reset database # make migrate-reset
migrate-reset:
	goose -dir $(MIGRATIONS_FOLDER) -certfile $(CA_FILE) reset

# Show current migration status # make migrate-status
migrate-status:
	goose -dir $(MAKEFILE_FOLDER) -certfile $(CA_FILE) status

# Redo last migration (Down then Up) # make migrate-redo
migrate-redo:
	goose -dir $(MIGRATIONS_FOLDER) -certfile $(CA_FILE) redo

# Run specific migration version # make migrate-version version=20200101120000
migrate-to:
	goose -dir $(MAKEFILE_FOLDER) -certfile $(CA_FILE) up-to $(version)

# Rollback to a specific migration version # make migrate-down-to version=20200101120000
migrate-down-to:
	goose -dir $(MAKEFILE_FOLDER) -certfile $(CA_FILE) down-to $(version)

# Force a specific migration version # make migrate-force version=20200101120000
migrate-force:
	goose -dir $(MAKEFILE_FOLDER) -certfile $(CA_FILE) force $(version)

# Print Goose help # make migrate-help
migrate-help:
	goose -h