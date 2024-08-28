# maptalk

## Quick Start

```console
go run cmd/maptalk/main.go
```

## Test

```console
go generate ./...
```

```console
go test ./...
```

## Coverage

```console
go test -cover ./... -coverprofile=cover.out
go tool cover -html=cover.out -o cover.html
open cover.html
```
