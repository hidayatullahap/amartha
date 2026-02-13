.PHONY: wire
wire:
	go mod tidy
	go run -mod=mod github.com/google/wire/cmd/wire src/wire/wire.go

.PHONY: run
run:
	go run main.go

.PHONY: sqlc-gen
sqlc-gen:
	.\bin\sqlc.exe generate