.PHONY: run build init initdb

run: build
	# go run ./cmd/tontine/main.go
	./bin/tontine

build:
	env GOOS=darwin GOARCH=amd64 go build -o ./bin/tontine ./cmd/tontine

initdb:
	touch ./bin/db.db

init: build initdb

db-migrate-up:
	sql-migrate up
	make gen-model

db-migrate-down:
	sql-migrate down
	make gen-model

gen-model:
	sqlboiler sqlite3 -o ./src/model/dbmodel -p dbmodel --wipe --no-tests