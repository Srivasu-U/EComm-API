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
- When I ran `make run`, everything seemed to run successfully but as soon as I sent a request to `/register` on Thunder Client, I got a huge panic, like this
```
2024/10/20 19:23:25 DB Ping successful
2024/10/20 19:23:25 Listening on :8080
2024/10/20 19:30:05 http: panic serving 127.0.0.1:50218: runtime error: invalid memory address or nil pointer dereference
goroutine 17 [running]:
net/http.(*conn).serve.func1()
        /usr/local/go/src/net/http/server.go:1947 +0xbe
panic({0x771b20?, 0xac1600?})
        /usr/local/go/src/runtime/panic.go:785 +0x132
database/sql.(*DB).conn(0x0, {0x874be0, 0xaf53a0}, 0x1)
        /usr/local/go/src/database/sql/sql.go:1309 +0x54
database/sql.(*DB).query(0x0, {0x874be0, 0xaf53a0}, {0x7e4d48, 0x23}, {0xc0001718b0, 0x1, 0x1}, 0x40?)
        /usr/local/go/src/database/sql/sql.go:1751 +0x57
database/sql.(*DB).QueryContext.func1(0x20?)
        /usr/local/go/src/database/sql/sql.go:1734 +0x4f
database/sql.(*DB).retry(0x10?, 0xc0001717f0)
        /usr/local/go/src/database/sql/sql.go:1568 +0x42
database/sql.(*DB).QueryContext(0xc0000a6590?, {0x874be0?, 0xaf53a0?}, {0x7e4d48?, 0xc0000ce240?}, {0xc0001718b0?, 0x7a2820?, 0xc0000ce240?})
        /usr/local/go/src/database/sql/sql.go:1733 +0xc5
database/sql.(*DB).Query(...)
        /usr/local/go/src/database/sql/sql.go:1747
github.com/Srivasu-U/EComm-API/service/user.(*Store).GetUserByEmail(0xc000164af0?, {0xc0000a6590?, 0xaf53a0?})
        /home/chiltu/Go-Proj/EComm-API/service/user/store.go:19 +0x76
github.com/Srivasu-U/EComm-API/service/user.(*Handler).handleRegister(0xc0000aa060, {0x874708, 0xc0000f8000}, 0xc0000c4640)
        /home/chiltu/Go-Proj/EComm-API/service/user/routes.go:48 +0x17b
net/http.HandlerFunc.ServeHTTP(0xc0000c4500?, {0x874708?, 0xc0000f8000?}, 0x10?)
        /usr/local/go/src/net/http/server.go:2220 +0x29
github.com/gorilla/mux.(*Router).ServeHTTP(0xc0000c0000, {0x874708, 0xc0000f8000}, 0xc0000c43c0)
        /home/chiltu/go/pkg/mod/github.com/gorilla/mux@v1.8.1/mux.go:212 +0x1e2
net/http.serverHandler.ServeHTTP({0x873430?}, {0x874708?, 0xc0000f8000?}, 0x6?)
        /usr/local/go/src/net/http/server.go:3210 +0x8e
net/http.(*conn).serve(0xc0000e8000, {0x874c88, 0xc0000a8540})
        /usr/local/go/src/net/http/server.go:2092 +0x5d0
created by net/http.(*Server).Serve in goroutine 1
        /usr/local/go/src/net/http/server.go:3360 +0x485
```
- I'm still on certain on how to "correctly" debug this, but the way I did it was following the stack trace, bottom-up, and checking the function calls
    - The culprit was `s.db.Query(...)` within the `store.go`. 
    - `s.db` was nil *facepalm* as it was in `cmd/main.go` => `server := api.NewApiServer(":8080", nil)` instead of `server := api.NewApiServer(":8080", db)`
    - So `Query()` was being called on a nil object. My bad but I am glad I was able to figure this out

## Migrations
- These are the tables used

![tablesUsed](/Notes/assets/tablesUsed.png)  

- Database migrations is a way to maintain a history of changes to the database so that it can be reproduced anywhere just using the files
    - Basically is a log of changes that can be executed
    - Kind of akin to git
- You can call `up` on the migrations to make the changes or you can call `down` on the migrations to revert the changes
    - That is, we run `main.go` through the terminal and pass in either `up` or `down` as the args to run the relevant migration file
    ```
    go run cmd/migrate/main.go up  // Run migration up file
    go run cmd/migrate/main.go down // Run migration down file
    ```
- To create new migration files, we can use the following command `make migration <migration-name>` which runs the relevant command from the `Makefile`
- We can run either `make migrate-up` or `make migrate-down` to run the relevant migration up or down files
- If there is an error in the migration which results in a dirty read, the following command can be run
```
migrate -path <PATH-TO-MIGRATIONS> -database "mysql://root:<password>@tcp(URL)/<db-name>" force <migration-number>
```
- `URL` is `localhost:3306` for me and `<db-name>` is `EComm`
- To get the `migration-number`, I just run `make migrate-up` which gives me an error message with the number in it *shrug*
- Also, if push comes to shove, just copy the entire up and down migration files, delete the entire DB from MySQL and start anew. 
- The tables must be created in the order of their dependencies. 
    - For example, in this project, the `order_items` table is dependent on both `orders` and `products` table existing as it uses the ids as foreign keys
    - This essentially means that the migrations for `orders` and `products` must be before `order_items`

## Questions
- *Why have a storage layer?*
- *How do you quantify what is a service and what isn't?*
    - *What are the design principles that go into this?*

## TODO
- TODO: Dockerize this application