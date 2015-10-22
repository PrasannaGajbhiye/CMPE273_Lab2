# CMPE273_Lab2
Making HTTP Requests using HTTP Router & Json format input in golang

## Usage

### Install

```
go get github.com/PrasannaGajbhiye/CMPE273_Lab2
```

Build & Run using following commands

```
cd CMPE273_Lab2
go get 
go build
./httpRouter.go
```

### New terminal to make requests 
```
curl -H "Content-Type: application/json" -X POST -d '{"name":"Foo Bar"}' http://127.0.0.1:8080/hello
```
Following will be the response for the above request:
```
{"Greeting":"Hello, Foo Bar!"}
```

