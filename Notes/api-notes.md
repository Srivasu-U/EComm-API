## General notes on this project
- The structure/APIs used are as follows:
![APIs](/Notes/assets/apiStructure.png)  

- DB used is MySQL. Application is Dockerized
- Using migrations to be able to use any database.

- In Golang, to build APIs, there are multiple frameworks, such as `Gin`, `Gorilla Mux`, `Chi` and `Fibre`
    - We are using `Gorilla Mux`
- We need to create an HTTP server and handler
    - This is done within the `cmd` folder. The convention in Golang is that `cmd` contains all the entry points.
        - The entry points are going to be the APIs themselves and the migrations
        - The entry point for the API is `main.go`. All the migrations are stored in `migrate` and all the apis in `api`.
- A server, for us, is a struct with an address and a DB pointer.
- We also need a router for the endpoints registration
    - Routers provide a centralized way to check API handlers and start from there for any modifications

## Practices
- All the config values as needed for things such as DB connection are stored in `config/env.go`
    - This has a `Config` struct with an initial configuration applied where the values are either lifted from the environment or have default values as fallback
- Such an approach is good practice in general
- Having only worked with `Mux`, I have to say that this is surprisingly verbose for Go. For example, having to do this for getting a JSON payload
```
var payload types.RegisterUserPayload
err := json.NewDecoder(r.Body).Decode(payload) 
```
- This is essentially taking a payload of type `RegisterUserPayload` (which is a struct) and decodes into.
    - `NewDecoder` takes in a `http.request` Body and returns a `Decoder` pointer
- Usage of `Repository Design Pattern`: there is an intermediate layer between the business logic and data storage
    - Provide standard way to access and manipulate data while abstracting away from the actual underlying data store tech.

## Questions
- *Why have a storage layer?*
- *How do you quantify what is a service and what isn't?*
    - *What are the design principles that go into this?*