module github.com/chiennguyen196/go-library/internal/lending

go 1.19

replace github.com/chiennguyen196/go-library/internal/common => ./../common

require (
	github.com/chiennguyen196/go-library/internal/common v0.0.0-00010101000000-000000000000
	github.com/deepmap/oapi-codegen v1.12.4
	github.com/go-chi/chi/v5 v5.0.8
	github.com/go-chi/render v1.0.2
	github.com/pkg/errors v0.9.1
)

require (
	github.com/ajg/form v1.5.1 // indirect
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/rs/zerolog v1.29.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
)
