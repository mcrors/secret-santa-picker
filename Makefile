include .env
export

DATABASE_URL ?= postgresql://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST_DEV}:5432/${DATABASE_NAME}?sslmode=disable

ifeq ($(ENV),test)
DATABASE_URL := postgresql://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST_TEST}:5432/${DATABASE_NAME}?sslmode=disable
endif
ifeq ($(ENV),prod)
DATABASE_URL := postgresql://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST_PROD}:5432/${DATABASE_NAME}?sslmode=disable
endif

GINKGO := go run github.com/onsi/ginkgo/v2/ginkgo

.PHONY: build migration_up migration_down migration_fix test test-unit

build:
	@go build -o bin/secret-santa cmd/main.go

run: build
	@./bin/secret-santa

migration_up:
	@echo ${DATABASE_URL}
	@migrate -path database/migration/ -database ${DATABASE_URL} -verbose up

migration_down:
	@echo ${DATABASE_URL}
	@migrate -path database/migration/ -database ${DATABASE_URL} -verbose down $(BACK)

migration_fix:
	@migrate -path database/migration/ -database ${DATABASE_URL} force $(VERSION)

test:
	@$(GINKGO) -r --race --cover --fail-fast

# Run only unit tests (requires Label("unit") in specs)
test-unit:
	@ENV=test $(GINKGO) -r --label-filter="unit" --race --cover --fail-fast
