.PHONY: openapi
openapi: openapi-http

.PHONY: openapi-http
openapi-http:
	@./scripts/openapi-http.sh lending internal/lending/ports ports