# Gin REST API

This is a basic REST API I implemented using Gin while learning more about the framework.
This project also uses gorm and Postgres (via Docker).

## Starting the project

Start the DB:

```bash
docker compose -f ./resources/docker-compose.yml up
```

Run the tests:

```bash
go test
```

Start the server:

```bash
air
```

Ps: air provides live-reload for the server

Access the endpoint:

```bash
http://localhost:8080/students
```

A Postman collection can be found inside /docs
