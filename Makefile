host ?= 103.93.58.89
port ?= 5432
user ?= postgres
pass ?= 1
db ?= event_organizer

migrate\:init:
	PGPASSWORD=$(pass) psql -U$(user) -d postgres -h $(host) postgres -p $(port) -c "create database $(db);"

migrate\:drop:
	PGPASSWORD=$(pass) psql -U$(user) -d postgres -h $(host) postgres -p $(port) -c "drop database if exists $(db) with (force);"

migrate\:up:
	migrate -database postgresql://$(user):$(pass)@$(host):$(port)/$(db)?sslmode=disable -path migrations up $(version)

migrate\:down:
	migrate -database postgresql://$(user):$(pass)@$(host):$(port)/$(db)?sslmode=disable -path migrations down $(version)

migrate\:reset: 
	$(MAKE) migrate:drop user=$(user) db=$(db)
	$(MAKE) migrate:init user=$(user) db=$(db)
	$(MAKE) migrate:up user=$(user) pass=$(pass) db=$(db)