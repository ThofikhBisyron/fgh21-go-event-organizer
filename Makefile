# host ?= 143.198.222.47
# port ?= 54321
# user ?= postgres
# pass ?= 1
# db ?= event_organizer

include .env

migrate\:init:
PGPASSWORD=$(DB_PASSWORD) psql -U$(DB_USER) -d postgres -h $(DB_HOST) -p $(DB_PORT) -c "create database $(DB_NAME);"

migrate\:drop:
PGPASSWORD=$(DB_PASSWORD) psql -U$(DB_USER) -d postgres -h $(DB_HOST) -p $(DB_PORT) -c "drop database if exists $(DB_NAME) with (force);"

migrate\:up:
	migrate -database postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable -path migrations up $(version)

migrate\:down:
	migrate -database postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable -path migrations down $(version)


migrate\:reset: 
	$(MAKE) migrate:drop user=$(DB_USER) db=$(DB_NAME)
	$(MAKE) migrate:init user=$(DB_USER) db=$(DB_NAME)
	$(MAKE) migrate:up user=$(DB_USER) pass=$(DB_PASSWORD) db=$(DB_NAME)



#LINUX
# migrate\:init:
# 	PGPASSWORD=$(DB_PASSWORD) psql -U$(DB_USER) -d postgres -h $(DB_HOST) -p $(DB_PORT) -c "create database $(DB_NAME);"

# migrate\:drop:
# 	PGPASSWORD=$(DB_PASSWORD) psql -U$(DB_USER) -d postgres -h $(DB_HOST) -p $(DB_PORT) -c "drop database if exists $(DB_NAME) with (force);"

# migrate\:up:
# 	migrate -database postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable -path migrations up $(version)

# migrate\:down:
# 	migrate -database postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable -path migrations down $(version)

# migrate\:reset: 
# 	$(MAKE) migrate:drop
# 	$(MAKE) migrate:init
# 	$(MAKE) migrate:up