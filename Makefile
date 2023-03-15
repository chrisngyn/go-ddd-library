.PHONY: openapi
openapi: openapi-http

.PHONY: openapi-http
openapi-http:
	@./scripts/openapi-http.sh lending internal/lending/ports ports
	@./scripts/openapi-http.sh catalogue internal/catalogue main


.PHONY: migrate-up
migrate-up:
	@./scripts/migrate-up.sh lending lending
	@./scripts/migrate-up.sh lending_test lending
	@./scripts/migrate-up.sh catalogue catalogue
	@./scripts/migrate-up.sh catalogue_test catalogue


.PHONY: sqlboiler
sqlboiler:
	@./scripts/sqlboiler.sh lending
	@./scripts/sqlboiler.sh catalogue

.PHONE: go-vet
go-vet:
	@(cd ./internal/common && go vet ./...)
	@(cd ./internal/lending && go vet ./...)
	@(cd ./internal/catalogue && go vet ./...)


.PHONE: go-download
go-download:
	@(cd ./internal/common && go mod download)
	@(cd ./internal/lending && go mod download)
	@(cd ./internal/catalogue && go mod download)

.PHONE: go-tidy
go-tidy:
	@(cd ./internal/common && go mod tidy)
	@(cd ./internal/lending && go mod tidy)
	@(cd ./internal/catalogue && go mod tidy)
