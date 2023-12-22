include .env
export

DATABASE_URL ?= postgresql://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST_DEV}:5432/${DATABASE_NAME}?sslmode=disable

ifeq ($(ENV),test)
DATABASE_URL := postgresql://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST_TEST}:5432/${DATABASE_NAME}?sslmode=disable
endif
ifeq ($(ENV),prod)
DATABASE_URL := postgresql://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST_PROD}:5432/${DATABASE_NAME}?sslmode=disable
endif

.PHONY: migration_up migration_down migration_fix

migration_up:
	@echo ${DATABASE_URL}
	migrate -path database/migration/ -database ${DATABASE_URL} -verbose up

migration_down:
	@echo ${DATABASE_URL}
	migrate -path database/migration/ -database ${DATABASE_URL} -verbose down $(BACK)

migration_fix:
	migrate -path database/migration/ -database ${DATABASE_URL} force $(VERSION)
