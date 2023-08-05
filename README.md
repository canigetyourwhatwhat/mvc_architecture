# mvc_architecture_without_view
It is dedicated to demonstrating the MVC software architecture without View 
> note: this repository is referenced from [this](https://github.com/gieart87/gotoko) repository

## Directory structure
The root directory contains an "app" directory which contains all the back-end codes.
A directory "database" contains methods and queries to set up a database.
Inside the directory, it has the below directories.
- controllers
  -   it contains all the business logic as Controller in MVC architecture
- models
  -   it contains all the models as Model in MVC architecture 


## Prerequisite:
-  You have installed Go version 1.20
-  You have installed MySQL locally or you run MySQL docker container.


## To run the server
```go
go run main.go
```
