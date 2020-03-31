# Kodingworks: Go Boilerplate Test

This apps was deployed at heroku: https://kodingworks-test.herokuapp.com/

Test the api with (postman):
- Download the file kodingworks.postman_collection.json for api docs
- Smport to postman
- Setup your environtment with key `BASE_URL` and value to https://kodingworks-test.herokuapp.com
- Try to send request of each module

## Clone & Run in your machine

Preparation, Create database with:
```
Db Name: kodingworks
DB pass: kodingworks
DB user: kodingworks
```

Makefile Command:

- `make serve` -> run the application
- `make migrate` -> migrate with migrate image
- `make native_migrate` -> migrate with migrate binary
- `make migrate_clean` -> clean the migrate dir
- `make test` -> run the test

```
NOTE: Modify the config.json and adjust it to your needs
```

# Project structure

**cmd/kodingworks**

    Directory for main.go file

**db/migrations**

    Collecting all migrations from every module and put in here together

**module example `guests`**

    This is the module, inside this dir we have api.go,handler.go, and model.go

**test**

    Here where we put our test from all module ~not ready yet~

**utils**

    Here is out utils/extensions to reduce the code and not repeat the same code


# Stack

Reference Go project layout https://github.com/golang-standards/project-layout

Database: postgres 1.12

Go: go1.12.17 darwin/amd64