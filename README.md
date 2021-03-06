# linqs

## 🚀 How to Setup

### Backend (database & api)
```go
// cd database/src/
$ make migrate/init
$ make migrate/up
$ docker-compose up
```
```go
// cd api/src/
$ make .env
$ make run
```

### Frontend (ui)
```sh
# cd ui
# install dependencies
$ yarn add

# start localhost
$ yarn start
```

## Test (api)
Test files are in the same directory with the files you need to test.

like this.

```sh
controller/
  fizz.go
  fizz_test.go
  buzz.go
  buzz_test.go
```

```go
# /api/src/
$ go test ./... # run all tests
```

## DB Structure

https://hackmd.io/UwTMbawbRMqKyKqBBn1p8Q


## Endpoint Management (Swagger)


```sh
# /swagger
$ docker-compose up
$ open localhost:8080
```

Check this URL.
https://qiita.com/rllllho/items/53a0023b32f4c0f8eabb

Write the endpoint.yml file.

## Prefix Commit Message
Coomit template will appear after below command.
```sh
$ git config commit.template .gitmessage 
```
