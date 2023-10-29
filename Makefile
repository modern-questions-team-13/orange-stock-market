ifeq ($(POSTGRES_SETUP_TEST),)
	POSTGRES_SETUP_TEST := user=test password=test dbname=test host=localhost port=5433 sslmode=disable
endif

ifeq ($(POSTGRES_SETUP_PROD),)
	POSTGRES_SETUP_PROD := user=user password=postgrespw dbname=postgres host=localhost port=5432 sslmode=disable
endif

INTERNAL_PKG_PATH=$(CURDIR)/internal/pkg
MIGRATION_FOLDER=$(CURDIR)/migrations
TEST_DOCKER_FILE_PATH=$(CURDIR)/tests/repository/docker-compose.yaml

.PHONY: test-env-up
test-env-up:
	docker-compose -f "$(TEST_DOCKER_FILE_PATH)" up -d

.PHONY: migration-create
migration-create:
	goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql

.PHONY: test-migration-up
test-migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" up

.PHONY: test-migration-down
test-migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" down

.PHONY: unit-test
unit-test:
	go test ./internal/pkg/repository/pgx

.PHONY: unit-test-coverage
unit-test-coverage:
	go test ./internal/pkg/repository/pgx -cover

.PHONY: integration-test
integration-test:
	go test ./tests/... -tags=integration