run:
	go run cmd/app/main.go

test:
	go test ./...

sqlite:
	sqlite3 app.db