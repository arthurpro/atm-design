# ATM Design Test

The program has been written and tested using **Golang 1.18** on macOS.

This program uses SQLite 3 for the database stored in `db.sqlite` (once initialized).
Logs will be written to `app.log`.

### Commands

Using `make`:

```
# Run program
make run

# Run unit tests
make test

# Build executable
make build
```

Without `make`:

```
# Run program
go run main.go

# Run unit tests
go test -v

# Build executable
go build
```