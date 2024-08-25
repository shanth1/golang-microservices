.PHONY: gen

run-dev:
	go run cmd/sso/main.go --config="./config/development.yaml"

run-migrations:
	go run cmd/migrator/main.go --storage-path="./storage/sso.db" --migrations-path="migrations"

run-test-migrations:
	go run cmd/migrator/main.go --storage-path="./storage/sso.db" --migrations-path="tests/migrations" --migrations-table="migrations_test"
