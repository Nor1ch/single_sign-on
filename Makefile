.PHONY: migrate migrate_test appgo

migrate:
	go run ./cmd/migrator --storage-path=./storage/sso.db --migrations-path=./migrations

migrate_test:
	go run ./cmd/migrator/main.go --storage-path=./storage/sso.db --migrations-path=./tests/migrations --migrations-table=migrations_test

appgo:
	go run ./cmd/sso/main.go --config=./config/local.yaml



