PATH_THIS:=$(realpath $(dir $(lastword ${MAKEFILE_LIST})))
DIR:=$(PATH_THIS)

include base.env
VARS:=$(shell sed -ne 's/ *\#.*$$//; /./ s/=.*$$// p' base.env )
$(foreach v,$(VARS),$(eval $(shell echo export $(v)="$($(v))")))

include .env
VARS:=$(shell sed -ne 's/ *\#.*$$//; /./ s/=.*$$// p' .env )
$(foreach v,$(VARS),$(eval $(shell echo export $(v)="$($(v))")))

CODE=\033[0;33m
NAME=\033[0;32m
NC=\033[0m # No Color

#    |\__/,|   (`\
#  _.|o o  |_   ) )
#-(((---(((--------

help:
	@echo "${NAME}  /')   |,\__/|${NC}"
	@echo "${NAME} ( (   _|  o o|._   go-api-boilerplate${NC}"
	@echo "${NAME}--------)))---)))-----------------------${NC}"
	@echo ""
	@echo "${NAME}server${NC}"
	@echo "    Run server"
	@echo "${NAME}new_migration${NC}"
	@echo "    Create a new migration."
	@echo "    Example: ${CODE}make new_migration type=sql name=MigrationName${NC}"
	@echo "${NAME}migrate${NC}"
	@echo "    Run migrations"
	@echo "${NAME}rollback${NC}"
	@echo "    Rollback last migration"
	@echo "${NAME}seed${NC}"
	@echo "    Run seeders"
	@echo "${NAME}test${NC}"
	@echo "    Run tests"
	@echo "${NAME}swagger${NC}"
	@echo "    Generate Swagger documentation"
	@echo ""


.PHONY: server
server:
	@cd $(DIR) \
	&& air server

.PHONY: new_migration
new_migration:
	@cd $(DIR) \
	&& goose create -type $(type) $(name)

.PHONY: migrate
migrate:
	@cd $(DIR) \
	&& goose up

.PHONY: rollback
rollback:
	@cd $(DIR) \
	&& goose down

.PHONY: seed
seed:
	@cd $(DIR) \
	&& go run main.go seed

.PHONY: test
test:
	@cd $(DIR) \
	&& go test ./...

.PHONY: swagger
swagger:
	@cd $(DIR) \
	&& swagger generate spec -o doc/swagger.json
