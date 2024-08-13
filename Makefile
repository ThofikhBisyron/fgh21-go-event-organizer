host ?= localhost
port ?= 5432
user ?= postgres
pass ?= 12345678
db ?= event_organizer

migrate\:init:
	psql -U$(user) -d postgres -p $(port) -c "create database $(db);"

migrate\:drop:
	psql -U$(user) -d postgres -p $(port) -c "drop database if exists $(db) with (force);"

migrate\:up:
	migrate -database postgresql://$(user):$(pass)@$(host):$(port)/$(db)?sslmode=disable -path migrations up $(version)

migrate\:down:
	migrate -database postgresql://$(user):$(pass)@$(host):$(port)/$(db)?sslmode=disable -path migrations down $(version)

migrate\:reset: 
	$(MAKE) migrate:drop user=$(user) db=$(db)
	$(MAKE) migrate:init user=$(user) db=$(db)
	$(MAKE) migrate:up user=$(user) pass=$(pass) db=$(db)