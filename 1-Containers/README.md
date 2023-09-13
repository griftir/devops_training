# Containers

## Commands

### Running the toy app
```
go run main.go
```

### Testing the toy app
```
go test ./...
```

### Building the toy app
```
go build -o hello ./main.go
```



### Handy Docker commands
#### Postres Database
```
docker run --name pg -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres
```
#### Connect to the db
```
psql -h localhost -U postgres
```
#### Get a shell on the container running postgres
```
docker exec -it pg bash
```