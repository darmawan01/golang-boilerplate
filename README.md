# Kodingworks: Go Boilerplate Test

Reference Go project layout https://github.com/golang-standards/project-layout

Database: postgres 1.12

Go: go1.12.17 darwin/amd64

## Manual

Preparation, Create database with:
```
Db Name: kodingworks
DB pass: kodingworks
DB user: kodingworks
```

Makefile Command

- `make serve` -> run the application
- `make migrate` -> migrate with migrate image
- `make native_migrate` -> migrate with migrate binary
- `make migrate_clean` -> clean the migrate dir
- `make test` -> run the test