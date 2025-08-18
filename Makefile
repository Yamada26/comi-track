run:
	go run main.go

test:
	go test ./...

sqlite:
	sqlite3 app.db