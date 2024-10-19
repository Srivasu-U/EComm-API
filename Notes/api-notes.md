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