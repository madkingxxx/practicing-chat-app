include .env
export
MIGRATION_NAME = $(shell $1)

migrate-create:
	# echo "Creating migration $(name)"
	migrate create -seq -ext sql -dir  migrations $(name)

swag-init:
	swag init -g ./internal/controller/http/v1/router.go