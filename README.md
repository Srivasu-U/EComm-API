# EComm-API
REST APIs in Golang with JWT Auth and MySQL datastore. The structure/APIs used are as follows:  

![APIs](/Notes/assets/apiStructure.png)  

## Requirements
- Golang, naturally
- MySQL server (should be up and running)
- Golang Migrate

## Running the migrations
- After making certain that MySQL or any SQL database is running, create the DB needed with the server
- Run the migrations
```
make migrate-up
```
- To undo the migrations in any case, run
```
make migrate-down
```
- Run the project using 
```
make run
```
- Run tests using 
``` 
make test
```
- All the relevant commands can be seen under the [Makefile](/Makefile)
