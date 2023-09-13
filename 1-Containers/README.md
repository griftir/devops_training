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


## Dockerizing

### Building the image from the Dockerfile
```
docker build -t hello .
```
### Run the image as a container
```
docker run --name hello1 -e -p 8080:8080 hello
```

### Run the image with environment variable
```
docker run --name hello1 -e NAME=Vince -p 8080:8080 hello
```

### Showing built images
```
docker image ls
```

### Showing running containers
```
docker container ls
```

### Stopping and removing 
```
docker stop hello1
docker remove hello1
```

### Running detached





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