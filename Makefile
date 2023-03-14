.PHONY: openapi
openapi: openapi-http

.PHONY: openapi-http
openapi-http:
	@./scripts/openapi-http.sh lending internal/lending/ports ports


.PHONY: migrate-up
migrate-up:
	@./scripts/migrate-up.sh lending lending
	@./scripts/migrate-up.sh lending_test lending


.PHONY: sqlboiler
sqlboiler:
	@./scripts/sqlboiler.sh lending

.PHONE: go-vet
go-vet:
	@(cd ./internal/common && go vet ./...)
	@(cd ./internal/lending && go vet ./...)


.PHONE: go-download
go-download:
	@(cd ./internal/common && go mod download)
	@(cd ./internal/lending && go mod download)

.PHONE: go-tidy
go-tidy:
	@(cd ./internal/common && go mod tidy)
	@(cd ./internal/lending && go mod tidy)
